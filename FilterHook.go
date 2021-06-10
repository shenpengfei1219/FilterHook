package FilterHook

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type FilterHook struct{}

var _ redis.Hook = (*FilterHook)(nil)

func NewFilterHook() *FilterHook {
	return new(FilterHook)
}

func (FilterHook) BeforeProcess(ctx context.Context, cmd redis.Cmder) (context.Context, error) {
	return ctx, nil
}

func (FilterHook) AfterProcess(ctx context.Context, cmd redis.Cmder) error {
	return nil
}

func (FilterHook) BeforeProcessPipeline(ctx context.Context, cmds []redis.Cmder) (context.Context, error) {
	return ctx, nil
}

func (FilterHook) AfterProcessPipeline(ctx context.Context, cmds []redis.Cmder) error {
	return nil
}
