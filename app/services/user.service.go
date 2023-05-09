package services

import (
	"context"
	"database/sql"
	"fmt"
	"geoloc/app/models"
	"geoloc/app/stores"
)

type UserService struct {
	store *stores.Store
}

func NewUserService(store *stores.Store) *UserService {
	return &UserService{store}
}

func (s *UserService) ListUserService(ctx context.Context) ([]models.User, error) {
	res, err := s.store.UserStore.ListStores(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *UserService) GetBYIDService(ctx context.Context, id int64) (*models.User, error) {
	return s.store.UserStore.GetByIDStore(ctx, id)
}

func (s *UserService) InsertService(ctx context.Context, tx *sql.Tx, req models.User) (*models.User, error) {

	user := models.User{}
	count, err := s.store.UserStore.GetCountStore(ctx)
	if err != nil {
		return nil, err
	}

	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.Email = req.Email
	user.Phone = req.Phone
	user.PasswordHash = req.PasswordHash

	code := fmt.Sprintf("U%05d", count+1)
	req.Code = code

	obj, err := s.store.UserStore.InsertStore(ctx, req)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *UserService) UpdateService(ctx context.Context, tx *sql.Tx, id int64, req models.User) (*models.User, error) {
	user, err := s.store.UserStore.GetByIDStore(ctx, id)
	if err != nil {
		return nil, err
	}

	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.Email = req.Email
	user.Phone = req.Phone

	fmt.Println(user)
	if err := s.store.UserStore.UpdateUser(ctx, tx, *user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) DeleteUserService(ctx context.Context, tx *sql.Tx, id int64) error {
	if err := s.store.UserStore.DeleteUser(ctx, tx, id); err != nil {
		return err
	}
	return nil
}
