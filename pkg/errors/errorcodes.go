package errors

import (
	"strconv"

	"github.com/pedidopago/pagarme/pkg/util"
)

var errAgent = map[string]string{
	"0000": "Número do cartão inválido.",
	"1000": "Transação não autorizada.",
	"1001": "Cartão vencido ou data de vencimento incorreta.",
	"1002": "Cartão não autorizado. Solicitar contato com central do cartão.",
	"1003": "Cartão não autorizado. Solicitar contato com central do cartão.",
	"1004": "Cartão não autorizado. Solicitar contato com central do cartão.",
	"1005": "Cartão não autorizado. Solicitar contato com central do cartão.",
	"1006": "O portador do cartão excedeu o número de tentativas de senha",
	"1007": "Transação recusada pelo banco, oriente o portador a contatar o banco/emissor",
	"1008": "Transação recusada pelo banco, oriente o portador a contatar o banco/emissor",
	"1009": "Número de afiliação do estabelecimento comercial está inválido",
	"1010": "O valor solicitado para captura não é válido",
	"1011": "Número do cartão invalido.",
	"1012": "Senha necessária para efetuar a transação",
	"1013": "Taxa da adquirente inválida.",
	"1014": "Conta requerida não encontrada",
	"1015": "Função não suportada",
	"1016": "Saldo insuficiente.",
	"1017": "Senha incorreta.",
	"1018": "Não há registros deste cartão.",
	"1019": "Sistema de prevenção do banco não autorizou a compra.",
	"1020": "Transação não permitida para este terminal.",
	"1021": "Limite de saque excedido.",
	"1022": "Transação não autorizada por violação de segurança.",
	"1023": "Número de saques excedido.",
	"1024": "Transação não autorizada por violação das leis.",
	"1025": "Cartão desabilitado. Oriente o usuário a contatar o banco/emissor do cartão.",
	"1026": "Senha de bloqueio inválida.",
	"1027": "Senha excedeu o limite de caracteres.",
	"1028": "Erro na sincronização da senha.",
	"1029": "Cartão com suspeita de falsificação",
	"1030": "Moeda não suportada pelo emissor",
	"1031": "Transação não autorizada, com disputa de taxas.",
	"1032": "Cartão bloqueado por perda ou roubo.",
	"1033": "Período para autorização solicitada não aceito",
	"1040": "Cartão não autorizado.",
	"1045": "Código de segurança inválido",
	"1047": "Senha não autorizada.",
	"1048": "Nova senha inválida",
	"1058": "Informações pessoais não encontradas",
	"2000": "Transação recusada pelo banco, oriente o portador a contatar o banco/emissor do cartão",
	"2001": "Cartão vencido ou data de vencimento incorreta. Oriente o usuário a contatar o banco/emissor do cartão",
	"2002": "Transação com suspeita de fraude",
	"2004": "Cartão com restrição. Oriente o usuário a contatar o banco/emissor do cartão",
	"2005": "Estabelecimento, favor contatar a área de risco da operadora de cartão",
	"2006": "O número de tentativas de senha foi excedido",
	"2007": "Cartão com restrição. Oriente o usuário a contatar o banco/emissor do cartão",
	"2008": "Cartão bloqueado por perda ou roubo",
	"2009": "Cartão bloqueado por perda ou roubo",
	"2010": "Cartão com suspeita de falsificação",
	"9103": "Retente a transação.",
	"9124": "Número do cartão (ou código de segurança) inválido.",
}

var errClient = map[string]string{
	"0000": "Número do cartão invalido. Verifique o número do cartão e digite-o novamente. Caso o erro se repita, por favor, entre em contato com a central de seu cartão.",
	"1000": "Transação não autorizada. Por favor, entre em contato com a central de seu cartão.",
	"1001": "Cartão vencido ou data de vencimento incorreta. Verifique a data de vencimento e digite novamente. Caso o cartão esteja vencido, por favor, utilize um novo cartão.",
	"1002": "Cartão não autorizado. Por favor, entre em contato com a central de seu cartão.",
	"1003": "Cartão não autorizado. Por favor, entre em contato com a central de seu cartão.",
	"1004": "Cartão não autorizado. Por favor, entre em contato com a central de seu cartão.",
	"1005": "Cartão não autorizado. Por favor, entre em contato com a central de seu cartão.",
	"1006": "Você excedeu o número de tentativas de senha. Por favor, entre em contato com a central de seu cartão.",
	"1007": "Transação não autorizada. Por favor, entre em contacto com a central de seu cartão.",
	"1008": "Transação não autorizada. Por favor, entre em contacto com a central de seu cartão.",
	"1009": "Adiliação do estabelecimento inválida. Por favor, solicite ao gerente do estabelecimento contatar a operadora do cartão.",
	"1010": "O valor solicitado na transação parece não ser válido. Por favor, verifique-o e tente novamente.",
	"1011": "Número do cartão invalido. Por favor, verifique o número do cartão e digite-o novamente. Caso o erro se repita, por favor, entre em contato com a central de seu cartão.",
	"1012": "Por favor, digite a senha para efetuar a transação.",
	"1013": "Por favor, solicite ao gerente do estabelecimento contatar a operadora do cartão.",
	"1014": "Por favor, solicite ao gerente do estabelecimento contatar a operadora do cartão.",
	"1015": "Por favor, solicite ao gerente do estabelecimento contatar a operadora do cartão.",
	"1016": "Verifique saldo do cartão antes de realizar nova tentativa de compra. Caso tenha dúvidas, por favor, entre em contato com a central de seu cartão.",
	"1017": "Senha incorreta. Por favor, verifique a senha e digite novamente.",
	"1018": "Transação não autorizada. Por favor, entre em contato com a central de seu cartão.",
	"1019": "Transação recusada. Por favor, entre em contato com a central de seu cartão.",
	"1020": "Solicite ao gerente do estabelecimento contatar a operadora do cartão.",
	"1021": "Limite de saque excedido. Por favor, entre em contato com a central de seu cartão.",
	"1022": "Transação não autorizada. Por favor, entre em contato com a central de seu cartão.",
	"1023": "Número de saques excedidos. Por favor, entre em contato com a central de seu cartão.",
	"1024": "Transação não autorizada. Por favor, entre em contato com a central de seu cartão.",
	"1025": "Cartão não habilitado. Por favor, entre em contato com a central de seu cartão.",
	"1026": "Senha incorreta. Por favor, verifique a senha e digite novamente.",
	"1027": "Senha incorreta. Por favor, verifique a senha e digite novamente.",
	"1028": "Senha não verificada. Por favor, digite a senha novamente.",
	"1029": "Cartão não autorizado. Por favor, entre em contato com a central de seu cartão.",
	"1030": "Moeda não suportada. Por favor, entre em contato com a nossa gerência.",
	"1031": "Transação não autorizada. Por favor, entre em contato com a nossa gerência.",
	"1032": "Cartão não autorizado. Por favor, entre em contato com a central de seu cartão.",
	"1033": "Transação não autorizada. Por favor, entre em contato com a nossa gerência.",
	"1040": "Cartão não autorizado. Por favor, entre em contato com a central de seu cartão.",
	"1045": "Código de segurança incorreto.  Por favor, verifique a senha e digite novamente.",
	"1047": "Senha incorreta. Por favor, entre em contato com a central de seu cartão.",
	"1048": "Senha incorreta. Por favor, entre em contato com a central de seu cartão.",
	"1058": "Informações pessoais não encontradas. Verifique que seu nome, endereço, CPF, email e número de celular foram preenchidos corretamente.",
	"2000": "Transação não autorizada. Por favor, entre em contato com o banco.",
	"2001": "Cartão vencido ou data de vencimento incorreta. Verifique a data de vencimento e digite novamente. Caso o cartão esteja vencido, por favor, utilize um novo cartão.",
	"2002": "Transação não autorizado. Por favor, entre em contato com a central de seu cartão.",
	"2004": "Cartão não autorizado. Por favor, entre em contato com a central de seu cartão.",
	"2005": "Não autorizado. Por favor, solicite ao gerente do estabelecimento contatar a operadora do cartão.",
	"2006": "Você excedeu o número de tentativas de senha. Por favor, entre em contato com a central de seu cartão.",
	"2007": "Cartão não autorizado. Por favor, entre em contato com a central de seu cartão.",
	"2008": "Cartão bloqueado. Por favor, entre em contato com a central de seu cartão.",
	"2009": "Cartão bloqueado. Por favor, entre em contato com a central de seu cartão.",
	"2010": "Cartão não autorizado. Por favor, entre em contato com a central de seu cartão.",
	"9103": "Esta tentativa de transação não foi efetivada. Por favor, retente a transação.",
	"9124": "Código de segurança inválido.",
}

var errPagarme = map[string]string{
	"0000": "Número do cartão invalido.",
	"1000": "Transação não autorizada.",
	"1001": "Verificar vencimento.",
	"1002": "Cartão não autorizado.",
	"1003": "Cartão não autorizado.",
	"1004": "Cartão não autorizado.",
	"1005": "Cartão não autorizado.",
	"1006": "Tentativas de senha excedidas.",
	"1007": "Transação não autorizada.",
	"1008": "Transação não autorizada.",
	"1009": "Afiliação do estabelecimento inválida.",
	"1010": "O valor não é valido.",
	"1011": "Número do cartão invalido.",
	"1012": "Senha não preenchida.",
	"1013": "Taxa da adquirente inválida.",
	"1014": "Conta não encontrada.",
	"1015": "Função não suportada.",
	"1016": "Saldo insuficiente.",
	"1017": "Senha incorreta.",
	"1018": "Não há registros deste cartão",
	"1019": "Banco não autorizou.",
	"1020": "Terminal não permitido.",
	"1021": "Limite de saque excedido.",
	"1022": "Transação não autorizada.",
	"1023": "Número de saques excedido",
	"1024": "Transação não autorizada.",
	"1025": "Cartão desabilitado.",
	"1026": "Senha de bloqueio inválida.",
	"1027": "Senha incorreta.",
	"1028": "Senha não verificada.",
	"1029": "Cartão não autorizado.",
	"1030": "Moeda não suportada.",
	"1031": "Transação não autorizada.",
	"1032": "Cartão boqueado.",
	"1033": "Período para autorização não aceito.",
	"1040": "Cartão não autorizado.",
	"1045": "Código de segurança inválido",
	"1047": "Senha não autorizada.",
	"1048": "Senha não autorizada.",
	"1058": "Informações pessoais não encontradas.",
	"2000": "Transação não autorizada.",
	"2001": "Verificar vencimento.",
	"2002": "Transação não autorizada.",
	"2004": "Cartão não autorizado.",
	"2005": "Transação não autorizada.",
	"2006": "Tentativas de senha excedidas.",
	"2007": "Cartão bloqueado.",
	"2008": "Cartão bloqueado.",
	"2009": "Cartão bloqueado.",
	"2010": "Cartão bloqueado.",
	"9103": "Retente a transação.",
	"9124": "Código de segurança inválido.",
}

// PaymentError is a container of error messages.
type PaymentError struct {
	Code       string `json:"code"`
	MsgPagarme string `json:"msg_pagarme"`
	MsgClient  string `json:"msg_client"`
	MsgAgent   string `json:"msg_agent"`
}

// GetPaymentError retrieves the error messages of a payment error code.
func GetPaymentError(code string) PaymentError {
	if pcode, ok := errPagarme[code]; ok {
		return PaymentError{
			Code:       code,
			MsgPagarme: pcode,
			MsgClient:  errClient[code],
			MsgAgent:   errAgent[code],
		}
	}
	return PaymentError{
		Code:       "XXXX",
		MsgPagarme: "Erro desconhecido (" + code + ")",
		MsgAgent:   "Erro desconhecido (PME" + code + ")",
		MsgClient:  "Erro. Por favor, entre em contato com a nossa gerência. (PME" + code + ")",
	}
}

// GetPaymentErrorInt retrieves the error messages of a payment error code.
func GetPaymentErrorInt(code int) PaymentError {
	cstr := util.Lpad(strconv.Itoa(code), "0", 4)
	return GetPaymentError(cstr)
}
