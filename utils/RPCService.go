package utils

import (
	"BeegoBcRPCCode/models"
	"encoding/json"
	"fmt"
)

//获取当前节点的区块个数
func GetBlockcount()float64{
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