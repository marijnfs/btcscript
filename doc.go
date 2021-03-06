// Copyright (c) 2013 Conformal Systems LLC.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

/*
Package btcscript implements bitcoin transaction scripts.

A complete description of the script language used by bitcoin can be found at
https://en.bitcoin.it/wiki/Script.  The following only serves as a quick
overview to provide information on how to use the package.

This package provides data structures and functions to parse and execute
bitcoin transaction scripts.

Script Overview

Bitcoin transaction scripts are written in a stack-base, FORTH-like language.

The bitcoin script language consists of a number of opcodes which fall into
several categories such pushing and popping data to and from the stack,
performing basic and bitwise arithmetic, conditional branching, comparing
hashes, and checking cryptographic signatures.  Scripts are processed from left
to right and intentionally do not provide loops.

The vast majority of Bitcoin scripts at the time of this writing are of several
standard forms which consist of a spender providing a public key and a signature
which proves the spender owns the associated private key.  This information
is used to prove the the spender is authorized to perform the transaction.

One benefit of using a scripting language is added flexibility in specifying
what conditions must be met in order to spend bitcoins.

Usage

The usage of this package consists of creating a new script engine for a pair
of transaction inputs and outputs and using the engine to execute the scripts.

The following function is an example of how to create and execute a script
engine to validate a transaction.

	// ValidateTx validates the txIdx'th input of tx. The output transaction
	// corresponding to the this input is the txInIdx'th output of txIn. The
	// block timestamp of tx is timestamp and the protocol version involved
	// is pver.
	func ValidateTx(tx *btcwire.MsgTx, txIdx int, txIn *btcwire.MsgTx, txInIdx int, pver int, timestamp time.Time) {
		pkScript := txIn.TxOut[txInIdx].PkScript
		sigScript := tx.txIn[TxIdx]
		engine, err := btcscript.NewScript(sigScript, pkScript, txInIdx,
			tx, pver, timestamp.After(btcscript.Bip16Activation))
		return engine.Execute()
	}

Errors

Errors returned by this package are of the form btcscript.StackErrX where X
indicates the specific error.  See Variables in the package documentation for a
full list.
*/
package btcscript
