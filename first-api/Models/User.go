package Models

import (
	"first-api/Config"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// GetAllMobiles Fetch all user data
func GetAllMobiles(user *[]Mobile) (err error) {
	if err = Config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}

// CreateMobile ... Insert New data
func CreateMobile(user *Mobile) (err error) {
	if err = Config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// GetMobileByID ... Fetch only one user by Id
func GetMobileByID(user *Mobile, id string) (err error) {

	if err = Config.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

// UpdateMobile ... Update user
func UpdateMobile(user *Mobile, id string) (err error) {
	fmt.Println(user)
	Config.DB.Save(user)
	return nil
}

// DeleteMobile ... Delete user
func DeleteMobile(user *Mobile, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(user)
	return nil
}
