package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

type BindFile struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
	Ns   string                `form:"ns" binding:"required"`
}

func setup() {
	archiveFile := "webapp.tar.gz"
	b64archive := "/tmp/" + archiveFile + ".base64"
	archive := uploadDir() + "/app/" + archiveFile
	err := ioutil.WriteFile(b64archive, []byte(getFiles()), os.ModePerm)
	if err != nil {
		fmt.Println(err)
		log.Fatalln("Unable to initialize webapp")
	}
	c := fmt.Sprintf("mkdir -p %s/app;base64 -d %s > %s; cd %s/app; tar -xf %s; rm %s",
		uploadDir(),
		b64archive,
		archive,
		uploadDir(),
		archiveFile,
		archiveFile,
	)
	cmd := exec.Command("/bin/bash", "-c", c)
	err = cmd.Run()
}

func main() {
	setup()
	r := gin.Default()
	if os.Getenv("DEV_MODE") == "1" {
		r.Use(cors.New(cors.Config{
			AllowOrigins:  []string{"http://localhost:8021"},
			AllowMethods:  []string{"*"},
			AllowHeaders:  []string{"*"},
			ExposeHeaders: []string{"Content-Length"},

			MaxAge: 12 * time.Hour,
		}))
	}
	r.GET("/*path", handler)
	r.POST("/upload", upload)

	err := r.Run("0.0.0.0:8020")
	if err != nil {
		log.Fatalln(err)
		return
	}
}

func uploadDir() string {
	if os.Getenv("UPLOAD_FOLDER") != "" {
		return os.Getenv("UPLOAD_FOLDER")
	}
	return "./uploads"
}

func handler(c *gin.Context) {

	tokens := strings.Split(strings.Split(c.Request.Host, ":")[0], ".")
	subdomain := "app"
	if len(tokens) > 1 {
		subdomain = strings.Join(tokens[0:len(tokens)-1], ".")
		if subdomain == "app" {
			c.AbortWithStatus(403)
			return
		}
	}
	fmt.Println("subdomain", subdomain)
	path := c.Param("path")
	if path == "/" {
		path = "index.html"
	}
	filePath := uploadDir() + "/" + subdomain + "/" + path
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		c.AbortWithStatus(404)
		return
	}
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Data(200, mime.TypeByExtension(path), content)
}

func upload(c *gin.Context) {
	var bindFile BindFile

	if err := c.ShouldBind(&bindFile); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("err: %s", err.Error()))
		return
	}

	file := bindFile.File
	folder := uploadDir() + "/" + bindFile.Ns
	_, err := os.Stat(folder)
	if os.IsNotExist(err) {
		err := os.MkdirAll(folder, os.ModePerm)
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("err: %s", err.Error()))
			return
		}
	}
	path := folder + "/" + filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, path); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "File " + file.Filename + " uploaded successfully",
	})
}
