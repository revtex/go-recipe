package recipe

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/kkyr/go-recipe"
	"github.com/kkyr/go-recipe/internal/html/scrape/schema"
	"github.com/kkyr/go-recipe/internal/http"
	"github.com/kkyr/go-recipe/internal/url"

	"github.com/PuerkitoBio/goquery"
)

var defaultFetcher recipe.Fetcher = http.NewClient()

// NewHTTPFetcher returns the default HTTP fetcher used by ScrapeURL.
//
// It performs a simple GET with a desktop User-Agent and works for most
// recipe sites. Sites behind JavaScript-based bot protection (e.g. Vercel
// challenges, Cloudflare Turnstile) will fail with this fetcher; supply a
// custom recipe.Fetcher via ScrapeURLWithFetcher in that case.
func NewHTTPFetcher() recipe.Fetcher {
	return http.NewClient()
}

// ScrapeURL retrieves the source at the provided url and returns a
// Scraper that scrapes recipe data from the retrieved HTML.
//
// It uses the built-in HTTP fetcher. For sites behind bot protection or
// to provide custom transport behavior, use ScrapeURLWithFetcher.
func ScrapeURL(urlStr string) (recipe.Scraper, error) {
	return ScrapeURLWithFetcher(context.Background(), urlStr, defaultFetcher)
}

// ScrapeURLWithFetcher retrieves the source at the provided url using the
// supplied Fetcher and returns a Scraper that scrapes recipe data from the
// retrieved HTML.
//
// Use this when the default fetcher cannot reach the site (for example,
// sites behind JavaScript-based bot protection) and pass a Fetcher that
// can — typically one backed by a headless browser or a third-party
// scraping service.
func ScrapeURLWithFetcher(ctx context.Context, urlStr string, fetcher recipe.Fetcher) (recipe.Scraper, error) {
	body, err := fetcher.Fetch(ctx, urlStr)
	if err != nil {
		return nil, fmt.Errorf("unable to GET url: %w", err)
	}

	return ScrapeHTML(urlStr, bytes.NewReader(body))
}

// ScrapeHTML returns a Scraper that scrapes recipe data from the provided HTML.
// the urlStr is used to determine if a specific scraper should be used, otherwise
// a generic scraper is used.
func ScrapeHTML(urlStr string, body io.Reader) (recipe.Scraper, error) {
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return nil, fmt.Errorf("unable to parse HTML document: %w", err)
	}

	host := url.GetHost(urlStr)
	if scraper, ok := hostToScraper[host]; ok {
		return scraper(doc)
	}

	scraper, err := schema.NewRecipeScraper(doc)
	if err != nil {
		return nil, fmt.Errorf("unable to get new schema scraper: %w", err)
	}

	return scraper, nil
}
