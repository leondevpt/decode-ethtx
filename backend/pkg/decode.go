package pkg

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"strings"

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
	ChainId  *big.Int `json:"chainId,omitempty"`
	R        string   `json:"r,omitempty"`
	V        string   `json:"v,omitempty"`
	S        string   `json:"s,omitempty"`
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
	m.Data = strings.TrimPrefix(hexutil.Encode(msg.Data()), "0x")
	m.From = msg.From().Hex()
	m.ChainId = tx.ChainId()
	v, r, s := tx.RawSignatureValues()
	m.V = v.Text(16)
	m.R = r.Text(16)
	m.S = s.Text(16)
	return &m, nil
}
