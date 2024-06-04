package deps

import (
	"context"
	"fmt"
	_ "github.com/lib/pq"
	"go.uber.org/fx"
	"threads/ent"
)

func NewClient(lc fx.Lifecycle, config *Config) *ent.Client {
	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		config.DBHost,
		config.DBPort,
		config.DBUser,
		config.DBName,
		config.DBPass,
	)

	client, err := ent.Open("postgres", connectionString)
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if err != nil {
				return err
			}

			return client.Schema.Create(ctx)
		},
		OnStop: func(ctx context.Context) error {
			return client.Close()
		},
	})

	return client
}
