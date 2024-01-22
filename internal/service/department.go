package service

import (
	"context"
	"helloadmin/api"
	"helloadmin/internal/model"
	"helloadmin/internal/repository"
	"time"
)

type DepartmentService interface {
	GetDepartmentById(ctx context.Context, id int64) (*api.DepartmentResponseItem, error)
	SearchDepartment(ctx context.Context, request *api.DepartmentFindRequest) (*api.DepartmentResponse, error)
	CreateDepartment(ctx context.Context, request *api.DepartmentCreateRequest) error
	UpdateDepartment(ctx context.Context, id int64, request *api.DepartmentUpdateRequest) error
	DeleteDepartment(ctx context.Context, id int64) error
}

func NewDepartmentService(service *Service, departmentRepository repository.DepartmentRepository) DepartmentService {
	return &departmentService{
		Service:              service,
		departmentRepository: departmentRepository,
	}
}

type departmentService struct {
	*Service
	departmentRepository repository.DepartmentRepository
}

func (s *departmentService) GetDepartmentById(ctx context.Context, id int64) (*api.DepartmentResponseItem, error) {
	department, err := s.departmentRepository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &api.DepartmentResponseItem{
		ID:        department.ID,
		UpdateAt:  department.UpdatedAt.Format(time.RFC3339),
		Name:      department.Name,
		ParentId:  department.ParentId,
		Sort:      department.Sort,
		Leader:    department.Leader,
		CreatedAt: department.CreatedAt.Format(time.RFC3339),
	}, nil
}

func (s *departmentService) SearchDepartment(ctx context.Context, req *api.DepartmentFindRequest) (*api.DepartmentResponse, error) {
	var result api.DepartmentResponse
	count, departs, err := s.departmentRepository.Find(ctx, req)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		result.Items = make([]api.DepartmentResponseItem, 0, req.Size)
		for _, depart := range *departs {
			tmp := api.DepartmentResponseItem{
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

func (s *departmentService) CreateDepartment(ctx context.Context, req *api.DepartmentCreateRequest) error {
	department := model.Department{
		Name:     req.Name,
		ParentId: req.ParentId,
		Leader:   req.Leader,
		Sort:     req.Sort,
	}
	return s.departmentRepository.Create(ctx, &department)
}

func (s *departmentService) UpdateDepartment(ctx context.Context, id int64, req *api.DepartmentUpdateRequest) error {
	department := model.Department{
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
