package pagarme

type DocumentType string

const (
	DocCPF      DocumentType = "cpf"
	DocCNPJ     DocumentType = "cnpj"
	DocPassport DocumentType = "passport"
)

type CustomerType string

const (
	CustomerIndividual  CustomerType = "individual"
	CustomerCorporation CustomerType = "corporation"
)

type PaymentMethod string

const (
	PaymentCreditCard PaymentMethod = "credit_card"
	PaymentBoleto     PaymentMethod = "boleto"
	PaymentPix        PaymentMethod = "pix"
)

type ResponseCode int

const (
	// ResponseCodeOk 200 - Tudo ocorreu como deveria e sua requisição foi processada com sucesso.
	ResponseCodeOk ResponseCode = 200
	// ResponseCodeMissingParameter 400 - Algum parâmetro obrigatório não foi passado, ou os parâmetros passados não estão corretos.
	ResponseCodeMissingParameter ResponseCode = 400
	ResponseCodeAuthError        ResponseCode = 401
	ResponseCodeNotFound         ResponseCode = 404
	ResponseCodeInternalError    ResponseCode = 500
)

type TrStatus string

const (
	TrProcessing     TrStatus = "processing"
	TrAuthorized     TrStatus = "authorized"
	TrPaid           TrStatus = "paid"
	TrRefunded       TrStatus = "refunded"
	TrWaitingPayment TrStatus = "waiting_payment"
	TrPendingRefund  TrStatus = "pending_refund"
	TrRefused        TrStatus = "refused"
)

type PayableStatus string

const (
	PbsPaid         PayableStatus = "paid"
	PbsWaitingFunds PayableStatus = "waiting_funds"
)

type PayableType string

const (
	PtChargeback       PayableType = "chargeback"
	PtRefund           PayableType = "refund"
	PtChargebackRefund PayableType = "chargeback_refund"
	PtCredit           PayableType = "credit"
)

type BankAccountType string

const (
	BnkAccContaCorrente         BankAccountType = "conta_corrente"
	BnkAccContaPoupanca         BankAccountType = "conta_poupanca"
	BnkAccContaCorrenteConjunta BankAccountType = "conta_corrente_conjunta"
	BnkAccContaPoupancaConjunta BankAccountType = "conta_poupanca_conjunta"
)
