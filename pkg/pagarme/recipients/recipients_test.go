package recipients

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/pedidopago/pagarme/v2/internal/pkg/www"
	"github.com/pedidopago/pagarme/v2/pkg/pagarme"
)

func getKeysEnv() (apikey, cryptokey string) {
	apikey = os.Getenv("PME_API_KEY")
	cryptokey = os.Getenv("PME_CRYPTO_KEY")
	return
}

func TestPut(t *testing.T) {
	bk0 := pagarme.BankAccountRecipient{
		BankCode:       "341",
		Agencia:        "0932",
		AgenciaDv:      "5",
		Conta:          "58054",
		ContaDv:        "1",
		Type:           "conta_corrente",
		DocumentNumber: "92545278157",
		LegalName:      "API BANK ACCOUNT",
	}
	tr0 := &pagarme.CreateRecipient{
		TransferEnabled:               true,
		TransferDay:                   "0",
		TransferInterval:              "daily",
		AutomaticAnticipationEnabled:  true,
		AnticipatableVolumePercentage: 5,
		PostbackURL:                   "test",
		BankAcc:                       bk0,
		RegisterInfo: pagarme.RegisterInformation{
			Type:           "corporation",
			DocumentNumber: "92545278157",
			Email:          "someone@gmail.com",
			SiteURL:        "http://www.site.com",
			CompanyName:    "Buenos Ayres",
			PhoneNum: []pagarme.PhoneNumbers{
				pagarme.PhoneNumbers{
					Ddd:    "11",
					Number: "987654321",
					Type:   "mobile",
				},
			},
			ManagingPart: []pagarme.ManagingPartners{
				pagarme.ManagingPartners{
					Type:           "individual",
					DocumentNumber: "92545278157",
					Email:          "gustavo@pedidopaog.com.br",
					Name:           "Gustavo Lima",
				},
			},
		},
	}
	t9 := www.JSONReader(tr0)
	gg, _ := ioutil.ReadAll(t9)
	t.Log(string(gg))
	cfg := pagarme.Default(getKeysEnv())
	cfg.Trace = true
	recipients := New(cfg)
	z0, z1, err := recipients.NewRecipient(tr0)
	if err != nil {
		t.Fatal(err.Error())
	}
	if z0.Code != pagarme.ResponseCodeOk {
		t.Fatal(z0.String())
	}
	t.Log(z1)
}
