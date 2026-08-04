package fix

import "net/url"

// FixURLParameterLoss preserves URL parameters when path is rewritten
// by cloning the original query values before modification.
func PreserveQueryParams(originalURL *url.URL) url.Values {
    params := make(url.Values)
    for k, v := range originalURL.Query() {
        params[k] = make([]string, len(v))
        copy(params[k], v)
    }
    return params
}
