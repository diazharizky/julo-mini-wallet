package enum

type WithdrawalStatus string

const (
	WithdrawalStatusSuccess WithdrawalStatus = "success"
	WithdrawalStatusFailed  WithdrawalStatus = "failed"
)
