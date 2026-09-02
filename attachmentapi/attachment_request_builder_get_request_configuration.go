// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package attachmentapi

import abstractions "github.com/microsoft/kiota-abstractions-go"

// AttachmentRequestBuilderGetRequestConfiguration represents a set of options to be used when making HTTP requests.
type AttachmentRequestBuilderGetRequestConfiguration = abstractions.RequestConfiguration[AttachmentRequestBuilderGetQueryParameters]
