package main

import (
	"github.com/sensu/sensu-plugin-sdk/sensu"
	corev2 "github.com/sensu/core/v2"
)

// Config represents the mutator plugin config.
type Config struct {
	sensu.PluginConfig
}

var (
	mutatorConfig = Config{
		PluginConfig: sensu.PluginConfig{
			Name:     "sensu-event-streamliner",
			Short:    "Sensu Event Streamliner",
			Keyspace: "sensu.io/plugins/sensu-event-streamliner/config",
		},
	}
)

func main() {
	mutator := sensu.NewGoMutator(&mutatorConfig.PluginConfig, nil, checkArgs, executeMutator)
	mutator.Execute()
}

func checkArgs(_ *corev2.Event) error {
	return nil
}

func executeMutator(event *corev2.Event) (*corev2.Event, error) {

	event.Entity.Redact = nil
	event.Entity.System.Network.Interfaces = nil
	event.Entity.Subscriptions = nil
	event.Check.Handlers = nil
	event.Check.History = nil
	event.Check.RuntimeAssets = nil
	event.Check.Subscriptions = nil

	return event, nil
}
