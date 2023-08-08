package service

import (
	"context"
	"errors"
	"max-inventory/encryption"
	"max-inventory/internal/models"
)

var (
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrRoleAlreadyAdded = errors.New("role was already added for this user")
	ErrRoleNotFound = errors.New("role not found for this user")
)

func (s *serv) RegisterUser(ctx context.Context, email, name, password string) error {
	u, _ := s.repo.GetUserByEmail(ctx, email)
	if u != nil {
		return ErrUserAlreadyExists
	}

	//Hash password
	bb, err := encryption.Encrypt([]byte(password))
	if err != nil {
		return err
	}
	pass, _ := encryption.ToBase64(bb)

	return s.repo.SaveUser(ctx, email, name, pass)
}

func (s *serv) LoginUser(ctx context.Context, email, password string) (*models.User, error) {
	u, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	//Decrypt the password
	bb, err := encryption.FromBase64(u.Password)
	if err != nil {
		return nil, err
	}
	decriptedPassword, err := encryption.Decrypt(bb)
	if err != nil {
		return nil, err
	}

	if string(decriptedPassword) != password {
		return nil, ErrInvalidCredentials
	}
	return &models.User{
		ID:    u.ID,
		Email: u.Email,
		Name:  u.Name,
	}, nil
}

func (s *serv) AddUserRole(ctx context.Context, userID, roleID int64) error {
	roles, err := s.repo.GetUserRoles(ctx, userID)
	if err!= nil {
        return err
    }
	for _, r := range roles {
        if r.RoleID == roleID {
            return ErrRoleAlreadyAdded
        }
    }
	return s.repo.SaveUserRole(ctx, userID, roleID)
}

func (s *serv) RemoveUserRole(ctx context.Context, userID, roleID int64) error {

	roles, err := s.repo.GetUserRoles(ctx, userID)
	if err!= nil {
        return err
    }
	roleFound := false
	for _, r := range roles {
        if r.RoleID == roleID {
            roleFound = true
			break
        }
    }
	if !roleFound {
		return ErrRoleNotFound
	}

	return s.repo.RemoveUserRole(ctx, userID, roleID)
}
