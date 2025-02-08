package service

import (
	"context"
	"fmt"
	"server/internal/datastruct"
	"server/internal/entity"
	"server/internal/repository"
	"server/util"

	"github.com/casbin/casbin/v2"
)

type AuthService interface {
	Login(ctx context.Context, ent *entity.AuthLogin) (*datastruct.AuthLoginData, *datastruct.AuthToken, error)
	Refresh(ctx context.Context, ent *entity.AuthRefresh) (*datastruct.AuthToken, error)
	Me(ctx context.Context, ent *entity.AuthMe) (*datastruct.AuthMe, error)
	Logout(ctx context.Context, ent *entity.AuthLogout) error
	ValidateAccess(ctx context.Context, ent *entity.AuthValidateAccess) (*util.AccessTokenClaims, error)
	ValidatePermission(ctx context.Context, ent *entity.AuthValidatePermission) error
}

type authService struct {
	util     *util.Util
	dao      repository.DAO
	enforcer *casbin.Enforcer
}

func NewAuthService(util *util.Util, dao *repository.DAO, enforcer *casbin.Enforcer) AuthService {
	return &authService{
		util:     util,
		dao:      *dao,
		enforcer: enforcer,
	}
}

func (m *authService) Login(ctx context.Context, ent *entity.AuthLogin) (*datastruct.AuthLoginData, *datastruct.AuthToken, error) {
	token := new(datastruct.AuthToken)

	data, err := m.dao.NewAuthRepository().Login(ctx, ent.Username)
	if err != nil {
		return data, token, err
	}

	if !data.IsActive {
		return data, token, fmt.Errorf("13:user not active")
	}

	// verify password
	if !m.util.Security.VerifyHash(data.Password, ent.Password) {
		return data, token, fmt.Errorf("3:username or password you entered is incorrect")
	}

	accessToken, accessExpire, err := m.util.Token.CreateAccess(ctx, data.Id, data.RoleCode)
	if err != nil {
		return data, token, fmt.Errorf("13:%v", err)
	}
	refreshToken, refreshExpired, errT := m.util.Token.CreateRefresh(ctx, data.Id, data.RoleCode)
	if errT != nil {
		return data, token, fmt.Errorf("13:%v", err)
	}
	m.enforcer.AddRoleForUser(data.Id, data.RoleCode)
	return data, &datastruct.AuthToken{
		AccessToken:    accessToken,
		AccessExpired:  accessExpire,
		RefreshToken:   refreshToken,
		RefreshExpired: refreshExpired,
	}, nil
}

func (m *authService) Refresh(ctx context.Context, ent *entity.AuthRefresh) (*datastruct.AuthToken, error) {
	newRefresh := new(datastruct.AuthToken)
	claim, err := m.util.Token.ParseRefresh(ent.Token)
	if err != nil {
		return newRefresh, fmt.Errorf("16:%v", err)
	}

	if err := m.util.Token.ValidateRefresh(ctx, claim); err != nil {
		return newRefresh, fmt.Errorf("16:%v", err)
	}

	accessToken, accessExpire, err := m.util.Token.CreateAccess(ctx, claim.UserId, claim.RoleCode)
	if err != nil {
		return newRefresh, fmt.Errorf("13:%v", err)
	}
	refreshToken, refreshExpired, err := m.util.Token.CreateRefresh(ctx, claim.UserId, claim.RoleCode)
	if err != nil {
		return newRefresh, fmt.Errorf("13:%v", err)
	}

	return &datastruct.AuthToken{
		AccessToken:    accessToken,
		AccessExpired:  accessExpire,
		RefreshToken:   refreshToken,
		RefreshExpired: refreshExpired,
	}, nil
}

func (m *authService) Logout(ctx context.Context, ent *entity.AuthLogout) error {
	err := m.dao.NewAuthRepository().Logout(ctx, ent.UserId)
	if err != nil {
		return err
	}
	return nil
}

func (m *authService) Me(ctx context.Context, ent *entity.AuthMe) (*datastruct.AuthMe, error) {
	data, err := m.dao.NewAuthRepository().Me(ctx, ent.UserId)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (m *authService) ValidateAccess(ctx context.Context, ent *entity.AuthValidateAccess) (*util.AccessTokenClaims, error) {
	claims, err := m.util.Token.ParseAccess(ent.Token)
	if err != nil {
		return nil, fmt.Errorf("16:%v", err)
	}
	if err := m.util.Token.ValidateAccess(ctx, claims); err != nil {
		return nil, fmt.Errorf("16:%v", err)
	}

	return claims, nil
}

func (m *authService) ValidatePermission(ctx context.Context, ent *entity.AuthValidatePermission) error {
	if err := m.enforcer.LoadPolicy(); err != nil {
		return fmt.Errorf("13:failed to load policy")
	}

	accepted, err := m.enforcer.Enforce(ent.UserId, ent.FullMethod)
	if err != nil {
		return fmt.Errorf("13:error when authorizing user's accessibility")
	}

	if !accepted {
		return fmt.Errorf("7:you are not allowed")
	}

	return nil
}
