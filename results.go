package gophercloud

import "net/http"

// Result acts as a base struct that other results can embed.
type Result struct {
	// Resp is the deserialized JSON structure returned from the server.
	Resp map[string]interface{}

	// Headers contains the HTTP header structure from the original response.
	Headers http.Header

	// Err is an error that occurred during the operation. It's deferred until extraction to make
	// it easier to chain operations.
	Err error
}

// RFC3339Milli describes a time format used by API responses.
const RFC3339Milli = "2006-01-02T15:04:05.999999Z"

// Link represents a structure that enables paginated collections how to
// traverse backward or forward. The "Rel" field is usually either "next".
type Link struct {
	Href string `mapstructure:"href"`
	Rel  string `mapstructure:"rel"`
}

// ExtractNextURL attempts to extract the next URL from a JSON structure. It
// follows the common structure of nesting back and next links.
func ExtractNextURL(links []Link) (string, error) {
	var url string

	for _, l := range links {
		if l.Rel == "next" {
			url = l.Href
		}
	}

	if url == "" {
		return "", nil
	}

	return url, nil
}
