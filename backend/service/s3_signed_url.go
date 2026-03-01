package service

import (
	"context"
	"time"

	"vita-track-ai/config"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func GenerateSignedURL(bucket string, key string) (string, error) {

	presignClient := s3.NewPresignClient(config.S3Client)

	req, err := presignClient.PresignGetObject(context.TODO(),
		&s3.GetObjectInput{
			Bucket: &bucket,
			Key:    &key,
		},
		s3.WithPresignExpires(5*time.Minute),
	)

	if err != nil {
		return "", err
	}

	return req.URL, nil
}
