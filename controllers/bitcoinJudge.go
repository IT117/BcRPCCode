package controllers

import "BeegoBcRPCCode/utils"

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
		return Info

	}
	return nil
}
