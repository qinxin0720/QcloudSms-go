package main

import (
	"github.com/qinxin0720/QcloudSms-go/QcloudSms"
	"net/http"
	"fmt"
)

var appID = 122333333
var appKey = "111111111112132312xx"
var phoneNumber = "21212313123"
var templId = 7839
var params = []string{"指定模板单发", "深圳", "小明"}

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
	//单发短信
	//发送短信没有指定模板ID时，发送的内容需要与已审核通过的模板内容相匹配，才可能下发成功，否则返回失败。
	//如需发送海外短信，同样可以使用此接口，只需将国家码"86"改写成对应国家码号。
	qcloudsms.SmsSingleSender.Send(0, 86, phoneNumber,
		"测试短信，普通单发，深圳，小明，上学。", "", "", callback)

	//指定模板ID单发
	//无论单发短信还是指定模板 ID 单发短信都需要从控制台中申请模板并且模板已经审核通过，才可能下发成功，否则返回失败。
	qcloudsms.SmsSingleSender.SendWithParam(86, phoneNumber, templId, params, "", "", "", callback)
}
