package main

import (
	"os"

	handler "github.com/SzymonSkursrki/golang_gin_grom_example/internal/handlers"
	"github.com/SzymonSkursrki/golang_gin_grom_example/internal/handlers/albumHandler"
	"github.com/SzymonSkursrki/golang_gin_grom_example/internal/handlers/artistHandler"
	"github.com/gin-gonic/gin"
)

func main() {
	setDevEnv()
	router()
}

func router() {
	router := gin.Default()
	router.SetTrustedProxies([]string{"192.168.1.1:8080"})
	// router.GET("/migrate", handler.Migrate)
	// router.GET("/albums/artist/:id", albumHandler.GetAlbumsByArtistID)
	// router.GET("/albums/:needle", albumHandler.GetAlbumBy)
	// router.GET("/albums", albumHandler.GetAlbums)
	// router.POST("/albums", albumHandler.PostAlbums)
	// router.DELETE("albums/:id", albumHandler.Delete)
	// //Artists
	// router.GET("/artists/albums/:id", artistHandler.GetArtistByIDWithAlbums)
	// router.GET("/artists/:needle", artistHandler.GetArtistBy)
	// router.GET("/artists", artistHandler.GetArtists)
	// router.POST("/artists", artistHandler.PostArtists)
	// router.DELETE("artists/:id", artistHandler.Delete)
	router.Run("localhost:8080")
}

func getRoutesDefinition() ()

func setDevEnv() {
	os.Setenv("DB_PASSWORD", "")
	os.Setenv("DB_USER", "root")
	os.Setenv("DB_HOST", "127.0.0.1")
	os.Setenv("DB_NAME", "example")
	os.Setenv("DB_PORT", "3306")
}

// func isNumeric(s string) bool {
// 	_, err := strconv.ParseFloat(s, 64)
// 	return err == nil
// }
