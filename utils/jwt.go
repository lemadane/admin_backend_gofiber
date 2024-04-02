package utils

import (
	"strconv"
	"time"

	"github.com/lemadane/admin_backend_gofiber/db"
	"github.com/lemadane/admin_backend_gofiber/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

const SecretKey = "secret"

// GenerateJWT generates a JSON Web Token (JWT) with the specified issuer.
// It returns the generated token as a string and any error encountered during the process.
// The token is signed using the HS256 signing method and includes the standard claims.
// The issuer parameter specifies the issuer of the token.
func GenerateJWT(issuer string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    issuer,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})
	return claims.SignedString([]byte(SecretKey))
}

// UpdatePassword updates the password of a user.
// It expects a JSON object containing the new password and password confirmation in the request body.
// If the passwords do not match, it returns a JSON response with a status code of 400 and an error message.
// It retrieves the user ID from the JWT cookie and updates the password for the corresponding user in the database.
// Finally, it returns a JSON response with the updated user object.
func UpdatePassword(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "passwords do not match",
		})
	}
	cookie := c.Cookies("jwt")
	id, _ := ParseJwt(cookie)
	userId, _ := strconv.Atoi(*id)
	user := models.User{
		Id: uint(userId),
	}
	user.SetPassword(data["password"])
	db.Session().Model(&user).Updates(user)
	return c.JSON(user)
}

// ParseJwt parses the given JWT token and returns the issuer claim value.
// It takes a cookie string as input and returns the issuer claim value as a string,
// along with any error encountered during parsing.
func ParseJwt(cookie string) (*string, error) {
	token, err := jwt.ParseWithClaims(
		cookie,
		&jwt.StandardClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})

	if err != nil || !token.Valid {
		return nil, err
	}
	claims := token.Claims.(*jwt.StandardClaims)
	return &claims.Issuer, nil
}
