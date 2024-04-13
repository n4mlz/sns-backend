package userDomain

import "errors"

var Factory *UserFactory

type UserFactory struct {
	userRepository IUserRepository
}

func NewUserFactory(userRepository IUserRepository) *UserFactory {
	return &UserFactory{
		userRepository: userRepository,
	}
}

func SetDefaultUserFactory(userFactory *UserFactory) {
	Factory = userFactory
}

func (uf *UserFactory) NewUser(userId UserId, userName UserName, displayName DisplayName, biography string) (*User, error) {
	if !userName.IsValid() || !displayName.IsValid() {
		return nil, errors.New("invalid user name or display name")
	}

	return &User{
		UserRepository: uf.userRepository,
		UserId:         userId,
		UserName:       userName,
		DisplayName:    displayName,
		Biography:      biography,
	}, nil
}

func (uf *UserFactory) GetUser(userId UserId) (*User, error) {
	user, err := uf.userRepository.FindById(userId)
	if err != nil {
		return nil, err
	}

	return &User{
		UserRepository: uf.userRepository,
		UserId:         user.UserId,
		UserName:       user.UserName,
		DisplayName:    user.DisplayName,
		Biography:      user.Biography,
		CreatedAt:      user.CreatedAt,
	}, nil
}

func (uf *UserFactory) GetUserByUserName(userName UserName) (*User, error) {
	user, err := uf.userRepository.FindByUserName(userName)
	if err != nil {
		return nil, err
	}

	return &User{
		UserRepository: uf.userRepository,
		UserId:         user.UserId,
		UserName:       user.UserName,
		DisplayName:    user.DisplayName,
		Biography:      user.Biography,
		CreatedAt:      user.CreatedAt,
	}, nil
}
