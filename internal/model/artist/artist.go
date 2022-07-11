package artist

import (
	"fmt"
	"strings"
	"time"

	"github.com/SzymonSkurski/golang_gin_gorm_example/internal/model/album"
	"gorm.io/gorm"
)

// Artist could own many albums oneToMany
type Artist struct {
	ID        uint `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `gorm:"not null;check:name !=\"\"" json:"name"`
	Surname   string         `gorm:"not null;check:surname !=\"\"" json:"surname"`
	Slug      string         `gorm:"not null; index; unique" json:"slug"`
	BirthDate time.Time      `gorm:"not null" json:"birthDate"`
	DeathDate time.Time      `gorm:"default:null;check:death_date >= birth_date" json:"deathDate"`
	Albums    []album.Album  `gorm:"foreignKey:ArtistID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (a *Artist) BeforeCreate(tx *gorm.DB) (err error) {
	//create slug
	updateSlug(a)
	return
}

func (a *Artist) BeforeUpdate(tx *gorm.DB) (err error) {
	//update slug
	updateSlug(a)
	return
}

func updateSlug(a *Artist) {
	a.Slug = GetSlug(fmt.Sprintf("%v %v", a.Name, a.Surname))
}

func GetSlug(s string) string {
	return strings.ReplaceAll(s, " ", "_")
}
