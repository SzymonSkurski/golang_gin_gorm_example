package album

import (
	"strings"
	"time"

	"gorm.io/gorm"
)

type Album struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Title     string         `json:"title" gorm:"index"`
	Artist    string         `json:"artist"`
	Slug      string         `json:"slug" gorm:"unique"`
	Price     int            `json:"price"`
}

func (album *Album) BeforeCreate(tx *gorm.DB) (err error) {
	//additional data validation here
	//create clug
	album.Slug = GetSlug(album.Title)
	return
}

func (album *Album) BeforeUpdate(tx *gorm.DB) (err error) {
	//update slug
	album.Slug = GetSlug(album.Title)
	return
}

func GetSlug(t string) string {
	return strings.ReplaceAll(t, " ", "_")
}
