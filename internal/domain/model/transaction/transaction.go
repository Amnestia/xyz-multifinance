package transactionmodel

type Limit struct {
	ID         int64 `json:"id" db:"id"`
	ConsumerID int64 `json:"consumer_id" db:"consumer_id"`
	Duration   int   `json:"duration" db:"duration"`
	Amount     int64 `json:"amount" db:"amount"`
}

type Transaction struct {
	ID               int64  `json:"id" db:"id"`
	ContractNumber   string `json:"contract_number" db:"contract_number"`
	AssetName        string `json:"asset_name" db:"asset_name"`
	ConsumerID       int64  `json:"consumer_id" db:"consumer_id"`
	PartnerID        int64  `json:"partner_id" db:"partner_id"`
	OTR              int64  `json:"otr" db:"otr"`
	AdminFee         int64  `json:"admin_fee" db:"admin_fee"`
	TotalInstallment int64  `json:"total_installment" db:"total_installment"`
	Interest         int64  `json:"interest" db:"interest"`
}

type Payment struct {
	ID            int64  `json:"id" db:"id"`
	TransactionID int64  `json:"transaction_id" db:"transaction_id"`
	ConsumerID    int64  `json:"consumer_id" db:"consumer_id"`
	TotalAmount   int64  `json:"total_amount" db:"total_amount"`
	MonthlyAmount int64  `json:"monthly_amount" db:"monthly_amount"`
	Duration      int    `json:"duration" db:"duration"`
	Interest      int64  `json:"interest" db:"interest"`
	Status        int    `json:"status" db:"status"`
	CreatedBy     string `json:"-" db:"created_by"`
	UpdatedBy     string `json:"-" db:"updated_by"`
}

type PaymentInstallment struct {
	ID        int64  `json:"id" db:"id"`
	PaymentID int64  `json:"payment_id" db:"payment_id"`
	Amount    int64  `json:"amount" db:"amount"`
	Status    int    `json:"status" db:"status"`
	CreatedBy string `json:"-" db:"created_by"`
	UpdatedBy string `json:"-" db:"updated_by"`
}
