package handler

import (
	"github.com/SzymonSkurski/golang_gin_gorm_example/internal/DB"
	"github.com/SzymonSkurski/golang_gin_gorm_example/internal/handlers/albumHandler"
	"github.com/SzymonSkurski/golang_gin_gorm_example/internal/handlers/artistHandler"
	"github.com/gin-gonic/gin"
)

func Migrate(c *gin.Context) {
	db := DB.GetMainDB()
	artistHandler.Migrate(db)
	albumHandler.Migrate(db)
}
