package req

type LoginReq struct {
	Code     int    `json:"code"`
	Username string `json:"username"`
	Password string `json:"password"`
	Uuid     string `json:"uuid"`
}
