package login_record

import (
	"context"
	"helloadmin/internal/repository"
)

type Repository interface {
	Create(ctx context.Context, record *Model) error
	Search(ctx context.Context, request *FindRequest) (int64, *[]Model, error)
}

func NewRepository(r *repository.Repository) Repository {
	return &loginRecordRepository{r}
}

type loginRecordRepository struct {
	*repository.Repository
}

func (lr *loginRecordRepository) Create(ctx context.Context, record *Model) error {
	if err := lr.DB(ctx).Create(record).Error; err != nil {
		return err
	}
	return nil
}

func (lr *loginRecordRepository) Search(ctx context.Context, req *FindRequest) (int64, *[]Model, error) {
	var count int64
	var record []Model
	query := lr.DB(ctx)
	if req.Email != "" {
		query = query.Where("email = ?", req.Email)
	}
	if req.Ip != "" {
		query = query.Where("ip = ?", req.Ip)
	}
	if req.Page > 0 {
		query = query.Offset((req.Page - 1) * req.Size).Limit(req.Size)
	}
	query.Model(Model{}).Count(&count)
	if result := query.Order("id desc").Find(&record); result.Error != nil {
		return count, nil, result.Error
	}
	return count, &record, nil
}
