package service

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"helloadmin/api"
	"helloadmin/internal/ecode"
	"helloadmin/internal/model"
	"helloadmin/internal/repository"
	"time"
)

type UserService interface {
	Register(ctx context.Context, req *api.RegisterRequest) error
	Login(ctx context.Context, req *api.LoginRequest) (string, error)
	GetProfile(ctx context.Context, userId string) (*api.GetProfileResponseData, error)
	UpdateProfile(ctx context.Context, userId string, req *api.UpdateProfileRequest) error
}

func NewUserService(service *Service, userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
		Service:  service,
	}
}

type userService struct {
	userRepo repository.UserRepository
	*Service
}

func (s *userService) Register(ctx context.Context, req *api.RegisterRequest) error {
	// check username
	if user, err := s.userRepo.GetByEmail(ctx, req.Email); err == nil && user != nil {
		return ecode.ErrEmailAlreadyUse
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	// Generate user ID
	userId, err := s.sid.GenString()
	if err != nil {
		return err
	}
	user := &model.User{
		UserId:   userId,
		Email:    req.Email,
		Password: string(hashedPassword),
	}
	// Transaction demo
	err = s.tm.Transaction(ctx, func(ctx context.Context) error {
		// Create a user
		if err = s.userRepo.Create(ctx, user); err != nil {
			return err
		}
		// TODO: other repo
		return nil
	})
	return err
}

func (s *userService) Login(ctx context.Context, req *api.LoginRequest) (string, error) {
	user, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil || user == nil {
		return "", ecode.ErrUnauthorized
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", err
	}
	token, err := s.jwt.GenToken(user.UserId, time.Now().Add(time.Hour*24*90))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *userService) GetProfile(ctx context.Context, userId string) (*api.GetProfileResponseData, error) {
	user, err := s.userRepo.GetByID(ctx, userId)
	if err != nil {
		return nil, err
	}

	return &api.GetProfileResponseData{
		UserId:    user.UserId,
		Nickname:  user.Nickname,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format(time.DateTime),
	}, nil
}

func (s *userService) UpdateProfile(ctx context.Context, userId string, req *api.UpdateProfileRequest) error {
	user, err := s.userRepo.GetByID(ctx, userId)
	if err != nil {
		return err
	}

	user.Email = req.Email
	user.Nickname = req.Nickname

	if err = s.userRepo.Update(ctx, user); err != nil {
		return err
	}

	return nil
}
