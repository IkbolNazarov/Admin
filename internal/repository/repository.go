package repository

import (
	"admin/internal/db"
	"admin/internal/models"
	"log"

	"gorm.io/gorm"
)

type Repository struct {
	Connection *gorm.DB
}

func NewRepository(conn *gorm.DB) *Repository {
	return &Repository{Connection: conn}
}

func (r *Repository) AddData(data *models.UserInfo) error {

	conn := db.DataB.Table("user_info")
	tx := conn.Create(data)
	if tx.Error != nil {
		log.Println(tx.Error)
		return tx.Error
	}
	return nil
}

func (r *Repository) GetData(pagination *models.Pagination) ([]*models.UserInfo, error) {
	var UserInfo []*models.UserInfo
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := db.DataB.Table("user_info").Limit(pagination.Limit).Offset(offset)
	result := queryBuider.Find(&UserInfo) //TODO: 2 раза моделька и еще table отдельно указанно
	if result.Error != nil {              // убрал одну модельку, если table не указать, то вместо "user_info" будет искать "user_infos"
		msg := result.Error
		return nil, msg
	}
	return UserInfo, nil
}

func (r *Repository) CountRows(l int64) (int64, error) {
	amount := db.DataB.Table("user_info").Count(&l)
	if amount.Error != nil {
		return 0, amount.Error
	}
	return l, nil
}

func (r *Repository) UpdateData(userData *models.UserInfo) error {
	tx := db.DataB.Table("user_info").Model(&models.UserInfo{}).Where("id = ?", userData.Id).Updates(models.UserInfo{Name: userData.Name, Icon: userData.Icon, Sort: userData.Sort})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (r *Repository) DeleteData(id int) error {
	var data *models.UserInfo
	query := db.DataB.Table("user_info").Where("id =?", id).Delete(&data) //TODO: уверен что работает?
	if query.Error != nil {
		return query.Error // DONE
	}
	return nil
}
