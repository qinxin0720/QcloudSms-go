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
		panic(err)
	}
	//发送语音验证码
	//语音验证码发送只需提供验证码数字，例如在 msg=“123”，您收到的语音通知为“您的语音验证码是1 2 3”，如需自定义内容，可以使用语音通知。
	qcloudsms.SmsVoiceVerifyCodeSender.Send(86, phoneNumber, "1234", 2, "", callback)

	//发送语音通知
	qcloudsms.SmsVoicePromptSender.Send(86, phoneNumber, 2, "1234", 2, "", callback)
}
