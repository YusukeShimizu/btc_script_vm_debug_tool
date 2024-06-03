# BTC Script VM debug tool
BTC Script VM is a tool that verifies scripts of a specified txid using a virtual machine (VM).  
It uses the mempool.space API to extract transactions, and the obtained transactions are stored in a Bolt database.

## Installation

This project is written in Go, so you need to have Go installed. If you haven't installed Go yet, please install it from [here](https://golang.org/dl/).

## Usage

1. Clone this repository:
    ```sh
    git clone git@github.com:YusukeShimizu/btc_script_vm_debug_tool.git
    ```
2. Move to the project directory:
    ```sh
    cd btc_script_vm_debug_tool
    ```
3. Run the project. Specify the txid as a command line argument:
    ```sh
    go run ./... <txid> <loglevel:optional>
    ```

## Example

By running the following command, the transaction script with txid `975ec405ac9dc9fa5ab8009d94d6a1fe31dff8a8127ea90d023104e52754e4d7` will be verified:
```sh
go run ./... 975ec405ac9dc9fa5ab8009d94d6a1fe31dff8a8127ea90d023104e52754e4d7
```

## Dependencies
This project depends on the following libraries:

* [mempool.space API](https://mempool.space/docs/api/rest) - For extracting transaction data
* [Bolt](https://github.com/boltdb/bolt) - A key-value store to save transaction data
