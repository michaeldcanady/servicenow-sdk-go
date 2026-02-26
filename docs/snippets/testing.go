package snippets

import (
	"testing"

	"github.com/jarcoal/httpmock"
)

// [START testing_mocking]
func TestMyFeature(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://xSDK_SN_URLx/api/now/v1/table/xSDK_SN_TABLEx",
		httpmock.NewStringResponder(200, `{"result": []}`))

	// ... test logic ...
}
// [END testing_mocking]
