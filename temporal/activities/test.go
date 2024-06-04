package activities

import (
	"context"
	"fmt"
	"threads/deps"
)

type TestActivity struct {
	config *deps.Config
}

func NewTestActivity(config *deps.Config) *TestActivity {
	return &TestActivity{config: config}
}

func (a *TestActivity) Test(ctx context.Context) error {
	fmt.Printf("%+v", a.config)
	return nil
}
