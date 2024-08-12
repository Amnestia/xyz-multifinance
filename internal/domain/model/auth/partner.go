package authmodel

type Partner struct {
	ID       int64  `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	ClientID string `json:"client_id" db:"client_id"`
	APIKey   string `json:"api_key" db:"api_key"`
	Webhook  string `json:"webhook" db:"webhook"`
}
