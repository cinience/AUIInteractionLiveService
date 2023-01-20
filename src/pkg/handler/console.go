package handler

import (
	"ApsaraLive/pkg/models"
	"ApsaraLive/pkg/utils"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-chi/render"
	"net/http"
	"strings"
)

type ManagerHandler struct {
}

func NewManagerHandler() *ManagerHandler {
	return &ManagerHandler{}
}

type loginRequest struct {
	Username string `form:"username" json:"accountId" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (h *ManagerHandler) Login(w http.ResponseWriter, r *http.Request) {
	var in loginRequest
	b := binding.Default(r.Method, strings.Split(r.Header.Get("Content-Type"), ";")[0])
	err := b.Bind(r, &in)
	var rst interface{}
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		rst = &models.Status{Code: http.StatusBadRequest, Message: err.Error()}
		render.DefaultResponder(w, r, rst)
		return
	}
	m := make(map[string]interface{})
	m["scene"] = "live"
	rst = m
	render.DefaultResponder(w, r, rst)
}

type getUserInfoRequest struct {
	Username string `form:"userId" json:"userId" binding:"required"`
}

func (h *ManagerHandler) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	var in getUserInfoRequest
	b := binding.Default(r.Method, strings.Split(r.Header.Get("Content-Type"), ";")[0])
	err := b.Bind(r, &in)
	var rst interface{}
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		rst = &models.Status{Code: http.StatusBadRequest, Message: err.Error()}
		render.DefaultResponder(w, r, rst)
		return
	}
	m := make(map[string]interface{})

	m["appId"] = "admin"
	m["appKey"] = "admin"
	m["scene"] = "live"

	rst = m
	w.Header().Set("x-app-id", "admin")
	render.DefaultResponder(w, r, rst)
}

func (h *ManagerHandler) Action(w http.ResponseWriter, r *http.Request) {
	var err error
	var rst interface{}
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		rst = &models.Status{Code: http.StatusBadRequest, Message: err.Error()}
		render.DefaultResponder(w, r, rst)
		return
	}
	m := make(map[string]interface{})

	m["appId"] = "admin"
	m["appKey"] = "admin"
	m["scene"] = "live"

	rst = m
	w.Header().Set("x-app-id", "admin")

	utils.DefaultResponder(w, r, rst)
}
