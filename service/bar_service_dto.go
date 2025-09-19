package service

type DoSomethingReq struct {
	User string
	Pass string
}

type DoSomethingRes struct {
	Result string
}

type DoAnotherThingReq struct {
	Token string
}

type DoAnotherThingRes struct {
	Success bool
}

type GetAllDataReq struct {
	Token string
}

type GetAllDataRes struct {
	Data []string
}

type NotifyAllReq struct {
	Token   string
	Message string
}

type NotifyAllRes struct {
	Success bool
}
