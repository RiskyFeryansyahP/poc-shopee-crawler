# Shopee Crawler

This project is a web scraper for Shopee using `chromedp` in Go. It extracts shop statistics and product details for a specific shop URL and logs them to the console. The scraper runs periodically every 3 minutes to fetch the latest data.

## Features

- Fetches and logs:
  - Shop name and statistics
  - Product names, prices, and currencies
- Periodic data scraping every 3 minutes

## Limitation
- The code and scraper are limited to one specific shop.
- Cannot bypass Shopee's authentication when it auto-redirects to the login page.
    - Be cautious when running the scraper too frequently, as Shopee may detect and block you by forcing a login page.
- The code requires improvements for scalability, handling multiple shops, and avoiding login redirects

## Requirements

- Go 1.19 or later
- Google Chrome or Chromium installed

## Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/yourusername/shopee-crawler.git
   cd shopee-crawler
2. **Install dependencies**:
    ```bash
    go mod tidy
    ```

## Configuration
1. Set the Target URL: Update the targetURL variable in the main() function with the Shopee shop URL you want to scrape:
    ```go
    targetURL := "https://shopee.com.my/evolene_sumbar.my#product_list"
    ```
2. Adjust Scraping Frequency: Change the interval for periodic scraping by updating the time.NewTicker duration:
    ```go
    ticker := time.NewTicker(time.Minute * 3) // Scrape every 3 minutes
    ```

## Running the scrapper
```bash
go run main.go
```

The scrapper will:
- Fetch shop and product details.
- Print the extracted data to the console
- Repeat the scraping process every 3 minutes.
