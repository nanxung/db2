package main

import (
	"github.com/vladimirvivien/gowfs"
	"log"
	"io/ioutil"
	"fmt"
)

/*
字符流 读取数据放入缓存区
 */
func readFiles(uri string,offset int64,length int64,buffSize int) []byte{
	fs,err:=gowfs.NewFileSystem(gowfs.Configuration{Addr:"tdp01:50070",User:"work"})
	if err!=nil{
		log.Fatal(err)
	}
	checksum,err:=fs.GetFileChecksum(gowfs.Path{Name:uri})
	if err!=nil{
		log.Fatal(err)
	}
	data,err:=fs.Open(gowfs.Path{Name:uri},offset,length,buffSize)
	recvData,_:=ioutil.ReadAll(data)
	fmt.Println(checksum)
	fmt.Println(string(recvData))
	return recvData
}

func main(){
	fmt.Println(readFiles("/toclickhouse/file",0,512,2048))
}