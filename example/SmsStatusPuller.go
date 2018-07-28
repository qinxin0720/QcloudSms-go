package main

import (
	"github.com/qinxin0720/QcloudSms-go/QcloudSms"
	"net/http"
	"fmt"
)

var appID = 122333333
var appKey = "111111111112132312xx"
var phoneNumber = "21212313123"

func callback(err error, resp *http.Response, resData string) {
	if err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Println("response data: ", resData)
	}
}

func main() {
	qcloudsms, err := QcloudSms.NewQcloudSms(appID, appKey)
	if err != nil {
		panic("appKey is nil")
	}
	//拉取短信回执以及回复
	//拉取短信回执
	qcloudsms.SmsStatusPuller.PullCallBack(10, callback)

	//拉取回复
	qcloudsms.SmsStatusPuller.PullReply(10, callback)

	//拉取单个手机短信状态
	//拉取短信回执
	qcloudsms.SmsMobileStatusPuller.PullCallBack(86, phoneNumber, 1511125600, 1511841600, 10, callback)

	//拉取回复
	qcloudsms.SmsMobileStatusPuller.PullReply(86, phoneNumber, 1511125600, 1511841600, 10, callback)
}
