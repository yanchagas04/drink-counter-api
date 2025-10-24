package driver

import (
	postsModels "drink-counter-api/posts/models"
	usersModels "drink-counter-api/users/models"
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)


func RunMigrations(db *gorm.DB) {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "principal",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(
                    &usersModels.User{}, 
                    &postsModels.Post{}, 
                    &postsModels.Comment{},
                )
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable(
					&usersModels.User{}, 
					&postsModels.Post{}, 
					&postsModels.Comment{},
                )
			},
		},
	})

	// Roda as migrações
	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	log.Println("Migration run successfully")
}