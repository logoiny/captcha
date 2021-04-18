package modle

type CaptchaGet struct {
	ImgHigh    int  `json:"img_high"`
	ImgWidth   int  `json:"img_width"`
	DfWh       bool `json:"df_wh"`
}

type CaptchaGetRsp struct {
	Buf       string `json:"buf"`
	CaptchaId string `json:"captcha_id"`
}

type CaptchaReload struct {
	CaptchaId string `json:"captcha_id"`
	ImgHigh    int  `json:"img_high"`
	ImgWidth   int  `json:"img_width"`
	DfWh       bool `json:"df_wh"`
}

type CaptchaReloadRsp struct {
	Buf       string `json:"buf"`
}

type CaptchaVerify struct {
	CaptchaId string `json:"captcha_id"`
	Captcha string `json:"captcha"`
}

type CaptchaVerifyRsp struct {
	Passed bool `json:"passed"`
}