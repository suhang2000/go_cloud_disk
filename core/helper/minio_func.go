package helper

import (
	"errors"
	"github.com/minio/minio-go"
	"go_cloud_disk/core/define"
	"mime/multipart"
	"path"
)

func UploadFile(client *minio.Client, file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	filename := UUID() + path.Ext(fileHeader.Filename)
	_, err := client.PutObject(define.BucketName, filename, file, fileHeader.Size, minio.PutObjectOptions{})
	if err != nil {
		return "", errors.New("failed to upload file")
	}
	return define.BucketName + "/" + filename, nil
}
