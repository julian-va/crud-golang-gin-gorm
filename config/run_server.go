package config

import (
	"crud-golang-gin-gorm/helper"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func RunServer() {
	log.Info().Msg("Started server!")

	routes := gin.Default()
	routes.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "FUNCIONA SIII")
	})
	server := &http.Server{
		Addr:    ":8001",
		Handler: routes,
	}
	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
