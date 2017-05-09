package models

type Error struct {
	Status string
	Msg    string
}

type RetMsg struct {
	ErrorCode int    `json:"errCode"`
	ErrorMsg  string `json:"errMsg"`
}

func (this *RetMsg) setDefaultValue() {
	this.ErrorCode = 0
	this.ErrorMsg = "操作成功"
}

func (this *RetMsg) NewRetMsg(errCode int, errMsg string) *RetMsg {
	var retMsg RetMsg
	retMsg.ErrorCode = errCode
	retMsg.ErrorMsg = errMsg
	return &retMsg
}
