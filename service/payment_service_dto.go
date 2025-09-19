package service

type PayReq struct {
	Username string
	Password string
	Amount   float64
	Currency string
}

type PayRes struct {
	TransactionID string
}

type RefundReq struct {
	Username      string
	Password      string
	TransactionID string
}

type RefundRes struct {
	Success bool
}
