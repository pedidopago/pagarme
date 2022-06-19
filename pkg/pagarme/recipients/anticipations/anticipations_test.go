package anticipations

import (
	"os"
	"testing"
	"time"

	"github.com/pedidopago/pagarme/v2/pkg/pagarme"
	"github.com/stretchr/testify/require"
)

func getKeysEnv() (apikey, cryptokey string) {
	apikey = os.Getenv("PME_API_KEY")
	cryptokey = os.Getenv("PME_CRYPTO_KEY")
	return
}

func TestQuery(t *testing.T) {
	cfg := pagarme.Default(getKeysEnv())
	ant := New(cfg, os.Getenv("PME_RECIPIENT_ID"))

	z0, _, err := ant.Query(QueryInput{
		Count: 1,
	})
	require.NoError(t, err)
	require.NotNil(t, z0)
	require.Equal(t, pagarme.ResponseCodeOk, z0.Code)
}

func Test_GetLimits_Build_Confirm_Cancel_Anticipation(t *testing.T) {
	cfg := pagarme.Default(getKeysEnv())
	ant := New(cfg, os.Getenv("PME_RECIPIENT_ID"))

	// TODO: make sure we get a working day
	payday := time.Now().AddDate(0, 0, 2).UnixNano() / int64(time.Millisecond)
	tf := pagarme.AntTimeframeStart

	z0, limits, err := ant.GetLimits(GetLimitsInput{
		Timeframe:   tf,
		PaymentDate: payday,
	})
	require.NoError(t, err)
	require.NotNil(t, z0)
	require.Equal(t, pagarme.ResponseCodeOk, z0.Code)
	require.NotNil(t, limits)

	req := &pagarme.CreateAnticipation{
		PaymentDate:       payday,
		Timeframe:         tf,
		RequestedAmount:   limits.Minimum.Amount,
		Build:             true,
		AutomaticTransfer: false,
	}

	z0, anticipation, err := ant.NewAnticipation(req)
	require.NoError(t, err)
	require.NotNil(t, z0)
	require.Equal(t, pagarme.ResponseCodeOk, z0.Code)
	require.NotNil(t, anticipation)

	//t.Log(anticipation)

	z0, anticipation, err = ant.ConfirmNewAnticipation(anticipation.ID)
	require.NoError(t, err)
	require.NotNil(t, z0)
	require.Equal(t, pagarme.ResponseCodeOk, z0.Code)
	require.NotNil(t, anticipation)

	//t.Log(anticipation)

	z0, anticipation, err = ant.CancelPendingAnticipation(anticipation.ID)
	require.NoError(t, err)
	require.NotNil(t, z0)
	require.Equal(t, pagarme.ResponseCodeOk, z0.Code)
	require.NotNil(t, anticipation)

	//t.Log(anticipation)
}

func Test_GetLimits_Build_Delete_Anticipation(t *testing.T) {
	cfg := pagarme.Default(getKeysEnv())
	ant := New(cfg, os.Getenv("PME_RECIPIENT_ID"))

	// TODO: make sure we get a working day
	payday := time.Now().AddDate(0, 0, 2).UnixNano() / int64(time.Millisecond)
	tf := pagarme.AntTimeframeStart

	z0, limits, err := ant.GetLimits(GetLimitsInput{
		Timeframe:   tf,
		PaymentDate: payday,
	})
	require.NoError(t, err)
	require.NotNil(t, z0)
	require.Equal(t, pagarme.ResponseCodeOk, z0.Code)
	require.NotNil(t, limits)

	req := &pagarme.CreateAnticipation{
		PaymentDate:       payday,
		Timeframe:         tf,
		RequestedAmount:   limits.Minimum.Amount,
		Build:             true,
		AutomaticTransfer: false,
	}

	z0, anticipation, err := ant.NewAnticipation(req)
	require.NoError(t, err)
	require.NotNil(t, z0)
	require.Equal(t, pagarme.ResponseCodeOk, z0.Code)
	require.NotNil(t, anticipation)

	t.Log(anticipation)

	z0, err = ant.Delete(anticipation.ID)
	require.NoError(t, err)
	require.NotNil(t, z0)
	require.Equal(t, pagarme.ResponseCodeOk, z0.Code)
}
