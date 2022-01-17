package controllers

import (
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
)

type User struct {
	IdUser         int `gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	Cnpj           string
	Name           string
	Email          string
	Password       string
	Phone          string
	ContactEmail   string
	Address        string
	ProfilePicture string
	City           string
	RegisterDate   time.Time
}

type UserWithoutPassword struct {
	IdUser         int `gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	Cnpj           string
	Name           string
	Email          string
	Phone          string
	ContactEmail   string
	Address        string
	ProfilePicture string
	City           string
	RegisterDate   time.Time
}

func (UserWithoutPassword) TableName() string {
	return "users"
}

func GetUsers(c *gin.Context) {

	var user []User
	db.Find(&user)
	c.JSON(200, gin.H{
		"data": user,
	})
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user User
	db.Where("id_user = @idUser", sql.Named("idUser", id)).Find(&user)
	c.JSON(200, gin.H{
		"data": user,
	})
}

func PostUser(c *gin.Context) {
	var user User
	c.BindJSON(&user)
	user.RegisterDate = time.Now()
	result := db.Create(&user)
	if result.Error == nil {
		c.JSON(200, gin.H{
			"data": user,
		})
	}
}

func PutUser(c *gin.Context) {
	var user User
	c.BindJSON(&user)
	db.Save(&user)
	c.JSON(200, gin.H{
		"data": user,
	})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user User
	db.Where("id_user = @idUser", sql.Named("idUser", id)).Delete(&user)
	c.JSON(200, gin.H{
		"data": user,
	})
}
