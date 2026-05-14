package http

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const (
	userAgent   = "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"
	httpTimeout = 15 * time.Second
)

func NewClient() *Client {
	return &Client{
		c: &http.Client{Timeout: httpTimeout},
	}
}

type Client struct {
	c *http.Client
}

// Fetch retrieves the body of the given URL using the provided context.
func (c *Client) Fetch(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("create HTTP request failed: %w", err)
	}

	req.Header.Set("User-Agent", userAgent)

	resp, err := c.c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if isBotProtection(resp) {
			return nil, fmt.Errorf("site is behind bot protection (likely requires JavaScript/TLS fingerprinting); use a Fetcher backed by a headless browser or scraping service: status %d", resp.StatusCode)
		}

		return nil, fmt.Errorf("received non-200 status code in response: %d", resp.StatusCode)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read response body: %w", err)
	}

	return b, nil
}

// isBotProtection returns true if the response carries fingerprints of a
// well-known bot-protection service that blocks plain HTTP clients.
func isBotProtection(resp *http.Response) bool {
	// Vercel challenge
	if resp.Header.Get("x-vercel-mitigated") == "challenge" {
		return true
	}

	// Cloudflare challenge / Turnstile
	if resp.Header.Get("cf-mitigated") == "challenge" {
		return true
	}
	if strings.Contains(resp.Header.Get("server"), "cloudflare") &&
		(resp.StatusCode == http.StatusForbidden || resp.StatusCode == 503) {
		return true
	}

	// Akamai Bot Manager surfaces an `ak_p` entry in server-timing on blocks.
	if resp.StatusCode == http.StatusForbidden {
		for _, v := range resp.Header.Values("Server-Timing") {
			if strings.Contains(v, "ak_p") {
				return true
			}
		}
	}

	// DataDome
	if resp.Header.Get("x-datadome") != "" ||
		resp.Header.Get("x-dd-b") != "" {
		return true
	}

	// PerimeterX
	if resp.Header.Get("x-px-block") != "" {
		return true
	}

	return false
}
