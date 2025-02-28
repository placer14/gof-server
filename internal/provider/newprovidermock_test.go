package provider_test

import(
	"github.com/placer14/gof-server/internal/provider"
	"testing"
)

func TestNewProviderMock(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		want provider.MDUProvider
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := provider.NewProviderMock()
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("NewProviderMock() = %v, want %v", got, tt.want)
			}
		})
	}
}
