package transactionmodel

type TransactionRequest struct {
	NIK              string `json:"nik"`
	ContractNumber   string `json:"contract_number"`
	AssetName        string `json:"asset_name"`
	OTR              int64  `json:"otr"`
	AdminFee         int64  `json:"admin_fee"`
	TotalInstallment int64  `json:"total_installment"`
	Duration         int    `json:"duration"`
	Interest         int64  `json:"interest"`
}
