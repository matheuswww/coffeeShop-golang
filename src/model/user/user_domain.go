package user_model

type userDomain struct {
	id string
	email string
	name string
	password string
}

func (ud *userDomain) GetEmail() string {
	return ud.email
}

func (ud *userDomain) GetName() string {
	return ud.name
}

func (ud *userDomain) GetPassword() string {
	return ud.password
}

func (ud *userDomain) GetId() string {
	return ud.id
}

func (ud *userDomain) SetId(id string) {
	ud.id = id
}