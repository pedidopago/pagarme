package pagarme

import (
	"io"
	"net/http"
	"time"

	"go.uber.org/zap"
)

// Config holds the config of pagarme
type Config struct {
	Apikey    string
	Cryptokey string
	Client    *http.Client
	Trace     bool
	Logger    *zap.Logger
}

// Default configuration uses:
//
// HTTP Client with 60 seconds timeout
// os.stdout for logging
// loglevel: error
func Default(apikey, cryptokey string) *Config {
	zl, _ := zap.NewProduction()
	cfg := &Config{
		Apikey:    apikey,
		Cryptokey: cryptokey,
		Client: &http.Client{
			Timeout: time.Second * 60,
		},
		Trace:  false,
		Logger: zl,
	}
	return cfg
}

func (c *Config) Do(method, urlpart string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, "https://api.pagar.me/1"+urlpart, body)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.Apikey, "x")
	req.Header.Set("X-PagarMe-User-Agent", "github.com/pedidopago/pagarme Dev")
	req.Header.Set("X-PagarMe-Version", "2017-08-28")
	req.Header.Set("Content-Type", "application/json")
	if c.Client == nil {
		return http.DefaultClient.Do(req)
	}
	return c.Client.Do(req)
}
