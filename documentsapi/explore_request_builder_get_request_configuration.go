package documentsapi

import (
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

// ExploreRequestBuilderGetQueryParameters ...
type ExploreRequestBuilderGetQueryParameters struct {
	Page        *int    `uriparametername:"page"`
	Limit       *int    `uriparametername:"limit"`
	Query       *string `uriparametername:"query"`
	TableName   *string `uriparametername:"table_name"`
	FolderSysID *string `uriparametername:"folder_sys_id"`
	RecordSysID *string `uriparametername:"record_sys_id"`
}

// ExploreRequestBuilderGetRequestConfiguration ...
type ExploreRequestBuilderGetRequestConfiguration struct {
	Headers         *abstractions.RequestHeaders
	Options         []abstractions.RequestOption
	QueryParameters *ExploreRequestBuilderGetQueryParameters
}
