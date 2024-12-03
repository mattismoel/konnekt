package sqlite

type userService struct {
	repo *Repository
}

func NewUserService(repo *Repository) *userService {
	return &userService{repo: repo}
}
