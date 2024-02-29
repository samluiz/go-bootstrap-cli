package drivers

type DriverExtensionModule struct {
	Name    string
	Package string
	ID      string
}

var (
	NO_DRIVER_EXT = DriverExtensionModule{
		Name:    "None",
		Package: "",
		ID:      "0",
	}
	SQLX = DriverExtensionModule{
		Name:    "SQLX",
		Package: "github.com/jmoiron/sqlx",
		ID:      "1",
	}
)

var DriverExtensions = []DriverExtensionModule{
	NO_DRIVER_EXT,
	SQLX,
}

func GetDriverExtensionModuleById(id string) DriverExtensionModule {
	switch id {
	case "0":
		return NO_DRIVER_EXT
	case "1":
		return SQLX
	default:
		return DriverExtensionModule{}
	}
}
