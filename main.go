package main

import (
	"os"
	"bufio"
	"io"
	"strings"
	"fmt"
)

var (
	mob =[]string{
		"",
	}
)

func main() {
	contentLine,err:=readFileByLine("")
	if err != nil {
		fmt.Printf("运行失败!err: %s \n",err.Error())
		return
	}
	var (
		writeByteArray []byte
	)
	for _, value := range *contentLine {
		for _, item := range mob {
			if strings.Index(value, item) > 0 {
				writeByteArray=append(writeByteArray,[]byte(value+"\n")...)
			}
		}
	}
	err=fileCreateAndWrite(&writeByteArray,"")
	if err != nil {
		fmt.Printf("运行失败 err: %s \n",err.Error())
		return
	}
	fmt.Println("运行成功!")
}

/**
文件读取逐行进行读取
创建人:邵炜
创建时间:2016年9月20日10:23:41
输入参数: 文件路劲
输出参数: 字符串数组(数组每一项对应文件的每一行) 错误对象
*/
func readFileByLine(filePath string) (*[]string, error) {
	var (
		readAll     = false
		readByte    []byte
		line        []byte
		err         error
		contentLine []string
	)
	fs, err := os.Open(filePath)
	if err != nil {

		return nil, err
	}
	defer fs.Close()
	buf := bufio.NewReader(fs)
	for err != io.EOF {
		if err != nil {

		}
		if readAll {
			readByte, readAll, err = buf.ReadLine()
			line = append(line, readByte...)
		} else {
			readByte, readAll, err = buf.ReadLine()
			line = append(line, readByte...)
			if len(strings.TrimSpace(string(line))) == 0 {
				continue
			}
			contentLine = append(contentLine, string(line))
			line = line[:0]
		}
	}

	return &contentLine, nil
}

/**
写文件
创建人:邵炜
创建时间:2016年9月7日16:31:39
输入参数:文件内容 写入文件的路劲(包含文件名)
输出参数:错误对象
*/
func fileCreateAndWrite(content *[]byte, fileName string) error {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(*content)
	if err != nil {
		return err
	}
	return nil
}