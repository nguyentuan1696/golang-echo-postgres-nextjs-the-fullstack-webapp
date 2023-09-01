package service

type IHealthCheckService interface {
}

type HealthCheckService struct {
}

func NewHealthCheckService() IHealthCheckService {
	return &HealthCheckService{}
}
