package bank

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/pedidopago/pagarme/internal/pkg/www"
	"github.com/pedidopago/pagarme/pkg/pagarme"
)

func getKeysEnv() (apikey, cryptokey string) {
	apikey = os.Getenv("PME_API_KEY")
	cryptokey = os.Getenv("PME_CRYPTO_KEY")
	return
}

func Test_NewBankAccount(t *testing.T) {
	bank := &pagarme.BankAccount{
		BankCode:       "341",
		Agencia:        "0932",
		AgenciaDV:      "5",
		Conta:          "58054",
		ContaDV:        "1",
		Type:           "conta_corrente",
		DocumentNumber: "92545278157",
		LegalName:      "API BANK ACCOUNT",
	}

	t9 := www.JSONReader(bank)
	gg, _ := ioutil.ReadAll(t9)
	t.Log(string(gg))
	cfg := pagarme.Default(getKeysEnv())
	cfg.Trace = true
	bankAccount := New(cfg)
	z0, z1, err := bankAccount.NewBankAccount(bank)
	if err != nil {
		t.Fatal(err.Error())
	}
	if z0.Code != pagarme.ResponseCodeOk {
		t.Fatal(z0.String())
	}
	t.Log(z1)
}
