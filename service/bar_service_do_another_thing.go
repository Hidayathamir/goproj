package service

import "context"

func (s *BarServiceImpl) DoAnotherThing(ctx context.Context, req DoAnotherThingReq) (DoAnotherThingRes, error) {
	ok := s.authClient.ValidateToken(req.Token)
	return DoAnotherThingRes{Success: ok}, nil
}
