package browserSpider

import (
	"github.com/chromedp/chromedp"
	"testing"
)

func TestNew2(t *testing.T) {
	browserHandle := New()
	chromedp.Run(browserHandle.Ctx,
		browserHandle.Navigate("http://www.baidu.com", ""),
	)
}
