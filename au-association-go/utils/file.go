package utils

import (
	"au-golang/global"
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"strings"
)

func UploadFile(content []byte, filePath string, fileName string) error {
	// 没有该文件就创建
	_, err := MustOpen(fileName, filePath)
	if err != nil {

		return fmt.Errorf("Fail to CreateFile : %v ", err)
	}

	// 写入文件
	if !strings.Contains(filePath+fileName, "..") {
		err = ioutil.WriteFile(filePath+fileName, content, 0644)
		if err != nil {

			return fmt.Errorf("Fail to WriteFile : %v ", err)
		}
	} else {

		return errors.New("path contains .. ")
	}

	return nil
}

// MustOpen 最大限度尝试读取文件
func MustOpen(fileName, filePath string) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err: %v", err)
	}

	src := dir + "/" + filePath

	err = IsNotExistMkDir(src)
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
	}

	f, err := Open(src+fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, fmt.Errorf("Fail to open file: %v ", err)
	}

	return f, nil
}

// MkDir 自动生成目录
func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// Open 读取文件内容
func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	if !strings.Contains(name, "..") {
		f, err := os.OpenFile(name, flag, perm)
		if err != nil {
			return nil, err
		}

		return f, nil
	} else {

		return nil, errors.New("path contains .. ")
	}
}

// CheckNotExist 检测文件状态
func CheckNotExist(src string) bool {
	_, err := os.Stat(src)

	return os.IsNotExist(err)
}

// IsNotExistMkDir 自动生成空文件
func IsNotExistMkDir(src string) error {
	if notExist := CheckNotExist(src); notExist == true {
		if err := MkDir(src); err != nil {
			return err
		}
	}

	return nil
}

// SaveConfig 保存全局变量中的配置文件到yaml中
func SaveConfig() error {
	b, err := yaml.Marshal(global.GVA_CONFIG)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("config.yaml", b, 0644)
	if err != nil {
		// handle error
		fmt.Println("SaveConfig 写入配置文件失败: " + err.Error())
		return err
	}
	return nil
}
