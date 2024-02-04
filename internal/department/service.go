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

func NewDepartmentService(repo DeptRepository) Service {
	return &departmentService{
		departmentRepository: repo,
	}
}

type departmentService struct {
	departmentRepository DeptRepository
}

func (s *departmentService) GetDepartmentById(ctx context.Context, id int64) (*ResponseItem, error) {
	department, err := s.departmentRepository.GetById(ctx, id)
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

func (s *departmentService) SearchDepartment(ctx context.Context, req *FindRequest) (*Response, error) {
	var result Response
	count, departs, err := s.departmentRepository.Find(ctx, req)
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

func (s *departmentService) CreateDepartment(ctx context.Context, req *CreateRequest) error {
	department := Model{
		Name:     req.Name,
		ParentId: req.ParentId,
		Leader:   req.Leader,
		Sort:     req.Sort,
	}
	return s.departmentRepository.Create(ctx, &department)
}

func (s *departmentService) UpdateDepartment(ctx context.Context, id int64, req *UpdateRequest) error {
	department := Model{
		Name:     req.Name,
		ParentId: req.ParentId,
		Leader:   req.Leader,
		Sort:     req.Sort,
	}
	return s.departmentRepository.Update(ctx, id, &department)
}

func (s *departmentService) DeleteDepartment(ctx context.Context, id int64) error {
	return s.departmentRepository.Delete(ctx, id)
}
