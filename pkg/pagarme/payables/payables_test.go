package payables

import (
	"os"
	"testing"
	"time"

	"github.com/pedidopago/pagarme/pkg/pagarme"
)

func getKeysEnv() (apikey, cryptokey string) {
	apikey = os.Getenv("PME_API_KEY")
	cryptokey = os.Getenv("PME_CRYPTO_KEY")
	return
}

func TestQuery(t *testing.T) {
	cfg := pagarme.Default(getKeysEnv())
	api := New(cfg)
	pmeresp, items, err := api.Query(nil)
	if err != nil {
		t.Fatal(err)
	}
	if pmeresp.Code != pagarme.ResponseCodeOk {
		t.Fatal(pmeresp.String())
	}
	t.Log(items)
	t.Log(len(items))
}

func TestQuery2(t *testing.T) {
	cfg := pagarme.Default(getKeysEnv())
	api := New(cfg)
	tt, err := time.Parse(time.RFC3339Nano, "2020-04-05T18:22:09.056Z")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tt.Format(time.RFC3339Nano))
	pmeresp, items, err := api.Query((&QueryInput{}).CreatedAt(pagarme.QueryOpGreaterOrEqualThan, tt).CreatedAt(pagarme.QueryOpLessThan, tt.Add(time.Second)).Count(100))
	if err != nil {
		t.Fatal(err)
	}
	if pmeresp.Code != pagarme.ResponseCodeOk {
		t.Fatal(pmeresp.String())
	}
	t.Log(items)
	t.Log(len(items))
}

func TestQuery3(t *testing.T) {
	cfg := pagarme.Default(getKeysEnv())
	api := New(cfg)
	pmeresp, items, err := api.Query((&QueryInput{}).TransactionID("6142879").Count(100))
	if err != nil {
		t.Fatal(err)
	}
	if pmeresp.Code != pagarme.ResponseCodeOk {
		t.Fatal(pmeresp.String())
	}
	t.Log(items)
	t.Log(len(items))
}
