package models

type Student struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Grade   string `json:"grade"`
	Address string `json:"address"`
}
