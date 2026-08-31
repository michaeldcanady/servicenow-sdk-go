// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package documentsapi

import (
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

// DeleteRequestBuilderDeleteQueryParameters ...
type DeleteRequestBuilderDeleteQueryParameters struct {
	DocSysID    *string `uriparametername:"doc_sys_id"`
	RecordSysID *string `uriparametername:"record_sys_id"`
	TableName   *string `uriparametername:"table_name"`
}

// DeleteRequestBuilderDeleteRequestConfiguration ...
type DeleteRequestBuilderDeleteRequestConfiguration struct {
	Headers         *abstractions.RequestHeaders
	Options         []abstractions.RequestOption
	QueryParameters *DeleteRequestBuilderDeleteQueryParameters
}
