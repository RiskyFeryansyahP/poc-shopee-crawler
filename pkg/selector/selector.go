package selector

const (
	BaseSelector               = "div.shop-page_product-list > div.shop-all-product-view > div.shop-search-result-view > div"
	ShopNameSelector           = "div.section-seller-overview-horizontal__portrait-info > h1.section-seller-overview-horizontal__portrait-name"
	ShopStatisticKeySelector   = "div.section-seller-overview-horizontal__seller-info-list div.section-seller-overview__item-text-name"
	ShopStatisticValueSelector = "div.section-seller-overview-horizontal__seller-info-list div.section-seller-overview__item-text-value"
	ProductNameSelector        = `div.shop-search-result-view a.contents div > div:nth-child(2) > div.space-y-1.mb-1.flex-1.flex.flex-col.justify-between > div:nth-child(1)`
	CurrencySelector           = "div.shop-search-result-view a.contents div > div:nth-child(2) > div.space-y-1.mb-1.flex-1.flex.flex-col.justify-between > div:nth-child(2) > div > div > span:nth-child(1)"
	PriceSelector              = "div.shop-search-result-view a.contents div > div:nth-child(2) > div.space-y-1.mb-1.flex-1.flex.flex-col.justify-between > div:nth-child(2) > div > div > span:nth-child(2)"
)
