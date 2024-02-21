package StudentController

import (
	"SelectionSystem/model"
	"SelectionSystem/service/StudentService"
	"SelectionSystem/utility"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type AdviceData struct {
	Account   string `json:"account"`   //账号
	Anonymity int    `json:"anonymity"` //匿名为1，实名为0
	Question  string `json:"question"`  //问题
	Advice    string `json:"advice"`    //建议
}

// 意见箱
func PostAdvice(c *gin.Context) {
	var data AdviceData
	var advice model.Advice
	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Println(err)
		utility.JsonResponseInternalServerError(c)
		return
	}
	advice.Account = data.Account
	advice.Anonymity = data.Anonymity
	advice.Question = data.Question
	advice.Advice = data.Advice
	advice.Time = time.Now().Format("2006-01-02 15:04:05")

	//将意见存入数据库
	err = StudentService.CreateAdvice(&advice)
	if err != nil {
		utility.JsonResponseInternalServerError(c)
		return
	}
	utility.JsonSuccessResponse(c, nil)
}
