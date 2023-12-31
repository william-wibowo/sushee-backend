package utils

import (
	"net/http"
	"reflect"

	"sushee-backend/dto"
	"sushee-backend/httperror"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func ResponseSuccessJSONData(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func ResponseSuccessJSON(c *gin.Context, data dto.ResponseStruct) {
	c.JSON(http.StatusOK, data)
}

func ResponseSuccessNoContent(c *gin.Context) {
	c.Status(http.StatusOK)
}

func ResponseErrorJSON(c *gin.Context, err any) {
	log.Error().Msgf("Internal Error: %v", err)

	if appErr, isAppError := err.(httperror.AppError); isAppError {
		c.AbortWithStatusJSON(appErr.StatusCode, appErr)
		return
	}

	if reflect.TypeOf(err).String() == "error" {
		otherError := err.(error)
		serverErr := httperror.InternalServerError(otherError.Error())
		c.AbortWithStatusJSON(serverErr.StatusCode, serverErr)
		return
	}

	if reflect.TypeOf(err).String() == "string" {
		serverErr := httperror.InternalServerError(err.(string))
		c.AbortWithStatusJSON(serverErr.StatusCode, serverErr)
		return
	}

	serverErr := httperror.InternalServerError("Internal Error")
	c.AbortWithStatusJSON(serverErr.StatusCode, serverErr)
}

func AbortWithError(c *gin.Context, err httperror.AppError) {
	c.AbortWithStatusJSON(err.StatusCode, err)
}
