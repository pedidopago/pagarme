package pagarme

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
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
	//
	if c.Trace {
		buf := new(bytes.Buffer)
		buf.WriteString(fmt.Sprintf("%v %v %v\n", req.Method, req.URL, req.Proto))
		buf.WriteString(fmt.Sprintf("Host: %v\n", req.Host))
		// headers
		for name, headers := range req.Header {
			for _, v := range headers {
				buf.WriteString(fmt.Sprintf("%v: %v\n", name, v))
			}
		}

		if req.Body != nil {
			buf.WriteString("\n")
			buf2 := new(bytes.Buffer)
			io.Copy(buf2, req.Body)
			bbytes := buf2.Bytes()
			req.Body.Close()
			buf2.Reset()
			buf2.Write(bbytes)
			req.Body = ioutil.NopCloser(buf2)
			buf.Write(bbytes)
			buf.WriteRune('\n')
		}
		fmt.Println(buf.String())
	}
	fmt.Println("e aqui - ", req.Body)
	//
	if c.Client == nil {
		return http.DefaultClient.Do(req)
	}
	return c.Client.Do(req)
}
