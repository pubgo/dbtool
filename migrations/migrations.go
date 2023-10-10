package migrations

import (
	"gitlab.asants.com/bas/migrate/migrations/internal"
	"gitlab.asants.com/bas/migrate/pkg/migrates"
)

func New() []migrates.Migrate {
	return []migrates.Migrate{
		internal.M0001CreateMenuModel,
	}
}
