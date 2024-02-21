package router

import (
	"SelectionSystem/controller/ReasonsController"
	"SelectionSystem/controller/SignatureController"
	"SelectionSystem/controller/StudentController"
	"SelectionSystem/controller/TFormController"
	"SelectionSystem/controller/UserController"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetRouter() *gin.Engine {
	r := gin.Default()

	Api := r.Group("/m1/3977821-0-default")
	{
		// 登录
		Api.POST("/login", UserController.LoginCheck)
		//修改密码
		Api.POST("/resetpwd", UserController.ResetPwd)
		//上传签名
		Api.POST("/signature", SignatureController.CreateSignature)
		//修改签名
		Api.PUT("/signature", SignatureController.UpdateSignature)
		Student := Api.Group("/student")
		{
			//提交表单
			Student.POST("/Form", StudentController.PostApplyForm)
			//查询表单
			Student.GET("/Form", StudentController.QueryApplyForm)
			//更新表单
			Student.PUT("/Form", StudentController.UpdateForm)
			//建议箱
			Student.POST("/Advice", StudentController.PostAdvice)
			//提交头像
			Student.POST("/Avatar", StudentController.PostAvatar)
			//更新头像
			Student.PUT("/Avatar", StudentController.UpdateAvatar)
		}
		Teacher := Api.Group("/teacher")
		{
			//显示表单
			Teacher.GET("/Form", TFormController.ShowForm)
			//处理表单
			Teacher.PUT("/Form", TFormController.HandelForm)
			//显示理由
			Teacher.GET("/reason", ReasonsController.ShowReasons)
			//添加理由
			Teacher.POST("/reason", ReasonsController.AddReasons)
			//删除理由
			Teacher.DELETE("/reason", ReasonsController.DeleteReasons)
		}
	}

	r.StaticFS("/storeFile", http.Dir("./storeFile"))

	return r
}
