package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"solafune-be/middleware"
	"solafune-be/utils"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err.Error())
	}

	r := gin.Default()

	r.Use(middleware.CORS())

	r.GET("/", func(c *gin.Context) {
		utils.HttpRespSuccess(c, http.StatusOK, os.Getenv("MESSAGE"), nil)
		return
	})

	r.POST("/upload", func(c *gin.Context) {
		photo, err := c.FormFile("photo")
		if err != nil {
			utils.HttpRespFailed(c, http.StatusUnprocessableEntity, err.Error())
			return
		}

		newFileName := utils.GenerateFileName(filepath.Ext(photo.Filename))

		savePath := filepath.Join("assets", newFileName)

		if err := c.SaveUploadedFile(photo, savePath); err != nil {
			utils.HttpRespFailed(c, http.StatusInternalServerError, err.Error())
			return
		}

		utils.HttpRespSuccess(c, http.StatusOK, "File uploaded successfully", nil)
	})

	if err := r.Run(fmt.Sprintf(":%s", os.Getenv("PORT"))); err != nil {
		panic(err.Error())
	}
}
