package databases

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
	POSTGRES_PGX = DatabaseModule{
		Name:    "Postgres with PGX",
		Package: "github.com/jackc/pgx/v4",
		ID:      "3",
	}
	SQLITE = DatabaseModule{
		Name:    "SQLite",
		Package: "github.com/mattn/go-sqlite3",
		ID:      "4",
	}
	MONGODB = DatabaseModule{
		Name:    "MongoDB",
		Package: "go.mongodb.org/mongo-driver/mongo",
		ID:      "5",
	}
)

var Databases = []DatabaseModule{
	NO_DATABASE,
	MYSQL,
	POSTGRES,
	POSTGRES_PGX,
	SQLITE,
	MONGODB,
}

func GetDatabaseModuleById(id string) DatabaseModule {
	switch id {
	case "0":
		return NO_DATABASE
	case "1":
		return MYSQL
	case "2":
		return POSTGRES
	case "3":
		return POSTGRES_PGX
	case "4":
		return SQLITE
	case "5":
		return MONGODB
	}
	return DatabaseModule{}
}
