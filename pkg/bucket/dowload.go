package bucket

import (
	"bytes"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func (b *BucketClient) DownloadFile(key string) ([]byte, error) {
	result, err := b.Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(b.Bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}

	defer result.Body.Close()

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(result.Body)
	return buf.Bytes(), err
}
