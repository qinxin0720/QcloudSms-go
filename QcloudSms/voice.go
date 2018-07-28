package QcloudSms

import (
	"net/url"
	"strconv"
	"errors"
)

type promptVoiceSender struct {
	appID int
	appKey,
	url string
}

func newPromptVoiceSender(appID int, appKey string) *promptVoiceSender {
	return &promptVoiceSender{
		appID:  appID,
		appKey: appKey,
		url:    `https://cloud.tim.qq.com/v5/tlsvoicesvr/sendvoiceprompt`,
	}
}

//Send 发送语音通知和验证码
func (s *promptVoiceSender) Send(nationCode int, phoneNumber string, promptType int, msg string, playTimes int, ext string, callback callbackFunc) error {
	if playTimes <= 0 {
		return errors.New("playtimes must great than zero")
	}
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
		Tel        tel    `json:"tel"`
		PromptType string `json:"prompttype"`
		PromptFile string `json:"promptfile"`
		PlayTimes  string `json:"playtimes"`
		Sig        string `json:"sig"`
		Time       int64  `json:"time"`
		Ext        string `json:"ext"`
	}

	Tel := tel{
		NationCode: strconv.Itoa(nationCode),
		Mobile:     phoneNumber,
	}

	Body := body{
		Tel:        Tel,
		PromptType: strconv.Itoa(promptType),
		PromptFile: msg,
		PlayTimes:  strconv.Itoa(playTimes),
		Sig:        calculateSignature(s.appKey, random, now, phoneNumbers),
		Time:       now,
		Ext:        ext,
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

type codeVoiceSender struct {
	appID int
	appKey,
	url string
}

func newCodeVoiceSender(appID int, appKey string) *codeVoiceSender {
	return &codeVoiceSender{
		appID:  appID,
		appKey: appKey,
		url:    `https://cloud.tim.qq.com/v5/tlsvoicesvr/sendcvoice`,
	}
}

func (s *codeVoiceSender) Send(nationCode int, phoneNumber, msg string, playTimes int, ext string, callback callbackFunc) error {
	if playTimes <= 0 {
		return errors.New("playtimes must great than zero")
	}
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
		Tel       tel    `json:"tel"`
		Msg       string `json:"msg"`
		PlayTimes string `json:"playtimes"`
		Sig       string `json:"sig"`
		Time      int64  `json:"time"`
		Ext       string `json:"ext"`
	}

	Tel := tel{
		NationCode: strconv.Itoa(nationCode),
		Mobile:     phoneNumber,
	}

	Body := body{
		Tel:       Tel,
		Msg:       msg,
		PlayTimes: strconv.Itoa(playTimes),
		Sig:       calculateSignature(s.appKey, random, now, phoneNumbers),
		Time:      now,
		Ext:       ext,
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

type ttsVoiceSender struct {
	appID int
	appKey,
	url string
}

func newTtsVoiceSender(appID int, appKey string) *ttsVoiceSender {
	return &ttsVoiceSender{
		appID:  appID,
		appKey: appKey,
		url:    `https://cloud.tim.qq.com/v5/tlsvoicesvr/sendtvoice`,
	}
}

func (s *ttsVoiceSender) Send(nationCode int, phoneNumber string, templID int, params []string, playTimes int, ext string, callback callbackFunc) error {
	if playTimes <= 0 {
		return errors.New("playtimes must great than zero")
	}
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
		Tel       tel      `json:"tel"`
		TplID     string   `json:"tpl_id"`
		Params    []string `json:"params"`
		PlayTimes string   `json:"playtimes"`
		Sig       string   `json:"sig"`
		Time      int64    `json:"time"`
		Ext       string   `json:"ext"`
	}

	Tel := tel{
		NationCode: strconv.Itoa(nationCode),
		Mobile:     phoneNumber,
	}

	Body := body{
		Tel:       Tel,
		Params:    params,
		TplID:     strconv.Itoa(templID),
		PlayTimes: strconv.Itoa(playTimes),
		Sig:       calculateSignature(s.appKey, random, now, phoneNumbers),
		Time:      now,
		Ext:       ext,
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

type fileVoiceSender struct {
	appID int
	appKey,
	url string
}

func newFileVoiceSender(appID int, appKey string) *fileVoiceSender {
	return &fileVoiceSender{
		appID:  appID,
		appKey: appKey,
		url:    `https://cloud.tim.qq.com/v5/tlsvoicesvr/sendfvoice`,
	}
}

func (s *fileVoiceSender) Send(nationCode int, phoneNumber, fid string, playTimes int, ext string, callback callbackFunc) error {
	if playTimes <= 0 {
		return errors.New("playtimes must great than zero")
	}
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
		Tel       tel    `json:"tel"`
		Fid       string `json:"fid"`
		PlayTimes string `json:"playtimes"`
		Sig       string `json:"sig"`
		Time      int64  `json:"time"`
		Ext       string `json:"ext"`
	}

	Tel := tel{
		NationCode: strconv.Itoa(nationCode),
		Mobile:     phoneNumber,
	}

	Body := body{
		Tel:       Tel,
		Fid:       fid,
		PlayTimes: strconv.Itoa(playTimes),
		Sig:       calculateSignature(s.appKey, random, now, phoneNumbers),
		Time:      now,
		Ext:       ext,
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

type voiceFileUploader struct {
	appID int
	appKey,
	url string
}

func newVoiceFileUploader(appID int, appKey string) *voiceFileUploader {
	return &voiceFileUploader{
		appID:  appID,
		appKey: appKey,
		url:    `https://cloud.tim.qq.com/v5/tlsvoicesvr/uploadvoicefile`,
	}
}

func (u *voiceFileUploader) Upload(fileContent []byte, contentType string, callback callbackFunc) error {
	if contentType != "mp3" && contentType != "wav" {
		return errors.New(`contentType is invalid and should be 'mp3' or 'wav'`)
	}
	var Type string
	if contentType == "wav" {
		Type = "audio/wav"
	} else {
		Type = "audio/mpeg"
	}

	reqUrl, err := url.Parse(u.url)
	if err != nil {
		return err
	}
	random := getRandom()
	now := getCurrentTime()
	fileSha1Sum := sha1sum(fileContent)
	headers := make(map[string]string)
	headers["Content-Type"] = Type
	headers["x-content-sha1"] = fileSha1Sum
	headers["Authorization"] = calculateAuth(u.appKey, random, now, fileSha1Sum)

	option := option{
		Protocol: reqUrl.Scheme,
		Host:     reqUrl.Host,
		Path:     reqUrl.Path + "?sdkappid=" + strconv.Itoa(u.appID) + "&random=" + strconv.Itoa(random),
		Method:   "POST",
		Headers:  headers,
		Body:     fileContent,
	}
	return request(option, callback)
}
