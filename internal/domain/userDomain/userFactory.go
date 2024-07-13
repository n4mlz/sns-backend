package userDomain

import "errors"

var Factory *UserFactory

type UserFactory struct {
	userRepository      *IUserRepository
	userImageRepository *IUserImageRepository
}

func NewUserFactory(userRepository IUserRepository, userImageRepository IUserImageRepository) *UserFactory {
	return &UserFactory{
		userRepository:      &userRepository,
		userImageRepository: &userImageRepository,
	}
}

func SetDefaultUserFactory(userFactory *UserFactory) {
	Factory = userFactory
}

func (uf *UserFactory) SaveUserToRepository(userId UserId, userName UserName, displayName DisplayName, biography Biography, iconUrl ImageUrl, bgImageUrl ImageUrl) (*User, error) {
	if !userName.IsValid() || !displayName.IsValid() || !biography.IsValid() {
		return nil, errors.New("invalid profile")
	}

	user := &User{
		userRepository:      uf.userRepository,
		userImageRepository: uf.userImageRepository,
		UserId:              userId,
		UserName:            userName,
		DisplayName:         displayName,
		Biography:           biography,
		IconUrl:             iconUrl,
		BgImageUrl:          bgImageUrl,
	}

	if (*uf.userRepository).IsExistUserName(userName) {
		oldUser, err := (*uf.userRepository).FindById(userId)
		if err != nil {
			return nil, err
		}

		if oldUser.UserName != userName {
			return nil, errors.New("user name is already used")
		}
	}

	if (*uf.userRepository).IsExistUserId(userId) && !(*uf.userRepository).IsExistUserName(userName) {
		oldUser, err := (*uf.userRepository).FindById(userId)
		if err != nil {
			return nil, err
		}

		if oldUser.UserName != userName {
			user.IconUrl, user.BgImageUrl, err = (*uf.userImageRepository).MoveResources(oldUser, user)
			if err != nil {
				return nil, err
			}
		}
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

func (uf *UserFactory) SaveUserNameToRepository(userId UserId, userName UserName) (*User, error) {
	if !userName.IsValid() {
		return nil, errors.New("invalid user name")
	}

	var displayName DisplayName
	var biography Biography
	var iconUrl ImageUrl
	var bgImageUrl ImageUrl

	user, err := Factory.GetUser(userId)

	if err != nil {
		displayName = DisplayName(userName.String())
		biography = Biography("")
		iconUrl = ""
		bgImageUrl = ""
	} else {
		displayName = user.DisplayName
		biography = user.Biography
		iconUrl = user.IconUrl
		bgImageUrl = user.BgImageUrl
	}

	return uf.SaveUserToRepository(userId, userName, displayName, biography, iconUrl, bgImageUrl)
}

func (uf *UserFactory) SaveUserSettingsToRepository(userId UserId, displayName DisplayName, biography Biography) (*User, error) {
	user, err := Factory.GetUser(userId)
	if err != nil {
		return nil, err
	}

	return uf.SaveUserToRepository(userId, user.UserName, displayName, biography, user.IconUrl, user.BgImageUrl)
}

func (uf *UserFactory) GetUser(userId UserId) (*User, error) {
	user, err := (*uf.userRepository).FindById(userId)
	if err != nil {
		return nil, err
	}

	user.userRepository = uf.userRepository
	user.userImageRepository = uf.userImageRepository

	return user, nil
}

func (uf *UserFactory) GetUsers(userIds []UserId) ([]*User, error) {
	users, err := (*uf.userRepository).FindByIds(userIds)
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		user.userRepository = uf.userRepository
		user.userImageRepository = uf.userImageRepository
	}

	return users, nil
}

func (uf *UserFactory) GetUserByUserName(userName UserName) (*User, error) {
	user, err := (*uf.userRepository).FindByUserName(userName)
	if err != nil {
		return nil, err
	}

	user.userRepository = uf.userRepository
	user.userImageRepository = uf.userImageRepository

	return user, nil
}

func (uf *UserFactory) DeleteUser(user *User) error {
	return (*uf.userRepository).Delete(user)
}
