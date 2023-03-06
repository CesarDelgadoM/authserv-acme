package handler

import (
	"time"

	"github.com/CesarDelgadoM/authserv-acme/database/memory"
	"github.com/CesarDelgadoM/authserv-acme/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey []byte = []byte("authserv_jwt_secret_key")

type AuthHandler struct {
	app *fiber.App
	db  *memory.MemoryDB
}

func NewAuthHandler(app *fiber.App, db *memory.MemoryDB) *AuthHandler {

	ah := &AuthHandler{
		app: app,
		db:  db,
	}
	ah.Router()

	return ah
}

func (ah *AuthHandler) Router() {

	ah.app.Post("/api/auth/signin", ah.SignIn)
	ah.app.Patch("/api/auth/signout/:user", ah.Logout)
}

func (ah *AuthHandler) SignIn(ctx *fiber.Ctx) error {

	var credentials *model.Credentiales

	err := ctx.BodyParser(&credentials)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	user := credentials.User
	password := credentials.Password

	// Validate credentials user
	pass, exist := ah.db.Find(user)
	if !exist || pass != password {
		return ctx.Status(fiber.StatusUnauthorized).JSON(
			model.NewResponse("user or password incorrect", fiber.StatusUnauthorized))
	}

	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &model.Claims{
		User: user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	ctx.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   signedToken,
		Expires: expirationTime,
	})

	ctx.Status(fiber.StatusOK).JSON(
		model.NewResponse("welcome user: "+user, fiber.StatusOK))

	return nil
}

func (ah *AuthHandler) Logout(ctx *fiber.Ctx) error {

	user := ctx.Params("user")
	_, exist := ah.db.Find(user)
	if !exist {
		return ctx.Status(fiber.StatusUnauthorized).JSON(
			model.NewResponse("user or password incorrect", fiber.StatusUnauthorized))
	}

	key := ctx.Cookies("token")
	if key == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(
			model.NewResponse("expired token, restart session", fiber.StatusUnauthorized))
	}

	claims := &model.Claims{}

	tkn, _ := jwt.ParseWithClaims(key, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if !tkn.Valid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(
			model.NewResponse("token not valid", fiber.StatusUnauthorized))
	}

	if user != claims.User {
		return ctx.Status(fiber.StatusUnauthorized).JSON(
			model.NewResponse("user incorrect", fiber.StatusUnauthorized))
	}

	ctx.Cookie(&fiber.Cookie{
		Name:    "token",
		Expires: time.Now(),
	})

	ctx.Status(fiber.StatusOK).JSON(
		model.NewResponse("signout user: "+user, fiber.StatusOK))

	return nil
}
