package utils

import (
	"BcRPCCode/entity"
	"BeegoBcRPCCode/models"
	"encoding/json"
	"fmt"
)

//获取当前节点的区块个数
func GetBlockcount() float64 {
	resByte := RpcRequest("getblockcount")
	//fmt.Println(string(resByte))
	var rpcResult models.RPCResult
	err := json.Unmarshal(resByte, &rpcResult)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}
	return rpcResult.Result.(float64)

}

//获取最新的区块hash
func GetBestBlockHash() string {
	result := RpcRequest("getbestblockhash")
	var rpcResult models.RPCResult
	err := json.Unmarshal(result, &rpcResult)
	if err != nil {
		return ""
	}
	return rpcResult.Result.(string)
}

//获取当前节点的区块的信息
func GetBlockInfoByHash(hash string) *models.BlockInfo {
	result := RpcRequest("getblock", hash)
	var rpcResult entity.RPCResult

	err := json.Unmarshal(result, &rpcResult)
	if err != nil {
		return nil
	}
	blockStr := rpcResult.Result.(map[string]interface{})
	var block models.BlockInfo
	block.Hash = blockStr["hash"].(string)
	block.Size = blockStr["size"].(float64)
	block.Height = blockStr["height"].(float64)
	block.Version = blockStr["version"].(float64)
	return &block
}

//获取指定高度的区块哈希
func GteBlockhash(height int) string {
	result := RpcRequest("getblockhash", height)
	var rpcResult entity.RPCResult
	err := json.Unmarshal(result, &rpcResult)
	if err != nil {
		return ""
	}
	return rpcResult.Result.(string)

}
//获取指定的哈希区块头
func GetBlockHeader(hash string, bool bool) *models.BlockHeader{
	resule:=RpcRequest("getblockheader",hash,bool)
	var rpcResult entity.RPCResult
	err:=json.Unmarshal(resule,&rpcResult)
	if err != nil {
		return nil
	}
	blockHeaderStr:=rpcResult.Result.(map[string]interface{})
	var blockheader  models.BlockHeader
	blockheader.Hash=blockHeaderStr["hash"].(string)

	return &blockheader
}
