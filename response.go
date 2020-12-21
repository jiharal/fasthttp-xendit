package xenditfasthttp

type (
	// GetAvailableDisbursementBanksResponse is used to ...
	GetAvailableDisbursementBanksResponse struct {
		Name string `json:"name"`
		Code string `json:"code"`
	}

	// CreateBatchDisbursementResponse is used to ...
	CreateBatchDisbursementResponse struct {
		ID                  string `json:"id"`
		Created             string `json:"created"`
		Reference           string `json:"reference"`
		TotalUploadedAmount int    `json:"total_uploaded_amount"`
		TotalUploadedCount  int    `json:"total_uploaded_count"`
		Status              string `json:"status"`
	}

	// VirtualAccountBanksResponse is used to create response ..
	VirtualAccountBanksResponse struct {
		Name string `json:"name"`
		Code string `json:"code"`
	}

	// AvailableDisbursementBanksResponse is ...
	AvailableDisbursementBanksResponse struct {
		Name            string `json:"name"`
		Code            string `json:"code"`
		CanDisburse     bool   `json:"can_disburse"`
		CanNameValidate bool   `json:"can_name_validate"`
	}
	// DisbursementResponse is ...
	DisbursementResponse struct {
		ID                      string   `json:"id"`
		UserID                  string   `json:"user_id"`
		ExternalID              string   `json:"external_id"`
		Amount                  int      `json:"amount"`
		BankCode                string   `json:"bank_code"`
		AccountHolderName       string   `json:"account_holder_name"`
		DisbursementDescription string   `json:"disbursement_description"`
		Status                  string   `json:"status"`
		IdempotencyKey          string   `json:"idempotency_key,omitempty"`
		EmailTo                 []string `json:"email_to,omitempty"`
		EmailCC                 []string `json:"email_cc,omitempty"`
		EmailBCC                []string `json:"email_bcc,omitempty"`
	}

	// FixedPaymentCodeResponse is used to ....
	FixedPaymentCodeResponse struct {
		ID               string `json:"id"`
		OwnerID          string `json:"owner_id"`
		ExternalID       string `json:"external_id"`
		RetailOutletName string `json:"retail_outlet_name"`
		Prefix           string `json:"prefix"`
		Name             string `json:"name"`
		PaymentCode      string `json:"payment_code,omitempty"`
		Type             string `json:"type,omitempty"`
		ExpectedAmount   int    `json:"expected_amount,omitempty"`
		IsSingleUse      bool   `json:"is_single_use,omitempty"`
		ExpirationDate   string `json:"expiration_date,omitempty"`
	}

	// CreateVAResponse is used to...
	CreateVAResponse struct {
		ID              string `json:"id"`
		Name            string `json:"name"`
		Currency        string `json:"currency"`
		OwnerID         string `json:"owner_id"`
		ExternalID      string `json:"external_id"`
		BankCode        string `json:"bank_code"`
		MerchantCode    string `json:"merchant_code"`
		AccountNumber   string `json:"account_number"`
		ExpirationDate  string `json:"expiration_date"`
		IsClosed        bool   `json:"is_closed"`
		IsSingleUse     bool   `json:"is_single_use"`
		Status          string `json:"status"`
		SuggestedAmount int    `json:"suggested_amount,omitempty"`
		ExpectedAmount  int    `json:"expected_amount,omitempty"`
		Description     string `json:"description,omitempty"`
	}
	// ResponseError is used to..
	ResponseError struct {
		ErrorCode string `json:"error_code"`
		Message   string `json:"message"`
	}

	// SimulateResponse struct ..
	SimulateResponse struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}
)
