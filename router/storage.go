package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) craeteBucket(ctx *gin.Context) {
	
	ctx.JSON(http.StatusOK, "created")
}
