package controllers

import (
	"Project/dao"
	"Project/database"
	"Project/dto"
	"Project/models"
	"net/http"

	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "secret "

func RegisterUser(c *gin.Context) {
	var user map[string]string
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(user["password"]), 14)
	saveuser := models.User{
		Name:     user["name"],
		Email:    user["email"],
		Password: password,
	}

	database.DB.Create(&saveuser)
	c.JSON(200, gin.H{"msg": "user is created"})
}

func Login(c *gin.Context) {
	var data map[string]string
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	var user models.User
	database.DB.Where("email = ?", data["email"]).First(&user)
	if user.Id == 0 {
		c.String(http.StatusNotFound, "User not found: "+data["email"])
		return
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.String(http.StatusBadRequest, "Incorrect Password")
		return
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.String(http.StatusInternalServerError, "Could not login")
		return
	}
	cookie, _ := c.Cookie("jwt")
	c.SetCookie(
		"jwt",
		token,
		60*60*24,
		"/",
		"localhost",
		false,
		true,
	)

	c.Cookie(cookie)

	c.JSON(200, gin.H{"token": token}) // TODO: pass
}

func User(c *gin.Context) {
	cookie, _ := c.Cookie("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.String(http.StatusUnauthorized, "unauthorized")
		return
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User
	database.DB.Where("id=?", claims.Issuer).First(&user)
	c.JSON(200, gin.H{"user": user})
}

func Logout(c *gin.Context) {
	cookie, _ := c.Cookie("jwt")
	c.SetCookie(
		"jwt",
		"",
		-60*60*24,
		"/",
		"localhost",
		false,
		true,
	)
	c.Cookie(cookie)
	c.JSON(200, gin.H{"message": "Logout Succcess ful"})
}

func CreatePoll(c *gin.Context) {
	var data dao.CreatePoll
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	savequestion := models.Polls{
		Question: data.Question,
	}
	database.DB.Create(&savequestion)
	id := savequestion.Id
	options := data.Options
	for _, val := range options {
		saveOption := models.Options{
			Option: val,
			PollId: id,
			Votes:  0,
		}

		database.DB.Create(&saveOption)
	}
	c.JSON(200, gin.H{"msg": "success fully added"})
}

func GetAllPolls(c *gin.Context) {

	polls := []models.Polls{}
	database.DB.Find(&polls)
	var allQuestions = []dto.Questions{}

	for _, val := range polls {
		options := []models.Options{}
		database.DB.Where("poll_id=?", val.Id).Find(&options)
		var allOptions []string
		for _, option := range options {
			allOptions = append(allOptions, option.Option)
		}
		question := dto.Questions{
			Id:       int(val.Id),
			Question: val.Question,
			Options:  allOptions,
		}

		allQuestions = append(allQuestions, question)
	}

	c.JSON(200, gin.H{"questions": &allQuestions})
}

func GetPollById(c *gin.Context) {

	var poll models.Polls

	if err := database.DB.Where("id = ?", c.Param("id")).First(&poll).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	//var allQuestions = []dto.Questions{}
	options := []models.Options{}
	database.DB.Where("poll_id=?", poll.Id).Find(&options)
	var allOptions []string
	for _, option := range options {
		allOptions = append(allOptions, option.Option)
	}
	question := dto.Questions{
		Id:       int(poll.Id),
		Question: poll.Question,
		Options:  allOptions,
	}

	c.JSON(200, gin.H{"poll: ": &question})

}
