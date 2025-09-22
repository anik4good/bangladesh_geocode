package config

import (
	"log"
	"os"
	"time"

	"github.com/anik4good/go-echo-apiboilerplate/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Conf models.Config
var file *os.File
var (
	GormDBConn *gorm.DB
)

func Init() {

	// Ensure logs directory exists
	if err := os.MkdirAll("logs", 0755); err != nil {
		log.Fatal("Failed to create logs directory: ", err)
	}
	//log file
	logFileName := "logs/" + time.Now().Format("2006-01-02") + ".log"
	file, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)

	// Try to load .env, log warning if missing
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Warning: .env file not loaded:", err)
	}

	log.Println(os.Getenv("CONNECTION"))

	// connectDb
	db, err := gorm.Open(mysql.Open(os.Getenv("CONNECTION")), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("connected")
	//db.(&models.User{}, &models.Queue{})
	//db.AutoMigrate(&models.User{}, &models.Queue{})

	//db.AutoMigrate(&models.User{}, &models.SendSms{}, &models.ModelHasRoles{})
	//db.AutoMigrate(&models.Division{}, &models.District{}, &models.Upazila{}, &models.Union{})
	GormDBConn = db
}
