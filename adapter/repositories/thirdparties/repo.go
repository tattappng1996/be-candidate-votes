package thirdparties

import "be-cadidate-votes/models"

type ThirdPartyRepo interface {
	// Bcrypt
	CheckPasswordHash(password, hash string) bool
	HashPassword(password string) (string, error)

	// JWT
	CreateJWTToken(secret string, claimExpired uint32, user models.User) (string, error)
}

type thirdPartyRepo struct {
}

func NewThirdPartyRepository() ThirdPartyRepo {
	return &thirdPartyRepo{}
}
