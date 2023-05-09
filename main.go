package main

import (
	"context"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"geoloc/app/api/handlers"
	"geoloc/app/api/routes"
	"geoloc/app/services"
	"geoloc/app/stores"
	"geoloc/dbmigrate"
	"geoloc/settings"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
	"github.com/peterbourgon/ff/v3"
	"github.com/peterbourgon/ff/v3/ffcli"
)

const (
	envPrefix string = "newDB"
)

func main() {
	godotenv.Load(".env.local")
	godotenv.Load()

	var (
		rootFlagSet = flag.NewFlagSet("newDB", flag.ExitOnError)
		run         = rootFlagSet.Bool("run", false, "runs db run the server")
		migrateup   = rootFlagSet.Bool("migrateup", false, "migrates up db")
		migratedown = rootFlagSet.Bool("migratedown", false, "migrates down db")
	)
	rootCmd := &ffcli.Command{
		Name:      "root",
		Options:   []ff.Option{ff.WithEnvVarPrefix(envPrefix)},
		ShortHelp: "run root command",
		FlagSet:   rootFlagSet,
		Exec: func(_ context.Context, arg []string) error {
			if !*run && !*migrateup && !*migratedown {
				return errors.New("-dbseed or -run is required")
			}
			if *run {
				RunServer()
			}
			if *migrateup {
				_, _, dbSource := LoadConfig()
				dbmigrate.RunMigrateUp(dbSource)
			}
			if *migratedown {
				_, _, dbSource := LoadConfig()
				dbmigrate.RunMigrateDown(dbSource)
			}
			return nil
		},
	}
	if err := rootCmd.ParseAndRun(context.Background(), os.Args[1:]); err != nil {
		log.Fatal(err.Error())
	}

}

func RunServer() {
	dbClient, serverAddress, _ := LoadConfig()
	defer dbClient.Close()

	if dbClient == nil {
		log.Fatal("db connection failed")
	}

	store := stores.NewStore(dbClient)
	services := services.NewService(store)
	handler := handlers.NewHandler(services)
	rt := routes.NewRoute(handler)

	router := chi.NewRouter()

	router.Use(middleware.Logger)
	routes.Routes(router, rt)

	log.Printf("server running on port %s", serverAddress)
	if err := http.ListenAndServe(serverAddress, router); err != nil {
		log.Fatal(err)
	}
}

func LoadConfig() (*sql.DB, string, string) {
	serverAdd := os.Getenv("SERVER_ADDRESS")
	dbUserName := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbsource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUserName, dbPassword, dbName)
	dbClient, err := settings.ConnectPostgres(dbsource)
	if err != nil {
		log.Fatal(err)
	}
	return dbClient, serverAdd, dbsource

}
