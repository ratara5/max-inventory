package service

import (
	"context"
	"errors"
	"max-inventory/internal/models"
)

var validRolesToAddProducts []int64 = []int64{1, 2}

var ErrInvalidPermissions = errors.New("user does not have permission to add product")

func (s *serv) GetProducts(ctx context.Context) ([]models.Product, error) {

	pp, err := s.repo.GetProducts(ctx)
	if err != nil {
		return nil, err
	}

	products := []models.Product{}

	for _, p := range pp {
		products = append(products, models.Product{
			ID:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
		})
	}
	return products, nil
}

func (s *serv) GetProduct(ctx context.Context, id int64) (*models.Product, error) {
	p, err := s.repo.GetProduct(ctx, id)
	if err != nil {
		return nil, err
	}

	product := &models.Product{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
	}
	return product, nil
}

func (s *serv) AddProduct(ctx context.Context, product models.Product, email string) error {
	u, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}

	roles, err := s.repo.GetUserRoles(ctx, u.ID)
	if err != nil {
		return err
	}

	userCanAdd := false

	for _, r := range roles {
		for _, vr := range validRolesToAddProducts {
			if vr == r.RoleID {
				userCanAdd = true
				break
			}
		}
	}
	if !userCanAdd {
		return ErrInvalidPermissions
	}

	return s.repo.SaveProduct(ctx, product.Name, product.Description, product.Price, u.ID)
}
