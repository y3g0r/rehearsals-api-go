//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=types.cfg.yml fastapi-fullstack-openapi.json
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=server.cfg.yml fastapi-fullstack-openapi.json

package api

import "context"

type FastAPI struct{}

func NewFastAPI() *FastAPI {
	return &FastAPI{}
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

var _ StrictServerInterface = (*FastAPI)(nil)
