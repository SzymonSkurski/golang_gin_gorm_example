package DB

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetRootDB() *gorm.DB {
	return connect("", true)
}

func connect(prefix string, root bool) *gorm.DB {
	db, err := gorm.Open(mysql.Open(getDSN(prefix, root)), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect to the database: %v", err))
	}
	return db
}

func getDSN(postfix string, root bool) string {
	charset := "utf8mb4"
	parseTime := "true"
	loc := "Local"
	dbName := "/"
	if !root {
		dbName = fmt.Sprintf("/%v", getEnvParam("NAME", postfix))
	}
	return fmt.Sprintf("%v:%v@tcp(%v:%v)%v?charset=%v&parseTime=%v&loc=%v",
		getEnvParam("USER", postfix),
		getEnvParam("PASSWORD", postfix),
		getEnvParam("HOST", postfix),
		getEnvParam("PORT", postfix),
		dbName,
		charset,
		parseTime,
		loc)
}

func getEnvParam(p string, postfix string) string {
	if len(postfix) > 0 {
		postfix = fmt.Sprintf("_%v", postfix) // _dev for example
	}
	return os.Getenv(fmt.Sprintf("DB_%v%v", p, postfix))
}
