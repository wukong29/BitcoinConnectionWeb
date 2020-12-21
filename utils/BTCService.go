package utils

import (
	"BitcoinConnection/entity"
	"encoding/json"
	"fmt"
)

func GetTxOutSetInfo() *entity.TxOutSetInfo {
	body,err := JsonRpc("gettxoutsetinfo")
	if err!= nil {
		fmt.Println(err.Error())
		return nil
	}
	fmt.Println(string(body))
	rpcRuselt := new(entity.RPCResult)
	err = json.Unmarshal(body,rpcRuselt)

	TxOutSetInfoStr := rpcRuselt.Result.(map[string]interface{})
	TxOutSetInfo := new(entity.TxOutSetInfo)
	TxOutSetInfo.Height = TxOutSetInfoStr["height"].(float64)
	TxOutSetInfo.Bestblock = TxOutSetInfoStr["bestblock"].(string)
	return TxOutSetInfo
}


func GetBestBlockHash()interface{}{
	body,err := JsonRpc("getbestblockhash")
	if err!= nil {
		fmt.Println(err.Error())
		return nil
	}
	fmt.Println(string(body))
	Besthash := new(entity.Besthash)
	err = json.Unmarshal(body,Besthash)
	return Besthash.Result
}
func GetBlockHashByHeight(height int) interface{}{
	body,err := JsonRpc("getblockhash",height)
	//fmt.Println(string(body))
	BlockHash := new(entity.BlockHash)
	err = json.Unmarshal(body,BlockHash)
	if err != nil{
		fmt.Println(err.Error())
		return nil
	}
	return BlockHash.Result
}

func GetBlockCount()interface{}{
	body,err := JsonRpc("getblockcount")
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	fmt.Println(string(body))
	BlockCount := new(entity.BlockCount)
	err = json.Unmarshal(body,BlockCount)
	return BlockCount.Result
}
