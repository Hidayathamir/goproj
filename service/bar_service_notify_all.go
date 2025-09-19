package service

import (
	"context"
	"fmt"

	"github.com/Hidayathamir/goproj/pkg/http/httperror"
)

func (s *BarServiceImpl) NotifyAll(ctx context.Context, req NotifyAllReq) (NotifyAllRes, error) {
	if !s.authClient.ValidateToken(req.Token) {
		return NotifyAllRes{}, httperror.InvalidToken()
	}
	for _, channel := range s.slackClient.GetChannelList() {
		err := s.slackClient.SendMessage(channel, req.Message)
		if err != nil {
			return NotifyAllRes{}, fmt.Errorf("s.slackClient.SendMessage, channel=%s: %w", channel, err)
		}
	}
	return NotifyAllRes{Success: true}, nil
}
