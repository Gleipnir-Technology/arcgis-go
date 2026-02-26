package response

type AdvancedEditingCapabilities struct {
	SupportsSplit                        bool     `json:"supportsSplit"`
	SupportsReturnServiceEditsInSourceSR bool     `json:"supportsReturnServiceEditsInSourceSR"`
	SupportedSQLFormatsInCalculate       []string `json:"supportedSqlFormatesInCalculate,omitempty"`
	SupportsAsyncApplyEdits              bool     `json:"supportsAsyncApplyEdits,omitempty"`
	SupportsReturnEditResults            bool     `json:"supportsReturnEditResults,omitempty"`
	SupportsApplyEditsbyUploadID         bool     `json:"supportsApplyEditsbyUploadID,omitempty"`
	SupportedApplyEditsUploadIDFormats   string   `json:"supportedApplyEditsUploadIDFormats,omitempty"`
}
