package user_profile_model

type user_profile_domain struct {
	id                int64
	email             string
	name              string
}

func (ud *user_profile_domain) GetId() int64 {
	return ud.id
}

func (ud *user_profile_domain) GetEmail() string {
	return ud.email
}

func (ud *user_profile_domain) GetName() string {
	return ud.name
}