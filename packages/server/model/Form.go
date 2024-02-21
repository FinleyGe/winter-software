package model

type Form struct {
	State                     int     `json:"state"` //0未审批 1同意 2驳回
	Id                        int     `json:"id"`
	Name                      string  `json:"name"`
	Account                   string  `json:"account"`
	Score                     float64 `json:"score"`
	PhoneNumber               string  `json:"phone_number"`
	Position                  string  `json:"position"`
	Rank                      int     `json:"rank"`
	MathOf1                   int     `json:"math_of_1"`
	MathOf2                   int     `json:"math_of_2"`
	CppOf1                    int     `json:"cpp_of_1"`
	CppOf2                    int     `json:"cpp_of_2"`
	Profile                   string  `json:"profile"`
	AvatarAddress             string  `json:"avatar_address"`
	TeacherName               string  `json:"teacher_name"`
	TimeOfStudent             string  `json:"time_of_student"`
	TimeOfTeacher             string  `json:"time_of_teacher"`
	Advice                    string  `json:"advice"`
	Check                     int     `json:"check"`
	AutographAddressOfStudent string  `json:"autograph_address_of_student"`
	AutographAddressOfTeacher string  `json:"autograph_address_of_teacher"`
}
