package thirdparties

import (
	"be-cadidate-votes/models"
	"time"

	"github.com/golang-jwt/jwt"
)

func (tpRepo *thirdPartyRepo) CreateJWTToken(secret string, claimExpired uint32, user models.User) (string, error) {
	tNow := time.Now()
	privateKey := []byte(secret)

	myClaims := jwt.MapClaims{
		"iat":       tNow.Unix(),
		"id":        user.ID,
		"username":  user.Username,
		"is_active": user.IsActive,
		"exp":       tNow.Unix() + int64(claimExpired),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, myClaims)
	return token.SignedString(privateKey)
}
