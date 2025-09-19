package service

import (
	"context"
	"fmt"

	"github.com/Hidayathamir/goproj/pkg/http/httperror"
)

func (s *BarServiceImpl) GetAllData(ctx context.Context, req GetAllDataReq) (GetAllDataRes, error) {
	if !s.authClient.ValidateToken(req.Token) {
		return GetAllDataRes{}, httperror.InvalidToken()
	}
	data, err := s.barRepository.GetAll()
	if err != nil {
		return GetAllDataRes{}, fmt.Errorf("s.barRepository.GetAll: %w", err)
	}
	return GetAllDataRes{Data: data}, nil
}
