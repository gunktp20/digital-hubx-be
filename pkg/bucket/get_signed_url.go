package bucket

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func (b *BucketClient) GetSignedURL(key string, expires time.Duration) (string, error) {
	req, _ := b.Client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(b.Bucket),
		Key:    aws.String(key),
	})

	// สร้าง URL พร้อมกำหนดเวลาหมดอายุ
	signedURL, err := req.Presign(expires)
	if err != nil {
		return "", fmt.Errorf("failed to generate signed URL: %v", err)
	}

	return signedURL, nil
}
