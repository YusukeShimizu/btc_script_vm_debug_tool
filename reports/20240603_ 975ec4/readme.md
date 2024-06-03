tx id: [975ec405ac9dc9fa5ab8009d94d6a1fe31dff8a8127ea90d023104e52754e4d7](https://blockstream.info/tx/975ec405ac9dc9fa5ab8009d94d6a1fe31dff8a8127ea90d023104e52754e4d7)

* 各入力は9あるが、送付元アドレスはすべて同じ
* 送金額: 450,290,000,000サトシ（4,502.9 BTC）

## script

対象のtxのinputは、2 of 3のmulti sig  
OP CHECKMULTISIGを検証するため、[opcodeCheckMultiSig](./opcodeCheckMultiSig.md) を利用

### 参考
[OP CHECKMULTISIG](https://en.bitcoin.it/wiki/OP_CHECKMULTISIG)

```mermaid
graph LR
    A[スタート] --> B[OP_DATA_34<br>スタックにデータを追加]
    B --> C[OP_HASH160<br>スタック上のデータをハッシュ化]
    C --> D[OP_DATA_20<br>20バイトのデータを追加]
    D --> E[OP_EQUAL<br>スタック上の値を比較<br>結果が真の場合は01を追加]
    E --> F[OP_0<br>スタックに0を追加]
    F --> G[OP_DATA_32<br>32バイトのデータを追加]
    G --> H[署名1<br>0x3045022100...]
    H --> I[署名2<br>0x3044022041...]
    I --> J[OP_2<br>スタックに2を追加]
    J --> K[公開鍵1<br>0x022a3ec3...]
    K --> L[公開鍵2<br>0x03e8b765...]
    L --> M[公開鍵3<br>0x038228f8...]
    M --> N[OP_3<br>スタックに3を追加]
    N --> O[OP_CHECKMULTISIG<br>3つの公開鍵と2つの署名の検証]
    O --> P[検証結果]
```

各入力がP2SH形式をとっており、そのスクリプトが `220020...` のように始まるため、P2SH-P2WPKHである


## それぞれの署名組み合わせ
※: chat gpt出力のため間違いの可能性あり。正確な情報は[log trace](./trace)を参照

| tx in | signature                                                                                                                                      | pubkey                                                             | 検証ステータス                |
| ----- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------ | ----------------------------- |
| 0     | 3044022041f0589bb772c9e6078e3f10871523cf5ae2e305c14748eb69e18297885b587202204e75eb3186db434a4ed4c41d434c28cf137cf17048e0c20f3c8a7a584e1667aa   | 028228f84daf8d397fcf252af0114c53db7a354c7ffcf293e8a6997cb1b253ac8f | Signature verification failed |
| 0     | 3044022041f0589bb772c9e6078e3f10871523cf5ae2e305c14748eb69e18297885b587202204e75eb3186db434a4ed4c41d434c28cf137cf17048e0c20f3c8a7a584e1667aa   | 03e8b76584f2a41aadd97e461721875a7ced12ab0ba9266181eae68e95bc03e190 | Signature verified            |
| 0     | 3045022100e3592b14c96e3eefdf8e3bea8343533ed3b1cfe55e2df4d4b0e62973f4a5b3760220047ca50da99cceda5ce4ca50b281323cb03ff3c01f78926b4c457e885cad0e86 | 022a3ec327cd0b967b53e25439d6ab05a57e55ba82c3cdf03c94b83106711d43ff | Signature verified            |
| 1     | 3045022100ebf65adb8716d19f37f47dc597f22e7bf657e2df93284f29ebba06709856604702205d01a5d42c08cc6d4a8da4094d76d1cab4a7ddc537370919a56ddb9c6f64a379 | 028228f84daf8d397fcf252af0114c53db7a354c7ffcf293e8a6997cb1b253ac8f | Signature verification failed |
| 1     | 3045022100ebf65adb8716d19f37f47dc597f22e7bf657e2df93284f29ebba06709856604702205d01a5d42c08cc6d4a8da4094d76d1cab4a7ddc537370919a56ddb9c6f64a379 | 03e8b76584f2a41aadd97e461721875a7ced12ab0ba9266181eae68e95bc03e190 | Signature verified            |
| 1     | 3045022100a5e59ad2090bfac9151258feb01c8edd8970da1a97d20cc6ae421a4210745f6802202ec9eed06dddab59f5fa096946a3906b2ec367514e4444c1b81f7fc35ea9e136 | 022a3ec327cd0b967b53e25439d6ab05a57e55ba82c3cdf03c94b83106711d43ff | Signature verified            |
| 2     | 3045022100df4f310653902226bd91cc5b1c1a2dd3fa5afb133914c54574149f5d41bb8ce1022022e06249a4aac37d0cd096789a0d8359a5ca550d72d646e70674ea912a3d121e | 028228f84daf8d397fcf252af0114c53db7a354c7ffcf293e8a6997cb1b253ac8f | Signature verification failed |
| 2     | 3045022100df4f310653902226bd91cc5b1c1a2dd3fa5afb133914c54574149f5d41bb8ce1022022e06249a4aac37d0cd096789a0d8359a5ca550d72d646e70674ea912a3d121e | 03e8b76584f2a41aadd97e461721875a7ced12ab0ba9266181eae68e95bc03e190 | Signature verified            |
| 2     | 3044022028d544cb7076538f3c9fe0ab1261518261ae81b7fc96040f96376999f23df1530220552b7dd3fa777806ab8d5b8944b19e39ea49e5712f364d81632dd5680aabc27e   | 022a3ec327cd0b967b53e25439d6ab05a57e55ba82c3cdf03c94b83106711d43ff | Signature verified            |
| 3     | 304402206824a841606529de18fcc3443deef536212a298974e4ce26cf397b421929298f02206288a5c4be779d24d5e3c93bfed8b356850e67db343ae0cf2adf8bcd0ad54585   | 028228f84daf8d397fcf252af0114c53db7a354c7ffcf293e8a6997cb1b253ac8f | Signature verification failed |
| 3     | 304402206824a841606529de18fcc3443deef536212a298974e4ce26cf397b421929298f02206288a5c4be779d24d5e3c93bfed8b356850e67db343ae0cf2adf8bcd0ad54585   | 03e8b76584f2a41aadd97e461721875a7ced12ab0ba9266181eae68e95bc03e190 | Signature verified            |
| 3     | 30440220638141e7b6a81cd070010041d93e997b02b57d9c35184b9044af8fa55703fda902200a676455687b586cd10a0ba7cf79265275c3713d04aaaad6741f4bb6ed612be2   | 022a3ec327cd0b967b53e25439d6ab05a57e55ba82c3cdf03c94b83106711d43ff | Signature verified            |
| 4     | 304402202f395bd4641c679a392513fe7ccf27d626434dd2acc24b165cc96547e09c19a502207fdca8a6353ba21a21a8d5d6f6a1c6fed5a70dac4b4480478ced5b35525c3ced   | 028228f84daf8d397fcf252af0114c53db7a354c7ffcf293e8a6997cb1b253ac8f | Signature verification failed |
| 4     | 304402202f395bd4641c679a392513fe7ccf27d626434dd2acc24b165cc96547e09c19a502207fdca8a6353ba21a21a8d5d6f6a1c6fed5a70dac4b4480478ced5b35525c3ced   | 03e8b76584f2a41aadd97e461721875a7ced12ab0ba9266181eae68e95bc03e190 | Signature verified            |
| 4     | 3044022031a476de5088ea45cb9c33b16c03a46edd7a204ff4e5bc348422a4c1efc56020022052eaaed997a70549750989a2dd35f58184cea229c5ac68f77e074f58d7aebc2e   | 022a3ec327cd0b967b53e25439d6ab05a57e55ba82c3cdf03c94b83106711d43ff | Signature verified            |
| 5     | 3045022100c52d6520b6721cdfc3208b0d51fd8ed68219ca01894621723c461e976c24af8c02206e2aa94b8f2633a63ae1f30b39025ebbc490b2abb85c91cc3fcc49969ccc7c32 | 028228f84daf8d397fcf252af0114c53db7a354c7ffcf293e8a6997cb1b253ac8f | Signature verification failed |
| 5     | 3045022100c52d6520b6721cdfc3208b0d51fd8ed68219ca01894621723c461e976c24af8c02206e2aa94b8f2633a63ae1f30b39025ebbc490b2abb85c91cc3fcc49969ccc7c32 | 03e8b76584f2a41aadd97e461721875a7ced12ab0ba9266181eae68e95bc03e190 | Signature verified            |
| 5     | 30440220337197c6e6298c2bab38f217e0927c2acc8a85c60d082e52a31979dc155771c1022008aabc406b5438ff5d34e4083bcff5ae251553de876ea35b3af1a091859ee8f8   | 022a3ec327cd0b967b53e25439d6ab05a57e55ba82c3cdf03c94b83106711d43ff | Signature verified            |
| 6     | 30440220474aef21b3d44241ad6b05fd13346c27a8d6ca92883292ca3334e4bc859a1ad402204877f6246366ef91789bf0fffcc668f46627f8bf067fc171aaa01b6d0c57ead7   | 028228f84daf8d397fcf252af0114c53db7a354c7ffcf293e8a6997cb1b253ac8f | Signature verification failed |
| 6     | 30440220474aef21b3d44241ad6b05fd13346c27a8d6ca92883292ca3334e4bc859a1ad402204877f6246366ef91789bf0fffcc668f46627f8bf067fc171aaa01b6d0c57ead7   | 03e8b76584f2a41aadd97e461721875a7ced12ab0ba9266181eae68e95bc03e190 | Signature verified            |
| 6     | 3045022100e45f145bc9a4f9a931ad8bfffab0ca3ab81eceaf1bdda2115596d678e4529f400220194f19170217443aaacdd80c1675760de71196f274a296e675091e7af2072778 | 022a3ec327cd0b967b53e25439d6ab05a57e55ba82c3cdf03c94b83106711d43ff | Signature verified            |
| 7     | 3045022100ebf01fbd32da349001bcb4d1e6e859d801e339a06cf8f9fcbe92adb8205318a40220664e5c3b8638123368f45b2e6131ea02bf9e1d217e22ec37c8d874db7d691466 | 028228f84daf8d397fcf252af0114c53db7a354c7ffcf293e8a6997cb1b253ac8f | Signature verification failed |
| 7     | 3045022100ebf01fbd32da349001bcb4d1e6e859d801e339a06cf8f9fcbe92adb8205318a40220664e5c3b8638123368f45b2e6131ea02bf9e1d217e22ec37c8d874db7d691466 | 03e8b76584f2a41aadd97e461721875a7ced12ab0ba9266181eae68e95bc03e190 | Signature verified            |
| 7     | 3045022100f456b62fbc2c67cfdcdb4d84b8ab40b142d26f23f57c47faedae2c975a01663a022035bb8e1293e53930c5c7a79f84e33c64fc4defbac546575e27d2b13960cdbeb5 | 022a3ec327cd0b967b53e25439d6ab05a57e55ba82c3cdf03c94b83106711d43ff | Signature verified            |
| 8     | 30440220310826373953b43f6f094d2b172a7e088306778a9eae3eef6ddcf6484729d16902207e06bb7730a955533b66251c17cebaa30e0625d8c5230f32082f5d70784ddb79   | 028228f84daf8d397fcf252af0114c53db7a354c7ffcf293e8a6997cb1b253ac8f | Signature verification failed |
| 8     | 30440220310826373953b43f6f094d2b172a7e088306778a9eae3eef6ddcf6484729d16902207e06bb7730a955533b66251c17cebaa30e0625d8c5230f32082f5d70784ddb79   | 03e8b76584f2a41aadd97e461721875a7ced12ab0ba9266181eae68e95bc03e190 | Signature verified            |
| 8     | 3044022017dee30b98fd415c73c388fec22f7e00f96c3b0fd084ba6a65746816ddf29048022004eb04198f426acf6668aa01351a28fc70cd2060f4bc1cbec7da08d5719defcf   | 022a3ec327cd0b967b53e25439d6ab05a57e55ba82c3cdf03c94b83106711d43ff | Signature verified            |

## decoded
https://live.blockcypher.com/btc/decodetx/ でdecodeしたもの
### 流出tx
`curl -sSL "https://mempool.space/api/tx/975ec405ac9dc9fa5ab8009d94d6a1fe31dff8a8127ea90d023104e52754e4d7/hex"`で取得

```sh
{
    "addresses": [
        "3P8MfdM4pULv7ozdQvfwAqNF29zAjmnUYD",
        "1B6rJRfjTXwEy36SCs5zofGMmdv2kdZw7P"
    ],
    "block_height": -1,
    "block_index": -1,
    "confirmations": 0,
    "double_spend": false,
    "fees": 10000000,
    "hash": "975ec405ac9dc9fa5ab8009d94d6a1fe31dff8a8127ea90d023104e52754e4d7",
    "inputs": [
        {
            "addresses": [
                "3P8MfdM4pULv7ozdQvfwAqNF29zAjmnUYD"
            ],
            "age": 837660,
            "output_index": 29,
            "output_value": 50300000000,
            "prev_hash": "c03acb80562302248f8be621e765315345a85cd3856bf49583e97b0428342908",
            "script": "220020cb374fed2759417dbafd9af9eaafedececd464ef0bc22193b52012efe26159f5",
            "script_type": "pay-to-script-hash",
            "sequence": 4294967295
        },
        {
            "addresses": [
                "3P8MfdM4pULv7ozdQvfwAqNF29zAjmnUYD"
            ],
            "age": 837660,
            "output_index": 21,
            "output_value": 50000000000,
            "prev_hash": "c03acb80562302248f8be621e765315345a85cd3856bf49583e97b0428342908",
            "script": "220020cb374fed2759417dbafd9af9eaafedececd464ef0bc22193b52012efe26159f5",
            "script_type": "pay-to-script-hash",
            "sequence": 4294967295
        },
        {
            "addresses": [
                "3P8MfdM4pULv7ozdQvfwAqNF29zAjmnUYD"
            ],
            "age": 837660,
            "output_index": 26,
            "output_value": 50000000000,
            "prev_hash": "c03acb80562302248f8be621e765315345a85cd3856bf49583e97b0428342908",
            "script": "220020cb374fed2759417dbafd9af9eaafedececd464ef0bc22193b52012efe26159f5",
            "script_type": "pay-to-script-hash",
            "sequence": 4294967295
        },
        {
            "addresses": [
                "3P8MfdM4pULv7ozdQvfwAqNF29zAjmnUYD"
            ],
            "age": 837660,
            "output_index": 23,
            "output_value": 50000000000,
            "prev_hash": "c03acb80562302248f8be621e765315345a85cd3856bf49583e97b0428342908",
            "script": "220020cb374fed2759417dbafd9af9eaafedececd464ef0bc22193b52012efe26159f5",
            "script_type": "pay-to-script-hash",
            "sequence": 4294967295
        },
        {
            "addresses": [
                "3P8MfdM4pULv7ozdQvfwAqNF29zAjmnUYD"
            ],
            "age": 837660,
            "output_index": 24,
            "output_value": 50000000000,
            "prev_hash": "c03acb80562302248f8be621e765315345a85cd3856bf49583e97b0428342908",
            "script": "220020cb374fed2759417dbafd9af9eaafedececd464ef0bc22193b52012efe26159f5",
            "script_type": "pay-to-script-hash",
            "sequence": 4294967295
        },
        {
            "addresses": [
                "3P8MfdM4pULv7ozdQvfwAqNF29zAjmnUYD"
            ],
            "age": 837660,
            "output_index": 25,
            "output_value": 50000000000,
            "prev_hash": "c03acb80562302248f8be621e765315345a85cd3856bf49583e97b0428342908",
            "script": "220020cb374fed2759417dbafd9af9eaafedececd464ef0bc22193b52012efe26159f5",
            "script_type": "pay-to-script-hash",
            "sequence": 4294967295
        },
        {
            "addresses": [
                "3P8MfdM4pULv7ozdQvfwAqNF29zAjmnUYD"
            ],
            "age": 837660,
            "output_index": 27,
            "output_value": 50000000000,
            "prev_hash": "c03acb80562302248f8be621e765315345a85cd3856bf49583e97b0428342908",
            "script": "220020cb374fed2759417dbafd9af9eaafedececd464ef0bc22193b52012efe26159f5",
            "script_type": "pay-to-script-hash",
            "sequence": 4294967295
        },
        {
            "addresses": [
                "3P8MfdM4pULv7ozdQvfwAqNF29zAjmnUYD"
            ],
            "age": 837660,
            "output_index": 28,
            "output_value": 50000000000,
            "prev_hash": "c03acb80562302248f8be621e765315345a85cd3856bf49583e97b0428342908",
            "script": "220020cb374fed2759417dbafd9af9eaafedececd464ef0bc22193b52012efe26159f5",
            "script_type": "pay-to-script-hash",
            "sequence": 4294967295
        },
        {
            "addresses": [
                "3P8MfdM4pULv7ozdQvfwAqNF29zAjmnUYD"
            ],
            "age": 837660,
            "output_index": 22,
            "output_value": 50000000000,
            "prev_hash": "c03acb80562302248f8be621e765315345a85cd3856bf49583e97b0428342908",
            "script": "220020cb374fed2759417dbafd9af9eaafedececd464ef0bc22193b52012efe26159f5",
            "script_type": "pay-to-script-hash",
            "sequence": 4294967295
        }
    ],
    "outputs": [
        {
            "addresses": [
                "1B6rJRfjTXwEy36SCs5zofGMmdv2kdZw7P"
            ],
            "script": "76a9146ecc7dd816248cd5919fcd86a0832c8553f4296e88ac",
            "script_type": "pay-to-pubkey-hash",
            "value": 450290000000
        }
    ],
    "preference": "high",
    "received": "2024-06-03T00:47:49.224899793Z",
    "relayed_by": "44.195.32.254",
    "size": 3006,
    "total": 450290000000,
    "ver": 1,
    "vin_sz": 9,
    "vout_sz": 1,
    "vsize": 1298
}
```

prev tx
`curl -sSL "https://mempool.space/api/tx/c03acb80562302248f8be621e765315345a85cd3856bf49583e97b0428342908/hex"`
```sh
{
    "addresses": [
        "3P8MfdM4pULv7ozdQvfwAqNF29zAjmnUYD",
        "3Dhk8F6KYvMyqHN1r4kAT8t7SceK7yDjA4"
    ],
    "block_height": -1,
    "block_index": -1,
    "confirmations": 0,
    "double_spend": false,
    "fees": 126700,
    "hash": "c03acb80562302248f8be621e765315345a85cd3856bf49583e97b0428342908",
    "inputs": [
        {
            "addresses": [
                "3Dhk8F6KYvMyqHN1r4kAT8t7SceK7yDjA4"
            ],
            "age": 837343,
            "output_index": 1,
            "output_value": 470301946745,
            "prev_hash": "fb5ff69a87f15881ca320d3561f478a8bb294a72a657924ee3036d46056afd34",
            "script": "00483045022100b9e3d29899a3978ea793ab0385026d453541aa93322c7c796b2d5b677fa08e6202202bfda89747fa2751afb5b21b991550f855acf926e72e6df15ced38f6000a72cf014730440220629d9dfbc38a89696b9e35297c62fe15ae0cedb97d5ba40437673e81f2ea5f25022042186666d76a9758c0686f7e05f39bc140fe5c11064b1ac6058083aa814baed5014c695221022279489fa3e9886b25f0c455e510f8b0f0d636930768d7d440dd103077597ffe210344402b39b5efbd49f88400cdf041d1cd15bee795e672c8fefdc9790f78233c00210381b8c97ac3e501fc41e1e3b9eb07acba5b6cded8fbeb32324bda23ec3200ba2e53ae",
            "script_type": "pay-to-script-hash",
            "sequence": 4294967294
        }
    ],
    "lock_time": 837658,
    "outputs": [
        {
            "addresses": [
                "3Dhk8F6KYvMyqHN1r4kAT8t7SceK7yDjA4"
            ],
            "script": "a91483c4bcd560f09a59ff0f4499f4345561ed01152887",
            "script_type": "pay-to-script-hash",
            "value": 1820045
        },
        {
            "addresses": [
                "3P8MfdM4pULv7ozdQvfwAqNF29zAjmnUYD"
            ],
            "script": "a914eb25855f1c231b6899e15e5e8aabf5a62f6c4d6387",
            "script_type": "pay-to-script-hash",
            "value": 1000000000
        },
        {
            "addresses": [
                "3P8MfdM4pULv7ozdQvfwAqNF29zAjmnUYD"
            ],
            "script": "a914eb25855f1c231b6899e15e5e8aabf5a62f6c4d6387",
            "script_type": "pay-to-script-hash",
            "value": 1000000000
        },
        {
            "addresses": [
                "3P8MfdM4pULv7ozdQvfwAqNF29zAjmnUYD"
            ],
            "script": "a914eb25855f1c231b6899e15e5e8aabf5a62f6c4d6387",
            "script_type": "pay-to-script-hash",
            "value": 1000000000
        },
        {
            "addresses": [
                "3P8MfdM4pULv7ozdQvfwAqNF29zAjmnUYD"
            ],
            "script": "a914eb25855f1c231b6899e15e5e8aabf5a62f6c4d6387",
            "script_type": "pay-to-script-hash",
            "value": 1000000000
        },
        {
            "addresses": [
                "3P8MfdM4pULv7ozdQvfwAqNF29zAjmnUYD"
            ],
            "script": "a914eb25855f1c231b6899e15e5e8aabf5a62f6c4d6387",
            "script_type": "pay-to-script-hash",
            "value": 1000000000
        },
        {
            "addresses": [
                "3P8MfdM4pULv7ozdQvfwAqNF29zAjmnUYD"
            ],
            "script": "a914eb25855f1c231b6899e15e5e8aabf5a62f6c4d6387",
            "script_type": "pay-to-script-hash",
            "value": 1000000000
        },
        {
            "addresses": [
                "3P8MfdM4pULv7ozdQvfwAqNF29zAjmnUYD"
            ],
            "script": "a914eb25855f1c231b6899e15e5e8aabf5a62f6c4d6387",
            "script_type": "pay-to-script-hash",
            "value": 1000000000
        },
        {
            "addresses": [
                "3P8MfdM4pULv7ozdQvfwAqNF29zAjmnUYD"
            ],
            "script": "a914eb25855f1c231b6899e15e5e8aabf5a62f6c4d6387",
            "script_type": "pay-to-script-hash",
            "value": 1000000000
        },
        {
            "addresses": [
                "3P8MfdM4pULv7ozdQvfwAqNF29zAjmnUYD"
            ],
            "script": "a914eb25855f1c231b6899e15e5e8aabf5a62f6c4d6387",
            "script_type": "pay-to-script-hash",
            "value": 1000000000
        },
        {
            "addresses": [
                "3P8MfdM4pULv7ozdQvfwAqNF29zAjmnUYD"
            ],
            "script": "a914eb25855f1c231b6899e15e5e8aabf5a62f6c4d6387",
            "script_type": "pay-to-script-hash",
            "value": 1000000000
        },
        {
            "addresses": [
                "3P8MfdM4pULv7ozdQvfwAqNF29zAjmnUYD"
            ],
            "script": "a914eb25855f1c231b6899e15e5e8aabf5a62f6c4d6387",
            "script_type": "pay-to-script-hash",
            "value": 1000000000
        },
        {
            "addresses": [
                "3P8MfdM4pULv7ozdQvfwAqNF29zAjmnUYD"
            ],
            "script": "a914eb25855f1c231b6899e15e5e8aabf5a62f6c4d6387",
            "script_type": "pay-to-script-hash",
            "value": 1000000000
        },
        {
            "addresses": [
                "3P8MfdM4pULv7ozdQvfwAqNF29zAjmnUYD"
            ],
            "script": "a914eb25855f1c231b6899e15e5e8aabf5a62f6c4d6387",
            "script_type": "pay-to-script-hash",
            "value": 1000000000
        },
        {
            "addresses": [
                "3P8MfdM4pULv7ozdQvfwAqNF29zAjmnUYD"
            ],
            "script": "a914eb25855f1c231b6899e15e5e8aabf5a62f6c4d6387",
            "script_type": "pay-to-script-hash",
            "value": 1000000000
        },
        {
            "addresses": [
                "3P8MfdM4pULv7ozdQvfwAqNF29zAjmnUYD"
            ],
            "script": "a914eb25855f1c231b6899e15e5e8aabf5a62f6c4d6387",
            "script_type": "pay-to-script-hash",
            "value": 1000000000
        },
        {
            "addresses": [
                "3P8MfdM4pULv7ozdQvfwAqNF29zAjmnUYD"
            ],
            "script": "a914eb25855f1c231b6899e15e5e8aabf5a62f6c4d6387",
            "script_type": "pay-to-script-hash",
            "value": 1000000000
        },
        {
            "addresses": [
                "3P8MfdM4pULv7ozdQvfwAqNF29zAjmnUYD"
            ],
            "script": "a914eb25855f1c231b6899e15e5e8aabf5a62f6c4d6387",
            "script_type": "pay-to-script-hash",
            "value": 1000000000
        },
        {
            "addresses": [
                "3P8MfdM4pULv7ozdQvfwAqNF29zAjmnUYD"
            ],
            "script": "a914eb25855f1c231b6899e15e5e8aabf5a62f6c4d6387",
            "script_type": "pay-to-script-hash",
            "value": 1000000000
        },
        {
            "addresses": [
                "3P8MfdM4pULv7ozdQvfwAqNF29zAjmnUYD"
            ],
            "script": "a914eb25855f1c231b6899e15e5e8aabf5a62f6c4d6387",
            "script_type": "pay-to-script-hash",
            "value": 1000000000
        },
        {
            "addresses": [
                "3P8MfdM4pULv7ozdQvfwAqNF29zAjmnUYD"
            ],
            "script": "a914eb25855f1c231b6899e15e5e8aabf5a62f6c4d6387",
            "script_type": "pay-to-script-hash",
            "value": 1000000000
        },
        {
            "addresses": [
                "3P8MfdM4pULv7ozdQvfwAqNF29zAjmnUYD"
            ],
            "script": "a914eb25855f1c231b6899e15e5e8aabf5a62f6c4d6387",
            "script_type": "pay-to-script-hash",
            "value": 50000000000
        },
        {
            "addresses": [
                "3P8MfdM4pULv7ozdQvfwAqNF29zAjmnUYD"
            ],
            "script": "a914eb25855f1c231b6899e15e5e8aabf5a62f6c4d6387",
            "script_type": "pay-to-script-hash",
            "value": 50000000000
        },
        {
            "addresses": [
                "3P8MfdM4pULv7ozdQvfwAqNF29zAjmnUYD"
            ],
            "script": "a914eb25855f1c231b6899e15e5e8aabf5a62f6c4d6387",
            "script_type": "pay-to-script-hash",
            "value": 50000000000
        },
        {
            "addresses": [
                "3P8MfdM4pULv7ozdQvfwAqNF29zAjmnUYD"
            ],
            "script": "a914eb25855f1c231b6899e15e5e8aabf5a62f6c4d6387",
            "script_type": "pay-to-script-hash",
            "value": 50000000000
        },
        {
            "addresses": [
                "3P8MfdM4pULv7ozdQvfwAqNF29zAjmnUYD"
            ],
            "script": "a914eb25855f1c231b6899e15e5e8aabf5a62f6c4d6387",
            "script_type": "pay-to-script-hash",
            "value": 50000000000
        },
        {
            "addresses": [
                "3P8MfdM4pULv7ozdQvfwAqNF29zAjmnUYD"
            ],
            "script": "a914eb25855f1c231b6899e15e5e8aabf5a62f6c4d6387",
            "script_type": "pay-to-script-hash",
            "value": 50000000000
        },
        {
            "addresses": [
                "3P8MfdM4pULv7ozdQvfwAqNF29zAjmnUYD"
            ],
            "script": "a914eb25855f1c231b6899e15e5e8aabf5a62f6c4d6387",
            "script_type": "pay-to-script-hash",
            "value": 50000000000
        },
        {
            "addresses": [
                "3P8MfdM4pULv7ozdQvfwAqNF29zAjmnUYD"
            ],
            "script": "a914eb25855f1c231b6899e15e5e8aabf5a62f6c4d6387",
            "script_type": "pay-to-script-hash",
            "value": 50000000000
        },
        {
            "addresses": [
                "3P8MfdM4pULv7ozdQvfwAqNF29zAjmnUYD"
            ],
            "script": "a914eb25855f1c231b6899e15e5e8aabf5a62f6c4d6387",
            "script_type": "pay-to-script-hash",
            "value": 50300000000
        }
    ],
    "preference": "high",
    "received": "2024-06-03T00:48:53.843176225Z",
    "relayed_by": "54.145.160.164",
    "size": 1266,
    "total": 470301820045,
    "ver": 2,
    "vin_sz": 1,
    "vout_sz": 30,
    "vsize": 1266
}
```