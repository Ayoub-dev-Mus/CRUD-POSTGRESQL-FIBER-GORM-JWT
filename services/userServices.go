package services

import (
	"FIVERR/database"
	"FIVERR/models"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers(c *fiber.Ctx) {
	db := database.DBConn
	var users []models.User
	db.Find(&users)
	c.JSON(users)
}

func GetUser(c *fiber.Ctx) {
	db := database.DBConn
	id := c.Params("Id")

	var user []models.User
	db.Find(&user, id)
	c.JSON(user)
}

func NewUser(c *fiber.Ctx) {
	db := database.DBConn
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&user)
	c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) {
	user := new(models.User)
	id := c.Params("Id")
	db := database.DBConn
	if err := c.BodyParser(user); err != nil {
		c.Status(503).SendString(err.Error())
	}
	db.Where("Id = ?", id).Updates(&user)
	c.Status(200).JSON(user)
}

func DeleteUser(c *fiber.Ctx) {
	id := c.Params("Id")
	db := database.DBConn

	var user models.User
	db.First(&user, id)
	if user.First_name == "" {
		c.Status(500).Send("No user Found with ID")
		return
	}
	db.Delete(&user)
	c.Send("User Successfully deleted")
}

const SecretKey = "ThisIsMySecretKey"

func Register(c *fiber.Ctx) {
	db := database.DBConn
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 12)

	user := models.User{
		Username: data["username"],
		Email:    data["email"],
		Password: password,
	}

	db.Create(&user)

	c.JSON(&user)
}

func Login(c *fiber.Ctx) {
	db := database.DBConn
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return
	}

	var user models.User

	db.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		c.JSON(fiber.Map{
			"message": "Invalid Credentials",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		c.JSON(fiber.Map{
			"message": "Invalid Credentials",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		c.JSON(fiber.Map{
			"message": "could not login",
		})
	}

	c.JSON(fiber.Map{
		"message": "Login Successfully",
	})

	c.JSON(token)
}
