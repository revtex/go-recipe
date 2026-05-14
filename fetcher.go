package recipe

import "context"

// Fetcher retrieves the raw bytes of a page at the given URL.
//
// The default implementation in pkg/recipe uses Go's net/http package and
// works for most recipe sites. Sites behind JavaScript-based bot protection
// (e.g. Vercel challenges, Cloudflare Turnstile) require a custom Fetcher
// that can execute JavaScript, such as one backed by a headless browser
// (chromedp, rod) or a third-party scraping service.
type Fetcher interface {
	Fetch(ctx context.Context, url string) ([]byte, error)
}
