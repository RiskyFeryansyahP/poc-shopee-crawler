package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"

	"github.com/confus1on/poc-shopee-crawler/pkg/selector"
)

func main() {
	fmt.Println("Shopee Crawler")

	opts := append(
		chromedp.DefaultExecAllocatorOptions[:],
		// chromedp.Flag("headless", false),
		chromedp.Flag("blink-settings", "imagesEnabled=false"),
		chromedp.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36"),
	)

	ctx, _ := chromedp.NewExecAllocator(context.Background(), opts...)

	targetURL := "https://shopee.co.id/evoleneofficialshop#product_list"

	err := scrapeShopeeByShopName(ctx, targetURL)

	if err != nil {
		log.Fatalf("failed scrappe data: %+v\n", err)
	}

	ticker := time.NewTicker(time.Minute * 3)

	// scrap data frequently every 3 minutes
	for range ticker.C {
		err := scrapeShopeeByShopName(ctx, targetURL)

		if err != nil {
			log.Fatalf("failed scrappe data: %+v\n", err)
		}
	}
}

func scrapeShopeeByShopName(
	ctx context.Context,
	target string,
) error {
	ctx, cancel := chromedp.NewContext(ctx)
	defer cancel()

	var productNameNodes []*cdp.Node
	var currencyNodes []*cdp.Node
	var priceNodes []*cdp.Node
	var shopStatisticKey []*cdp.Node
	var shopStatisticValue []*cdp.Node
	var shopName string

	err := chromedp.Run(
		ctx,
		chromedp.Navigate(target),
		chromedp.WaitVisible(selector.BaseSelector),

		chromedp.Text(selector.ShopNameSelector, &shopName),
		chromedp.Nodes(selector.ShopStatisticKeySelector, &shopStatisticKey, chromedp.ByQueryAll),
		chromedp.Nodes(selector.ShopStatisticValueSelector, &shopStatisticValue, chromedp.ByQueryAll),

		chromedp.Nodes(selector.ProductNameSelector, &productNameNodes, chromedp.ByQueryAll),
		chromedp.Nodes(selector.CurrencySelector, &currencyNodes, chromedp.ByQueryAll),
		chromedp.Nodes(selector.PriceSelector, &priceNodes, chromedp.ByQueryAll),
	)

	if err != nil {
		return err
	}

	fmt.Println("Shop Name:", shopName)

	for idx := range shopStatisticKey {
		var statisticKey, statisticValue string

		err := chromedp.Run(
			ctx,
			chromedp.Text(shopStatisticKey[idx].FullXPath(), &statisticKey),
			chromedp.Text(shopStatisticValue[idx].FullXPath(), &statisticValue),
		)

		if err != nil {
			return err
		}

		fmt.Printf("%s%s\n\n", statisticKey, statisticValue)
	}

	fmt.Println("========================================================")

	for i := range productNameNodes {
		var productName, currency, price string

		err := chromedp.Run(
			ctx,
			chromedp.Text(productNameNodes[i].FullXPath(), &productName),
			chromedp.Text(currencyNodes[i].FullXPath(), &currency),
			chromedp.Text(priceNodes[i].FullXPath(), &price),
		)

		if err != nil {
			return err
		}

		fmt.Printf("%s - %s %s\n", productName, currency, price)
	}

	return nil
}
