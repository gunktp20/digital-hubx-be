package bucket

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gunktp20/digital-hubx-be/pkg/config"
	"github.com/gunktp20/digital-hubx-be/pkg/constant"
)

type BucketClientService interface {
	DownloadFile(key string) ([]byte, error)
	GetSignedURL(key string, expires time.Duration) (string, error)
	UploadFile(key string, body []byte, contentType string) error
}

// BucketClient ตัวช่วยในการเชื่อมต่อ Bucket
type BucketClient struct {
	Client *s3.S3
	Bucket string
}

// NewBucketClient สร้าง Client ใหม่
func NewBucketClient(pctx context.Context, conf *config.Config) BucketClientService {

	_, cancel := context.WithTimeout(pctx, 20*time.Second)
	defer cancel()

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(conf.Bucket.BucketRegion),
		Credentials: credentials.NewStaticCredentials(
			conf.Bucket.AccessKey,       // Access Key
			conf.Bucket.SecretAccessKey, // Secret Key
			"",                          // Token (ถ้าไม่มี ให้ใส่เป็นค่าว่าง)
		),
	})
	if err != nil {
		log.Fatalf("failed to create bucket session: %v", err)
	}

	fmt.Println(constant.Green + "Bucket connection successful" + constant.Reset)

	return &BucketClient{
		Client: s3.New(sess),
		Bucket: conf.Bucket.BucketName,
	}
}
