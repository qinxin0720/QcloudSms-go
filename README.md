# 腾讯云短信SDK Go语言实现

## 说明
> 此 SDK 为非官方版本，命名和结构上与官方版本有一些区别。

> 海外短信和国内短信使用同一接口，只需替换相应的国家码与手机号码，每次请求群发接口手机号码需全部为国内或者海外手机号码。

> 语音通知目前支持语音验证码以及语音通知功能。

## 功能

## Features

##### 短信
- [x] 单发短信
- [x] 指定模板单发短信
- [x] 群发短信
- [x] 群发模板短信
- [ ] 短信下发状态通知
- [ ] 短信回复
- [x] 拉取短信状态
- [x] 拉取单个手机短信状态

##### 语音
- [x] 发送语音验证码
- [x] 发送语音通知
- [ ] 语音验证码状态通知
- [ ] 语音通知状态通知
- [ ] 语音通知按键通知
- [ ] 语音送达失败原因推送

## 如何使用

### 单发短信

```go
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
		panic("appKey is nil")
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
```

### 群发短信

```go
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
		panic("appKey is nil")
	}
	//群发短信
	qcloudsms.SmsMultiSender.Send(0, 86, phoneNumbers,
		"测试短信，普通群发，深圳，小明，上学。", "", "", callback)

	//指定模板ID群发
	//群发一次请求最多支持 200 个号码
	qcloudsms.SmsMultiSender.SendWithParam(86, phoneNumbers, templId, params, "", "", "", callback)
}
```

### 发送语音验证码

```go
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
	//发送语音验证码
	//语音验证码发送只需提供验证码数字，例如在 msg=“123”，您收到的语音通知为“您的语音验证码是1 2 3”，如需自定义内容，可以使用语音通知。
	qcloudsms.SmsVoiceVerifyCodeSender.Send(86, phoneNumber, "1234", 2, "", callback)

	//发送语音通知
	qcloudsms.SmsVoicePromptSender.Send(86, phoneNumber, 2, "1234", 2, "", callback)
}
```

### 拉取短信回执以及回复

```go
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
```

## 许可

这个项目使用MIT许可.