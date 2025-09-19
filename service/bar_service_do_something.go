package service

import "context"

func (s *BarServiceImpl) DoSomething(ctx context.Context, req DoSomethingReq) (DoSomethingRes, error) {
	token := s.authClient.Login(req.User, req.Pass)
	return DoSomethingRes{Result: "BarService got token: " + token}, nil
}
