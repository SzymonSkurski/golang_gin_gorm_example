package artist

import (
	"fmt"
	"time"

	"github.com/SzymonSkursrki/golang_gin_grom_example/internal/model/album"
	"gorm.io/gorm"
)

type Artist struct {
	ID        uint `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `gorm:"not null;check:name !=\"\"" json:"name"`
	Surname   string         `gorm:"not null;check:surname !=\"\"" json:"surname"`
	Slug      string         `gorm:"not null; index; unique" json:"slug"`
	BirthDate time.Time      `gorm:"not null" json:"birthDate"`
	DeathDate time.Time      `gorm:"default:null" json:"deathDate"`
	Albums    []album.Album
}

func (a *Artist) BeforeCreate(tx *gorm.DB) (err error) {
	//create slug
	a.Slug = GetSlug(a)
	return
}

func (a *Artist) BeforeUpdate(tx *gorm.DB) (err error) {
	//update slug
	a.Slug = GetSlug(a)
	return
}

func GetSlug(a *Artist) string {
	return fmt.Sprintf("%v_%v", a.Name, a.Surname)
}
