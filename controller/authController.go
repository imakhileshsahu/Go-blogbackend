package controller

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/imakhileshsahu/blogbackend/database"
	"github.com/imakhileshsahu/blogbackend/models"
	"github.com/imakhileshsahu/blogbackend/util"
)

func validateEmail(email string) bool {
	Re := regexp.MustCompile(`[a-z0-9._%+\-]+@[a-z0-9._%+\-]+\.[a-z0-9._%+\-]`)
	return Re.MatchString(email)
}
func Register(c *fiber.Ctx) error {
	var data map[string]interface{}
	var userData models.User
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable to parse body")
	}
	//check if password less than 6 characters
	if len(data["password"].(string)) <= 6 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Password must be greater than 6 characters",
		})

	}

	if !validateEmail(strings.TrimSpace(data["email"].(string))) {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid Email Address",
		})
	}
	//check email is already exist in database
	database.DB.Where("email=?", strings.TrimSpace(data["email"].(string))).First(&userData)
	if userData.ID != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": " Email Address already exist",
		})
	}

	user := models.User{
		FirstName: data["first_name"].(string),
		LastName:  data["last_name"].(string),
		Phone:     data["phone"].(string),
		// Username:  data["username"].(string),
		// <<<<<<< HEAD
		// Email:strings.TrimSpace(data["email"].(string)),
		// =======
		Email: strings.TrimSpace(data["email"].(string)),
		// >>>>>>> f45aa6b (migrated to postgres + auth fixes)
	}
	user.SetPassword(data["password"].(string))
	err := database.DB.Create(&user)
	if err != nil {
		log.Println(err)
	}
	c.Status(200)
	return c.JSON(fiber.Map{
		"message": "Account Registration created successful",
		"user":    user,
	})
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable to parse body")
	}
	var user models.User
	database.DB.Where("email=?", data["email"]).First(&user)
	if user.ID == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "Email Address doesn't exit,kindly create an account",
		})
	}
	if err := user.ComparePassword(data["password"]); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}
	token, err := util.GenerateJwt(strconv.Itoa(int(user.ID)))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		Secure:   false,  // ❌ keep false in localhost, ✅ true in production (HTTPS)
		SameSite: "None", // 🔑 required for cross-site cookies
	}

	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "you have successfully login",
		"user":    user,
	})
}

type Claims struct {
	jwt.StandardClaims
}

func Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour), // expire immediately
		HTTPOnly: true,
		Secure:   false,  // ❌ false for localhost, ✅ true in production (HTTPS)
		SameSite: "None", // must match login cookie
	})

	return c.JSON(fiber.Map{
		"message": "Logout successful",
	})
}
