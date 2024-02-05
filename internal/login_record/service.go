package login_record

import (
	"context"
	"helloadmin/internal/api"
	"time"
)

type Service interface {
	Create(ctx context.Context, record *CreateRequest) error
	Search(ctx context.Context, request *FindRequest) (*Response, error)
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

type service struct {
	repo Repository
}

func (lrs *service) Create(ctx context.Context, req *CreateRequest) error {
	model := Model{
		Ip:           req.Ip,
		Os:           req.Os,
		Email:        req.Email,
		Browser:      req.Browser,
		Platform:     req.Platform,
		ErrorMessage: req.ErrorMessage,
		UpdatedAt:    time.Now(),
		CreatedAt:    time.Now(),
	}
	if err := lrs.repo.Create(ctx, &model); err != nil {
		return err
	}
	return nil
}

func (lrs *service) Search(ctx context.Context, request *FindRequest) (*Response, error) {
	var result Response
	count, records, err := lrs.repo.Search(ctx, request)
	if err != nil {
		return nil, err
	}
	result.Items = make([]LoginRecordItem, 0)
	if count > 0 {
		for _, record := range *records {
			result.Items = append(result.Items, LoginRecordItem{
				Ip:           record.Ip,
				Os:           record.Os,
				Email:        record.Email,
				Browser:      record.Browser,
				Platform:     record.Platform,
				ErrorMessage: record.ErrorMessage,
				UpdatedAt:    record.UpdatedAt.Format(time.DateTime),
				CreatedAt:    record.CreatedAt.Format(time.DateTime),
			})
		}
	}
	result.Pagination = api.Pagination{
		Page:  request.Page,
		Size:  request.Size,
		Count: int(count),
	}
	return &result, nil
}
