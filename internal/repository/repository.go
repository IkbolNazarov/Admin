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

	conn := db.DataB.Table(models.GetTableName())
	tx := conn.Create(data)
	if tx.Error != nil {
		log.Println(tx.Error)
		return tx.Error
	}
	return nil
}

func (r *Repository) GetData(pagination *models.Pagination, l int64) ([]*models.UserInfo, error) {
	var UserInfo []*models.UserInfo
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := db.DataB.Limit(pagination.Limit).Offset(offset)
	result := queryBuider.Find(&UserInfo) //TODO: 2 раза моделька и еще table отдельно указанно
	if result.Error != nil {              // DONE
		msg := result.Error
		return nil, msg
	}
	amount := db.DataB.Table(models.GetTableName()).Count(&l)
	if amount.Error != nil {
		return nil, amount.Error
	}
	return UserInfo, nil
}

func (r *Repository) CountRows(l int64) error {
	amount := db.DataB.Table(models.GetTableName()).Count(&l)
	if amount.Error != nil {
		return amount.Error
	}
	return nil
}

func (r *Repository) UpdateData(userData *models.UserInfo) error {
	tx := db.DataB.Table(models.GetTableName()).Model(&models.UserInfo{}).Where("id = ?", userData.Id).Updates(models.UserInfo{Name: userData.Name, Icon: userData.Icon, Sort: userData.Sort})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (r *Repository) DeleteData(id int) error {
	var data *models.UserInfo
	query := db.DataB.Table(models.GetTableName()).Where("id =?", id).Delete(&data) //TODO: уверен что работает?
	if query.Error != nil {
		return query.Error //DONE
	}
	return nil
}
