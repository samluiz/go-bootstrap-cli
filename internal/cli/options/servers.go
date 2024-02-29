package options

type WebServerModule struct {
	Name    string
	Package string
	ID      string
}

var (
	STDLIB = WebServerModule{
		Name:    "Standard Library",
		Package: "",
		ID:      "0",
	}
	FIBER = WebServerModule{
		Name:    "Fiber",
		Package: "github.com/gofiber/fiber/v2",
		ID:      "1",
	}
	GIN = WebServerModule{
		Name:    "Gin",
		Package: "github.com/gin-gonic/gin",
		ID:      "2",
	}
	CHI = WebServerModule{
		Name:    "Chi",
		Package: "github.com/go-chi/chi",
		ID:      "3",
	}
	ECHO = WebServerModule{
		Name:    "Echo",
		Package: "github.com/labstack/echo/v4",
		ID:      "4",
	}
	GORILLA = WebServerModule{
		Name:    "Gorilla",
		Package: "github.com/gorilla/mux",
		ID:      "5",
	}
	IRIS = WebServerModule{
		Name:    "Iris",
		Package: "github.com/kataras/iris/v12",
		ID:      "6",
	}
	MUX = WebServerModule{
		Name:    "Mux",
		Package: "github.com/gorilla/mux",
		ID:      "7",
	}
	AERO = WebServerModule{
		Name:    "Aero",
		Package: "github.com/aerogo/aero",
		ID:      "8",
	}
	FASTHTTP = WebServerModule{
		Name:    "FastHTTP",
		Package: "github.com/valyala/fasthttp",
		ID:      "9",
	}
	BEEGO = WebServerModule{
		Name:    "Beego",
		Package: "github.com/astaxie/beego",
		ID:      "10",
	}
)

func GetServerById(id string) WebServerModule {
	switch id {
	case "0":
		return STDLIB
	case "1":
		return FIBER
	case "2":
		return GIN
	case "3":
		return CHI
	case "4":
		return ECHO
	case "5":
		return GORILLA
	case "6":
		return IRIS
	case "7":
		return MUX
	case "8":
		return AERO
	case "9":
		return FASTHTTP
	case "10":
		return BEEGO
	default:
		return WebServerModule{}
	}
}

func GetServerNameById(id string) string {
	return GetServerById(id).Name
}

func (m *WebServerModule) GetServerPackageByServerId(id string) string {
	return GetServerById(id).Package
}
