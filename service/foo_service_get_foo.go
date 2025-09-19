package service

import (
	"context"
	"strconv"
)

func (s *FooServiceImpl) GetFoo(ctx context.Context, req GetFooReq) (GetFooRes, error) {
	data := s.fooRepository.FindByID(req.ID)
	return GetFooRes{Name: s.formatFooData(req.ID, data)}, nil
}

func (s *FooServiceImpl) formatFooData(id int, data string) string {
	if data == "" {
		return ""
	}
	return "foo[" + strconv.Itoa(id) + "]: " + data
}
