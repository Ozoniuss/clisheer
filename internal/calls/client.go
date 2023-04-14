package calls

import (
	"net/http"
	"time"
)

// client is a custom http client used by this package. Having this client
// allows for both customizing the http client.
//
// It's also helpful to reuse an http client becuase of the client's ability to
// cache TCP connection and reuse them accross multiple requests via the
// Keep-Alive mechanism. However, this would have also been achieved with
// the default http client.
var client = http.Client{
	// Todo: customize client behaviour.
	Timeout: 30 * time.Second,
}
