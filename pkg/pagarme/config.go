package pagarme

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
	"time"

	"go.uber.org/zap"
)

// Config holds the config of pagarme
type Config struct {
	Apikey        string
	Cryptokey     string
	Version       string
	Client        *http.Client
	Trace         bool
	SkipUnmarshal bool
	Logger        *zap.Logger

	HandleRateLimit  bool
	TimeoutRetry     time.Duration
	ExpBackOffParams *ExponentialBackoffParams
}

type ExponentialBackoffParams struct {
	MaxRetries   int
	IntervalBase time.Duration
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
		Version:   "2017-08-28",
		Client: &http.Client{
			Timeout: time.Second * 60,
		},
		Trace:         false,
		SkipUnmarshal: false,
		Logger:        zl,
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
	req.Header.Set("X-PagarMe-Version", c.Version)
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
	//
	if c.Client == nil {
		return http.DefaultClient.Do(req)
	}
	resp, err := c.Client.Do(req)

	if err == nil && resp != nil && resp.StatusCode == http.StatusTooManyRequests && c.HandleRateLimit {
		return c.rateLimitReachedHandler(req, resp, err)
	}

	return resp, err
}

func (c *Config) rateLimitReachedHandler(req *http.Request, inresp *http.Response, inerr error) (*http.Response, error) {
	c.Logger.Debug("rateLimitReachedHandler")
	retryAfterSec, _ := strconv.Atoi(inresp.Header.Get("Retry-After"))
	retryAfter := time.Duration(retryAfterSec) * time.Second
	if c.TimeoutRetry != 0 {
		if retryAfter > c.TimeoutRetry {
			c.Logger.Debug("rateLimitReachedHandler: timeout retry non sufficient (1)")
			return inresp, inerr
		}
	}

	retryStartedAt := time.Now()

	time.Sleep(retryAfter)
	if c.ExpBackOffParams == nil {
		return c.Client.Do(req)
	}

	var resp = inresp
	var rerr = inerr
	var intervalBase = float64(c.ExpBackOffParams.IntervalBase.Milliseconds())
	for attempt := 1; attempt <= c.ExpBackOffParams.MaxRetries; attempt++ {
		var waitFor = time.Millisecond * time.Duration(intervalBase*math.Pow(2, float64(attempt-1)))
		if c.TimeoutRetry != 0 {
			if time.Now().Sub(retryStartedAt)+waitFor > c.TimeoutRetry {
				c.Logger.Debug("rateLimitReachedHandler: timeout retry non sufficient (2)")
				return resp, inerr
			}
		}
		c.Logger.Debug(fmt.Sprint("attempt ", attempt, "sleeping for ", waitFor))
		time.Sleep(waitFor)
		resp, rerr = c.Client.Do(req)
		if rerr != nil {
			return nil, rerr
		}
		if resp != nil && resp.StatusCode != http.StatusTooManyRequests {
			break
		}
	}

	return resp, rerr
}
