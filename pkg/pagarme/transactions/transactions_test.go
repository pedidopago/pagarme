package transactions

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/pedidopago/pagarme/internal/pkg/www"
	"github.com/pedidopago/pagarme/pkg/pagarme"
)

func getKeysEnv() (apikey, cryptokey string) {
	apikey = os.Getenv("PME_API_KEY")
	cryptokey = os.Getenv("PME_CRYPTO_KEY")
	return
}

func TestPut(t *testing.T) {
	adr0 := &pagarme.Address{
		Zipcode:       "01232000",
		Street:        "Alameda Barros",
		StreetNumber:  "400",
		Complementary: "Apto 1000",
		State:         "SP",
		City:          "Sao Paulo",
		Country:       "br",
	}
	tr0 := &pagarme.Transaction{
		Amount:             4569, // R$ 45.69
		CardHolderName:     "John Smith",
		CardExpirationDate: "1250",
		CardNumber:         "4485038201734841", // test card generated (not real)
		CardCVV:            "111",
		PaymentMethod:      pagarme.PaymentCreditCard,
		Async:              false,
		Installments:       1,
		SoftDescriptor:     "teste",
		Capture:            "true",
		Customer: &pagarme.Customer{
			ExternalID: "TESTEJOHN1",
			Name:       "John Smith",
			Country:    pagarme.Brazil,
			Type:       pagarme.CustomerIndividual,
			Email:      "john.smith@gmail.com",
			PhoneNumbers: []string{
				"+5511900000000",
			},
			Documents: []*pagarme.Document{
				&pagarme.Document{
					Type:   pagarme.DocCPF,
					Number: "82700962028",
				},
			},
		},
		Billing: &pagarme.Billing{
			Name:    "John Smith",
			Address: adr0,
		},
		Shipping: &pagarme.Shipping{
			Name:         "John Smith",
			Address:      adr0,
			DeliveryDate: pagarme.YYYYMMDDFromTime(time.Now().AddDate(0, 0, 2)),
		},
		Items: []*pagarme.Item{
			&pagarme.Item{
				ID:        "test_1",
				Title:     "Vitamina C 100mg (100 cap)",
				UnitPrice: 4568,
				Tangible:  true,
				Quantity:  1,
			},
		},
	}
	t9 := www.JSONReader(tr0)
	gg, _ := ioutil.ReadAll(t9)
	t.Log(string(gg))
	cfg := pagarme.Default(getKeysEnv())
	//cfg.Trace = true
	trs := New(cfg)
	z0, z1, err := trs.Put(tr0)
	if err != nil {
		t.Fatal(err.Error())
	}
	if z0.Code != pagarme.ResponseCodeOk {
		t.Fatal(z0.String())
	}
	t.Log(z1)
}

func TestPutRefuse(t *testing.T) {
	adr0 := &pagarme.Address{
		Zipcode:       "01232000",
		Street:        "Alameda Barros",
		StreetNumber:  "400",
		Complementary: "Apto 1000",
		State:         "SP",
		City:          "Sao Paulo",
		Country:       "br",
	}
	tr0 := &pagarme.Transaction{
		Amount:             4569, // R$ 45.69
		CardHolderName:     "John Smith",
		CardExpirationDate: "1250",
		CardNumber:         "4485038201734841", // test card generated (not real)
		CardCVV:            "611",              // test env, CVV starting with 6 = authorization error
		PaymentMethod:      pagarme.PaymentCreditCard,
		Async:              false,
		Installments:       1,
		SoftDescriptor:     "teste",
		Capture:            "true",
		Customer: &pagarme.Customer{
			ExternalID: "TESTEJOHN1",
			Name:       "John Smith",
			Country:    pagarme.Brazil,
			Type:       pagarme.CustomerIndividual,
			Email:      "john.smith@gmail.com",
			PhoneNumbers: []string{
				"+5511900000000",
			},
			Documents: []*pagarme.Document{
				&pagarme.Document{
					Type:   pagarme.DocCPF,
					Number: "82700962028",
				},
			},
		},
		Billing: &pagarme.Billing{
			Name:    "John Smith",
			Address: adr0,
		},
		Shipping: &pagarme.Shipping{
			Name:         "John Smith",
			Address:      adr0,
			DeliveryDate: pagarme.YYYYMMDDFromTime(time.Now().AddDate(0, 0, 2)),
		},
		Items: []*pagarme.Item{
			&pagarme.Item{
				ID:        "test_1",
				Title:     "Vitamina C 100mg (100 cap)",
				UnitPrice: 4568,
				Tangible:  true,
				Quantity:  1,
			},
		},
	}
	t9 := www.JSONReader(tr0)
	gg, _ := ioutil.ReadAll(t9)
	t.Log(string(gg))
	cfg := pagarme.Default(getKeysEnv())
	//cfg.Trace = true
	trs := New(cfg)
	z0, z1, err := trs.Put(tr0)
	if err != nil {
		t.Fatal(err.Error())
	}
	if z0.Code != pagarme.ResponseCodeOk {
		t.Fatal(z0.String())
	}
	t.Log(z1)
}

func TestRefundCard(t *testing.T) {
	adr0 := &pagarme.Address{
		Zipcode:       "01232000",
		Street:        "Alameda Barros",
		StreetNumber:  "400",
		Complementary: "Apto 1000",
		State:         "SP",
		City:          "Sao Paulo",
		Country:       "br",
	}
	tr0 := &pagarme.Transaction{
		Amount:             4569, // R$ 45.69
		CardHolderName:     "John Smith",
		CardExpirationDate: "1250",
		CardNumber:         "4485038201734841", // test card generated (not real)
		CardCVV:            "123",              // test env, CVV starting with 6 = authorization error
		PaymentMethod:      pagarme.PaymentCreditCard,
		Async:              false,
		Installments:       1,
		SoftDescriptor:     "teste",
		Capture:            "true",
		Customer: &pagarme.Customer{
			ExternalID: "TESTEJOHN1",
			Name:       "John Smith",
			Country:    pagarme.Brazil,
			Type:       pagarme.CustomerIndividual,
			Email:      "john.smith@gmail.com",
			PhoneNumbers: []string{
				"+5511900000000",
			},
			Documents: []*pagarme.Document{
				&pagarme.Document{
					Type:   pagarme.DocCPF,
					Number: "82700962028",
				},
			},
		},
		Billing: &pagarme.Billing{
			Name:    "John Smith",
			Address: adr0,
		},
		Shipping: &pagarme.Shipping{
			Name:         "John Smith",
			Address:      adr0,
			DeliveryDate: pagarme.YYYYMMDDFromTime(time.Now().AddDate(0, 0, 2)),
		},
		Items: []*pagarme.Item{
			&pagarme.Item{
				ID:        "test_1",
				Title:     "Vitamina C 100mg (100 cap)",
				UnitPrice: 4568,
				Tangible:  true,
				Quantity:  1,
			},
		},
	}
	t9 := www.JSONReader(tr0)
	gg, _ := ioutil.ReadAll(t9)
	t.Log(string(gg))
	cfg := pagarme.Default(getKeysEnv())
	//cfg.Trace = true
	trs := New(cfg)
	z0, z1, err := trs.Put(tr0)
	if err != nil {
		t.Fatal(err.Error())
	}
	if z0.Code != pagarme.ResponseCodeOk {
		t.Fatal(z0.String())
	}
	if z1.Status != pagarme.TrPaid {
		t.Fatal(z1.Status, "!=", pagarme.TrPaid)
	}
	//t.Log(z1)
	z0, z1, err = trs.Refund(z1.ID, &RefundInput{})
	if err != nil {
		t.Fatal(err.Error())
	}
	if z1.Status != pagarme.TrRefunded {
		t.Fatal(z1.Status, "!=", pagarme.TrPaid, z1)
	}
}

func TestPutBoleto(t *testing.T) {
	adr0 := &pagarme.Address{
		Zipcode:       "01232000",
		Street:        "Alameda Barros",
		StreetNumber:  "400",
		Complementary: "Apto 1000",
		State:         "SP",
		City:          "Sao Paulo",
		Country:       "br",
	}
	tr0 := &pagarme.Transaction{
		Amount:               15669, // R$ 156.69
		CardHolderName:       "John Smith",
		PaymentMethod:        pagarme.PaymentBoleto,
		BoletoExpirationDate: time.Now().AddDate(0, 0, 2).Format("2006-01-02"),
		BoletoInstructions:   "test billet please ignore",
		Async:                false,
		Installments:         1,
		SoftDescriptor:       "teste boleto",
		Capture:              "true",
		Customer: &pagarme.Customer{
			ExternalID: "TESTEJOHN1",
			Name:       "John Smith",
			Country:    pagarme.Brazil,
			Type:       pagarme.CustomerIndividual,
			Email:      "john.smith@gmail.com",
			PhoneNumbers: []string{
				"+5511900000000",
			},
			Documents: []*pagarme.Document{
				&pagarme.Document{
					Type:   pagarme.DocCPF,
					Number: "82700962028",
				},
			},
		},
		Billing: &pagarme.Billing{
			Name:    "John Smith",
			Address: adr0,
		},
		Shipping: &pagarme.Shipping{
			Name:         "John Smith",
			Address:      adr0,
			DeliveryDate: pagarme.YYYYMMDDFromTime(time.Now().AddDate(0, 0, 2)),
		},
		Items: []*pagarme.Item{
			&pagarme.Item{
				ID:        "test_1",
				Title:     "Vitamina C 100mg (100 cap)",
				UnitPrice: 4568,
				Tangible:  true,
				Quantity:  1,
			},
		},
	}
	t9 := www.JSONReader(tr0)
	gg, _ := ioutil.ReadAll(t9)
	t.Log(string(gg))
	cfg := pagarme.Default(getKeysEnv())
	//cfg.Trace = true
	trs := New(cfg)
	z0, z1, err := trs.Put(tr0)
	if err != nil {
		t.Fatal(err.Error())
	}
	if z0.Code != pagarme.ResponseCodeOk {
		t.Fatal(z0.String())
	}
	t.Log(z1)
}
