package utils

import (
	"BcRPCCode/entity"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

/**
准备json——rpc通信的数据格式
*/

func RpcRequest(method string, params ...interface{}) []byte {
	rpcRequest := entity.RPCRequest{
		Id:      time.Now().Unix(),
		Method:  method,
		Jsonrpc: "2.0",
	}
	if params != nil {
		rpcRequest.Params = params
	}
	reqBytes, err := json.Marshal(&rpcRequest) //序列化
	if err != nil {
		fmt.Println(err.Error())
		return nil

	}

	//2.发送post请求

	client := &http.Client{}
	request, err := http.NewRequest("POST", RPCURL, bytes.NewBuffer(reqBytes))
	if err != nil {
		fmt.Println(err.Error())
		return nil

	}
	//请求头设置

	request.Header.Add("Encoding", "UTF-8")
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Basic "+Base64())
	//java:HttpResponse reponse = client.exwcute(post);
	//java返回响应类：HttpResonse

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return nil

	}
	resByte, _ := ioutil.ReadAll(response.Body)

	code := response.StatusCode
	if code == 200 {
		fmt.Println("请求成功！")

	} else {
		fmt.Println("请求失败", code)

	}

	return resByte
}