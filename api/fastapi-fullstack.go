//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=types.cfg.yml fastapi-fullstack-openapi.json
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=server.cfg.yml fastapi-fullstack-openapi.json

package api

import (
	"context"
	"github.com/y3g0r/rehearsals-api-go/internal/domain"
	"github.com/y3g0r/rehearsals-api-go/internal/service"
	"time"
)

var tokenType = "bearer"

type UserService interface {
	CreateUser(ctx context.Context, p service.CreateUserParams) error
	Authenticate(ctx context.Context, p service.AuthenticationParams) (domain.User, error)
}

type CryptoService interface {
	CreateAccessToken(subject string, expiresDelta time.Duration) (string, error)
}

type FastAPI struct {
	userService UserService
	cryptoSvc   CryptoService
}

var _ StrictServerInterface = (*FastAPI)(nil)

func NewFastAPI(userSvc UserService, cryptoSvc CryptoService) *FastAPI {
	return &FastAPI{
		userService: userSvc,
		cryptoSvc:   cryptoSvc,
	}
}

func (f *FastAPI) ItemsReadItems(ctx context.Context, request ItemsReadItemsRequestObject) (ItemsReadItemsResponseObject, error) {
	//TODO implement me
	return ItemsReadItems422JSONResponse{}, nil
}

func (f *FastAPI) ItemsCreateItem(ctx context.Context, request ItemsCreateItemRequestObject) (ItemsCreateItemResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (f *FastAPI) ItemsDeleteItem(ctx context.Context, request ItemsDeleteItemRequestObject) (ItemsDeleteItemResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (f *FastAPI) ItemsReadItem(ctx context.Context, request ItemsReadItemRequestObject) (ItemsReadItemResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (f *FastAPI) ItemsUpdateItem(ctx context.Context, request ItemsUpdateItemRequestObject) (ItemsUpdateItemResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (f *FastAPI) LoginLoginAccessToken(ctx context.Context, request LoginLoginAccessTokenRequestObject) (LoginLoginAccessTokenResponseObject, error) {
	form, err2 := request.Body.ReadForm(18590499)
	if err2 != nil {
		detail := "Couldn't read form"
		return LoginLoginAccessToken400JSONResponse{Detail: &detail}, nil
	}

	params := service.AuthenticationParams{
		Username: form.Value["username"][0],
		Password: form.Value["password"][0],
	}
	user, err := f.userService.Authenticate(ctx, params)
	if err != nil {
		detail := "Invalid username or password"
		return LoginLoginAccessToken400JSONResponse{Detail: &detail}, nil
	}

	if !user.IsActive() {
		detail := "User is not active"
		return LoginLoginAccessToken400JSONResponse{Detail: &detail}, nil
	}

	accessToken, err := f.cryptoSvc.CreateAccessToken(user.ID(), 15*time.Minute)
	if err != nil {
		detail := "Couldn't create access token"
		// TODO: this should be 500, but I don't know how to create one properly yet
		return LoginLoginAccessToken400JSONResponse{Detail: &detail}, nil
	}

	return LoginLoginAccessToken200JSONResponse{
		AccessToken: accessToken,
		TokenType:   &tokenType,
	}, nil
}

func (f *FastAPI) LoginTestToken(ctx context.Context, request LoginTestTokenRequestObject) (LoginTestTokenResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (f *FastAPI) LoginRecoverPasswordHtmlContent(ctx context.Context, request LoginRecoverPasswordHtmlContentRequestObject) (LoginRecoverPasswordHtmlContentResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (f *FastAPI) LoginRecoverPassword(ctx context.Context, request LoginRecoverPasswordRequestObject) (LoginRecoverPasswordResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (f *FastAPI) LoginResetPassword(ctx context.Context, request LoginResetPasswordRequestObject) (LoginResetPasswordResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (f *FastAPI) UsersReadUsers(ctx context.Context, request UsersReadUsersRequestObject) (UsersReadUsersResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (f *FastAPI) UsersCreateUser(ctx context.Context, request UsersCreateUserRequestObject) (UsersCreateUserResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (f *FastAPI) UsersDeleteUserMe(ctx context.Context, request UsersDeleteUserMeRequestObject) (UsersDeleteUserMeResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (f *FastAPI) UsersReadUserMe(ctx context.Context, request UsersReadUserMeRequestObject) (UsersReadUserMeResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (f *FastAPI) UsersUpdateUserMe(ctx context.Context, request UsersUpdateUserMeRequestObject) (UsersUpdateUserMeResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (f *FastAPI) UsersUpdatePasswordMe(ctx context.Context, request UsersUpdatePasswordMeRequestObject) (UsersUpdatePasswordMeResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (f *FastAPI) UsersRegisterUser(ctx context.Context, request UsersRegisterUserRequestObject) (UsersRegisterUserResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (f *FastAPI) UsersDeleteUser(ctx context.Context, request UsersDeleteUserRequestObject) (UsersDeleteUserResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (f *FastAPI) UsersReadUserById(ctx context.Context, request UsersReadUserByIdRequestObject) (UsersReadUserByIdResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (f *FastAPI) UsersUpdateUser(ctx context.Context, request UsersUpdateUserRequestObject) (UsersUpdateUserResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (f *FastAPI) UtilsHealthCheck(ctx context.Context, request UtilsHealthCheckRequestObject) (UtilsHealthCheckResponseObject, error) {
	return UtilsHealthCheck200JSONResponse(true), nil
}

func (f *FastAPI) UtilsTestEmail(ctx context.Context, request UtilsTestEmailRequestObject) (UtilsTestEmailResponseObject, error) {
	//TODO implement me
	panic("implement me")
}
