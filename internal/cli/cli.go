package cli

import (
	"bufio"
	"fmt"
	"io"

	"github.com/buger/goterm"
	"github.com/samluiz/go-bootstrap-cli/internal/cli/commands"
	"github.com/samluiz/go-bootstrap-cli/internal/cli/options"
)

func Run(in io.Reader) {
	scanner := bufio.NewScanner(in)

	fmt.Println("Welcome to the Go Starter CLI!")

	fmt.Printf("%s\n", goterm.Color(goterm.Bold("Enter your new module name:"), goterm.CYAN))
	scanner.Scan()
	module := scanner.Text()

	server := getServer()
	database := getDatabaseModule()
	sqlTool, queryBuilder := getSqlTools(database.ID)

	if sqlTool.ID == options.PGX.ID {
		database.Package = ""
	}

	fmt.Println("Scaffolding your project...")

	packages := []string{server.Package, database.Package, sqlTool.Package, queryBuilder.Package}

	commands.GenerateGoProject(module, packages)
}

func getServer() options.WebServerModule {
	menu := NewMenu("Select your web server")

	menu.addOption(options.STDLIB.Name, options.STDLIB.ID)
	menu.addOption(options.FIBER.Name, options.FIBER.ID)
	menu.addOption(options.GIN.Name, options.GIN.ID)
	menu.addOption(options.CHI.Name, options.CHI.ID)
	menu.addOption(options.ECHO.Name, options.ECHO.ID)
	menu.addOption(options.GORILLA.Name, options.GORILLA.ID)
	menu.addOption(options.IRIS.Name, options.IRIS.ID)
	menu.addOption(options.MUX.Name, options.MUX.ID)
	menu.addOption(options.AERO.Name, options.AERO.ID)
	menu.addOption(options.FASTHTTP.Name, options.FASTHTTP.ID)
	menu.addOption(options.BEEGO.Name, options.BEEGO.ID)

	choice := menu.Display()

	return options.GetServerById(choice)
}

func getDatabaseModule() options.DatabaseModule {
	menu := NewMenu("Select your database")

	menu.addOption(options.NO_DATABASE.Name, options.NO_DATABASE.ID)
	menu.addOption(options.MYSQL.Name, options.MYSQL.ID)
	menu.addOption(options.POSTGRES.Name, options.POSTGRES.ID)
	menu.addOption(options.SQLITE.Name, options.SQLITE.ID)
	menu.addOption(options.MONGODB.Name, options.MONGODB.ID)
	menu.addOption(options.REDIS.Name, options.REDIS.ID)

	choice := menu.Display()

	return options.GetDatabaseModuleById(choice)
}

func getSqlTools(databaseId string) (options.SQLToolModule, options.SQLToolModule) {
	if databaseId == options.NO_DATABASE.ID {
		return options.SQLToolModule{}, options.SQLToolModule{}
	}

	menu := NewMenu("Select your SQL tool")

	menu.addOption(options.NO_SQL_TOOL.Name, options.NO_SQL_TOOL.ID)
	menu.addOption(options.SQLX.Name, options.SQLX.ID)
	if databaseId == options.POSTGRES.ID {
		menu.addOption(options.PGX.Name, options.PGX.ID)
	}
	menu.addOption(options.GORM.Name, options.GORM.ID)
	menu.addOption(options.SQLBOILER.Name, options.SQLBOILER.ID)
	menu.addOption(options.ENT.Name, options.ENT.ID)

	sqlTool := options.GetSQLToolById(menu.Display())

	menu = NewMenu("Select a query builder")

	if sqlTool.ID == options.GORM.ID || sqlTool.ID == options.ENT.ID || sqlTool.ID == options.SQLBOILER.ID {
		return sqlTool, options.SQLToolModule{}
	}

	menu.addOption(options.NO_SQL_TOOL.Name, options.NO_SQL_TOOL.ID)
	menu.addOption(options.SQLC.Name, options.SQLC.ID)
	menu.addOption(options.SQUIRREL.Name, options.SQUIRREL.ID)
	menu.addOption(options.SQLZ.Name, options.SQLZ.ID)

	queryBuilder := options.GetSQLToolById(menu.Display())

	return sqlTool, queryBuilder
}
