package main

import (
	"fmt"
	"log"
	"os"

	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btclog"
)

func init() {
	// Set up a logger for the script package to use.
	logger := btclog.NewBackend(os.Stdout).Logger("SCRIPT")
	logger.SetLevel(btclog.LevelTrace)
	txscript.UseLogger(logger)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("txid is required as a command line argument.")
		os.Exit(1)
	}

	txid := os.Args[1]
	if txid == "" {
		fmt.Println("txid is required as a command line argument.")
		os.Exit(1)
	}

	fetcher := NewFetcher()
	msgTx, err := fetcher.getTransaction(txid)
	if err != nil {
		log.Fatalf("Failed to fetch data: %v", err)
	}

	for i, txIn := range msgTx.TxIn {
		fmt.Println("Input:", i)
		prevout := fetcher.FetchPrevOutput(txIn.PreviousOutPoint)
		vm, err := txscript.NewEngine(
			prevout.PkScript,
			msgTx,
			i,
			txscript.StandardVerifyFlags,
			nil,
			nil,
			prevout.Value,
			fetcher)
		if err != nil {
			log.Fatalf("failed to create script engine for input %d: %v", i, err)
		}
		// Execute the script.
		// Debugging allows you to follow the details of the script's movement.
		err = vm.Execute()
		if err != nil {
			log.Fatalf("failed to execute script for input %d: %v", i, err)
		}
	}
}
