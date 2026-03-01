package service

import (
	"context"
	"mime/multipart"

	"vita-track-ai/config"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func UploadToS3(file *multipart.FileHeader, key string, bucket string) error {

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	_, err = config.S3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: &bucket,
		Key:    &key,
		Body:   src,
	})

	return err
}
