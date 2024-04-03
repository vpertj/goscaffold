package service

import (
	"fmt"
	"os"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/valyala/fasthttp"
)

const (
	accessKey = "UsM-wojWOGvo0aXSJl2ld6IwjUWLP3qFfh131qu0"
	secretKey = "WJ1CCzzOnsYN2NLVB_lPWp_pTHGhW0OPQj3Ci3-_"
	bucket    = "bigfileuplad"
)

type UploadService struct {
}
type ReturnValue struct {
	KEY  string
	HASH string
	MSG  string
}

func (*UploadService) UploadFile(ctx *fasthttp.RequestCtx, filePath string) (ReturnValue, error) {

	// 七牛云配置
	mac := qbox.NewMac(accessKey, secretKey)
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	key := "video/upload/"
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	//formUploader := storage.NewFormUploader(&cfg)
	//putRet := storage.PutRet{}
	recorder, err := storage.NewFileRecorder(os.TempDir())
	// 创建七牛云上传表单
	ret := storage.PutRet{}
	//putExtra := storage.PutExtra{}
	putExtra := storage.RputV2Extra{
		Recorder: recorder,
	}
	formUploader := storage.NewResumeUploaderV2(&cfg)

	err = formUploader.PutFile(ctx, &ret, upToken, key, filePath, &putExtra)
	if err != nil {
		return ReturnValue{
			KEY:  "",
			HASH: "",
			MSG:  "文件上传失败",
		}, err
	}
	fmt.Println(ret.Key, ret.Hash)
	return ReturnValue{
		KEY:  ret.Key,
		HASH: ret.Hash,
		MSG:  "文件上传成功",
	}, nil
}
