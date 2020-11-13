package types

import (
	"bytes"
	"encoding/binary"
	"github.com/tybc/crypto"
	"github.com/tybc/errors"
	"github.com/tybc/vm/vmcommon"
	"github.com/tybc/wallet"
)

var (
	inputErr = errors.New("tx input")
)

type TxInput struct {
	ID Hash `json:"id"`
	Spend
	ScriptSig []byte `json:"scriptSig"`
}

func (txInput *TxInput) SetSpend(utxo *UTXO) {
	txInput.SoureId = BytesToHash(utxo.SoureId)
	txInput.SourcePos = utxo.SourcePos
	txInput.Amount = utxo.Amount
	txInput.ScriptPk = utxo.ScriptPk
}

func (txInput *TxInput) GenerateID() Hash {
	//TODO maybe need: len(txInput.ScriptPk) >0
	split := []byte(":")
	var sourcePosByte = make([]byte, 4)
	binary.LittleEndian.PutUint32(sourcePosByte, txInput.SourcePos)

	var amountByte = make([]byte, 8)
	binary.LittleEndian.PutUint64(amountByte, txInput.Amount)

	b := bytes.Join([][]byte{
		txInput.SpendOutputId.Bytes(),
		split,
		txInput.SoureId.Bytes(),
		split,
		sourcePosByte,
		split,
		amountByte,
		split,
		txInput.ScriptPk,
	}, []byte{})

	h := BytesToHash(crypto.Sha256(b))
	return h
}

func (txInput *TxInput) GenerateScriptSig(wt *wallet.Wallet, txId *Hash) (*[]byte, error) {
	message := bytes.Join([][]byte{
		txInput.ID.Bytes(),
		(*txId).Bytes(),
	}, []byte{})

	sig := crypto.Sign(wt.Priv, message)

	//scriptSig = <signature> <public key>
	scriptSig := bytes.Join([][]byte{
		{byte(vmcommon.OpPushData64)},
		sig,
		{byte(vmcommon.OpPushData32)},
		wt.Pub,
	}, []byte{})
	return &scriptSig, nil
}
