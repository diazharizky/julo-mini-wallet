package repositories

type walletRepository struct{}

func NewWalletRepository() walletRepository {
	return walletRepository{}
}
