package QcloudSms

import (
	"net/url"
	"strconv"
)

type smsSingleSender struct {
	appID int
	appKey,
	url string
}

func newSmsSingleSender(appID int, appKey string) *smsSingleSender {
	return &smsSingleSender{
		appID:  appID,
		appKey: appKey,
		url:    `https://yun.tim.qq.com/v5/tlssmssvr/sendsms`,
	}
}

func (s *smsSingleSender) Send(msgType, nationCode int, phoneNumber, msg, extend, ext string, callback callbackFunc) error {
	reqUrl, err := url.Parse(s.url)
	if err != nil {
		return err
	}
	random := getRandom()
	now := getCurrentTime()
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	var phoneNumbers []string
	phoneNumbers = append(phoneNumbers, phoneNumber)

	type body struct {
		Tel    tel    `json:"tel"`
		Type   int    `json:"type"`
		Msg    string `json:"msg"`
		Sig    string `json:"sig"`
		Time   int64  `json:"time"`
		Extend string `json:"extend"`
		Ext    string `json:"ext"`
	}

	Tel := tel{
		NationCode: strconv.Itoa(nationCode),
		Mobile:     phoneNumber,
	}

	Body := body{
		Tel:    Tel,
		Type:   msgType,
		Msg:    msg,
		Sig:    calculateSignature(s.appKey, random, now, phoneNumbers),
		Time:   now,
		Extend: extend,
		Ext:    ext,
	}

	option := option{
		Protocol: reqUrl.Scheme,
		Host:     reqUrl.Host,
		Path:     reqUrl.Path + "?sdkappid=" + strconv.Itoa(s.appID) + "&random=" + strconv.Itoa(random),
		Method:   "POST",
		Headers:  headers,
		Body:     Body,
	}
	return request(option, callback)
}

func (s *smsSingleSender) SendWithParam(nationCode int, phoneNumber string, templID int, params []string, sign, extend, ext string, callback callbackFunc) error {
	reqUrl, err := url.Parse(s.url)
	if err != nil {
		return err
	}
	random := getRandom()
	now := getCurrentTime()
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	var phoneNumbers []string
	phoneNumbers = append(phoneNumbers, phoneNumber)

	type body struct {
		Tel    tel      `json:"tel"`
		Sign   string   `json:"sign"`
		TplID  int      `json:"tpl_id"`
		Params []string `json:"params"`
		Sig    string   `json:"sig"`
		Time   int64    `json:"time"`
		Extend string   `json:"extend"`
		Ext    string   `json:"ext"`
	}

	Tel := tel{
		NationCode: strconv.Itoa(nationCode),
		Mobile:     phoneNumber,
	}

	Body := body{
		Tel:    Tel,
		Sign:   sign,
		TplID:  templID,
		Params: params,
		Sig:    calculateSignature(s.appKey, random, now, phoneNumbers),
		Time:   now,
		Extend: extend,
		Ext:    ext,
	}

	option := option{
		Protocol: reqUrl.Scheme,
		Host:     reqUrl.Host,
		Path:     reqUrl.Path + "?sdkappid=" + strconv.Itoa(s.appID) + "&random=" + strconv.Itoa(random),
		Method:   "POST",
		Headers:  headers,
		Body:     Body,
	}
	return request(option, callback)
}

type smsMultiSender struct {
	appID int
	appKey,
	url string
}

func newSmsMultiSender(appID int, appKey string) *smsMultiSender {
	return &smsMultiSender{
		appID:  appID,
		appKey: appKey,
		url:    `https://yun.tim.qq.com/v5/tlssmssvr/sendmultisms2`,
	}
}

func (s *smsMultiSender) Send(msgType, nationCode int, phoneNumbers []string, msg, extend, ext string, callback callbackFunc) error {
	reqUrl, err := url.Parse(s.url)
	if err != nil {
		return err
	}
	random := getRandom()
	now := getCurrentTime()
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	type body struct {
		Tel    []tel  `json:"tel"`
		Type   int    `json:"type"`
		Msg    string `json:"msg"`
		Sig    string `json:"sig"`
		Time   int64  `json:"time"`
		Extend string `json:"extend"`
		Ext    string `json:"ext"`
	}

	var Tel []tel
	for _, v := range phoneNumbers {
		Tel = append(Tel, tel{
			NationCode: strconv.Itoa(nationCode),
			Mobile:     v,
		})
	}

	Body := body{
		Tel:    Tel,
		Type:   msgType,
		Msg:    msg,
		Sig:    calculateSignature(s.appKey, random, now, phoneNumbers),
		Time:   now,
		Extend: extend,
		Ext:    ext,
	}

	option := option{
		Protocol: reqUrl.Scheme,
		Host:     reqUrl.Host,
		Path:     reqUrl.Path + "?sdkappid=" + strconv.Itoa(s.appID) + "&random=" + strconv.Itoa(random),
		Method:   "POST",
		Headers:  headers,
		Body:     Body,
	}
	return request(option, callback)
}

func (s *smsMultiSender) SendWithParam(nationCode int, phoneNumbers []string, templID int, params []string, sign, extend, ext string, callback callbackFunc) error {
	reqUrl, err := url.Parse(s.url)
	if err != nil {
		return err
	}
	random := getRandom()
	now := getCurrentTime()
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	type body struct {
		Tel    []tel    `json:"tel"`
		Sign   string   `json:"sign"`
		TplID  int      `json:"tpl_id"`
		Params []string `json:"params"`
		Sig    string   `json:"sig"`
		Time   int64    `json:"time"`
		Extend string   `json:"extend"`
		Ext    string   `json:"ext"`
	}

	var Tel []tel
	for _, v := range phoneNumbers {
		Tel = append(Tel, tel{
			NationCode: strconv.Itoa(nationCode),
			Mobile:     v,
		})
	}

	Body := body{
		Tel:    Tel,
		Sign:   sign,
		TplID:  templID,
		Params: params,
		Sig:    calculateSignature(s.appKey, random, now, phoneNumbers),
		Time:   now,
		Extend: extend,
		Ext:    extend,
	}

	option := option{
		Protocol: reqUrl.Scheme,
		Host:     reqUrl.Host,
		Path:     reqUrl.Path + "?sdkappid=" + strconv.Itoa(s.appID) + "&random=" + strconv.Itoa(random),
		Method:   "POST",
		Headers:  headers,
		Body:     Body,
	}
	return request(option, callback)
}

type smsStatusPuller struct {
	appID int
	appKey,
	url string
}

func newSmsStatusPuller(appID int, appKey string) *smsStatusPuller {
	return &smsStatusPuller{
		appID:  appID,
		appKey: appKey,
		url:    `https://yun.tim.qq.com/v5/tlssmssvr/pullstatus`,
	}
}

func (s *smsStatusPuller) Pull(msgType, max int, callback callbackFunc) error {
	reqUrl, err := url.Parse(s.url)
	if err != nil {
		return err
	}
	random := getRandom()
	now := getCurrentTime()
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	type body struct {
		Sig  string `json:"sig"`
		Type int    `json:"type"`
		Time int64  `json:"time"`
		Max  int    `json:"max"`
	}

	Body := body{
		Sig:  calculateSignature(s.appKey, random, now, []string{}),
		Type: msgType,
		Time: now,
		Max:  max,
	}

	option := option{
		Protocol: reqUrl.Scheme,
		Host:     reqUrl.Host,
		Path:     reqUrl.Path + "?sdkappid=" + strconv.Itoa(s.appID) + "&random=" + strconv.Itoa(random),
		Method:   "POST",
		Headers:  headers,
		Body:     Body,
	}
	return request(option, callback)
}

func (s *smsStatusPuller) PullCallBack(max int, callback callbackFunc) error {
	return s.Pull(0, max, callback)
}

func (s *smsStatusPuller) PullReply(max int, callback callbackFunc) error {
	return s.Pull(1, max, callback)
}

type smsMobileStatusPuller struct {
	appID int
	appKey,
	url string
}

func newSmsMobileStatusPuller(appID int, appKey string) *smsMobileStatusPuller {
	return &smsMobileStatusPuller{
		appID:  appID,
		appKey: appKey,
		url:    `https://yun.tim.qq.com/v5/tlssmssvr/pullstatus4mobile`,
	}
}

func (s *smsMobileStatusPuller) Pull(msgType, nationCode int, mobile string, beginTime, endTime, max int, callback callbackFunc) error {
	reqUrl, err := url.Parse(s.url)
	if err != nil {
		return err
	}
	random := getRandom()
	now := getCurrentTime()
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	type body struct {
		Sig        string `json:"sig"`
		Type       int    `json:"type"`
		Time       int64  `json:"time"`
		Max        int    `json:"max"`
		BeginTime  int    `json:"begin_time"`
		EndTime    int    `json:"end_time"`
		NationCode string `json:"nationcode"`
		Mobile     string `json:"mobile"`
	}

	Body := body{
		Sig:        calculateSignature(s.appKey, random, now, []string{}),
		Type:       msgType,
		Time:       now,
		Max:        max,
		BeginTime:  beginTime,
		EndTime:    endTime,
		NationCode: strconv.Itoa(nationCode),
		Mobile:     mobile,
	}

	option := option{
		Protocol: reqUrl.Scheme,
		Host:     reqUrl.Host,
		Path:     reqUrl.Path + "?sdkappid=" + strconv.Itoa(s.appID) + "&random=" + strconv.Itoa(random),
		Method:   "POST",
		Headers:  headers,
		Body:     Body,
	}
	return request(option, callback)
}

func (s *smsMobileStatusPuller) PullCallBack(nationCode int, mobile string, beginTime, endTime, max int, callback callbackFunc) error {
	return s.Pull(0, nationCode, mobile, beginTime, endTime, max, callback)
}

func (s *smsMobileStatusPuller) PullReply(nationCode int, mobile string, beginTime, endTime, max int, callback callbackFunc) error {
	return s.Pull(0, nationCode, mobile, beginTime, endTime, max, callback)
}
