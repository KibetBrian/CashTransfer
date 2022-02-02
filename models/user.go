package models
import(
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	PhoneNumber  int`gorm:"primaryKey; "notNull"`
	Email string `gorm:type:varchar(255); json: "email" `
	Username string `gorm:"type:varchar(255);" json: "title"`
	Password string `gorm: "type:varchar(255);`
}

func getUser() User {
	var user User
	return user
}

func getUsers () []User{
	var users []User;
	return users;
}