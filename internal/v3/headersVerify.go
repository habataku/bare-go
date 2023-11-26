package v3

import (
    "fmt"
    "net/http"
    "net/url"
)

func verifyHeaders(headers http.Header) (err error) {
    if headers.Get("X-Bare-Headers") == "" {
        return fmt.Errorf("Missing X-Bare-Headers")
    }
    if headers.Get("X-Bare-Url") == "" {
        return fmt.Errorf("Missing X-Bare-Url")
    }
    _, urlErr := url.ParseRequestURI(headers.Get("X-Bare-Url"))
    if urlErr != nil {
        return fmt.Errorf("Invalid X-Bare-Url")
    }
    return nil
}
