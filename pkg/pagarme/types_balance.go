package pagarme

import "time"

type BalanceOperation struct {
	Object           string                 `json:"object"`
	Id               int                    `json:"id"`
	Status           BalanceOperationStatus `json:"status"`
	BalanceAmount    int                    `json:"balance_amount"`
	BalanceOldAmount int                    `json:"balance_old_amount"`
	Type             string                 `json:"type"`
	Amount           int                    `json:"amount"`
	Fee              int                    `json:"fee"`
	DateCreated      time.Time              `json:"date_created"`
	MovementObject   map[string]interface{} `json:"movement_object"`
}
