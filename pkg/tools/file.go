package tools

import (
	"bytes"
	"fmt"
	"github.com/satori/go.uuid"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

func IsFileExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsExist(err) {
		return true
	}
	return false
}

func IsImageFile(data []byte) bool {
	return strings.Contains(http.DetectContentType(data),"image/")
}

func IsPngImageFile(data []byte) bool {
	return strings.Contains(http.DetectContentType(data), "image/png")
}

func IsJpegImageFile(data []byte) bool {
	return strings.Contains(http.DetectContentType(data), "image/jpeg")
}

func IsGifImageFile(data []byte) bool {
	return strings.Contains(http.DetectContentType(data), "image/gif")
}

func FileExtension(fileheader multipart.FileHeader) string {
	filename := fileheader.Filename
	s := strings.Split(filename, ",")
	return s[len(s)-1]
}

func CreateImgFile(file multipart.File, fileheader multipart.FileHeader) error {
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return err
	}
	if !IsImageFile(buf.Bytes()) {
		return fmt.Errorf("sorry, the file type is unsupporsed.")
	}
	newFilename := fmt.Sprintf("%s.%s", uuid.Must(uuid.NewV4()), FileExtension(fileheader))
	out, err := os.OpenFile(newFilename, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer out.Close()
	if _, err = io.Copy(out, file); err != nil {
		return err
	}
	return nil
}