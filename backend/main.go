package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	h := gin.New()
	corsCfg := cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodPut, http.MethodDelete, http.MethodOptions},
		ExposeHeaders:   []string{"*"},
	}

	h.Use(
		gin.Logger(),
		cors.New(corsCfg),
	)

	RegisterHandlers(h, &serverImpl{})

	srv := &http.Server{
		Addr:              ":8080",
		Handler:           h,
		ReadHeaderTimeout: 1 * time.Second,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}
}

type serverImpl struct{}

func (s *serverImpl) GetIds(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": uuid.New(),
	})
}
