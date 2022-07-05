package albumHandler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/SzymonSkursrki/golang_gin_grom_example/internal/DB/mainDB"
	"github.com/SzymonSkursrki/golang_gin_grom_example/internal/model/album"
	"github.com/SzymonSkursrki/golang_gin_grom_example/internal/paginator"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func getDBModel(db *gorm.DB) (tx *gorm.DB) {
	return db.Model(album.Album{})
}

func GetAlbums(c *gin.Context) {
	// Get all records
	albums := []album.Album{}
	db := mainDB.GetDB()

	result := db.Scopes(paginator.Paginate(c.Request)).Find(&albums)
	// SELECT * FROM users;

	// result.RowsAffected // returns found records count, equals `len(users)`
	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": result.Error})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"albums": albums, "paginator": paginator.PaginateInfo(c.Request, getDBModel(db))})
	}
}

func GetAlbumsByArtistID(c *gin.Context) {
	// Get all records related to artist
	id := c.Param("id")
	albums := []album.Album{}
	db := mainDB.GetDB()
	result := db.Where("artist_id = ?", id).Find(&albums)
	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": result.Error})
	} else {
		c.IndentedJSON(http.StatusOK, albums)
	}
}

// postAlbums adds an album from JSON received in the request body.
func PostAlbums(c *gin.Context) {
	var newAlbum album.Album
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, err.Error())
		return
	}
	db := mainDB.GetDB()

	// Create & insert
	if result := db.Create(&newAlbum); result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, result.Error)
		return
	}

	// newAlbum.ID             // returns inserted data's primary key
	// result.Error        		// returns error
	// result.RowsAffected // returns inserted records count
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func GetAlbumBy(c *gin.Context) {
	albums := []album.Album{}
	needle := c.Param("needle")
	db := mainDB.GetDB()
	if id, err := strconv.ParseUint(needle, 0, 64); err == nil {
		//get by ID
		a := album.Album{ID: uint(id)}
		if result := db.First(&a); result.Error == nil {
			albums = append(albums, a)
			c.IndentedJSON(http.StatusOK, gin.H{"albums": albums})
			return
		}
	}
	//try get by title or slug
	a := album.Album{}
	if result := db.Where(&album.Album{Title: needle}).Or(&album.Album{Slug: album.GetSlug(needle)}).First(&a); result.Error == nil {
		albums = append(albums, a)
		c.IndentedJSON(http.StatusOK, gin.H{"albums": albums})
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func Delete(c *gin.Context) {
	id := c.Param("id")
	db := mainDB.GetDB()

	if res := db.Delete(&album.Album{}, id); res.Error == nil {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "deleted"})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"error": res.Error})
	}
}

func Migrate(db *gorm.DB) {
	// Migrate the schema
	e := db.AutoMigrate(&album.Album{})
	if e != nil {
		panic(fmt.Sprintf("failed to migrate database: %v", e))
	}
}
