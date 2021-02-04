package handler

import (
	"net/http"

	rest "github.com/SagirovVitaliy/rest_api"
	"github.com/gin-gonic/gin"
)


func (h *Handler) createList(c *gin.Context) {
	id, _ := c.Get(userCtx)
	if !ok {
		NewErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}

	var input rest.List
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	
}

func (h *Handler) getAllList(c *gin.Context) {
	
}

func (h *Handler) getListByID(c *gin.Context) {
	
}

func (h *Handler) updateList(c *gin.Context) {

}

func (h *Handler) deleteList(c *gin.Context) {

}