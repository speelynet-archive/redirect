package main

import "net/url"

// Redirect takes an HTTP url and converts it to an HTTPS url.
func Redirect(url *url.URL) string {
	r := "https://" + url.Host + url.Path

	if url.RawQuery != "" {
		r += "?" + url.RawQuery
	}

	return r
}
