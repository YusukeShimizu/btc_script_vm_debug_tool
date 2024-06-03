package main

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/boltdb/bolt"
	"github.com/btcsuite/btcd/wire"
)

const dbName = "db"
const bucketName = "txs"

type fetcher struct {
	db *bolt.DB
}

func NewFetcher() *fetcher {
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	return &fetcher{
		db: db,
	}
}

func (f *fetcher) Close() {
	f.db.Close()
}

var (
	bucketNotFoundError = errors.New("bucket not found")
	notFoundError       = errors.New("key not found")
)

func (f *fetcher) fetch(key string) (string, error) {
	value := ""
	err := f.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return bucketNotFoundError
		}
		value = string(bucket.Get([]byte(key)))
		if value == "" {
			return notFoundError
		}
		return nil
	})
	return value, err
}

func (f *fetcher) getTransaction(txid string) (*wire.MsgTx, error) {
	txHex, err := f.fetch(txid)
	if err != nil {
		log.Printf("failed to fetch transaction from db %s: %v\n", txid, err)
	}
	if txHex == "" {
		fmt.Printf("fetching transaction from api: %s\n", txid)
		url := fmt.Sprintf("https://mempool.space/api/tx/%s/hex", txid)
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		txHexb, err := io.ReadAll(io.Reader(resp.Body))
		if err != nil {
			return nil, err
		}
		err = f.db.Update(func(tx *bolt.Tx) error {
			b, err := tx.CreateBucketIfNotExists([]byte(bucketName))
			if err != nil {
				return err
			}
			return b.Put([]byte(txid), txHexb)
		})
		if err != nil {
			return nil, err
		}
		txHex = string(txHexb)
	}
	txBytes, err := hex.DecodeString(txHex)
	if err != nil {
		return nil, err
	}
	msgTx := wire.NewMsgTx(wire.TxVersion)
	err = msgTx.Deserialize(bytes.NewReader(txBytes))
	if err != nil {
		return nil, err
	}
	return msgTx, nil
}

// FetchPrevOutput attempts to fetch the previous output referenced by
// the passed outpoint. A nil value will be returned if the passed
// outpoint doesn't exist.
func (f *fetcher) FetchPrevOutput(outPoint wire.OutPoint) *wire.TxOut {
	prevTx, err := f.getTransaction(outPoint.Hash.String())
	if err != nil {
		fmt.Printf("failed to get transaction %s: %v\n", outPoint.Hash.String(), err)
		return nil
	}
	return prevTx.TxOut[outPoint.Index]
}
