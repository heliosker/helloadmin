package server

import (
	"context"
	"go.uber.org/zap"
	"gorm.io/gorm"
	login_log "helloadmin/internal/login_record"
	"helloadmin/internal/model"
	"helloadmin/internal/operation_record"
	"helloadmin/pkg/log"
	"os"
)

type Migrate struct {
	db  *gorm.DB
	log *log.Logger
}

func NewMigrate(db *gorm.DB, log *log.Logger) *Migrate {
	return &Migrate{
		db:  db,
		log: log,
	}
}
func (m *Migrate) Start(ctx context.Context) error {
	if err := m.db.AutoMigrate(&model.User{}); err != nil {
		m.log.Error("user migrate error", zap.Error(err))
		return err
	}
	if err := m.db.AutoMigrate(&model.Role{}); err != nil {
		m.log.Error("role migrate error", zap.Error(err))
		return err
	}
	if err := m.db.AutoMigrate(&model.Menu{}); err != nil {
		m.log.Error("menu migrate error", zap.Error(err))
		return err
	}
	if err := m.db.AutoMigrate(&model.Department{}); err != nil {
		m.log.Error("department migrate error", zap.Error(err))
		return err
	}
	if err := m.db.AutoMigrate(&login_log.Model{}); err != nil {
		m.log.Error("sign_record migrate error", zap.Error(err))
		return err
	}
	if err := m.db.AutoMigrate(&operation_record.Model{}); err != nil {
		m.log.Error("operation_record migrate error", zap.Error(err))
		return err
	}

	m.log.Info("AutoMigrate success")
	os.Exit(0)
	return nil
}
func (m *Migrate) Stop(ctx context.Context) error {
	m.log.Info("AutoMigrate stop")
	return nil
}
