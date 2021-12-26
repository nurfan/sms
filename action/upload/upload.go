package action

import (
	"context"
	"io"
	"log"
	"math/rand"
	"mime/multipart"
	"os"
	"time"

	"github.com/nurfan/sms/model"
	"github.com/nurfan/sms/util/errors"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type Upload struct {
	e errors.UniError
}

func (u *Upload) Handle(ctx context.Context, file *multipart.FileHeader) (model.Attachment, *errors.UniError) {
	var result model.Attachment
	var formatFile string

	// check format file
	switch file.Header.Get("Content-Type") {
	case "image/png":
		formatFile = ".png"
	case "application/pdf":
		formatFile = ".pdf"
	default:
		return result, u.e.BadRequest("document format not allow for upload")
	}

	// reading file in request
	src, err := file.Open()
	if err != nil {
		log.Println("error open file : ", err)
		return result, u.e.SystemError(err)
	}
	defer src.Close()

	// setup destination
	currentData := time.Now().Format("2006/01/02")
	if _, err := os.Stat("./storage/" + currentData); os.IsNotExist(err) {
		err := os.MkdirAll("./storage/"+currentData, 0777)
		if err != nil {
			log.Println("error create directory", err)
		}
	}

	// create file
	randomStr := StringWithCharset()
	path := "./storage/" + currentData + "/" + randomStr + formatFile
	dst, err := os.Create(path)
	if err != nil {
		log.Println("error create file : ", err)
		return result, u.e.SystemError(err)
	}
	defer dst.Close()

	// copy resource
	if _, err = io.Copy(dst, src); err != nil {
		log.Println("error copy file : ", err)
		return result, u.e.SystemError(err)
	}

	// mapping response
	result.Path = os.Getenv("FILE_SERVER") + "/storage/" + currentData + "/" + randomStr + formatFile
	result.Type = file.Header.Get("Content-Type")

	return result, nil
}

func StringWithCharset() string {
	var charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, 17)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func NewUpload() *Upload {
	return &Upload{}
}
