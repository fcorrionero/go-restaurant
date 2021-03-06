//+build wireinject

package main

import (
	"github.com/fcorrionero/go-restaurant/src/application/command/add_allergen"
	"github.com/fcorrionero/go-restaurant/src/application/command/add_dish"
	"github.com/fcorrionero/go-restaurant/src/application/command/add_ingredient"
	"github.com/fcorrionero/go-restaurant/src/application/query/find_dish_by_id"
	"github.com/fcorrionero/go-restaurant/src/application/query/find_dish_by_name"
	"github.com/fcorrionero/go-restaurant/src/application/query/find_dishes_by_allergen"
	"github.com/fcorrionero/go-restaurant/src/domain"
	"github.com/fcorrionero/go-restaurant/src/infrastructure/persistence/mongo"
	"github.com/fcorrionero/go-restaurant/src/infrastructure/persistence/mysql"
	"github.com/fcorrionero/go-restaurant/src/infrastructure/ui/dishes_http"
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

func InitializeCrudHttpController() dishes_http.CrudHttpController {
	db := StartGormDB()
	rDRepo := NewMongoDishesRepository()
	dRepo := NewMysqlDishesRepository(db)
	iRepo := NewMysqlIngredientsRepository(db)
	aRepo := NewMysqlAllergensRepository(db)
	return dishes_http.NewCrudController(add_allergen.New(aRepo), add_ingredient.New(iRepo, aRepo), add_dish.New(dRepo, rDRepo, iRepo))
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
