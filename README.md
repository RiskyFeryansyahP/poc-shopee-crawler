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

Example Result:
```
Shop Name: Evolene Official Shop
Produk: 42

Pengikut: 153,3RB

Mengikuti: 1

Penilaian: 4.9 (132,3RB Penilaian)

Performa Chat: 100% (Hitungan Jam)

Bergabung: 5 Tahun Lalu

========================================================
Evomass Evolene 2lbs - Suplemen Fitness - Suplemen Workout - Gainer - Rp 237.500
[NEW] Crevolene Monohydrate Evolene - Creatine Monohydrate - Rp 199.500
Crevolene Creapure Evolene Creatine - Rp 142.500
Evomass Evolene 10lbs - Suplemen Fitness - Suplemen Workout - Mass Gainer - Rp 760.000
New Prevo Evolene 7s - Pre-Workout - Energy Drink Halal dan BPOM - 2x More Performance - Rp 53.200
Mini Bundling Bulking Muscle Up - Rp 305.550
EvoWhey Evolene Whey Protein Sachet - Suplemen Fitness - Suplemen Workout - Rp 18.525
Isolene Evolene Sachet - Suplemen Fitness - Suplemen Workout - Whey Protein Isolate - Rp 20.900
New Prevo Evolene 25s - Pre-Workout - Energy Drink Halal dan BPOM - 2x More Performance - Rp 185.250
Isolene Evolene 396gr/12 Sachet - Suplemen Fitness - Suplemen Workout - Whey Protein Isolate - Rp 228.000
Isolene Evolene 50 Sachet - Suplemen Fitness - Suplemen Workout - Whey Protein Isolate - Rp 817.000
EvoWhey Evolene Whey Protein 420gr/12 Sachet - Suplemen Fitness - Suplemen Workout - Rp 185.250
Evomass Evolene Cokelat 5lbs - Suplemen Fitness - Suplemen Workout - Mass Gainer - Rp 522.500
Evoboost Evolene 60 kapsul - Suplemen Fitness - Testosteron Booster - Rp 313.500
Bundling Bulking Muscle Up - Rp 1.028.200
Prevo Evolene Sachet - Pre-Workout - Energy Drink Halal dan BPOM - No Caffeine - Rp 8.075
EvoWhey Evolene Whey Protein 50 Sachet - Suplemen Fitness - Suplemen Workout - Rp 707.750
Shaker Evolene 500ml Food Grade Material - Rp 75.000
Evolene - Protein Bar - Evobar Coklat - Rp 96.900
EvoWhey Evolene Whey Protein Cokelat 6 Sachet - Suplemen Fitness - Suplemen Workout - Rp 101.650
Prevo Evolene 7 Sachet - Pre-Workout - Energy Drink Halal dan BPOM - No Caffeine - Rp 53.200
New Prevo Evolene 1s - Pre-Workout - Energy Drink Halal dan BPOM - 2x More Performance - Rp 8.075
Prevo Evolene 225gr Pre-Workout - Energy Drink Halal dan BPOM - No Caffeine - Rp 332.500
Bundling Bulking Power Up - Rp 1.115.500
BCAA Evolene 375gram - Suplemen Fitness - Suplemen Workout - Rp 261.250
Evogreen Evolene 12 Serving 100% Plant Based Protein - Vegan - Rp 170.000
Bundling Cutting Premium Muscle Up - Rp 1.086.400
Evolene Evogreen 50s 100% Plant Based Isolate Protein - IVS Certified - Rp 617.500
Bundling Double Muscle Up - Rp 315.250
Mini Bundling Bulking Power Up - Rp 533.500
```