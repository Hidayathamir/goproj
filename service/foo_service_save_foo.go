package service

import (
	"context"
	"fmt"
)

func (s *FooServiceImpl) SaveFoo(ctx context.Context, req SaveFooReq) (SaveFooRes, error) {
	err := s.fooRepository.Save(req.Data)
	if err != nil {
		return SaveFooRes{}, fmt.Errorf("s.fooRepository.Save: %w", err)
	}
	return SaveFooRes{Success: true}, nil
}
