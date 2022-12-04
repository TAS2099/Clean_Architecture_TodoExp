package domain

type Task struct {
	ID   int    `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Task string `json:"task"`
	Exp  int    `json:"exp"`
	Due  string `json:"due"`
	Done bool   `json:"done"`
}
