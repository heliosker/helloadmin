package department

import (
	"context"
	"helloadmin/internal/api"
	"time"
)

type Service interface {
	GetDepartmentById(ctx context.Context, id int64) (*ResponseItem, error)
	SearchDepartment(ctx context.Context, request *FindRequest) (*Response, error)
	CreateDepartment(ctx context.Context, request *CreateRequest) error
	UpdateDepartment(ctx context.Context, id int64, request *UpdateRequest) error
	DeleteDepartment(ctx context.Context, id int64) error
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

type service struct {
	repo Repository
}

func (s *service) GetDepartmentById(ctx context.Context, id int64) (*ResponseItem, error) {
	department, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &ResponseItem{
		ID:        department.ID,
		UpdateAt:  department.UpdatedAt.Format(time.RFC3339),
		Name:      department.Name,
		ParentId:  department.ParentId,
		Sort:      department.Sort,
		Leader:    department.Leader,
		CreatedAt: department.CreatedAt.Format(time.RFC3339),
	}, nil
}

func (s *service) SearchDepartment(ctx context.Context, req *FindRequest) (*Response, error) {
	var result Response
	count, departs, err := s.repo.Find(ctx, req)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		result.Items = make([]ResponseItem, 0, req.Size)
		for _, depart := range *departs {
			tmp := ResponseItem{
				ID:        depart.ID,
				Name:      depart.Name,
				ParentId:  depart.ParentId,
				Sort:      depart.Sort,
				Leader:    depart.Leader,
				CreatedAt: depart.CreatedAt.Format(time.RFC3339),
			}
			result.Items = append(result.Items, tmp)
		}
	}
	result.Pagination = api.Pagination{
		Page:  req.Page,
		Size:  req.Size,
		Count: int(count),
	}
	return &result, nil
}

func (s *service) CreateDepartment(ctx context.Context, req *CreateRequest) error {
	department := Model{
		Name:      req.Name,
		ParentId:  req.ParentId,
		Leader:    req.Leader,
		Sort:      req.Sort,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}
	return s.repo.Create(ctx, &department)
}

func (s *service) UpdateDepartment(ctx context.Context, id int64, req *UpdateRequest) error {
	department := Model{
		Name:      req.Name,
		ParentId:  req.ParentId,
		Leader:    req.Leader,
		Sort:      req.Sort,
		UpdatedAt: time.Now(),
	}
	return s.repo.Update(ctx, id, &department)
}

func (s *service) DeleteDepartment(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
