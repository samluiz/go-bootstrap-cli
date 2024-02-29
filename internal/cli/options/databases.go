package options

type DatabaseModule struct {
	Name    string
	Package string
	ID      string
}

var (
	NO_DATABASE = DatabaseModule{
		Name:    "None",
		Package: "",
		ID:      "0",
	}
	MYSQL = DatabaseModule{
		Name:    "MySQL",
		Package: "github.com/go-sql-driver/mysql",
		ID:      "1",
	}
	POSTGRES = DatabaseModule{
		Name:    "Postgres",
		Package: "github.com/lib/pq",
		ID:      "2",
	}
	SQLITE = DatabaseModule{
		Name:    "SQLite",
		Package: "github.com/mattn/go-sqlite3",
		ID:      "3",
	}
	MONGODB = DatabaseModule{
		Name:    "MongoDB",
		Package: "go.mongodb.org/mongo-driver/mongo",
		ID:      "4",
	}
	REDIS = DatabaseModule{
		Name:    "Redis",
		Package: "github.com/go-redis/redis/v8",
		ID:      "5",
	}
)

func GetDatabaseModuleById(id string) DatabaseModule {
	switch id {
	case "0":
		return NO_DATABASE
	case "1":
		return MYSQL
	case "2":
		return POSTGRES
	case "3":
		return SQLITE
	case "4":
		return MONGODB
	case "5":
		return REDIS
	}
	return DatabaseModule{}

}

func GetDatabaseModuleNameById(id string) string {
	return GetDatabaseModuleById(id).Name
}

func (m *DatabaseModule) GetDatabaseModulePackageById(id string) string {
	return GetDatabaseModuleById(id).Package
}
