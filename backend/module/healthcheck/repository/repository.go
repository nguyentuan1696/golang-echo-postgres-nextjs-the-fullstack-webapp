package repository

type IHealthCheckRepository interface {
}

type HealthCheckRepository struct {
}

func NewHealthCheckRepository() IHealthCheckRepository {
	return &HealthCheckRepository{}
}
