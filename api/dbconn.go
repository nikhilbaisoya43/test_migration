package api

import (
	"04_pro/model"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func INitializeServer() {
	db := INitdb()
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	router.GET("/health-checkup", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Server is running...",
		})
	})
	appPort := os.Getenv("APP_PORT")
	router.Run(fmt.Sprintf(":%s", appPort))
}

func INitdb() *gorm.DB {
	dbConnectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", os.Getenv("DBHOST"), os.Getenv("DBPORT"), os.Getenv("DBUSER"), os.Getenv("DBNAME"), os.Getenv("DBPASSWORD"))
	dbInstance, err := gorm.Open("postgres", dbConnectionString)

	if err != nil {
		panic(fmt.Sprintf("Failed to connect database for %s, %s", dbConnectionString, err))
	} else {
		fmt.Println("Successfully connect to database")
	}
	//defer db.Close()

	if os.Getenv("IS_DEBUG") == "1" {
		dbInstance.LogMode(true)
	}

	// Run Migrations
	eeee := dbInstance.AutoMigrate(&model.App{}).Error
	if eeee != nil {
		log.Println(eeee.Error())
	}

	fmt.Println("-----------------------_______________________")
	errors := dbInstance.AutoMigrate(&model.User{}).Error
	if errors != nil {
		log.Println(errors.Error())
	}
	fmt.Println("-----------------------_______________________")

	return dbInstance

}
