package service

type Service struct {
	ColdStorage
	HotStorage
	Storage
}

var s Service

func GetService() Service {
	return s
}

func InitializeService(service Service) {
	s = service
}
