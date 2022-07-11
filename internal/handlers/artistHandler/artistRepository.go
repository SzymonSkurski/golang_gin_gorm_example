package artistHandler

import (
	"fmt"

	"github.com/SzymonSkurski/golang_gin_grom_example/internal/DB"
	"github.com/SzymonSkurski/golang_gin_grom_example/internal/model/artist"
)

func GetArtistByID(id uint) (error, artist.Artist) {
	a := artist.Artist{ID: id}
	db := DB.GetMainDB()
	if result := db.First(&a); result.Error != nil {
		return fmt.Errorf("cannot get artist #%v; error:%v", id, result.Error), a
	}
	return nil, a
}
