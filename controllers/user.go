package controllers

import (
	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/mhdiiilham/ginorm/db"
	h "github.com/mhdiiilham/ginorm/helpers"
	m "github.com/mhdiiilham/ginorm/models"
)

// CreateUser ...
func CreateUser(c *gin.Context) {
	var body m.UserSignUpInput

	// Validate user's input
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"errors": err.Error()})
		return
	}

	user := m.User{
		FirstName:      body.FirstName,
		LastName:       body.LastName,
		Email:          body.Email,
		PasswordHashed: h.HashPassword([]byte(body.Password)),
	}

	creatingUser := db.MySQL().Save(&user)

	if creatingUser.Error != nil {
		c.JSON(500, gin.H{"errors": creatingUser.Error.Error()})
		return
	}

	token, err := h.CreateJWTToken(user.ID, user.Email)

	if err != nil {
		log.Warn("Error at creating token after signup. Error: ", err.Error())
		c.JSON(500, gin.H{"errors": "Internal Server Error"})
		return
	}

	c.JSON(201, gin.H{
		"message": "Sign Up Success",
		"email":   user.Email,
		"token":   token,
	})

}

// Login ...
func Login(c *gin.Context) {
	var body m.UserLoginInput
	var user m.User

	// Validate user's Input
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"errors": err.Error()})
		return
	}

	// Find user with given email
	find := db.MySQL().Where("email = ?", body.Email).Find(&user)

	if find.Error != nil {
		c.JSON(400, gin.H{"errors": "Email / Password is Wrong!"})
		return
	}

	isValid := h.ComparePassword(user.PasswordHashed, []byte(body.Password))
	if !isValid {
		c.JSON(400, gin.H{"errors": "Email / Password is Wrong!"})
		return
	}

	token, err := h.CreateJWTToken(user.ID, user.Email)
	if err != nil {
		log.Warn("Error at creating token after signup. Error: ", err.Error())
		c.JSON(500, gin.H{"errors": "Internal Server Error"})
		return
	}

	c.JSON(200, gin.H{
		"message": "Sign In Success",
		"email":   user.Email,
		"token":   token,
	})

}
