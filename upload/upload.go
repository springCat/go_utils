package upload

import (
	"github.com/gin-gonic/gin"
	"go_utils/id"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strings"
	"errors"
)

var ImageType = map[string]string{
	"image/gif":     ".gif",
	"image/jpeg":    ".jpg",
	"image/png":     ".png",
	"image/svg+xml": ".svg",
}

var TextType = map[string]string{
	".htm":  "text/html; charset=utf-8",
	".html": "text/html; charset=utf-8",
	".pdf":  "application/pdf",
	".xml":  "text/xml; charset=utf-8",
}

var (
	ErrorMediaType = errors.New(http.StatusText(http.StatusUnsupportedMediaType))
	OutOfContentLength = errors.New(http.StatusText(http.StatusRequestEntityTooLarge))
	MaxFileSize int64 = 1024 * 1024
)


func IsImage(f multipart.File, header *multipart.FileHeader) ([]byte, bool, error) {
	contentType := header.Header.Get("Content-Type")
	if ImageType[contentType] == "" {
		return nil, false, nil
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, false, err
	}
	fileType := http.DetectContentType(b)
	if ImageType[fileType] == "" {
		return nil, false, nil
	}
	return b, true, nil
}

func IsOutOfLength(c *gin.Context) bool {
	return c.Request.ContentLength > MaxFileSize
}

func NewFileName(fileName string) string {
	l := len(fileName)
	i := strings.LastIndex(fileName, ".")
	newFileName := make([]byte, (l - i + 32))
	copy(newFileName[:32], id.UUID32())
	copy(newFileName[32:], fileName[i:])
	return string(newFileName)
}

func Save(path string, fileName string, data []byte) error {
//	newFileName := NewFileName(fileName)
	return ioutil.WriteFile(path+"/"+fileName, data, 0666)
}

func SaveImage(path string,param string,c *gin.Context) (int,error) {

	if IsOutOfLength(c) {
		return http.StatusRequestEntityTooLarge,OutOfContentLength
	}

	f, header, err := c.Request.FormFile(param)

	if err != nil {
		return http.StatusInternalServerError,err
	}

	data,ok,err := IsImage(f,header)
	if err != nil {
		return http.StatusInternalServerError,err
	}

	if !ok {
		return http.StatusUnsupportedMediaType,ErrorMediaType
	}

	err = Save(path,header.Filename,data)
	if err != nil {
		return http.StatusInternalServerError,err
	}

	return 0,nil
}

