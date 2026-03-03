package seed

import (
	"fmt"
	"go-fiber-gorm/database"
)

func RunSeed() {
	UserSeed(database.DB)

	fmt.Println("Seeding complete")
}
