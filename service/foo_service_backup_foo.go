package service

import (
	"context"
	"fmt"
	"strconv"
)

func (s *FooServiceImpl) BackupFoo(ctx context.Context, req BackupFooReq) (BackupFooRes, error) {
	data := s.fooRepository.FindByID(req.ID)
	key := s.buildS3Key(req.ID)
	err := s.s3Client.Upload("dummy-bucket", key, data)
	if err != nil {
		return BackupFooRes{}, fmt.Errorf("s.s3Client.Upload: %w", err)
	}
	return BackupFooRes{Success: true}, nil
}

func (s *FooServiceImpl) buildS3Key(id int) string {
	return "foo-backup-" + strconv.Itoa(id)
}
