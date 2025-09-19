package service

type GetFooReq struct {
	ID int
}

type GetFooRes struct {
	Name string
}

type SaveFooReq struct {
	Data string
}

type SaveFooRes struct {
	Success bool
}

type BackupFooReq struct {
	ID int
}

type BackupFooRes struct {
	Success bool
}

type PayForFooReq struct {
	ID       int
	Username string
	Password string
	Amount   float64
	Currency string
}

type PayForFooRes struct {
	TransactionId string
}
