package main

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gfile"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/officialaccount"
	miniConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/material"
)


//set GOARCH=amd64
//set GOOS=linux
//go build -o gzh_image_tool main.go


//var gzhUrl string ="https://api.weixin.qq.com/cgi-bin/media/uploadimg?access_token="
var appid = "wx6c12f09296b8900e"
var appsecret = "32177ebc8e8f50c22dfb5258a7576c9a"

func main() {
	fmt.Println("开始运行->")

	//dao.UserImage.OmitEmpty().OrderDesc("id")
	img := "https://qiniu.storage.heifengni.com/tmp_3ece585eb1ef1339d03fb19c2111d0215e87e2538681369a.jpg"
	savePath := "pinguo.png"
	saveDir := "/webapp/img/gzh_img_tool/"
	down(img,savePath)
	upload(saveDir+savePath)
	_ = gfile.Remove(saveDir+savePath)
}


func upload(path string) string {
	gzh := newOfficialAccount(appid, appsecret)
	u, err := gzh.addImg(path)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return u
}
func down(url, path string) {
	if r, err := g.Client().Get(url); err != nil {
		panic(err)
	} else {
		defer r.Close()
		_ = gfile.PutBytes(path, r.ReadAll())
	}
}


type OffAcc struct {
	*officialaccount.OfficialAccount
}
func newOfficialAccount(appid, secret string) *OffAcc {
	wc := wechat.NewWechat()
	cacheClient := cache.NewMemory()
	wc.SetCache(cacheClient)
	cfg := &miniConfig.Config{
		AppID:     appid,
		AppSecret: secret,
		Cache:     nil,
	}
	return &OffAcc{wc.GetOfficialAccount(cfg)}
}

func (oa *OffAcc) addImg(filename string) (url string, err error) {
	_, u, err := oa.GetMaterial().AddMaterial(material.MediaTypeImage, filename)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	return u, nil

}
func (oa *OffAcc) addVideo(filename, title, introduction string) (url string, err error) {
	_, u, err := oa.GetMaterial().AddVideo(filename, title, introduction)
	if err != nil {
		return "", err
	}
	return u, nil
}
