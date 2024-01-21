package repositories

import (
	"be-cadidate-votes/adapter/repositories/tb"
	"be-cadidate-votes/adapter/repositories/thirdparties"

	"gorm.io/gorm"
)

type Repo struct {
	TbRepo         tb.TbRepo
	ThirdPartyRepo thirdparties.ThirdPartyRepo
}

func NewRepository(db *gorm.DB) *Repo {
	return &Repo{
		tb.NewTbRepository(db),
		thirdparties.NewThirdPartyRepository(),
	}
}
