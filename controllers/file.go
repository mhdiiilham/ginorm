package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	// "github.com/mhdiiilham/ginorm/models"
	log "github.com/sirupsen/logrus"
)

// UploadSingleImage ...
func UploadSingleImage(c *gin.Context) {
	// Declare varible needed
	var imageData map[string]interface{}
	var buf = new(bytes.Buffer)
	writer := multipart.NewWriter(buf)

	// Read the image from request body
	fileHeader, err := c.FormFile("image")
	if err != nil {
		log.Info("Error on line 18 controller of file, err: ", err)
		c.JSON(500, gin.H{"errors": err.Error()})
		return
	}

	// Create image file
	imageToUpload, err := writer.CreateFormFile("image", fileHeader.Filename)
	if err != nil {
		log.Info("Error on line 28 controller of file, err: ", err)
		c.JSON(500, gin.H{"errors": "Internal Server Error"})
		return
	}

	// Get the image file and assign to `imageToUpload`
	imageFile, err := fileHeader.Open()
	if err != nil {
		log.Info("Error on line 37 controller of file, err: ", err)
		c.JSON(500, gin.H{"errors": "Internal Server Error"})
		return
	}
	io.Copy(imageToUpload, imageFile)
	writer.Close()

	//Prepare for the HTTP Request
	req, err := http.NewRequest("POST", os.Getenv("IMGUR_URI"), buf)
	if err != nil {
		log.Info("Error on line 48 controller of file, err: ", err)
		c.JSON(500, gin.H{"errors": "Internal Server Error"})
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Client-ID "+os.Getenv("IMGUR_CLIENT_ID"))

	// Do the HTTP Request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Info("Error on line 60 controller of file, err: ", err)
		c.JSON(500, gin.H{"errors": "Internal Server Error. Uploading image fail, please try again!"})
		return
	}
	defer resp.Body.Close()

	imgurResp, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal([]byte(imgurResp), &imageData); err != nil {
		log.Info("Error on line 72 controller of file, err: ", err)
		c.JSON(500, gin.H{"errors": "Internal Server Error"})
		return
	}

	c.JSON(200, gin.H{
		"message": "Image uploaded",
		"image": imageData,
	})
	
}