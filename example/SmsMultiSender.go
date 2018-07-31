package main

import (
	"github.com/qinxin0720/QcloudSms-go/QcloudSms"
	"net/http"
	"fmt"
)

var appID = 122333333
var appKey = "111111111112132312xx"
var phoneNumbers = []string{"21212313123", "12345678902", "12345678903"}
var templId = 7839
var params = []string{"指定模板群发", "深圳", "小明"}

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
		panic(err)
	}
	//群发短信
	qcloudsms.SmsMultiSender.Send(0, 86, phoneNumbers,
		"测试短信，普通群发，深圳，小明，上学。", "", "", callback)

	//指定模板ID群发
	//群发一次请求最多支持 200 个号码
	qcloudsms.SmsMultiSender.SendWithParam(86, phoneNumbers, templId, params, "", "", "", callback)
}
