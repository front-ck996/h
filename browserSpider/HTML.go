package browserSpider

import (
	"context"
	"github.com/chromedp/chromedp"
)

// 获取HTML 数据
func (c *BrowserHandle) Html(_ctx2 context.Context) string {
	var _ht string
	chromedp.Run(_ctx2, chromedp.OuterHTML("html", &_ht))
	//	chromedp.Run(_ctx2, chromedp.Evaluate(`document.querySelector("html").outerHTML`, &_ht))
	return _ht
}
