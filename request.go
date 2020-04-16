package xenditfasthttp

type (
	// CreateVARequest is ...
	CreateVARequest struct {
		ExternalID  string `json:"external_id"`
		BankCode    string `json:"bank_code"`
		Name        string `json:"name"`
		IsSingleUse bool   `json:"is_single_use"`
	}
)
