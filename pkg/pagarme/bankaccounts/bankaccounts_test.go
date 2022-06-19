package bankaccounts

import (
	"os"
	"testing"

	"github.com/pedidopago/pagarme/v2/pkg/pagarme"
	"github.com/stretchr/testify/require"
)

func getKeysEnv() (apikey, cryptokey string) {
	apikey = os.Getenv("PME_API_KEY")
	cryptokey = os.Getenv("PME_CRYPTO_KEY")
	return
}

func TestAPI_Query(t *testing.T) {
	cfg := pagarme.Default(getKeysEnv())
	api := New(cfg)
	z0, z1, err := api.Query(QueryInput{})
	require.NoError(t, err)
	require.NotNil(t, z0)
	require.Equal(t, pagarme.ResponseCodeOk, z0.Code)
	require.NotNil(t, z1)
	t.Log(z1)
}
