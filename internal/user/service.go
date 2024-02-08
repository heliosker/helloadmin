package user

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"helloadmin/internal/api"
	"helloadmin/internal/ecode"
	"helloadmin/pkg/helper/generate"
	"helloadmin/pkg/helper/sid"
	"helloadmin/pkg/jwt"
	"time"
)

type Service interface {
	Register(ctx context.Context, req *RegisterRequest) error
	Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error)
	GetProfile(ctx context.Context, userId string) (*ProfileData, error)
	UpdateProfile(ctx context.Context, userId string, req *UpdateProfileRequest) error
	Search(ctx context.Context, request *FindRequest) (*Response, error)
}

func NewService(sid *sid.Sid, jwt *jwt.JWT, repo Repository) Service {
	return &userService{
		repo: repo,
		sid:  sid,
		jwt:  jwt,
	}
}

type userService struct {
	repo Repository
	sid  *sid.Sid
	jwt  *jwt.JWT
}

func (s *userService) Register(ctx context.Context, req *RegisterRequest) error {
	// check username
	if user, err := s.repo.GetByEmail(ctx, req.Email); err == nil && user != nil {
		return ecode.ErrEmailAlreadyUse
	}

	salt := generate.RandomString(16)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password+salt), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	// Generate user ID
	userId, err := s.sid.GenString()
	if err != nil {
		return err
	}

	user := &Model{
		UserId:    userId,
		Email:     req.Email,
		Password:  string(hashedPassword),
		Salt:      salt,
		Nickname:  req.Nickname,
		RoleId:    req.RoleId,
		DeptId:    req.DeptId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	// Transaction
	if err = s.repo.Create(ctx, user); err != nil {
		return err
	}
	return nil
}

func (s *userService) Search(ctx context.Context, req *FindRequest) (*Response, error) {
	var response Response
	total, items, err := s.repo.Search(ctx, req)
	if err != nil {
		return nil, err
	}
	response.Items = make([]ProfileData, 0)
	if total > 0 {
		for _, item := range *items {
			response.Items = append(response.Items, ProfileData{
				UserId:    item.UserId,
				Nickname:  item.Nickname,
				Email:     item.Email,
				RoleId:    item.RoleId,
				DeptId:    item.DeptId,
				CreatedAt: item.CreatedAt.Format(time.RFC3339),
				UpdatedAt: item.UpdatedAt.Format(time.RFC3339),
			})
		}
	}
	response.Pagination = api.Pagination{
		Page:  req.Page,
		Size:  req.Size,
		Count: int(total),
	}
	return &response, nil
}

func (s *userService) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	user, err := s.repo.GetByEmail(ctx, req.Email)
	if err != nil || user == nil {
		return nil, ecode.ErrUserNotFound
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password+user.Salt), []byte(req.Password+user.Salt)); err != nil {
		return nil, ecode.ErrPasswordIncorrect
	}
	expiresAt := time.Now().Add(time.Hour * 24 * 90)
	token, err := s.jwt.GenToken(user.UserId, expiresAt)
	if err != nil {
		return nil, err
	}
	return &LoginResponse{
		AccessToken: token,
		ExpiresAt:   expiresAt.Format(time.RFC3339),
		TokenType:   "Bearer",
	}, nil
}

func (s *userService) GetProfile(ctx context.Context, userId string) (*ProfileData, error) {
	user, err := s.repo.GetByID(ctx, userId)
	if err != nil {
		return nil, err
	}

	return &ProfileData{
		UserId:    user.UserId,
		Nickname:  user.Nickname,
		Email:     user.Email,
		RoleId:    user.RoleId,
		DeptId:    user.DeptId,
		CreatedAt: user.CreatedAt.Format(time.DateTime),
		UpdatedAt: user.UpdatedAt.Format(time.DateTime),
	}, nil
}

func (s *userService) UpdateProfile(ctx context.Context, userId string, req *UpdateProfileRequest) error {
	user, err := s.repo.GetByID(ctx, userId)
	if err != nil {
		return err
	}
	user.Email = req.Email
	user.Nickname = req.Nickname
	user.RoleId = req.RoleId
	user.DeptId = req.DeptId
	user.UpdatedAt = time.Now()
	if err = s.repo.Update(ctx, user); err != nil {
		return err
	}
	return nil
}
