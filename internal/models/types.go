package models

import "time"

type DbKey struct {
	DbConnection struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		Dbname   string `json:"dbname"`
	}
}

type Config struct {
	LocalHost struct {
		Host string `json:"host"`
		Port string `json:"port"`
	}
}

type UserInfo struct {
	Id        int    `gorm:"column:id"` 	    //TODO: gorm: column надо делать так несработает
	Name      string `gorm:"column:name"`
	Icon      string `gorm:"column:icon"`
	UpdatedAt time.Time `gorm: "column: updated at"`
	Active    bool   `gorm:"column:active"`		//Done 
	Info_Type string `gorm:"cloumn:info_type"`
	Sort      int    `gorm:"column:sort"`
}

func GetTableName()(string){
	tableName := "user_info"
	return tableName
}

type Pagination struct {
	Limit      int         `json:"limit,omitempty;query:limit"`
	Page       int         `json:"page,omitempty;query:page"`
	TotalPages int64       `json:"total_pages"`
	Records    interface{} `json:"records"`
}
