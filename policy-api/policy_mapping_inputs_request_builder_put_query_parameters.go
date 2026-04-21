package policyapi

// PolicyMappingInputsRequestBuilderPutQueryParameters represents the query parameters for Put
type PolicyMappingInputsRequestBuilderPutQueryParameters struct {

	// AppName represents the appName query parameter
	AppName *string `url:"appName,omitempty"`

	// DeployableName represents the deployableName query parameter
	DeployableName *string `url:"deployableName,omitempty"`

	// InputName represents the inputName query parameter
	InputName *string `url:"inputName,omitempty"`

	// InputValue represents the inputValue query parameter
	InputValue *string `url:"inputValue,omitempty"`

	// PolicyName represents the policyName query parameter
	PolicyName *string `url:"policyName,omitempty"`

	// ReturnFields represents the returnFields query parameter
	ReturnFields *string `url:"returnFields,omitempty"`
}
