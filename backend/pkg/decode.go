package pkg

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

type EthTx struct {
	Nonce    uint64   `json:"nonce"`
	GasPrice *big.Int `json:"gasPrice"`
	GasLimit uint64   `json:"gasLimit"`
	To       string   `json:"to"`
	Value    *big.Int `json:"value"`
	Data     string   `json:"data"`
	From     string   `json:"from"`
	ChainId  *big.Int `json:"chainId"`
	R        *big.Int `json:"r"`
	V        *big.Int `json:"v"`
	S        *big.Int `json:"s"`
}

func (e *EthTx) ToJSON() ([]byte, error) {
	return json.Marshal(e)
}

func DecodeTx(hexTx string) (*EthTx, error) {
	var m EthTx
	rawTxBytes := common.FromHex(hexTx)
	var tx = new(types.Transaction)
	err := rlp.DecodeBytes(rawTxBytes, tx)
	if err != nil {
		return &m, err
	}
	msg, err := tx.AsMessage(types.LatestSignerForChainID(tx.ChainId()), nil)
	if err != nil {
		return &m, err
	}

	m.Nonce = msg.Nonce()
	m.GasPrice = msg.GasPrice()
	m.GasLimit = msg.Gas()
	m.To = msg.To().Hex()
	m.Value = msg.Value()
	m.Data = hexutil.Encode(msg.Data())
	m.From = msg.From().Hex()
	m.V, m.R, m.S = tx.RawSignatureValues()
	return &m, nil
}
