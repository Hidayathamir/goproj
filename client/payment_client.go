package client

type PaymentClient interface {
	Charge(amount float64, currency string) string
	Refund(transactionID string) bool
	GetStatus(transactionID string) string
}

// compile-time check
var _ PaymentClient = (*PaymentClientImpl)(nil)

type PaymentClientImpl struct {
}

func NewPaymentClient() *PaymentClientImpl {
	return &PaymentClientImpl{}
}

func (c *PaymentClientImpl) Charge(amount float64, currency string) string {
	return "txn_dummy_123"
}

func (c *PaymentClientImpl) Refund(transactionID string) bool {
	return true
}

func (c *PaymentClientImpl) GetStatus(transactionID string) string {
	return "SUCCESS"
}
