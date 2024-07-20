package userDomain

var Service *UserService

type UserService struct {
	UserRepository *IUserRepository
}

func NewUserService(userRepository IUserRepository) *UserService {
	return &UserService{UserRepository: &userRepository}
}

func SetDefaultUserService(userService *UserService) {
	Service = userService
}

func (us *UserService) UserNameToUserId(userName UserName) (UserId, error) {
	user, err := (*us.UserRepository).FindByUserName(userName)
	if err != nil {
		return "", err
	}
	return user.UserId, nil
}

func (us *UserService) IsExistUserName(userName UserName) bool {
	return (*us.UserRepository).IsExistUserName(userName)
}

func (us *UserService) ExtractMutualUsers(sourceUser *User, targetUsers []*User) []*User {
	var result []*User
	for _, targetUser := range targetUsers {
		if sourceUser.IsMutual(targetUser) {
			result = append(result, targetUser)
		}
	}
	return result
}
