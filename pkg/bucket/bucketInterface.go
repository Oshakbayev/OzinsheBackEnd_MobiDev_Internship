package bucket

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
)

type BucketInterface interface {
	MovieBucketServiceInterface
}

type Bucket struct {
	log          *log.Logger
	bucketClient *s3.Client
}

func CreateBucket(log *log.Logger, bucketClient *s3.Client) BucketInterface {
	return &Bucket{log: log, bucketClient: bucketClient}
}
