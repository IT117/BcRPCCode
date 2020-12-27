package controllers

import (
	"BeegoBcRPCCode/utils"
	"fmt"
)

func BictoinJudge(bitcoin string, target interface{})interface{} {
	if bitcoin == "getbestblockhash" {
		hash := utils.GetBestBlockHash()
		return hash
	}
	if bitcoin == "getblockcount" {
		count := utils.GetBlockcount()
		return count

	}
	if bitcoin == "getblock" {
		Info := utils.GetBlockInfoByHash(target.(string))
       InfoMap := make(map[string]interface{})
       InfoMap["该区块的哈希："]=Info.Hash
       InfoMap["该区块的大小："]=Info.Size
       InfoMap["该区块的版本："]=Info.Version
       InfoMap["该区块的高度："]=Info.Height

		fmt.Println("区块的内容：",InfoMap)
       for k,v :=range InfoMap{
       	fmt.Println(k,v)
	   }
		return InfoMap

	}
	return nil
}
