package service

func GetHealthStatus() map[string]string{
	return  map[string]string{
		"status": "ok",
		"service":"auth-service",
	}
}