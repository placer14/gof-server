package provider

import (
	"context"

	"github.com/open-feature/go-sdk/openfeature"
)

type MDUProviderMock struct {
	stringFlagValues map[string]string
	boolFlagValues   map[string]bool
	floatFlagValues  map[string]float64
	intFlagValues    map[string]int64
}

func (m *MDUProviderMock) SetFloat(name string, value float64) {
	m.floatFlagValues[name] = value
}
func (m *MDUProviderMock) SetString(name string, value string) {
	m.stringFlagValues[name] = value

}
func (m *MDUProviderMock) SetBool(name string, value bool) {
	m.boolFlagValues[name] = value

}
func (m *MDUProviderMock) SetInt(name string, value int64) {
	m.intFlagValues[name] = value

}

func (m *MDUProviderMock) Metadata() openfeature.Metadata { return openfeature.Metadata{} }
func (m *MDUProviderMock) BooleanEvaluation(ctx context.Context, flag string, defaultValue bool, evalCtx openfeature.FlattenedContext) openfeature.BoolResolutionDetail {
	return openfeature.BoolResolutionDetail{Value: m.boolFlagValues[flag]}
}
func (m *MDUProviderMock) StringEvaluation(ctx context.Context, flag string, defaultValue string, evalCtx openfeature.FlattenedContext) openfeature.StringResolutionDetail {
	return openfeature.StringResolutionDetail{Value: m.stringFlagValues[flag]}
}
func (m *MDUProviderMock) FloatEvaluation(ctx context.Context, flag string, defaultValue float64, evalCtx openfeature.FlattenedContext) openfeature.FloatResolutionDetail {
	return openfeature.FloatResolutionDetail{Value: m.floatFlagValues[flag]}
}
func (m *MDUProviderMock) IntEvaluation(ctx context.Context, flag string, defaultValue int64, evalCtx openfeature.FlattenedContext) openfeature.IntResolutionDetail {
	return openfeature.IntResolutionDetail{Value: m.intFlagValues[flag]}
}
func (m *MDUProviderMock) ObjectEvaluation(ctx context.Context, flag string, defaultValue interface{}, evalCtx openfeature.FlattenedContext) openfeature.InterfaceResolutionDetail {
	return openfeature.InterfaceResolutionDetail{}
}
func (m *MDUProviderMock) Hooks() []openfeature.Hook { return []openfeature.Hook{} }

