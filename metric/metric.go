package metric

import "github.com/amazingguni/meerkat/model"

type Store interface {
	CreateMetric(*model.Metric) error
	List(int, int) ([]model.Metric, int, error)
}
