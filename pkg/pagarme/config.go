package pagarme

import (
	"net/http"
)

// Config holds the config of pagarme
type Config struct {
	Apikey    string
	Cryptokey string
	Client    *http.Client
	Trace     bool
}
