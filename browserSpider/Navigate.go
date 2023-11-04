package browserSpider

import (
	"context"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/front-ck996/csy"
	"net/url"
	"strings"
	"time"
)

func (c *BrowserHandle) Navigate(_url string, cookie string) chromedp.Tasks {
	cookies := []string{}
	split := strings.Split(cookie, ";")
	for _, v := range split {
		_one_cookie := strings.Split(strings.TrimSpace(v), "=")
		if len(_one_cookie) == 2 {
			cookies = append(cookies, _one_cookie...)
		}
	}
	if len(cookies)%2 != 0 {
		panic("length of cookies must be divisible by 2")
	}
	parse, _ := url.Parse(_url)
	name := "." + csy.DomainRootName(parse.Host)
	return chromedp.Tasks{
		chromedp.ActionFunc(func(ctx context.Context) error {
			// create cookie expiration
			expr := cdp.TimeSinceEpoch(time.Now().Add(180 * 24 * time.Hour))
			// add cookies to chrome
			for i := 0; i < len(cookies); i += 2 {
				err := network.SetCookie(cookies[i], cookies[i+1]).
					WithExpires(&expr).
					WithDomain(name).
					//WithHTTPOnly(true).
					Do(ctx)
				if err != nil {
					return err
				}
			}
			return nil
		}),
		// navigate to site
		chromedp.Navigate(_url),
		// read the returned values
		// read network values
		chromedp.ActionFunc(func(ctx context.Context) error {
			//cookies, err := network.GetAllCookies().Do(ctx)
			//if err != nil {
			//	return err
			//}
			//
			////for i, cookie := range cookies {
			////	log.Printf("chrome cookie %d: %+v", i, cookie)
			////}

			return nil
		}),
	}
}
