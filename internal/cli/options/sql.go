package options

type SQLToolModule struct {
	Name    string
	Package string
	ID      string
}

var (
	NO_SQL_TOOL = SQLToolModule{
		Name:    "None",
		Package: "",
		ID:      "0",
	}
	SQLX = SQLToolModule{
		Name:    "SQLX",
		Package: "github.com/jmoiron/sqlx",
		ID:      "1",
	}
	PGX = SQLToolModule{
		Name:    "PGX",
		Package: "github.com/jackc/pgx",
		ID:      "2",
	}
	GORM = SQLToolModule{
		Name:    "GORM",
		Package: "github.com/go-gorm/gorm",
		ID:      "3",
	}
	SQLBOILER = SQLToolModule{
		Name:    "SQLBoiler",
		Package: "github.com/volatiletech/sqlboiler",
		ID:      "4",
	}
	ENT = SQLToolModule{
		Name:    "Ent",
		Package: "github.com/ent/ent",
		ID:      "5",
	}
	SQLC = SQLToolModule{
		Name:    "SQLC",
		Package: "github.com/kyleconroy/sqlc",
		ID:      "6",
	}
	SQUIRREL = SQLToolModule{
		Name:    "Squirrel",
		Package: "github.com/Masterminds/squirrel",
		ID:      "7",
	}
	SQLZ = SQLToolModule{
		Name:    "SQLZ",
		Package: "github.com/ulule/sqlz",
		ID:      "8",
	}
)

func GetSQLToolById(id string) SQLToolModule {
	switch id {
	case "0":
		return NO_SQL_TOOL
	case "1":
		return SQLX
	case "2":
		return PGX
	case "3":
		return GORM
	case "4":
		return SQLBOILER
	case "5":
		return ENT
	case "6":
		return SQLC
	case "7":
		return SQUIRREL
	case "8":
		return SQLZ
	}
	return SQLToolModule{}
}

func GetSQLToolModuleNameById(id string) string {
	return GetSQLToolById(id).Name
}

func (m *SQLToolModule) GetSQLToolModulePackageById(id string) string {
	return GetSQLToolById(id).Package
}
