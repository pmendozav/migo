package partnerhandle

import (
	"mio/app/models"
	repo "mio/app/repositories"
)

func FindUsersWithWildcard(wildcard string) ([]models.Partner) {
	var (
		idpartner string
		name string
		code string
		key string
	)

	partners := make([]models.Partner, 0)
	repo := repo.BongoPartnerRepo{}
	
	rows, _ := repo.FindPartner(wildcard)
	for rows.Next() {
		rows.Scan(&idpartner, &name, &code, &key)
		partners = append(partners, models.Partner{IdPartner:idpartner, Name:name, Code:code, Key:key})
	}

	return partners
}