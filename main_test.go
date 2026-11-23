package main

import (
	"encoding/json"
	"testing"

	corev2 "github.com/sensu/core/v2"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
}

func TestExecuteMutator(t *testing.T) {
	assert := assert.New(t)
	event := corev2.FixtureEvent("entity1", "check1")
	event.Metrics = corev2.FixtureMetrics()
	ev, err := executeMutator(event)
	assert.NoError(err)
	assert.Equal(0, len(ev.Entity.Redact))
	assert.Equal(0, len(ev.Entity.System.Network.Interfaces))
	assert.Equal(0, len(ev.Entity.Subscriptions))
	assert.Equal(0, len(ev.Check.Handlers))
	assert.Equal(0, len(ev.Check.History))
	assert.Equal(0, len(ev.Check.RuntimeAssets))
	assert.Equal(0, len(ev.Check.Subscriptions))
}

// Regression test: the mutated event must still encode to JSON when it carries
// maps (labels, annotations). modern-go/reflect2 v1.0.1 panicked here on Go 1.18+.
func TestExecuteMutatorEncodesMaps(t *testing.T) {
	assert := assert.New(t)
	event := corev2.FixtureEvent("entity1", "check1")
	event.Entity.ObjectMeta.Labels = map[string]string{"site": "cadc"}
	event.Entity.ObjectMeta.Annotations = map[string]string{"note": "x"}
	event.Check.ObjectMeta.Labels = map[string]string{"team": "it"}
	ev, err := executeMutator(event)
	assert.NoError(err)
	out, err := json.Marshal(ev)
	assert.NoError(err)
	assert.Contains(string(out), `"site":"cadc"`)
}
