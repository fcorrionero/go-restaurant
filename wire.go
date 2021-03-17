//+build wireinject

package main

import (
	"database/sql"
	"github.com/fcorrionero/go-restaurant/application/query/find_dish_by_id"
	"github.com/fcorrionero/go-restaurant/application/query/find_dish_by_name"
	"github.com/fcorrionero/go-restaurant/application/query/find_dishes_by_allergen"
	"github.com/fcorrionero/go-restaurant/domain"
	"github.com/fcorrionero/go-restaurant/infrastructure/persistence/mongo"
	"github.com/fcorrionero/go-restaurant/infrastructure/persistence/mysql"
	"github.com/fcorrionero/go-restaurant/infrastructure/ui/dishes_http"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	dMysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func InitializeDishesRepository() domain.DishesRepository {
	wire.Build(NewMongoDishesRepository)
	return mongo.DishesRepository{}
}

func InitializeDishesHttpController(dishesRepository domain.DishesRepository) dishes_http.DishesHttpController {
	wire.Build(dishes_http.NewDishesHttpController, find_dish_by_id.New, find_dish_by_name.New, find_dishes_by_allergen.New)
	return dishes_http.DishesHttpController{}
}

func NewMongoDishesRepository() domain.DishesRepository {
	return mongo.New("root", "example", "0.0.0.0", "27017", "go-restaurant")
}

func NewMysqlAllergensRepository(db *gorm.DB) domain.AllergensRepository {
	return mysql.NewAllergensRepository(db)
}

func NewMysqlIngredientsRepository(db *gorm.DB) domain.IngredientsRepository {
	return mysql.NewIngredientsRepository(db)
}

func NewMysqlDishesRepository(db *gorm.DB) domain.DishesRepository {
	return mysql.NewDishesRepository(db)
}

func StartMysqlDB() *sql.DB {
	db, err := sql.Open("mysql", "root:example@/go_restaurant")
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func StartGormDB() *gorm.DB {
	dsn := "root:example@/go_restaurant?charset=utf8mb4&parseTime=True&loc=Local"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,  // Slow SQL threshold
			LogLevel:      logger.Error, // Log level
			Colorful:      false,        // Disable color
		},
	)
	db, err := gorm.Open(dMysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
