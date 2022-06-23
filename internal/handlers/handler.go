package handler

import (
	"github.com/SzymonSkursrki/golang_gin_grom_example/internal/DB/mainDB"
	"github.com/SzymonSkursrki/golang_gin_grom_example/internal/handlers/albumHandler"
	"github.com/SzymonSkursrki/golang_gin_grom_example/internal/handlers/artistHandler"
	"github.com/gin-gonic/gin"
)

func Migrate(c *gin.Context) {
	db := mainDB.GetDB()
	artistHandler.Migrate(db)
	albumHandler.Migrate(db)

}
