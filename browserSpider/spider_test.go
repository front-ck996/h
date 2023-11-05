package browserSpider

import (
	"github.com/chromedp/chromedp"
	"testing"
)

func TestNew2(t *testing.T) {
	browserHandle := New(BrowserHandleInit{})
	chromedp.Run(browserHandle.Ctx,
		browserHandle.Navigate("https://bot.sannysoft.com/", ""),
	)
}
