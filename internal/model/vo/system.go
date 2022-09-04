package vo

type LoginResp struct {
	Msg   string `json:"msg"`
	Code  int    `json:"code"`
	Token string `json:"token"`
}

type ImageResp struct {
	Msg            string `json:"msg"`
	Img            string `json:"img"`
	Code           int    `json:"code"`
	CaptchaEnabled bool   `json:"captchaEnabled"`
	Uuid           string `json:"uuid"` // 用户id
}
