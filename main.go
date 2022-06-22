package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/SzymonSkursrki/golang_gin_grom_example/internal/model/album"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	setDevEnv()
	router()
}

func router() {
	router := gin.Default()
	router.SetTrustedProxies([]string{"192.168.1.1:8080"})
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:needle", getAlbumBy)
	router.Run("localhost:8080")
}

func getAlbumBy(c *gin.Context) {
	albums := []album.Album{}
	needle := c.Param("needle")
	db := mainDB()
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
	//try related to artist
	if result := db.Where(&album.Album{Artist: needle}).Find(&albums); result.Error == nil {
		c.IndentedJSON(http.StatusOK, gin.H{"albums": albums})
		return
	}

	// for _, a := range albums {
	// 	if fmt.Sprintf("%v", a.ID) == needle || strings.EqualFold(a.Title, needle) || strings.EqualFold(getSlug(a.Title), getSlug(needle)) {
	// 		c.IndentedJSON(http.StatusOK, a)
	// 		return
	// 	}
	// }
	//nont found
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func getAlbums(c *gin.Context) {
	// Get all records
	albums := []album.Album{}
	db := mainDB()
	result := db.Find(&albums)
	// SELECT * FROM users;

	// result.RowsAffected // returns found records count, equals `len(users)`
	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": result.Error})
		return
	}
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album.Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	// TODO: Validation ?
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, err.Error())
		return
	}

	// Add the new album to the slice. 201 - status created
	// albums = append(albums, newAlbum)
	// c.IndentedJSON(http.StatusCreated, newAlbum)

	// dsn := "root:@tcp(127.0.0.1:3306)/example?charset=utf8mb4&parseTime=True&loc=Local"
	db := mainDB()
	migrateAlbums(db)

	// Create & insert
	result := db.Create(&newAlbum)
	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, result.Error)
		return
	}

	// newAlbum.ID             // returns inserted data's primary key
	// result.Error        // returns error
	// result.RowsAffected // returns inserted records count
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func mainDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(getMainDBDSN()), &gorm.Config{})
	// db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}
	return db
}

func migrateAlbums(db *gorm.DB) {
	// Migrate the schema
	e := db.AutoMigrate(&album.Album{})
	if e != nil {
		panic(fmt.Sprintf("failed to migrate database: %v", e))
	}
}

func getMainDBDSN() string {
	charset := "utf8mb4"
	parseTime := "true"
	loc := "Local"
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=%v&loc=%v",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		charset,
		parseTime,
		loc)
}

func setDevEnv() {
	os.Setenv("DB_PASSWORD", "")
	os.Setenv("DB_USER", "root")
	os.Setenv("DB_HOST", "127.0.0.1")
	os.Setenv("DB_NAME", "example")
	os.Setenv("DB_PORT", "3306")
}

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
