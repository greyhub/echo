package handler

import (
	"net/http"
	"strconv"

	"github.com/greyhub/echo/cookbook/twitter/model"
	"github.com/labstack/echo/v4"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func (h *Handler) CreatePost(c echo.Context) (err error) {
	u := &model.User{
		ID: bson.ObjectIdHex(userIDFromToken(c)),
	}
	p := &model.Post{
		ID:   bson.NewObjectId(),
		From: u.ID.Hex(),
	}
	if err = c.Bind(p); err != nil {
		return
	}

	// Validation
	if p.To == "" || p.Message == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Invalid to or message fields"}
	}

	// Find user from database
	db := h.DB.Clone()
	defer db.Close()
	err = db.DB("twitter").C("users").FindId(u.ID).One(u)
	if err != nil {
		if err == mgo.ErrNotFound {
			return echo.ErrNotFound
		}
		return
	}

	// Save post in database
	if err = db.DB("twitter").C("posts").Insert(p); err != nil {
		return
	}

	return c.JSON(http.StatusCreated, p)
}

func (h *Handler) FetchPost(c echo.Context) (err error) {
	userID := userIDFromToken(c)
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	// Defaults
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 100
	}

	// Retrieve posts from database
	posts := []*model.Post{}
	db := h.DB.Clone()
	err = db.DB("twitter").C("posts").
		Find(bson.M{"to": userID}).
		Skip((page - 1) * limit).
		Limit(limit).
		All(&posts)
	if err != nil {
		return
	}
	defer db.Close()

	return c.JSON(http.StatusOK, posts)
}
