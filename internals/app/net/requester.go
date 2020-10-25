package net

// Requester interface
type Requester interface {
	// RequestURL sends HTTP GET request to the URL
	// Input:
	//  - url to send the request
	// Output:
	//  - the body of the response, if request was successed
	//  - error, otherwise
	RequestURL(url string) (string, error)
}

// RequesterFunc is a helper type which allows
// to use functions with Requester interface
type RequesterFunc func(string) (string, error)

// RequestURL impl
func (f RequesterFunc) RequestURL(url string) (string, error) {
	return f(url)
}
