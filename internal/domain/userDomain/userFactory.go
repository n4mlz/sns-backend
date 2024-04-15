package userDomain

import "errors"

var Factory *UserFactory

type UserFactory struct {
	userRepository *IUserRepository
}

func NewUserFactory(userRepository IUserRepository) *UserFactory {
	return &UserFactory{
		userRepository: &userRepository,
	}
}

func SetDefaultUserFactory(userFactory *UserFactory) {
	Factory = userFactory
}

func (uf *UserFactory) SaveUserToRepository(userId UserId, userName UserName, displayName DisplayName, biography Biography) (*User, error) {
	if !userName.IsValid() || !displayName.IsValid() || !biography.IsValid() {
		return nil, errors.New("invalid profile")
	}

	user := &User{
		UserRepository: uf.userRepository,
		UserId:         userId,
		UserName:       userName,
		DisplayName:    displayName,
		Biography:      biography,
	}

	err := (*uf.userRepository).Save(user)
	if err != nil {
		return nil, err
	}

	if !user.IsFollowing(user) {
		err = user.Follow(user)
		if err != nil {
			return nil, err
		}
	}

	return user, nil
}

func (uf *UserFactory) GetUser(userId UserId) (*User, error) {
	user, err := (*uf.userRepository).FindById(userId)
	if err != nil {
		return nil, err
	}

	user.UserRepository = uf.userRepository

	return user, nil
}

func (uf *UserFactory) GetUsers(userIds []UserId) ([]*User, error) {
	users, err := (*uf.userRepository).FindByIds(userIds)
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		user.UserRepository = uf.userRepository
	}

	return users, nil
}

func (uf *UserFactory) GetUserByUserName(userName UserName) (*User, error) {
	user, err := (*uf.userRepository).FindByUserName(userName)
	if err != nil {
		return nil, err
	}

	user.UserRepository = uf.userRepository

	return user, nil
}