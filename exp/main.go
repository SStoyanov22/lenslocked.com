package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "lenslocked_dev"
	password = "lenslocked_dev"
	dbname   = "lenslocked_dev"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", host, port, user, password, dbname)
	db, err := gorm.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	user := User{
Name: "Michael Scott",
Email: "michael@dundermifflin.com",
}
if err := us.Create(&user); err != nil {
panic(err)
}
// NOTE: You may need to update the query code a bit as well
foundUser, err := us.ByID(1)
if err != nil {
panic(err)
}
fmt.Println(foundUser)
	// db.LogMode(true)

	// db.AutoMigrate(&User{}, &Order{})

	// var user User
	// db.Preload("Orders").First(&user)
	// if db.Error != nil {
	// 	panic(db.Error)
	// }

	// fmt.Println("name ", user.Name)
	// fmt.Println("email ", user.Email)
	// fmt.Println("orders", user.Orders)

}
func (us *User) Create(user *User) error {
	return us.db.Create(user).Error
	}

type User struct {
	gorm.Model
	Name   string
	Email  string `gorm:"not null;unique index"`
	Orders []Order
}

type Order struct {
	gorm.Model
	UserId      uint
	Amount      int
	Description string
}

func createOrdeR(db *gorm.DB, user User, amount int, description string) {
	db.Create(&Order{
		UserId:      user.ID,
		Amount:      amount,
		Description: description,
	})

	if db.Error != nil {
		panic(db.Error)
	}
}

func getInfo() (name, email string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What is your name?")
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Println("What is your email")
	email, _ = reader.ReadString('\n')
	email = strings.TrimSpace(email)
	return name, email
}
