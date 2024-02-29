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
	"github.com/samluiz/goselcli/menu"
)

func Run(in io.Reader) {
	scanner := bufio.NewScanner(in)

	fmt.Println(goterm.Color(goterm.Bold("\nWelcome to the Goinit CLI!\n"), goterm.CYAN))

	fmt.Printf("\n%s\n", goterm.Color(goterm.Bold("Enter your new module name (eg: github.com/owner/repo):"), goterm.CYAN))
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
	m := menu.NewMenu("Select your web server")

	for _, server := range servers.Servers {
		m.AddOption(server.Name, server.ID)
	}

	choice := m.Display()

	return servers.GetServerById(choice)
}

func getDatabaseModule() databases.DatabaseModule {
	m := menu.NewMenu("Select your database (if any)")

	for _, db := range databases.Databases {
		m.AddOption(db.Name, db.ID)
	}

	database := m.Display()

	return databases.GetDatabaseModuleById(database)
}

func getDriverExtension(database databases.DatabaseModule) drivers.DriverExtensionModule {
	if database.ID == databases.NO_DATABASE.ID ||
		database.ID == databases.POSTGRES_PGX.ID ||
		database.ID == databases.MONGODB.ID {
		return drivers.NO_DRIVER_EXT
	}

	m := menu.NewMenu("Select your database driver extension (if any)")

	for _, driver := range drivers.DriverExtensions {
		m.AddOption(driver.Name, driver.ID)
	}

	choice := m.Display()

	return drivers.GetDriverExtensionModuleById(choice)
}

func getSqlCodegens(database databases.DatabaseModule) sqlcodegens.SQLCodeGenModule {
	if database.ID == databases.NO_DATABASE.ID ||
		database.ID == databases.MONGODB.ID {
		return sqlcodegens.NO_SQL_CODEGEN
	}

	m := menu.NewMenu("Select your SQL code generator (if any)")

	for _, sqlCodeGen := range sqlcodegens.SQLCodeGens {
		m.AddOption(sqlCodeGen.Name, sqlCodeGen.ID)
	}

	sqlCodeGen := m.Display()

	return sqlcodegens.GetSQLCodeGenModuleById(sqlCodeGen)
}
