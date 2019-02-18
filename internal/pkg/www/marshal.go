package www

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

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

func Unmarshal(resp *http.Response, target interface{}) error {
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	return dec.Decode(target)
}

func UnmarshalTrace(logger *zap.Logger, resp *http.Response, target interface{}) error {
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	logger.Info(string(bs))
	resp.Body.Close()
	return json.Unmarshal(bs, target)
}
