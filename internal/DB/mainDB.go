package DB

import (
	"gorm.io/gorm"
)

func GetMainDB() *gorm.DB {
	return connect("", false)
}
