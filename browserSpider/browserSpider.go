package browserSpider

import (
	"context"
	"github.com/chromedp/chromedp"
	"github.com/front-ck996/csy"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type BrowserHandle struct {
	Ctx      context.Context
	Headless bool
	Close    context.CancelFunc
	UA       string
	TempDir  string
	NoDel    bool
}

func New() BrowserHandle {
	c := BrowserHandle{}
	if c.UA == "" {
		c.UA = "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.87 Safari/537.36"
	}
	if c.TempDir == "" {
		c.TempDir = "c:\\click-temp\\"
	}
	if c.NoDel {
		dir, _ := csy.NewFile().IsDir(c.TempDir)
		if !dir {
			os.MkdirAll(c.TempDir, os.ModePerm)
		}
	} else {
		os.RemoveAll(c.TempDir)
		os.MkdirAll(c.TempDir, os.ModePerm)
	}
	dir := c.TempDir
	if c.NoDel {

	} else {
		var err error
		dir, err = ioutil.TempDir(c.TempDir, "chromedp-example")
		if err != nil {
			panic(err)
		}
	}
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.UserAgent(c.UA),
		chromedp.DisableGPU,
		chromedp.NoDefaultBrowserCheck,
		chromedp.Flag("headless", c.Headless),
		//chromedp.Flag("start-maximized", true),
		chromedp.Flag("ignore-certificate-errors", true),
		//chromedp.Flag("incognito", true),
		chromedp.Flag("window-size", "1380,900"),
		chromedp.UserDataDir(dir),
		chromedp.Flag("disable-infobars", true),
		//chromedp.Flag("disable-infobars", true),
		chromedp.Flag("excludeSwitches", `['enable-automation', 'load-extension']`),
	)
	chromelist := []string{
		`C:\Users\123\Desktop\Chrome-bin\chrome.exe`,
		`C:\Users\Administrator\Desktop\chrome\Chrome-bin\chrome.exe`,
		`C:\Users\Administrator\AppData\Local\Google\Chrome\Application\chrome.exe`,
		`C:\Program Files (x86)\Google\Chrome\Application\chrome.exe`,
	}
	for _, chrome_v := range chromelist {
		if csy.NewFile().IsFile(chrome_v) {
			opts = append(opts, chromedp.ExecPath(chrome_v))
			break
		}
	}
	var allocCtx context.Context
	var ctxxx context.Context
	allocCtx, c.Close = chromedp.NewExecAllocator(context.Background(), opts...)

	// also set up a custom logger
	ctxxx, c.Close = chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	//程序允许得总耗时
	c.Ctx, c.Close = context.WithTimeout(ctxxx, time.Hour*90000)
	//page.AddScriptToEvaluateOnNewDocument()

	chromedp.Run(c.Ctx, DelWebdriver())
	return c
}

// 关闭浏览器
func (c *BrowserHandle) Off() {
	c.Close()
}
