package models

import (
	"context"
	"first-api/config"

	_ "github.com/go-sql-driver/mysql"
)

type MobileResourceStore interface {
	List(ctx context.Context) (*[]Mobile, error)
	Get(ctx context.Context) (*Mobile, error)
	Create(ctx context.Context) error
	Update(ctx context.Context, app *Mobile) (*Mobile, error)
	UpdateByID(ctx context.Context, id string) error
}

// GetAllMobiles Fetch all user data
//
//	func GetAllMobiles(user *[]Mobile) (err error) {
//			if err = config.DB.Find(user).Error; err != nil {
//				return err
//			}
//			return nil
//		}
func GetAllMobiles() (result *[]Mobile, err error) {
	mobile := []Mobile{}

	if err = config.DB.Find(&mobile).Error; err != nil {
		return nil, err
	}

	result = &mobile
	return result, nil
}

// CreateMobile ... Insert New data
func CreateMobile(user *Mobile) (err error) {
	if err = config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// GetMobileByID ... Fetch only one user by Id
// func GetMobileByID(user *Mobile, id string) (err error) {

// 	if err = config.DB.Where("id = ?", id).First(user).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

func GetMobileByID(id string) (result *Mobile, err error) {
	mobile := &Mobile{}

	if err = config.DB.Where("id = ?", id).First(mobile).Error; err != nil {
		return nil, err
	}

	result = mobile
	return result, nil
}

// UpdateMobile ... Update user
//
//	func UpdateMobile(user *Mobile, id string) (err error) {
//		fmt.Println(user)
//		config.DB.Save(user)
//		return nil
//	}
func UpdateMobile(id string) (result *Mobile, err error) {
	mobile := &Mobile{}
	if err = config.DB.Model(mobile).Where("id = ?", id).Updates(mobile).Error; err != nil {
		return nil, err
	}

	result = mobile
	return result, nil
}

// DeleteMobile ... Delete user
//
//	func DeleteMobile(user *Mobile, id string) (err error) {
//		config.DB.Where("id = ?", id).Delete(user)
//		return nil
//	}
func DeleteMobile(id string) (success bool, err error) {
	user := &Mobile{}

	if err = config.DB.Where("id = ?", id).First(user).Error; err != nil {
		return false, err
	}

	if err = config.DB.Delete(user).Error; err != nil {
		return false, err
	}

	return true, nil
}
