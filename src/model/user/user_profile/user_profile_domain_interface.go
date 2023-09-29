package user_profile_model

type UserProfileDomainInterface interface {
	GetId() int64
	GetEmail() string
	GetName() string
}

func NewUserProfileDomain(id int64, email, name string) UserProfileDomainInterface {
	return &user_profile_domain{
		id,
		email,
		name,
	}
}
