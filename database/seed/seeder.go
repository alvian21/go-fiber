package seed

import (
	"fmt"
	"go-fiber-gorm/database"
)

func RunSeed() {
	RoleSeed(database.DB)
	PermissionSeed(database.DB)
	RolePermissionSeed(database.DB)
	UserSeed(database.DB)

	fmt.Println("Seeding complete")
}
