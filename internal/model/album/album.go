package album

import (
	"strings"
	"time"

	"gorm.io/gorm"
)

type Album struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	Title     string         `json:"title" gorm:"index"`
	ArtistID  uint           `json:"artistID" gorm:"default:0;index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;check:artist_id > 0"`
	Slug      string         `json:"slug" gorm:"unique"`
	Price     int            `json:"price"`
}

func (album *Album) BeforeCreate(tx *gorm.DB) (err error) {
	//create slug
	updateSlug(album)
	return nil
}

func (album *Album) BeforeUpdate(tx *gorm.DB) (err error) {
	//update slug
	updateSlug(album)
	return nil
}

func updateSlug(album *Album) {
	album.Slug = GetSlug(album.Title)
}

func GetSlug(t string) string {
	return strings.ReplaceAll(t, " ", "_")
}
