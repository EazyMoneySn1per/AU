package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// 加密请求
func Encrypt(url, contentType string, data []byte) []byte {
	re, err := http.Post(url, contentType, bytes.NewBuffer(data))
	if err != nil {
		log.Println("加密请求出错：", err)
	}

	body, err := ioutil.ReadAll(re.Body)

	if err != nil {
		log.Println("加密结果出错：", err)
	}
	return body
}

//解密请求

func Decrypt(url, contentType string, body []byte) map[string]interface{} {
	//请求解密内容
	type st struct {
		Skey string `json:"skey"`
	}
	var s = &st{
		Skey: string(body),
	}

	ss, err := json.Marshal(s)
	//发起一个请求进行解码
	r, err := http.Post(url, contentType, bytes.NewBuffer(ss))
	if err != nil {
		log.Println("解密请求出错:", err)
	}

	body, err = ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("解密结果出错:", err)
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)

	if err != nil {
		log.Println("解密结果转化出错", err)
	}

	return result
}

var IV = []byte("ABCDEF1234123412")

// AES 加密 测试用
func Encrypt2(encodeStr string, key []byte) (string, error) {
	encodeBytes := []byte(encodeStr)
	//根据key 生成密文
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	encodeBytes = PKCS5Padding(encodeBytes, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, IV)
	crypted := make([]byte, len(encodeBytes))
	blockMode.CryptBlocks(crypted, encodeBytes)
	return base64.StdEncoding.EncodeToString(crypted), nil
}

// 16字节为一组，填充为16字节倍数
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	// 填充size
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// AES 解密
func Decrypt2(decodeStr string, key []byte) ([]byte, error) {
	//先解密base64
	decodeBytes, err := base64.StdEncoding.DecodeString(decodeStr)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, IV)
	origData := make([]byte, len(decodeBytes))
	blockMode.CryptBlocks(origData, decodeBytes)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

// 移除Padding
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	// 移除填充的size
	return origData[:(length - unpadding)]
}
