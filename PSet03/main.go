package main

import (
	"fmt"

	"github.com/btcsuite/btcd/chaincfg"
)

var (
	// we're running on testnet3
	testnet3Parameters = &chaincfg.TestNet3Params
)

func main() {
	fmt.Printf("mas.s62 pset03 - utxohunt\n")

	// Task #1 make an address pair
	// Call AddressFrom PrivateKey() to make a keypair
	toadd, _ := AddressFromPrivateKey("mpQQryVrYmGNPxVqNeE5Rgodvsk25Pihg")
	fmt.Printf("address is %s\n", toadd)
	var o uint32 = 19
	fromtxid := "1f497ac245eb25cd94157c290f62d042e3bdda1e57920b6d1d2c5cfa362c12da"
	key := "mas.s62"
	fromadd := "mpQQryVrYmGNPxVqNeE5RgoYAv2v66Psao"
	var q int64 = 495000

	// Task #2 make a transaction
	// Call EZTxBuilder to make a transaction
	tx1 := EZTxBuilder(fromtxid, o, toadd, key, fromadd, q)
	txid := TxToHex(tx1)
	fmt.Printf("txid is: %s\n", txid)

	// task 3, call OpReturnTxBuilder() the same way EZTxBuilder() was used
	fromtxid = "b9770e48c1b60b58bca667aecdc82262c9ba1ded4dffb1fcf6ec6037ca265028"
	q = 490000
	tx2 := OpReturnTxBuilder(fromtxid, q)
	txid2 := TxToHex(tx2)
	fmt.Printf("txid2 is: %s\n", txid2)

	//part II: More coins

	// Run first AddressFromPrivateKey to gather fromtxid
	fromadd, _ = AddressFromPrivateKey("mpQQryVrYmGNPxVqNeE5RgoYAv2v66Psao")
	fmt.Printf("address with more coins is %s\n", fromadd)

	fromtxid = "dda03859cdc9fe79d6bda5adfcfd74ebe6479211f7889841c55a9214326a52f0"
	o = 1
	q = 4995000
	key = "mpQQryVrYmGNPxVqNeE5RgoYAv2v66Psao"
	tx3 := EZTxBuilder(fromtxid, o, toadd, key, fromadd, q)
	txid3 := TxToHex(tx3)
	fmt.Printf("txid3 is: %s\n", txid3)

	fromtxid = "1012bd6d2ab420576681ae31cee8131452c217d0ba0f0d436a329cde21241fbd"
	q = 4990000
	tx4 := OpReturnTxBuilder(fromtxid, q)
	txid4 := TxToHex(tx4)
	fmt.Printf("txid4 is: %s\n", txid4)
	return

}
