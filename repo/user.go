package repo

import (
	"crypto/md5"
	"encoding/hex"
	"fiber-gorm/database"
	"fiber-gorm/models"
	"fiber-gorm/utils"
)

func Login(loginReq *models.LoginRequest) (*models.User, error) {
	var model *models.User
	hash := md5.Sum([]byte(loginReq.Password))
	encode := hex.EncodeToString(hash[:])
	err := database.DBConn.Where("user_name = ?", loginReq.UserName).
		Where("password = ?", encode).
		First(&model).Error

	if err != nil {
		return nil, err
	}

	return model, nil
}

func GetAllUsers(pagination *models.Pagination) (*models.Pagination, error) {
	var model []*models.User
	database.DBConn.Scopes(utils.Paginate(model, pagination, database.DBConn)).Find(&model)
	pagination.Rows = model

	return pagination, nil
}
