package enum

type DepositStatus string

const (
	DepositStatusSuccess DepositStatus = "success"
	DepositStatusFailed  DepositStatus = "failed"
)
