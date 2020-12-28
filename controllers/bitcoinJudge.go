package controllers

import (
	"BeegoBcRPCCode/utils"
	"fmt"
	"strconv"
)

func BictoinJudge(bitcoin string, target interface{},target1 interface{}) interface{} {
	//获取最新的区块hash
	if bitcoin == "getbestblockhash" {
		hash := utils.GetBestBlockHash()

		return hash
	}
	//获取当前节点的区块总数
	if bitcoin == "getblockcount" {
		count := utils.GetBlockcount()
		return count

	}
	//获取指定哈希区块的信息
	if bitcoin == "getblock" {
		Info := utils.GetBlockInfoByHash(target.(string))
		InfoMap := make(map[string]interface{})
		InfoMap["该区块的哈希："] = Info.Hash
		InfoMap["该区块的大小："] = Info.Size
		InfoMap["该区块的版本："] = Info.Version
		InfoMap["该区块的高度："] = Info.Height

		fmt.Println("区块的内容：", InfoMap)
		for k, v := range InfoMap {
			fmt.Println(k, v)
		}
		return InfoMap

	}
	//获取当前节点的最新的区块的hash
	if bitcoin == "gitbestblockhash" {
		Hash := utils.GetBestBlockHash()
		return Hash
	}
	//获取指定高度的区块哈希
	if bitcoin == "getblockhash" {
		targetStr := target.(string)
		targetHeight, _ := strconv.Atoi(targetStr)
		Hash := utils.GteBlockhash(targetHeight)

		return Hash
	}
	//获取指定区块的区块头信息
	if bitcoin=="getblockheader" {
		blockheader:=utils.GetBlockHeader(target.(string),target1.(bool))
		return blockheader
	}
	return nil
}
