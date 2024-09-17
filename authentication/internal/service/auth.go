package service

import (
	"context"
	"fmt"
	"server/config"
	"server/internal/datastruct"
	"server/internal/entity"
	"server/internal/repository"
	"server/util"
)

type AuthService interface {
	Login(ctx *context.Context, ent *entity.AuthLogin) (*datastruct.AuthLoginData, *datastruct.AuthToken, *util.Error)
	Refresh(ctx *context.Context, ent *entity.AuthRefresh) (*datastruct.AuthToken, *util.Error)
	Me(ctx *context.Context, ent *entity.AuthMe) (*datastruct.AuthMe, *util.Error)
	Logout(ctx *context.Context, ent *entity.AuthLogout) *util.Error
	ValidateAccess(ctx *context.Context, ent *entity.AuthValidateAccess) (*util.AccessTokenClaims, *util.Error)
	ValidatePermission(ctx *context.Context, ent *entity.AuthValidatePermission) *util.Error
}

type authService struct {
	dao repository.DAO
}

func NewAuthService(dao *repository.DAO) AuthService {
	return &authService{
		dao: *dao,
	}
}

func (m *authService) Login(ctx *context.Context, ent *entity.AuthLogin) (*datastruct.AuthLoginData, *datastruct.AuthToken, *util.Error) {
	token := new(datastruct.AuthToken)

	data, err := m.dao.NewAuthRepository().Login(ctx, ent.Username)
	if err.Errors != nil {
		return data, token, err
	}

	if !data.IsActive {
		return data, token, &util.Error{
			Errors:     fmt.Errorf("user not active"),
			StatusCode: 5,
		}
	}

	// verify password
	if !util.NewSecurity().VerifyHash(data.Password, *ent.Password) {
		return data, token, &util.Error{
			Errors:     fmt.Errorf("username or password you entered is incorrect, please enter the correct username and password"),
			StatusCode: 5,
		}
	}

	accessToken, accessExpire, errT := util.NewToken().CreateAccess(ctx, &data.Id, &data.RoleCode)
	if errT != nil {
		return data, token, &util.Error{
			Errors:     errT,
			StatusCode: 13,
		}
	}
	refreshToken, refreshExpired, errT := util.NewToken().CreateRefresh(ctx, &data.Id, &data.RoleCode)
	if errT != nil {
		return data, token, &util.Error{
			Errors:     errT,
			StatusCode: 13,
		}
	}
	enforce := config.NewEnforcer()
	enforce.AddRoleForUser(data.Id, data.RoleCode)
	return data, &datastruct.AuthToken{
		AccessToken:    accessToken,
		AccessExpired:  accessExpire,
		RefreshToken:   refreshToken,
		RefreshExpired: refreshExpired,
	}, &util.Error{}
}

func (m *authService) Refresh(ctx *context.Context, ent *entity.AuthRefresh) (*datastruct.AuthToken, *util.Error) {
	newRefresh := new(datastruct.AuthToken)
	claim, errT := util.NewToken().ParseRefresh(ent.Token)
	if errT != nil {
		return newRefresh, &util.Error{
			Errors:     errT,
			StatusCode: 16,
		}
	}

	if errT := util.NewToken().ValidateRefresh(ctx, claim); errT != nil {
		return newRefresh, &util.Error{
			Errors:     errT,
			StatusCode: 16,
		}
	}

	accessToken, accessExpire, errT := util.NewToken().CreateAccess(ctx, &claim.UserId, &claim.RoleCode)
	if errT != nil {
		return newRefresh, &util.Error{
			Errors:     errT,
			StatusCode: 13,
		}
	}
	refreshToken, refreshExpired, errT := util.NewToken().CreateRefresh(ctx, &claim.UserId, &claim.RoleCode)
	if errT != nil {
		return newRefresh, &util.Error{
			Errors:     errT,
			StatusCode: 13,
		}
	}

	return &datastruct.AuthToken{
		AccessToken:    accessToken,
		AccessExpired:  accessExpire,
		RefreshToken:   refreshToken,
		RefreshExpired: refreshExpired,
	}, &util.Error{}
}

func (m *authService) Logout(ctx *context.Context, ent *entity.AuthLogout) *util.Error {
	err := m.dao.NewAuthRepository().Logout(ctx, ent.UserId)
	if err.Errors != nil {
		return err
	}
	return &util.Error{}
}

func (m *authService) Me(ctx *context.Context, ent *entity.AuthMe) (*datastruct.AuthMe, *util.Error) {
	data, err := m.dao.NewAuthRepository().Me(ctx, ent.UserId)
	if err.Errors != nil {
		return data, err
	}

	return data, &util.Error{}
}

func (m *authService) ValidateAccess(ctx *context.Context, ent *entity.AuthValidateAccess) (*util.AccessTokenClaims, *util.Error) {
	claims, errT := util.NewToken().ParseAccess(ent.Token)
	if errT != nil {
		return nil, &util.Error{
			Errors:     errT,
			StatusCode: 16,
		}
	}
	if errT := util.NewToken().ValidateAccess(ctx, claims); errT != nil {
		return nil, &util.Error{
			Errors:     errT,
			StatusCode: 16,
		}
	}

	return claims, &util.Error{}
}

func (m *authService) ValidatePermission(ctx *context.Context, ent *entity.AuthValidatePermission) *util.Error {
	enforcer := config.NewEnforcer()
	if err := enforcer.LoadPolicy(); err != nil {
		return &util.Error{
			Errors:     fmt.Errorf("failed to load policy"),
			StatusCode: 13,
		}
	}

	accepted, err := enforcer.Enforce(*ent.UserId, *ent.FullMethod)
	if err != nil {
		return &util.Error{
			Errors:     fmt.Errorf("error when authorizing user's accessibility"),
			StatusCode: 13,
		}
	}

	if !accepted {
		return &util.Error{
			Errors:     fmt.Errorf("you are not allowed"),
			StatusCode: 7,
		}
	}

	return &util.Error{}
}
