package models

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	ErrNotFound  = errors.New("models: resource not found")
	ErrInvalidId = errors.New("models: ID provided was invalid")
)

type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"notnull;unique_index"	`
}

type UserService struct {
	db *gorm.DB
}

func NewUserService(connectionInfo string) (*UserService, error) {
	db, err := gorm.Open("postgres", connectionInfo)
	if err != nil {
		return nil, err
	}

	db.LogMode(true)

	return &UserService{
		db: db,
	}, nil
}

func (us *UserService) Close() error {
	return us.db.Close()
}

func (us *UserService) Automigrate() error {
	return us.db.AutoMigrate().Error
}

func (us *UserService) ById(id uint) (*User, error) {
	var user User
	db := us.db.Where("id=?", id)

	err := first(db, &user)

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (us *UserService) ByEmail(email string) (*User, error) {
	var user User
	db := us.db.Where("email=?", email)
	err := first(db, user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (us *UserService) Update(user *User) error {
	return us.db.Save(user).Error
}

func (us *UserService) Delete(id uint) error {
	if id == 0 {
		return ErrInvalidId
	}

	user := User{Model: gorm.Model{ID: id}}
	return us.db.Delete(&user).Error
}

func (us *UserService) Create(user *User) error {
	return us.db.Create(&user).Error
}

func (us *UserService) DestructiveReset() error {
	err := us.db.DropTableIfExists().Error
	if err != nil {
		return err
	}

	return us.db.AutoMigrate().Error
}

//first will query using the provided gorm.DB and it will
//get the first item return and place it into dst. If
//nothing is found in the query it will return ErrNotFound
func first(db *gorm.DB, dst interface{}) error {
	err := db.First(dst).Error
	if err != nil {
		return ErrNotFound
	}

	return err
}
