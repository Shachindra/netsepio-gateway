package delegatereviewsubmission

type DelegateReviewSubmissionRequest struct {
	Reviewer   string `json:"reviewer" binding:"required,hexadecimal"`
	Metadata   string `json:"metadata" binding:"required"`
	Category   string `json:"category" binding:"required"`
	Domain     string `json:"domain" binding:"required"`
	SiteUrl    string `json:"siteUrl" binding:"required"`
	SiteType   string `json:"siteType" binding:"required"`
	SiteTag    string `json:"siteTag" binding:"required"`
	SiteSafety string `json:"siteSafety" binding:"required"`
}

type DelegateReviewSubmissionPayload struct {
	Result TransactionResult `json:"Result"`
}

type TransactionResult struct {
	TransactionHash string `json:"transaction_hash"`
	GasUsed         int    `json:"gas_used"`
	GasUnitPrice    int    `json:"gas_unit_price"`
	Sender          string `json:"sender"`
	SequenceNumber  int    `json:"sequence_number"`
	Success         bool   `json:"success"`
	TimestampUS     int64  `json:"timestamp_us"`
	Version         int16  `json:"version"`
	VMStatus        string `json:"vm_status"`
}
