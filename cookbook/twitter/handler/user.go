package handler

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/greyhub/echo/cookbook/twitter/model"
	"github.com/labstack/echo/v4"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func (h *Handler) Signup(c echo.Context) (err error) {
	// Bind
	u := &model.User{ID: bson.NewObjectId()}
	if err = c.Bind(u); err != nil {
		return
	}

	// Validate
	if u.Email == "" || u.Password == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Invalid email or password"}
	}

	// Save user
	db := h.DB.Clone()
	defer db.Close()
	if err = db.DB("twitter").C("users").Insert(u); err != nil {
		return
	}

	return c.JSON(http.StatusCreated, u)
}

func (h *Handler) Login(c echo.Context) (err error) {
	// Bind
	u := new(model.User)
	if err = c.Bind(u); err != nil {
		return
	}

	// Find user
	db := h.DB.Clone()
	defer db.Close()
	err = db.DB("twitter").C("users").Find(bson.M{"email": u.Email, "password": u.Password}).One(u)
	if err != nil {
		if err == mgo.ErrNotFound {
			return &echo.HTTPError{Code: http.StatusUnauthorized,
				Message: "Invalid email or password"}
		}
		return
	}

	// ---- JWT ----

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = u.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token & send it as response
	u.Token, err = token.SignedString([]byte(Key))
	if err != nil {
		return err
	}

	u.Password = "" // Important: Don't send password
	return c.JSON(http.StatusOK, u)
}

func (h *Handler) Follow(c echo.Context) (err error) {
	userID := userIDFromToken(c)
	id := c.Param("id")

	// Add a follower to user
	db := h.DB.Clone()
	defer db.Close()
	err = db.DB("twitter").C("users").
		UpdateId(bson.ObjectIdHex(id),
			bson.M{"$addToSet": bson.M{"followers": userID}})
	if err != nil {
		if err == mgo.ErrNotFound {
			return echo.ErrNotFound
		}
	}
	return
}

func userIDFromToken(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims["id"].(string)
}
