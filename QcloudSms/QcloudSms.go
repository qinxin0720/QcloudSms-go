package QcloudSms

import "github.com/pkg/errors"

type qcloudsms struct {
	appid                    int
	appkey                   string
	SmsSingleSender          *smsSingleSender
	SmsMultiSender           *smsMultiSender
	SmsStatusPuller          *smsStatusPuller
	SmsMobileStatusPuller    *smsMobileStatusPuller
	SmsVoicePromptSender     *promptVoiceSender
	PromptVoiceSender        *promptVoiceSender
	SmsVoiceVerifyCodeSender *codeVoiceSender
	CodeVoiceSender          *codeVoiceSender
	TtsVoiceSender           *ttsVoiceSender
	VoiceFileUploader        *voiceFileUploader
	FileVoiceSender          *fileVoiceSender
}

//NewQcloudSms new一个qcloudsms实例
func NewQcloudSms(appid int, appkey string) (*qcloudsms, error) {
	if appkey == "" {
		return nil, errors.New("appkey is nil")
	}
	return &qcloudsms{
		appid:                    appid,
		appkey:                   appkey,
		SmsSingleSender:          newSmsSingleSender(appid, appkey),
		SmsMultiSender:           newSmsMultiSender(appid, appkey),
		SmsStatusPuller:          newSmsStatusPuller(appid, appkey),
		SmsMobileStatusPuller:    newSmsMobileStatusPuller(appid, appkey),
		SmsVoicePromptSender:     newPromptVoiceSender(appid, appkey),
		PromptVoiceSender:        newPromptVoiceSender(appid, appkey),
		SmsVoiceVerifyCodeSender: newCodeVoiceSender(appid, appkey),
		CodeVoiceSender:          newCodeVoiceSender(appid, appkey),
		TtsVoiceSender:           newTtsVoiceSender(appid, appkey),
		VoiceFileUploader:        newVoiceFileUploader(appid, appkey),
		FileVoiceSender:          newFileVoiceSender(appid, appkey),
	}, nil
}
