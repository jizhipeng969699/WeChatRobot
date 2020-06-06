package main

import (
	"fmt"
	"log"
	"robot/consts"
	"strings"
)

func main() {
	wechat := NewWeChat()

	wechat.GetUUID()
}

type weChat struct {
	uuid string
}

func NewWeChat() *weChat {
	return &weChat{}
}

func (this *weChat) GetUUID() error {
	if this.uuid != "" {
		return nil
	}

	uri := consts.LoginBaseURL + "/jslogin?appid=wx782c26e4c19acffb&fun=new&lang=zh_CN&_=" + this.timestamp()
	//result: window.QRLogin.code = 200; window.QRLogin.uuid = "xxx"; //wx782c26e4c19acffb  wxeb7ec651dd0aefa9
	data, err := this.get(uri)
	if err != nil {
		log.Printf("get uuid fail err:%v", err)
		return err
	}

	res := make(map[string]string)
	datas := strings.Split(string(data), ";")
	for _, d := range datas {
		kvs := strings.Split(d, " = ")
		if len(kvs) == 2 {
			res[strings.Trim(kvs[0], " ")] = strings.Trim(strings.Trim(kvs[1], " "), "\"")
		}
	}
	if res["window.QRLogin.code"] == "200" {
		if uuid, ok := res["window.QRLogin.uuid"]; ok {
			this.uuid = uuid
			return nil
		}
	}

	return fmt.Errorf(string(data))
}
