package provider_test

import (
	"context"
	"testing"

	"github.com/open-feature/go-sdk/openfeature"
	"github.com/placer14/gof-server/internal/provider"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStringEvaluation(t *testing.T) {
	mock := provider.NewProviderMock()
	assert.NoError(t, openfeature.SetProvider(mock))

	mockImpl, ok := mock.(*provider.MDUProviderMock)
	require.True(t, ok)

	expectedStr := "metal.v1"
	mockImpl.SetString("dataplane_generation", expectedStr)

	evalCtx := openfeature.NewEvaluationContext("", map[string]any{})
	client := openfeature.NewClient("stringEvalTests")
	generation, err := client.StringValue(context.Background(), "dataplane_generation", "k8s.v1", evalCtx)
	assert.NoError(t, err)
	assert.Equal(t, expectedStr, generation)
}

func TestBoolEvaluation(t *testing.T) {
	mock := provider.NewProviderMock()
	assert.NoError(t, openfeature.SetProvider(mock))

	mockImpl, ok := mock.(*provider.MDUProviderMock)
	assert.True(t, ok)

	err := openfeature.SetProvider(mock)
	assert.NoError(t, err)

	expectedStr := "metal.v1"
	mockImpl.SetString("dataplane_generation", expectedStr)

	evalCtx := openfeature.NewEvaluationContext("", map[string]any{})
	client := openfeature.NewClient("boolEvalTests")

	generation, err := client.BooleanValue(context.Background(), "grant_soil_access", true, evalCtx)
	assert.NoError(t, err)
	assert.Equal(t, false, generation)
}

func TestFloatEvaluation(t *testing.T) {
	mock := provider.NewProviderMock()
	assert.NoError(t, openfeature.SetProvider(mock))

	mockImpl, ok := mock.(*provider.MDUProviderMock)
	assert.True(t, ok)
	err := openfeature.SetProvider(provider.NewProvider())
	assert.NoError(t, err)

	evalCtx := openfeature.NewEvaluationContext("", map[string]any{})
	client := openfeature.NewClient("floatEvalTests")

	generation, err := client.FloatValue(context.Background(), "special_ability_buff_perc", 0.0, evalCtx)

	assert.NoError(t, err)
	assert.Equal(t, 0.23456, generation)
}

func TestIntEvaluation(t *testing.T) {
	mock := provider.NewProviderMock()
	assert.NoError(t, openfeature.SetProvider(mock))

	mockImpl, ok := mock.(*provider.MDUProviderMock)
	assert.True(t, ok)
	err := openfeature.SetProvider(provider.NewProvider())
	assert.NoError(t, err)

	evalCtx := openfeature.NewEvaluationContext("", map[string]any{})
	client := openfeature.NewClient("intEvalTests")

	generation, err := client.IntValue(context.Background(), "num_of_special_abilities", 0, evalCtx)

	assert.NoError(t, err)
	assert.Equal(t, int64(12), generation)
}
