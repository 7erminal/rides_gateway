package functions

import (
	"mime/multipart"
	"net/http"
	"path/filepath"
	"time"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

func SaveImage(c *beego.Controller, fieldName string, file multipart.File, header multipart.FileHeader) (respCode int, filePath_ string) {
	var filePath string = ""

	defer file.Close()

	// Save the uploaded file
	fileName := filepath.Base(header.Filename)
	filePath = "/tmp/" + time.Now().Format("20060102150405") + fileName // Define your file path
	err := c.SaveToFile(fieldName, "../images/"+filePath)

	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		logs.Error("Error saving file", err)
		// c.Data["json"] = map[string]string{"error": "Failed to save the image file."}
		errorMessage := "Error: Failed to save the image file"

		return 400, errorMessage
	}

	host, _ := beego.AppConfig.String("imagesBaseUrl")
	filePath = host + filePath

	logs.Info("Full file path is ", filePath)

	return 200, filePath

}
