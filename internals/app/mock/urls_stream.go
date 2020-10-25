package mock

import (
	"strings"
)

// UrlsStreamFive is an input URLs stream mock with five URLs
var UrlsStreamFive = strings.NewReader("Url 1\r\nUrl 2\r\nUrl 3\r\nUrl 4\r\nUrl 5\r\n")

// UrlsStreamTwo is an input URLs stream mock with two URLs
var UrlsStreamTwo = strings.NewReader("Url 1\r\nUrl 2")
