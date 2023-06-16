package store

import (
	"first-api/docker/db"
	"first-api/models"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type MobileResourceStore interface {
	GetAllMobiles() ([]models.Mobile, error)
	GetMobileByID(id string) (*models.Mobile, error)
	CreateMobile(*models.Mobile) error
	UpdateMobile(id string) error
	DeleteMobile(id string) (bool, error)
}

type Store struct {
	DB *gorm.DB
}

var mobileStore *Store

func NewStore() *Store {
	db, err := db.ReturnDB()
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
	return &Store{
		DB: db,
	}
}

func GetStore() *Store {
	if mobileStore == nil {
		mobileStore = NewStore()
	}
	return mobileStore
}

// GetAllMobiles Fetch all user data

func (s *Store) GetAllMobiles() (result []models.Mobile, err error) {
	mobile := []models.Mobile{}

	if err = s.DB.Find(&mobile).Error; err != nil {
		return nil, err
	}

	result = mobile
	return result, nil
}

// CreateMobile ... Insert New data
func (s *Store) CreateMobile(user *models.Mobile) (err error) {
	if err = s.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// GetMobileByID ... Fetch only one user by Id
func (s *Store) GetMobileByID(id string) (result *models.Mobile, err error) {
	mobile := &models.Mobile{}

	if err = s.DB.Where("id = ?", id).First(mobile).Error; err != nil {
		return nil, err
	}

	result = mobile
	return result, nil
}

// UpdateMobile ... Update user

func (s *Store) UpdateMobile(id string) (result *models.Mobile, err error) {
	mobile := &models.Mobile{}
	if err = s.DB.Model(mobile).Where("id = ?", id).Updates(mobile).Error; err != nil {
		return nil, err
	}

	result = mobile
	return result, nil
}

// DeleteMobile ... Delete user

func (s *Store) DeleteMobile(id string) (success bool, err error) {
	user := &models.Mobile{}

	if err = s.DB.Where("id = ?", id).First(user).Error; err != nil {
		return false, err
	}

	if err = s.DB.Delete(user).Error; err != nil {
		return false, err
	}

	return true, nil
}
