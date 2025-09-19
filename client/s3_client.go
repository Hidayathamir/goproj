package client

type S3Client interface {
	Upload(bucket, key, data string) error
	Download(bucket, key string) string
	DeleteObject(bucket, key string) bool
}

// compile-time check
var _ S3Client = (*S3ClientImpl)(nil)

type S3ClientImpl struct {
}

func NewS3Client() *S3ClientImpl {
	return &S3ClientImpl{}
}

func (c *S3ClientImpl) Upload(bucket, key, data string) error {
	return nil
}

func (c *S3ClientImpl) Download(bucket, key string) string {
	return "dummy data"
}

func (c *S3ClientImpl) DeleteObject(bucket, key string) bool {
	return true
}
