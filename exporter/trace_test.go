package exporter

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/exporter/exportertest"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

func TestTracesExporterNoErrors(t *testing.T) {
	f := NewFactory()
	lte, err := f.CreateTracesExporter(context.Background(), exportertest.NewNopCreateSettings(), f.CreateDefaultConfig())
	require.NotNil(t, lte)
	assert.NoError(t, err)

	assert.NoError(t, lte.ConsumeTraces(context.Background(), ptrace.NewTraces()))

	assert.NoError(t, lte.Shutdown(context.Background()))
}
