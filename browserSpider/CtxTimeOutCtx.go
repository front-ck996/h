package browserSpider

import (
	"context"
	"time"
)

func (c *BrowserHandle) CtxTimeOutCtx(duration time.Duration) context.Context {
	timeout, _ := context.WithTimeout(c.Ctx, duration)
	return timeout
}
