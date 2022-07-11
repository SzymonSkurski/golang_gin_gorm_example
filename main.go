package main

import (
	"os"

	"github.com/SzymonSkurski/golang_gin_gorm_example/internal/DB"
	handler "github.com/SzymonSkurski/golang_gin_gorm_example/internal/handlers"
	"github.com/SzymonSkurski/golang_gin_gorm_example/internal/handlers/albumHandler"
	"github.com/SzymonSkurski/golang_gin_gorm_example/internal/handlers/artistHandler"
	"github.com/gin-gonic/gin"
)

func main() {
	setDevEnv()
	migrate()
	router()
}

func router() {
	router := gin.Default() // default will use midleware
	api := router.Group("/api")
	apiArtist := api.Group("/albums")
	apiArtist.GET("artist/:id", albumHandler.GetAlbumsByArtistID)
	apiArtist.GET("/:needle", albumHandler.GetAlbumBy)
	apiArtist.GET("/", albumHandler.GetAlbums)
	apiArtist.POST("/", albumHandler.PostAlbums)
	apiArtist.DELETE("/:id", albumHandler.Delete)

	// router.SetTrustedProxies([]string{"192.168.1.1:8080"})
	router.GET("/migrate", handler.Migrate)

	//Artists
	router.GET("/artists/albums/:id", artistHandler.GetArtistsAlbums)
	router.GET("/artists/:needle", artistHandler.GetArtistBy)
	router.GET("/artists", artistHandler.GetArtists)
	router.POST("/artists", artistHandler.PostArtists)
	router.DELETE("artists/:id", artistHandler.Delete)
	router.Run(":8080") //don't use localhost:8080 to avoid (56) error in docker env
}

func migrate() {
	DB.GetRootDB().Exec("CREATE DATABASE IF NOT EXISTS example")
	db := DB.GetMainDB()
	artistHandler.Migrate(db)
	albumHandler.Migrate(db)
}

func setDevEnv() {
	env := map[string]string{
		"DB_PASSWORD": "root",
		"DB_USER":     "root",
		"DB_HOST":     "localhost",
		"DB_NAME":     "example",
		"DB_PORT":     "3306",
	}
	for e, v := range env {
		if os.Getenv(e) == "" {
			os.Setenv(e, v) // set default env
		}
	}
	// run docker container inspect mariadb-server bridge.IPAddress to get db i
}

// func isNumeric(s string) bool {
// 	_, err := strconv.ParseFloat(s, 64)
// 	return err == nil
// }
