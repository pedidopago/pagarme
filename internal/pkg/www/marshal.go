package www

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/pedidopago/pagarme/v2/pkg/pagarme"
	"go.uber.org/zap"
)

func JSONReader(data interface{}) io.Reader {
	r, w := io.Pipe()
	enc := json.NewEncoder(w)
	go func() {
		enc.Encode(data)
		w.Close()
	}()
	return r
}

func Unmarshal(config *pagarme.Config, resp *http.Response, target any) error {
	if config.SkipUnmarshal {
		return nil
	}
	if config.Trace {
		return unmarshalTrace(config.Logger, resp, target)
	} else {
		return unmarshal(resp, target)
	}
}

func unmarshal(resp *http.Response, target any) error {
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	return dec.Decode(target)
}

func unmarshalTrace(logger *zap.Logger, resp *http.Response, target any) error {
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	logger.Info(string(bs))
	resp.Body.Close()
	return json.Unmarshal(bs, target)
}
