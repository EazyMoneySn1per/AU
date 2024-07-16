package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//var key = []byte("Rh&z83I7X7G0vCLf")

type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Veri(username, password string) map[string]interface{} {
	url := "http://10.1.20.86/api/tripartite/auth"
	var u = &user{
		Username: username,
		Password: password,
	}

	uu, _ := json.Marshal(u)

	//发送http请求加密
	body := Encrypt("http://127.0.0.1:5000/encrypt", "application/json", uu)

	//设置一个认证请求,参数：请求方法，请求路径，请求体
	request, err := http.NewRequest("POST", url, bytes.NewBufferString(string(body)))
	if err != nil {
		log.Println("认证请求出错：", err)
	}

	//设置请求头
	request.Header.Set("content-type", "text/plain")

	//添加请求体
	client := &http.Client{}
	//执行请求对账号密码进行验证
	resp, err := client.Do(request)

	if err != nil {
		log.Println("认证出错:", err)
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("认证结果出错:", err)
	}

	//发送一个请求解密
	result := Decrypt("http://127.0.0.1:5000/decrypt", "application/json", body)

	return result
}
