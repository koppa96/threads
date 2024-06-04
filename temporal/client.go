package temporal

import (
	"context"
	"go.temporal.io/sdk/client"
	"go.uber.org/fx"
	"threads/deps"
)

func NewClient(lc fx.Lifecycle, config *deps.Config) client.Client {
	c, err := client.Dial(client.Options{
		HostPort: config.TemporalAddress,
	})

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return err
		},
		OnStop: func(ctx context.Context) error {
			c.Close()
			return nil
		},
	})

	return c
}
