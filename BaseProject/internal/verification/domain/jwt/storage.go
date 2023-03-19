package jwt

import (
	"github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/jwt/storage/io"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Storage interface {
	GenerateToken(input io.GenerateTokenInput, expiresAt time.Time) (*io.Token, error)
	GenerateSmtpCodeVerification(code string) (string, error)
	VerifyToken(tokenString string) (*jwt.Token, error)
	VerifySmtpToken(tokenString string, signature string) (*jwt.Token, error)
}
