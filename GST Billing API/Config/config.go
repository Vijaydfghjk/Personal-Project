package config

import (
	models "GST_billing_api/Models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect_DB() {
	var err error
	dsn := "root:Vijay@123@tcp(localhost:3306)/gst_billing_db?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	DB.AutoMigrate(&models.Product{}, &models.Bill{})

}
