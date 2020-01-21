package controllers

import (
	"net/http"
)

type Buy struct {
}

func (x *Buy) Get(c *gin.Context) {
	return http.StatusMethodNotAllowed, nil
}

func (req *Buy) Create(c *gin.Context) {
	return http.StatusNotImplemented, nil
}
