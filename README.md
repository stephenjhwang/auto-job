# auto-job
Automizing the job application process one job at a time

# Goals:

## 1. Get LinkedIn job postings

How?

1. LinkedIn API?

* Does not support retreiving jobs so not possible.

2. Web Scraping with Selenium? 

* Possible and most mature way of web automation at this point in time.

3. Web Scraping with [chromedp](https://github.com/chromedp/chromedp)? 

* Possible and boasts about being faster than selenium

## 2. Setup auto form filling to apply to jobs

1. Fill Login

* I use Bitwarden so would likely use `bw serve` (https://bitwarden.com/help/vault-management-api/) to have local bitwarden api server to access Bitwarden login information and password generation.

2. Fill form

* Either with selenium or chromedp. Will need specific functions to handle different sites (eg. workday, icims, greenhouse). May not even be possible because of so many unique fields especially in workday.

## 3. Interface

Will be done with CLI and then automate it with cron job on some cloud provider or locally.
