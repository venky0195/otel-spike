package exporter

import (
	"go.opentelemetry.io/collector/component"
)

// Config defines configuration for your exporter.
type config struct {
	// configgrpc.GRPCClientSettings `mapstructure:",squash"` // squash ensures fields are correctly decoded in embedded struct.
}

var _ component.Config = (*config)(nil)

// Validate the configuration for errors to implement the configvalidator interface.
// You can skip this if you do not want to validate your config
func (c *config) Validate() error {
	return nil

}
