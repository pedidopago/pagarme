package pagarme

import (
	"net/http"
	"os"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

type roundTripFunc func(r *http.Request) (*http.Response, error)

func (s roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return s(r)
}

func TestExceedRateLimitNoHandling(t *testing.T) {
	zl, err := zap.NewDevelopment()
	require.NoError(t, err)

	cfg := &Config{
		Client: &http.Client{
			Timeout: time.Second * 60,
		},
		Trace:           false,
		Logger:          zl,
		HandleRateLimit: false,
	}

	var returnOkAfter time.Time
	var actuallyReturnOkAfter *time.Time

	cfg.Client.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		require.Equal(t, "api.pagar.me", r.URL.Host)
		var statusCode int
		var header = make(http.Header)

		if time.Now().Before(returnOkAfter) {
			statusCode = http.StatusTooManyRequests
			retryAfter := int(returnOkAfter.Sub(time.Now()).Seconds())
			header.Set("Retry-After", strconv.Itoa(retryAfter))
		} else if actuallyReturnOkAfter != nil && time.Now().Before(*actuallyReturnOkAfter) {
			statusCode = http.StatusTooManyRequests
			retryAfter := int(actuallyReturnOkAfter.Sub(time.Now()).Seconds())
			header.Set("Retry-After", strconv.Itoa(retryAfter))
		} else {
			statusCode = http.StatusOK
		}
		return &http.Response{
			Status:     http.StatusText(statusCode),
			StatusCode: statusCode,
			Header:     header,
		}, nil
	})

	returnOkAfter = time.Now().Add(time.Hour * 100)

	resp, err := cfg.Do(http.MethodGet, "/transactions", nil, nil)
	require.NoError(t, err)
	require.Equal(t, http.StatusTooManyRequests, resp.StatusCode)

	cfg.HandleRateLimit = true

	returnOkAfter = time.Now().Add(time.Second * 5)

	resp, err = cfg.Do(http.MethodGet, "/transactions", nil, nil)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	returnOkAfter = time.Now().Add(time.Second * 15)
	cfg.TimeoutRetry = time.Second * 10

	resp, err = cfg.Do(http.MethodGet, "/transactions", nil, nil)
	require.NoError(t, err)
	require.Equal(t, http.StatusTooManyRequests, resp.StatusCode)

	returnOkAfter = time.Now().Add(time.Second * 5)
	{
		v := time.Now().Add(time.Second * 7)
		actuallyReturnOkAfter = &v
	}

	cfg.ExpBackOffParams = &ExponentialBackoffParams{
		MaxRetries:   5,
		IntervalBase: time.Millisecond * 500,
	}

	resp, err = cfg.Do(http.MethodGet, "/transactions", nil, nil)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
}

func getKeysEnv() (apikey, cryptokey string) {
	apikey = os.Getenv("PME_API_KEY")
	cryptokey = os.Getenv("PME_CRYPTO_KEY")
	return
}

func Test_RateLimitPg(t *testing.T) {
	const parallelRequests = 8

	cfg := Default(getKeysEnv())

	var wg sync.WaitGroup
	wg.Add(parallelRequests)
	for i := 0; i < parallelRequests; i++ {
		go func(id int) {
			resp, err := cfg.Do(http.MethodGet, "/balance", nil, nil)
			require.NoError(t, err)
			require.NotNil(t, resp)
			wg.Done()
			t.Logf("%+v\n", struct {
				Id                 int
				HttpStatus         int
				RateLimitRemaining string
			}{
				Id:                 id,
				HttpStatus:         resp.StatusCode,
				RateLimitRemaining: resp.Header.Get("X-RateLimit-Remaining"),
			})
		}(i + 1)
	}
	wg.Wait()
	resp, err := cfg.Do(http.MethodGet, "/balance", nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	t.Log("RateLimitRemaining: " + resp.Header.Get("X-RateLimit-Remaining"))
	require.Equal(t, http.StatusTooManyRequests, resp.StatusCode)
}
