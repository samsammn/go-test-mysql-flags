package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/samsammn/project7-test/config"
	"github.com/samsammn/project7-test/controller"
	"github.com/samsammn/project7-test/database"
	"github.com/samsammn/project7-test/repository"
	"github.com/samsammn/project7-test/router"
)

func main() {
	if len(os.Args) < 1 {
		fmt.Println("your arguments is nil, please type run or config.")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "run":
		flagRun()
	case "config":
		flagConfig()
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func appRun() {
	db := database.ConnectDB()
	userRepo := repository.NewUserRepository(db)
	userController := controller.NewUserController(userRepo)

	app := router.NewHttpRouter()

	app.GET("/users", userController.FindAll)
	app.POST("/user", userController.Store)

	app.SERVE(":8888")
}

func flagRun() {
	runCommand := flag.NewFlagSet("run", flag.ExitOnError)
	appDebug := runCommand.Bool("debug", true, "set your project mode to debug")

	runCommand.Parse(os.Args[2:])

	if *appDebug == true {
		fmt.Println("Your application in mode Debug")
	}

	appRun()
}

func flagConfig() {
	configCommand := flag.NewFlagSet("config", flag.ExitOnError)
	switch os.Args[2] {
	case "set":
		appName := configCommand.String("app_name", "Simple App Zam", "set your application name")
		appVersion := configCommand.String("app_version", "v0.0.0", "set your application version")
		dbDriver := configCommand.String("db_driver", "mysql", "set database application driver")
		dbName := configCommand.String("db_name", "test", "set application database name")
		dbUser := configCommand.String("db_user", "root", "set application database user")
		dbPass := configCommand.String("db_pass", "", "set application database pass")
		dbHost := configCommand.String("db_host", "localhost", "set application database host")
		dbPort := configCommand.Int("db_port", 3306, "set application database port")

		configCommand.Parse(os.Args[3:])

		configs := config.Config{}

		if *appName != "Simple App Zam" {
			configs.AppName = *appName
		}

		if *appVersion != "v0.0.0" {
			configs.AppVersion = *appVersion
		}

		if *dbDriver != "mysql" {
			configs.DbDriver = *dbDriver
		}

		if *dbName != "test" {
			configs.DbName = *dbName
		}

		if *dbUser != "root" {
			configs.DbUser = *dbUser
		}

		if *dbPass != "" {
			configs.DbPass = *dbPass
		}

		if *dbHost != "localhost" {
			configs.DbHost = *dbHost
		}

		if *dbPort != 3306 {
			configs.DbPort = *dbPort
		}

		config.Set(configs)

	case "get":
		config.Get()

	default:
		fmt.Println("please type command is set or get only")
		os.Exit(1)
	}
}
