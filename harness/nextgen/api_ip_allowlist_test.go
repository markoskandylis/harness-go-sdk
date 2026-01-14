package nextgen

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetIpAllowlist(t *testing.T) {
	client := NewAPIClient(NewConfiguration())
	require.NotNil(t, client)

	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	apiKey := os.Getenv("HARNESS_PLATFORM_API_KEY")

	if accountId == "" || apiKey == "" {
		t.Skip("Skipping test because HARNESS_ACCOUNT_ID or HARNESS_PLATFORM_API_KEY is not set.")
	}

	_, _, err := client.IpAllowlistApi.GetIpAllowlist(context.Background(), accountId)
	require.NoError(t, err)
}