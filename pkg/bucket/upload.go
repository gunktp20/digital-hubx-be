package bucket

import (
	"bytes"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

// UploadFile อัปโหลดไฟล์ไปยัง S3
func (b *BucketClient) UploadFile(key string, body []byte, contentType string) error {
	_, err := b.Client.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(b.Bucket),
		Key:         aws.String(key),
		Body:        bytes.NewReader(body),
		ContentType: aws.String(contentType),
	})
	return err
}
