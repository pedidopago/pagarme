package pagarme

import "time"

type BalanceField struct {
	Amount int `json:"amount"`
}

type Balance struct {
	Object       string       `json:"object"`
	WaitingFunds BalanceField `json:"waiting_funds"`
	Available    BalanceField `json:"available"`
	Transferred  BalanceField `json:"transferred"`
}

type BalanceOperation struct {
	Object           string                 `json:"object"`
	Id               int                    `json:"id"`
	Status           BalanceOperationStatus `json:"status"`
	BalanceAmount    int                    `json:"balance_amount"`
	BalanceOldAmount int                    `json:"balance_old_amount"`
	Type             BalanceOperationType   `json:"type"`
	Amount           int                    `json:"amount"`
	Fee              int                    `json:"fee"`
	DateCreated      time.Time              `json:"date_created"`
	MovementObject   interface{}            `json:"movement_object"`
}
