package user

import (
	"context"
	"time"

	"golang.org/x/crypto/bcrypt"
	"helloadmin/internal/api"
	"helloadmin/internal/ecode"
	"helloadmin/pkg/helper/generate"
	"helloadmin/pkg/helper/sid"
	"helloadmin/pkg/jwt"
)

type Service interface {
	Register(ctx context.Context, req *RegisterRequest) error
	Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error)
	GetProfileByUserId(ctx context.Context, userId string) (*ProfileData, error)
	GetProfileById(ctx context.Context, id int64) (*ProfileData, error)
	Update(ctx context.Context, id int64, req *UpdateRequest) error
	Search(ctx context.Context, request *FindRequest) (*Response, error)
}

func NewService(sid *sid.Sid, jwt *jwt.JWT, repo Repository) Service {
	return &userService{
		sid:  sid,
		jwt:  jwt,
		repo: repo,
	}
}

type userService struct {
	sid  *sid.Sid
	jwt  *jwt.JWT
	repo Repository
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
				Id:       item.ID,
				UserId:   item.UserId,
				Nickname: item.Nickname,
				Email:    item.Email,
				Department: struct {
					Id   uint   `json:"id"`
					Name string `json:"name"`
				}{
					Id:   item.Department.ID,
					Name: item.Department.Name,
				},
				Role: struct {
					Id   uint   `json:"id"`
					Name string `json:"name"`
				}{
					Id:   item.Role.ID,
					Name: item.Role.Name,
				},
				RoleId:    item.RoleId,
				DeptId:    item.DeptId,
				CreatedAt: item.CreatedAt.Format(time.DateTime),
				UpdatedAt: item.UpdatedAt.Format(time.DateTime),
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
		ExpiresAt:   expiresAt.Format(time.DateTime),
		TokenType:   "Bearer",
	}, nil
}

func (s *userService) GetProfileByUserId(ctx context.Context, userId string) (*ProfileData, error) {
	user, err := s.repo.GetByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}
	return profile(user), nil
}

func (s *userService) GetProfileById(ctx context.Context, id int64) (*ProfileData, error) {
	user, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return profile(user), nil
}

func (s *userService) Update(ctx context.Context, id int64, req *UpdateRequest) error {
	user, err := s.repo.GetById(ctx, id)
	if err != nil {
		return err
	}
	if user.RoleId == 0 && req.RoleId != 0 {
		return ecode.ErrAdminUserCanNotModify
	}
	user.Email = req.Email
	user.Nickname = req.Nickname
	user.RoleId = uint(req.RoleId)
	user.DeptId = uint(req.DeptId)
	user.UpdatedAt = time.Now()
	if err = s.repo.Update(ctx, user); err != nil {
		return err
	}
	return nil
}

func profile(m *Model) *ProfileData {
	if m == nil {
		return nil
	}
	return &ProfileData{
		Id:       m.ID,
		UserId:   m.UserId,
		Nickname: m.Nickname,
		Email:    m.Email,
		RoleId:   m.RoleId,
		DeptId:   m.DeptId,
		Department: struct {
			Id   uint   `json:"id"`
			Name string `json:"name"`
		}{
			Id:   m.Department.ID,
			Name: m.Department.Name,
		},
		Role: struct {
			Id   uint   `json:"id"`
			Name string `json:"name"`
		}{
			Id:   m.Role.ID,
			Name: m.Role.Name,
		},
		CreatedAt: m.CreatedAt.Format(time.DateTime),
		UpdatedAt: m.UpdatedAt.Format(time.DateTime),
	}
}
