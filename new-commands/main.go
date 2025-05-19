package newcommands

import (
	"fmt"

	"github.com/ElianDev55/goapi/new-commands/seed"
)

func Seed () error {
	content := `
	
	package db

import (
	"fmt"
	"log"

	"github.com/ElianDev55/service-sign-language/internal/user"
	"github.com/ElianDev55/service-sign-language/tools/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)




func ConnectionDb() (*gorm.DB,error) {
	envDb, errEnvDb := config.LoadEnvDb()

	if errEnvDb != nil {
		fmt.Println(errEnvDb)
	}

	dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
        envDb.DatabaseHost,
        envDb.DatabaseUser,
        envDb.DatabasePassword,
        envDb.DatabaseName,
        envDb.DatabasePort,
    )

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if envDb.DatabaseDebug == "TRUE" {
		db = db.Debug()
	}

	if envDb.DatabaseAutoMigrate == "TRUE"{
		
		errMigate := db.AutoMigrate(&user.User{})
		
		if errMigate != nil{
			log.Println(errMigate)
		}

	}

	return db, nil

}
	`
	err := seed.Create(content)

	if err != nil {
		fmt.Println("Error creating seed file:", err)
		return err
	}

	return nil

}
