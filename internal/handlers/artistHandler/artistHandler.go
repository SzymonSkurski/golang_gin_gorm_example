package artistHandler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/SzymonSkurski/golang_gin_gorm_example/internal/DB"
	"github.com/SzymonSkurski/golang_gin_gorm_example/internal/model/artist"
	"github.com/SzymonSkurski/golang_gin_gorm_example/internal/paginator"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func getDBModel(db *gorm.DB) (tx *gorm.DB) {
	return db.Model(artist.Artist{})
}

func GetArtists(c *gin.Context) {
	// Get all records
	artists := []artist.Artist{}
	db := DB.GetMainDB()
	result := db.Scopes(paginator.Paginate(c.Request)).Find(&artists)
	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": result.Error})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"artists": artists, "paginator": paginator.PaginateInfo(c.Request, getDBModel(db))})
	}
}

func GetArtistBy(c *gin.Context) {
	artists := []artist.Artist{}
	needle := c.Param("needle")
	db := DB.GetMainDB()
	if id, err := strconv.ParseUint(needle, 0, 64); err == nil {
		//get by ID
		a := artist.Artist{ID: uint(id)}
		if result := db.First(&a); result.Error == nil {
			artists = append(artists, a)
			c.IndentedJSON(http.StatusOK, gin.H{"artists": artists})
			return
		}
	}
	//try get by name or surname or by slug
	a := artist.Artist{}
	if result := db.Where(&artist.Artist{Name: needle}).
		Or(&artist.Artist{Surname: needle}).
		Or(&artist.Artist{Slug: artist.GetSlug(needle)}).
		First(&a); result.Error == nil {
		artists = append(artists, a)
		c.IndentedJSON(http.StatusOK, gin.H{"artists": artists})
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "artist not found"})
}

func GetArtistsAlbums(c *gin.Context) {
	artists := []artist.Artist{}
	sID := c.Param("id")
	id, _ := strconv.ParseUint(sID, 0, 64)
	db := DB.GetMainDB()
	//get artist by ID and preload related
	a := artist.Artist{ID: uint(id)}
	if result := db.Preload("Albums").First(&a); result.Error == nil {
		artists = append(artists, a)
	}
	c.IndentedJSON(http.StatusOK, gin.H{"artists": artists})
}

// postAlbums adds an album from JSON received in the request body.
func PostArtists(c *gin.Context) {
	var newArtist artist.Artist

	// Call BindJSON to bind the received JSON to
	if err := c.BindJSON(&newArtist); err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, err.Error())
		return
	}

	db := DB.GetMainDB()

	// Create & insert
	result := db.Create(&newArtist)
	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, result.Error)
		return
	}
	c.IndentedJSON(http.StatusCreated, newArtist)
}

func Delete(c *gin.Context) {
	id := c.Param("id")
	db := DB.GetMainDB()

	if res := db.Delete(&artist.Artist{}, id); res.Error == nil {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "deleted"})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"error": res.Error})
	}
}

func Migrate(db *gorm.DB) {
	// Migrate the schema
	e := db.AutoMigrate(&artist.Artist{})
	if e != nil {
		panic(fmt.Sprintf("failed to migrate database: %v", e))
	}
}
