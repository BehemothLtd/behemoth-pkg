package uploaders

import (
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/BehemothLtd/behemoth-pkg/golang/constants"
	"github.com/BehemothLtd/behemoth-pkg/golang/exceptions"
	translator "github.com/BehemothLtd/behemoth-pkg/golang/translators"
	"github.com/BehemothLtd/behemoth-pkg/golang/utils"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Uploader struct {
	Ctx            *gin.Context
	StorageService string
	UploadPath     string
}

type ClientUploader struct {
	cl         *storage.Client
	bucketName string
	uploadPath string
}

type UploadedBlob struct {
	Url string `json:"url"`
	Key string `json:"key"`
}

func (u *Uploader) Upload() ([]*UploadedBlob, error) {
	err := u.Ctx.Request.ParseMultipartForm(10 << 20)
	if err != nil {
		return nil, err
	}

	var uploadedBlobs []*UploadedBlob

	form := u.Ctx.Request.MultipartForm

	files := form.File["files[]"]

	for _, file := range files {
		if err := u.validate(file); err != nil {
			return nil, err
		}

		filename := file.Filename

		var uploadedBlob *UploadedBlob
		var uploadErr error

		switch u.StorageService {
		case "local":
			uploadedBlob, uploadErr = u.uploadLocally(file, filename)
		case "google":
			uploadedBlob, uploadErr = u.uploadToGCS(file, filename)
		default:
			return nil, fmt.Errorf("unsupported storage service: %s", u.StorageService)
		}

		if uploadErr != nil {
			return nil, uploadErr
		}

		uploadedBlobs = append(uploadedBlobs, uploadedBlob)
	}

	return uploadedBlobs, nil
}

func (u *Uploader) validate(fileHeader *multipart.FileHeader) error {
	if err := u.validateFileSize(fileHeader); err != nil {
		return err
	}

	if err := u.validateFileType(fileHeader); err != nil {
		return err
	}

	return nil
}

func (u *Uploader) validateFileSize(fileHeader *multipart.FileHeader) error {
	if fileHeader.Size > constants.FileMaxSize {
		message := translator.Translate(nil, "errValidation_maxSizeImg", constants.FileMaxSize/1024/1024)
		return exceptions.NewBadRequestError(&message)
	}
	return nil
}

func (u *Uploader) validateFileType(fileHeader *multipart.FileHeader) error {
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	buffer := make([]byte, 512)

	_, err = file.Read(buffer)
	if err != nil && err != io.EOF {
		return err
	}
	file.Seek(0, io.SeekStart)

	filetype := http.DetectContentType(buffer)
	allowedTypes := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
		"image/gif":  true,
	}

	if !allowedTypes[filetype] {
		message := translator.Translate(nil, "ValidateFileFormat")
		return exceptions.NewBadRequestError(&message)
	}

	return nil
}

func (u *Uploader) uploadLocally(file *multipart.FileHeader, fileName string) (*UploadedBlob, error) {
	blobKey := uuid.New().String()

	uploadDst := filepath.Join(".", "tmp_app", "uploads", blobKey)
	err := u.Ctx.SaveUploadedFile(file, uploadDst)
	if err != nil {
		return nil, err
	}

	return &UploadedBlob{
		Url: "http://localhost:" + utils.GetEnv("API_PORT", "3000") + "/uploads/" + blobKey,
		Key: blobKey,
	}, nil
}

func (u *Uploader) uploadToGCS(file *multipart.FileHeader, fileName string) (*UploadedBlob, error) {
	bucketName := os.Getenv("GCS_BUCKET_NAME")
	projectId := os.Getenv("GCS_PROJECT_ID")
	gcsAccountService := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")

	if bucketName == "" || projectId == "" || gcsAccountService == "" {
		return nil, errors.New("invalid Setting for Upload")
	}

	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", gcsAccountService)
	client, err := storage.NewClient(context.Background())
	if err != nil {
		return nil, err
	}

	uploadClient := &ClientUploader{
		cl:         client,
		bucketName: bucketName,
		uploadPath: u.UploadPath,
	}

	blobFile, err := file.Open()
	if err != nil {
		return nil, err
	}

	defer blobFile.Close()
	blobKey := uuid.New().String()

	err = uploadClient.UploadFile(blobFile, blobKey)
	if err != nil {
		return nil, err
	}

	filePublicUrl := fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucketName, blobKey)

	return &UploadedBlob{
		Url: filePublicUrl,
		Key: blobKey,
	}, nil

}

func (c *ClientUploader) UploadFile(file multipart.File, object string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	wc := c.cl.Bucket(c.bucketName).Object(c.uploadPath + object).NewWriter(ctx)
	defer wc.Close()

	if _, err := io.Copy(wc, file); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}
	return nil
}

func (c *ClientUploader) DeleteFile(key string) {
	ctx := context.Background()
	o := c.cl.Bucket(c.bucketName).Object(key)
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	o.Delete(ctx)
}
