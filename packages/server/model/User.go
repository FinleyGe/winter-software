package model

type User struct {
	Account      string `json:"account"`
	Password     string `json:"password"`
	Identity     string `json:"identity"`
	NumOfStudent int    `json:"num_of_student"`
}
