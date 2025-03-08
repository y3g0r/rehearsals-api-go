package domain

type User interface {
	ID() string
	HashedPassword() string
	Email() string
	IsActive() bool
	IsSuperuser() bool
	FullName() string
}

type SimpleUser struct {
	id             string
	hashedPassword string
	email          string
	isActive       bool
	isSuperuser    bool
	fullName       string
}

func NewSimpleUser(id string, hashedPassword string) *SimpleUser {
	return &SimpleUser{
		id:             id,
		hashedPassword: hashedPassword,
	}
}

func NewSimpleUserWithEmail(id string, hashedPassword string, email string) *SimpleUser {
	return &SimpleUser{
		id:             id,
		hashedPassword: hashedPassword,
		email:          email,
		isActive:       true,
	}
}

func (s *SimpleUser) ID() string {
	return s.id
}

func (s *SimpleUser) HashedPassword() string {
	return s.hashedPassword
}

func (s *SimpleUser) Email() string {
	return s.email
}

func (s *SimpleUser) IsActive() bool {
	return s.isActive
}

func (s *SimpleUser) IsSuperuser() bool {
	return s.isSuperuser
}

func (s *SimpleUser) FullName() string {
	return s.fullName
}

var _ User = (*SimpleUser)(nil)
