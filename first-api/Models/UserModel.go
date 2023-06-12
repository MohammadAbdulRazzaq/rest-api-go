package Models

type User struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}
type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (b *User) TableName() string {
	return "mobiles"
}
