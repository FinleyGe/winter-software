package model

type Reason struct {
	Value   int    `gorm:"primaryKey;autoIncrement" json:"value"`
	Account string `json:"account"`
	Label   string `json:"label"`
}
