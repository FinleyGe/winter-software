package StudentController

import (
	"SelectionSystem/model"
	"SelectionSystem/service/StudentService"
	"SelectionSystem/utility"
	"github.com/gin-gonic/gin"
	"log"
)

//type UpdateData struct {
//	Name        string  `json:"name"`
//	Account     string  `json:"account"`
//	Score       float64 `json:"score"`
//	PhoneNumber string  `json:"phone_number"`
//	Position    string  `json:"position"`
//	Rank        int     `json:"rank"`
//	MathOf1     int     `json:"math_of_1"`
//	MathOf2     int     `json:"math_of_2"`
//	CppOf1      int     `json:"cpp_of_1"`
//	CppOf2      int     `json:"cpp_of_2"`
//	Profile     string  `json:"profile"`
//	TeacherName string  `json:"teacher_name"`
//}

func UpdateForm(c *gin.Context) {
	var data model.Form
	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Println(err)
		utility.JsonResponseInternalServerError(c)
	}

	err = StudentService.UpdateForm(data)
	if err != nil {
		log.Println(err)
		return
	}
	utility.JsonResponse(200, "OK", nil, c)
}
