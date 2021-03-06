// +build unit

package metrics

import (
	"fmt"
	"testing"

	metrics1 "github.com/consensys/orchestrate/pkg/toolkit/app/metrics"
	"github.com/consensys/orchestrate/pkg/toolkit/app/metrics/testutils"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListenerMetrics(t *testing.T) {
	m := NewListenerMetrics()

	registry := prometheus.NewRegistry()
	err := registry.Register(m)
	assert.NoError(t, err, "Registering TransactionSchedulerMetrics should not fail")

	m.BlockCounter().
		With("chain_uuid", "chain_uuid").
		Add(1)

	families, err := registry.Gather()
	require.NoError(t, err, "Gathering metrics should not error")
	require.Len(t, families, 1, "Count of metrics families should be correct")

	testutils.AssertCounterFamily(t, families[0], fmt.Sprintf("%s_%s", metrics1.Namespace, Subsystem), BlockName, []float64{1}, "Current block processed", nil)
}
