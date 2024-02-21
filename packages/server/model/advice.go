package model

type Advice struct {
	Id        int    `json:"id"`        //主键
	Account   string `json:"account"`   //账号
	Anonymity int    `json:"anonymity"` //匿名为1，实名为0
	Time      string `json:"time"`      //时间
	Question  string `json:"question"`  //问题
	Advice    string `json:"advice"`    //建议
}
