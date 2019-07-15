package interfaces

import (
	repo "github.com/salihkemaloglu/gignoxqc-beta-001/repositories"
)

//IUserRepository ..
type IUserRepository interface {
	Login() (*repo.User, error)
}
