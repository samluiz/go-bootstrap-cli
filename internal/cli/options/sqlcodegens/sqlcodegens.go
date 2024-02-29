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
	SQLC = SQLCodeGenModule{
		Name:    "SQLC",
		Package: "github.com/kyleconroy/sqlc",
		ID:      "1",
	}
	SQUIRREL = SQLCodeGenModule{
		Name:    "Squirrel",
		Package: "github.com/Masterminds/squirrel",
		ID:      "2",
	}
	SQLZ = SQLCodeGenModule{
		Name:    "SQLZ",
		Package: "github.com/ulule/sqlz",
		ID:      "3",
	}
	GOQU = SQLCodeGenModule{
		Name:    "Goqu",
		Package: "github.com/doug-martin/goqu",
		ID:      "4",
	}
	HOOD = SQLCodeGenModule{
		Name:    "Hood",
		Package: "github.com/eaigner/hood",
		ID:      "5",
	}
)

var SQLCodeGens = []SQLCodeGenModule{
	NO_SQL_CODEGEN,
	SQLC,
	SQUIRREL,
	SQLZ,
}

func GetSQLCodeGenModuleById(id string) SQLCodeGenModule {
	switch id {
	case "0":
		return NO_SQL_CODEGEN
	case "1":
		return SQLC
	case "2":
		return SQUIRREL
	case "3":
		return SQLZ
	default:
		return SQLCodeGenModule{}
	}
}
