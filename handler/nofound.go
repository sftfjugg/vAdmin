package handler

import (
	"github.com/gin-gonic/gin"
	jwt "vAdmin/pkg/jwtauth"
	"log"
	"net/http"
)

func NoFound(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	log.Printf("NoRoute claims: %#v\n", claims)
	c.JSON(http.StatusOK, gin.H{
		"code":    "404",
		"message": "not found",
	})
}
