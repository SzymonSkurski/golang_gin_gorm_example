package main

import (
	"os"

	"github.com/SzymonSkursrki/golang_gin_grom_example/internal/DB/mainDB"
	handler "github.com/SzymonSkursrki/golang_gin_grom_example/internal/handlers"
	"github.com/SzymonSkursrki/golang_gin_grom_example/internal/handlers/albumHandler"
	"github.com/SzymonSkursrki/golang_gin_grom_example/internal/handlers/artistHandler"
	"github.com/gin-gonic/gin"
)

func main() {
	setDevEnv()
	migrate()
	router()
}

func router() {
	router := gin.Default()
	// router.SetTrustedProxies([]string{"192.168.1.1:8080"})
	router.GET("/migrate", handler.Migrate)
	router.GET("/albums/artist/:id", albumHandler.GetAlbumsByArtistID)
	router.GET("/albums/:needle", albumHandler.GetAlbumBy)
	router.GET("/albums", albumHandler.GetAlbums)
	router.POST("/albums", albumHandler.PostAlbums)
	router.DELETE("albums/:id", albumHandler.Delete)
	//Artists
	router.GET("/artists/albums/:id", artistHandler.GetArtistsAlbums)
	router.GET("/artists/:needle", artistHandler.GetArtistBy)
	router.GET("/artists", artistHandler.GetArtists)
	router.POST("/artists", artistHandler.PostArtists)
	router.DELETE("artists/:id", artistHandler.Delete)
	router.Run("0.0.0.0:8080") //don't use localhost:8080 to avoid (56) error in docker environment
}

func migrate() {
	db := mainDB.GetDB()
	db.Exec("CREATE DATABASE IF NOT EXISTS example")
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
