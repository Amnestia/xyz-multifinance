package transactionmodel

type LookupPayload struct {
	ConsumerID int64
	PaymentID  []int64
	Duration   int
	Status     int
}
