package service

type CreateUserParams struct {
	Email       string
	Password    string
	FullName    string
	IsActive    bool
	IsSuperuser bool
}

type AuthenticationParams struct {
	Username string
	Password string
}
