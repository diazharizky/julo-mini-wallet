package repositories

type userRepository struct{}

func NewUserRepository() userRepository {
	return userRepository{}
}
