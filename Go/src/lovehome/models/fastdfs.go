package models

import (
	"astaxie/beego/logs"
	"github.com/tedcy/fdfs_client"
)

func UploadByName(filename string) {
	f,err := fdfs_client.NewClientWithConfig("../conf/client.go")
	defer f.Destory()
	if err != nil{
		logs.Error("创建fdfs_client失败",err.Error())
		return
	}

	fileId,errupload := f.UploadByFilename(filename)
	if fileId != nil{
		logs.Error("上传文件失败",errupload.Error())
		return
	}

	logs.Info(fileId)
}
