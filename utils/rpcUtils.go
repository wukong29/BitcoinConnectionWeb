package utils

import (
	"BitcoinConnection/entity"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

/**
 * 准备json格式的数据
 */
func JsonRpc(Order string,params ...interface{}) ([]byte,error) {
	rpcReq := entity.RPCRequest{
		Id:      time.Now().Unix(),
		Method:  Order,
		Jsonrpc: RPCVERSION,
	}
	if params != nil {
		rpcReq.Params = params
	}
	reqBytes, err := json.Marshal(rpcReq)
	if err != nil {
		fmt.Println(err.Error())
		return nil,err
	}
	//发送一个请求并设置请求头
	request := RpcPost(reqBytes)
	//发送post请求操作
	return ClientPost(request)
}

/**
 * 发送一个post请求,并设置请求头
 */
func RpcPost(reqBytes []byte) *http.Request {
	//client := &http.Client{}
	request, err := http.NewRequest("POST", RPCURL, bytes.NewBuffer(reqBytes))
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	authStr := RPCUSER + ":" + RPCPASSWORD
	authBase64 := Base64EncodeString(authStr)
	request.Header.Add("Encoding", "UTF-8")
	request.Header.Add("Authorization", "Basic " + authBase64)
	request.Header.Add("Content-Type", "application/json")
	return request
}

func ClientPost(request *http.Request) ([]byte,error){
	//发起post请求操作
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return nil,err
	}
	body, err := ioutil.ReadAll(response.Body)
	code := response.StatusCode
	fmt.Println(code)
	if code == 200 {
		fmt.Println("请求成功")
	} else {
		fmt.Println("请求失败")
	}
	return body,err
}



