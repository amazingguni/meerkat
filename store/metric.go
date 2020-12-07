package store

import (
	"github.com/amazingguni/meerkat/model"
	"github.com/jinzhu/gorm"
)

type MetricStore struct {
	db *gorm.DB
}

func NewMetricStore(db *gorm.DB) *MetricStore {
	return &MetricStore{
		db: db,
	}
}

func (ms *MetricStore) CreateMetric(m *model.Metric) error {
	tx := ms.db.Begin()
	if err := tx.Create(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (ms *MetricStore) List(offset int, limit int) ([]model.Metric, int, error) {
	var (
		metrics []model.Metric
		count   int
	)
	ms.db.Model(&metrics).Count(&count)
	ms.db.Offset(offset).Limit(limit).Order("created_at desc").Find(&metrics)
	return metrics, count, nil
}
