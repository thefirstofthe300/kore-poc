//go:generate goagen bootstrap -d github.com/thefirstofthe300/kore-poc/koredata-goa/design

package main

import (
	"log"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/hegemone/kore-poc/koredata-goa/app"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	// Create service
	service := goa.New("koredata")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.LogResponse())
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	jwtMiddleware, _ := NewJWTMiddleware()
	app.UseJWTMiddleware(service, jwtMiddleware)

	app.UseBasicAuthMiddleware(service, NewBasicAuthMiddleware())

	db, err := gorm.Open("sqlite3", ":memory:")

	if err != nil {
		log.Fatalf("unable to open database: %s", err)
	}

	defer db.Close()

	db.AutoMigrate(&app.Quote{}, &app.Suggestion{})

	name1 := "Danny"
	name2 := "Jack"
	quote1 := "Wassup"
	quote2 := "Hit the road"

	danny := app.Quote{Name: &name1, Quote: &quote1}
	db.Create(&danny)
	jack := app.Quote{Name: &name2, Quote: &quote2}
	db.Create(&jack)

	// Mount "quote" controller
	c, err := NewQuoteController(service, db)

	if err != nil {
		service.LogError("Unable to create QuoteController: ", err)
	}

	d := NewSuggestionController(service, db)

	if err != nil {
		service.LogError("Unable to create QuoteController: ", err)
	}

	app.MountQuoteController(service, c)
	app.MountSuggestionController(service, d)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}
