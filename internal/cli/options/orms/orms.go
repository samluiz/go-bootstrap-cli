package orms

type ORMModule struct {
	Name    string
	Package string
	ID      string
}

var (
	NO_ORM = ORMModule{
		Name:    "No",
		Package: "",
		ID:      "0",
	}
	GORM = ORMModule{
		Name:    "GORM",
		Package: "gorm.io/gorm",
		ID:      "1",
	}
	SQLBOILER = ORMModule{
		Name:    "SQLBoiler",
		Package: "github.com/volatiletech/sqlboiler",
		ID:      "2",
	}
	ENT = ORMModule{
		Name:    "Ent",
		Package: "github.com/ent/ent",
		ID:      "3",
	}
	MGM = ORMModule{
		Name:    "MGM",
		Package: "go.mongodb.org/mongo-driver/mongo",
		ID:      "4",
	}
)

var ORMs = []ORMModule{
	NO_ORM,
	GORM,
	SQLBOILER,
	ENT,
}

var MongoORMs = []ORMModule{
	NO_ORM,
	MGM,
}

func GetORMModuleById(id string) ORMModule {
	switch id {
	case "0":
		return NO_ORM
	case "1":
		return GORM
	case "2":
		return SQLBOILER
	case "3":
		return ENT
	case "4":
		return MGM
	default:
		return ORMModule{}
	}
}
