package artistHandler

import (
	"fmt"
	"net/http"

	"github.com/SzymonSkursrki/golang_gin_grom_example/internal/DB/mainDB"
	"github.com/SzymonSkursrki/golang_gin_grom_example/internal/model/artist"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAlbums(c *gin.Context) {
	// Get all records
	artists := []artist.Artist{}
	db := mainDB.GetDB()
	result := db.Find(&artists)
	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": result.Error})
	} else {
		c.IndentedJSON(http.StatusOK, artists)
	}
}

// postAlbums adds an album from JSON received in the request body.
func PostArtists(c *gin.Context) {
	var newArtist artist.Artist

	// Call BindJSON to bind the received JSON to
	if err := c.BindJSON(&newArtist); err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, err.Error())
		return
	}

	db := mainDB.GetDB()
	migrate(db)

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
	db := mainDB.GetDB()

	if res := db.Delete(&artist.Artist{}, id); res.Error == nil {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "deleted"})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"error": res.Error})
	}
}

func migrate(db *gorm.DB) {
	// Migrate the schema
	e := db.AutoMigrate(&artist.Artist{})
	if e != nil {
		panic(fmt.Sprintf("failed to migrate database: %v", e))
	}
}
