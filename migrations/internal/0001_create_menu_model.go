package internal

import (
	"gitlab.asants.com/bas/migrate/pkg/migrates"
	"gorm.io/gorm"
)

func M0001CreateMenuModel() *migrates.Migration {
	type TestUser struct {
		ID       uint
		Username string `gorm:"not null;uniqueIndex"`
		Email    string
		Language string
		Timezone string
	}

	return &migrates.Migration{
		ID: "0001_create_model",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&TestUser{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(&TestUser{})
		},
	}
}
