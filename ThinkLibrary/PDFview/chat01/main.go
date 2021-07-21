package main

import (
	"context"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"io/ioutil"
	"log"
)

func main() {

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var buf []byte
	if err := chromedp.Run(ctx, savePDF("https://colobu.com/2021/05/05/generate-pdf-for-a-web-page-by-using-chromedp/", &buf)); err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile("baidu.pdf", buf, 0666); err != nil {
		log.Fatal(err)
	}

	log.Print("succ.")

}

func savePDF(urlstr string, res *[]byte) chromedp.Tasks {

	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, err := page.PrintToPDF().WithPrintBackground(false).Do(ctx)
			if err != nil {
				return err
			}
			*res = buf
			return nil
		}),
	}
}
