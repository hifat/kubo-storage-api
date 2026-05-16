package middlewaremld

type Service interface {
	ValidateToken(token string) error
}
