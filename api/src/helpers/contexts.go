package helpers

import (
	"context"
	"time"
)

func TimeoutCtx(s int) context.Context {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	return ctx
}
