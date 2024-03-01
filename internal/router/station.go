package router

import (
	"net/http"
	request "sola-test-task/internal/dto/request/station"

	"github.com/gin-gonic/gin"
)

func (r *Router) CreateStation(ginCtx *gin.Context) {
	c := r.context(ginCtx)

	var req request.CreateStation
	if bindErr := transBindJSON(c, ginCtx, &req, r.transServ); bindErr != nil {
		sendResponse(c, ginCtx, http.StatusBadRequest, nil, bindErr)
		return
	}

	resp, httpErr := r.stCont.CreateStation(c, &req)
	if httpErr != nil {
		sendResponse(c, ginCtx, httpErr.Code(), nil, httpErr)
		return
	}

	sendResponse(c, ginCtx, http.StatusCreated, resp, nil)
}
