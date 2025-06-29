package jwt

type JWTGenerationError struct{}

func (e *JWTGenerationError) Error() string {
	return "failed to generate authentication token"
}
