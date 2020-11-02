package main

import (
	"log"
	"net/http"
	"testinguser/user"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	router := gin.Default()
	router.GET("/handler", handler)
	router.Run()
}

func handler(c *gin.Context) {
	dsn := "root:@tcp(127.0.0.1:3306)/belajar-adonis?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//cek error variable
	if err != nil {
		log.Fatal(err.Error())
	}
	var users []user.User
	db.Find(&users)

	c.JSON(http.StatusOK, users)

}
