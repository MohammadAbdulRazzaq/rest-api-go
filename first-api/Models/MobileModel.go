package models

type Mobile struct {
	ID     uint   `json:"id"`
	Brand  string `json:"brand"`
	Model  string `json:"model"`
	Price  uint   `json:"price"`
	MEMORY uint   `json:"memory"`
	CAMERA uint   `json:"camera"`
}
