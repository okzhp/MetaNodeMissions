package config

import (
	"auctionBackEnd/model"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	fmt.Println("开始初始化数据库🤖...")

	err := godotenv.Load()
	if err != nil {
		log.Fatalln("读取.env文件失败❌")
	}

	// 从环境变量获取MySQL连接配置
	dbHost := GetEnv("DB_HOST", "")
	dbPort := GetEnv("DB_PORT", "")
	dbUser := GetEnv("DB_USERNAME", "")
	dbPassword := GetEnv("DB_PASSWORD", "")
	dbName := GetEnv("DB_NAME", "")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("连接MySQL失败❌")
	}
	fmt.Println("连接MySQL成功✌️")

	fmt.Println("开始数据表迁移🤖...")
	DB.AutoMigrate(&model.Auction{}, &model.Bid{})
	fmt.Println("数据表迁移成功✌️")

	fmt.Println("数据库初始化成功✌️")

	return DB
}

func GetEnv(key string, defaultValue string) string {
	if value := os.Getenv(key); len(value) != 0 {
		return value
	}
	return defaultValue
}
