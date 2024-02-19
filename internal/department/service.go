package department

import (
	"context"
	//"helloadmin/internal/user"
	"time"

	"helloadmin/internal/ecode"
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
		Name:      department.Name,
		ParentId:  department.ParentId,
		Sort:      department.Sort,
		Leader:    department.Leader,
		CreatedAt: department.CreatedAt.Format(time.DateTime),
		UpdateAt:  department.UpdatedAt.Format(time.DateTime),
	}, nil
}

func (s *service) SearchDepartment(ctx context.Context, req *FindRequest) (*Response, error) {
	var result Response
	_, departs, err := s.repo.Find(ctx, req)
	if err != nil {
		return nil, err
	}
	result.Items = buildTree(departs, 0)
	return &result, nil
}

func buildTree(deptList *[]Model, parentId uint) []ResponseItem {
	result := make([]ResponseItem, 0)
	if len(*deptList) > 0 {
		for _, depart := range *deptList {
			if depart.ParentId == parentId {
				child := ResponseItem{
					ID:        depart.ID,
					Name:      depart.Name,
					ParentId:  depart.ParentId,
					Sort:      depart.Sort,
					Leader:    depart.Leader,
					CreatedAt: depart.CreatedAt.Format(time.DateTime),
					UpdateAt:  depart.UpdatedAt.Format(time.DateTime),
				}
				child.Children = buildTree(deptList, depart.ID)
				result = append(result, child)
			}
		}
	}
	return result
}

func (s *service) CreateDepartment(ctx context.Context, req *CreateRequest) error {
	if req.ParentId > 0 {
		if dept, _ := s.repo.GetById(ctx, int64(req.ParentId)); dept == nil {
			return ecode.ErrDeptParentNotFound
		}
	}
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
	if req.ParentId > 0 {
		if dept, _ := s.repo.GetById(ctx, int64(req.ParentId)); dept == nil {
			return ecode.ErrDeptParentNotFound
		}
	}
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
	if departments, _ := s.repo.GetByParentId(ctx, id); len(*departments) > 0 {
		return ecode.ErrDeptHasChild
	}
	return s.repo.Delete(ctx, id)
}
