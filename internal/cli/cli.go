package cli

import (
	"bufio"
	"fmt"
	"io"

	"github.com/buger/goterm"
	"github.com/samluiz/goinit/internal/cli/commands"
	"github.com/samluiz/goinit/internal/cli/options/databases"
	"github.com/samluiz/goinit/internal/cli/options/drivers"
	"github.com/samluiz/goinit/internal/cli/options/orms"
	"github.com/samluiz/goinit/internal/cli/options/servers"
	"github.com/samluiz/goinit/internal/cli/options/sqlcodegens"
)

func Run(in io.Reader) {
	scanner := bufio.NewScanner(in)

	fmt.Println(goterm.Color(goterm.Bold("Welcome to the Go Starter CLI!\n"), goterm.CYAN))

	fmt.Printf("\n%s\n", goterm.Color(goterm.Bold("Enter your new module name:"), goterm.CYAN))
	scanner.Scan()
	module := scanner.Text()

	fullPath, dir, err := commands.CreateProjectDir(module)
	if err != nil {
		commands.Fatal(fullPath, err)
	}

	fmt.Println(goterm.Color(fmt.Sprintf("Created directory: %s\n", fullPath), goterm.BLUE))

	var server servers.ServerModule
	var database databases.DatabaseModule
	var driver drivers.DriverExtensionModule
	var orm orms.ORMModule
	var sqlCodeGen sqlcodegens.SQLCodeGenModule

	server = getServer()
	database = getDatabaseModule()
	driver = getDriverExtension(database)
	sqlCodeGen = getSqlCodegens(database)

	fmt.Println(goterm.Color(goterm.Bold("Generating your project..."), goterm.CYAN))

	packages := []string{server.Package, database.Package, driver.Package, orm.Package, sqlCodeGen.Package}

	commands.GenerateGoProject(module, dir, fullPath, packages)
}

func getServer() servers.ServerModule {
	menu := NewMenu("Select your web server")

	for _, server := range servers.Servers {
		menu.addOption(server.Name, server.ID)
	}

	choice := menu.Display()

	return servers.GetServerById(choice)
}

func getDatabaseModule() databases.DatabaseModule {
	menu := NewMenu("Select your database (if any)")

	for _, db := range databases.Databases {
		menu.addOption(db.Name, db.ID)
	}

	database := menu.Display()

	return databases.GetDatabaseModuleById(database)
}

func getDriverExtension(database databases.DatabaseModule) drivers.DriverExtensionModule {
	if database.ID == databases.NO_DATABASE.ID ||
		database.ID == databases.POSTGRES_PGX.ID ||
		database.ID == databases.MONGODB.ID {
		return drivers.NO_DRIVER_EXT
	}

	menu := NewMenu("Select your database driver extension (if any)")

	for _, driver := range drivers.DriverExtensions {
		menu.addOption(driver.Name, driver.ID)
	}

	choice := menu.Display()

	return drivers.GetDriverExtensionModuleById(choice)
}

func getSqlCodegens(database databases.DatabaseModule) sqlcodegens.SQLCodeGenModule {
	if database.ID == databases.NO_DATABASE.ID ||
		database.ID == databases.MONGODB.ID {
		return sqlcodegens.NO_SQL_CODEGEN
	}

	menu := NewMenu("Select your SQL code generator (if any)")

	for _, sqlCodeGen := range sqlcodegens.SQLCodeGens {
		menu.addOption(sqlCodeGen.Name, sqlCodeGen.ID)
	}

	sqlCodeGen := menu.Display()

	return sqlcodegens.GetSQLCodeGenModuleById(sqlCodeGen)
}
