package sqlcodegens

type SQLCodeGenModule struct {
	Name    string
	Package string
	ID      string
}

var (
	NO_SQL_CODEGEN = SQLCodeGenModule{
		Name:    "No",
		Package: "",
		ID:      "0",
	}
	SQUIRREL = SQLCodeGenModule{
		Name:    "Squirrel",
		Package: "github.com/Masterminds/squirrel",
		ID:      "1",
	}
	SQLZ = SQLCodeGenModule{
		Name:    "SQLZ",
		Package: "github.com/ido50/sqlz",
		ID:      "2",
	}
	GOQU = SQLCodeGenModule{
		Name:    "Goqu",
		Package: "github.com/doug-martin/goqu/v9",
		ID:      "3",
	}
)

var SQLCodeGens = []SQLCodeGenModule{
	NO_SQL_CODEGEN,
	SQUIRREL,
	SQLZ,
	GOQU,
}

func GetSQLCodeGenModuleById(id string) SQLCodeGenModule {
	switch id {
	case "0":
		return NO_SQL_CODEGEN
	case "1":
		return SQUIRREL
	case "2":
		return SQLZ
	case "3":
		return GOQU
	default:
		return SQLCodeGenModule{}
	}
}
