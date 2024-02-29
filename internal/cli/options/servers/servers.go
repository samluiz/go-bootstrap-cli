package servers

type ServerModule struct {
	Name    string
	Package string
	ID      string
}

var (
	STDLIB = ServerModule{
		Name:    "Standard Library",
		Package: "",
		ID:      "0",
	}
	FIBER = ServerModule{
		Name:    "Fiber",
		Package: "github.com/gofiber/fiber/v2",
		ID:      "1",
	}
	GIN = ServerModule{
		Name:    "Gin",
		Package: "github.com/gin-gonic/gin",
		ID:      "2",
	}
	CHI = ServerModule{
		Name:    "Chi",
		Package: "github.com/go-chi/chi/v5",
		ID:      "3",
	}
	ECHO = ServerModule{
		Name:    "Echo",
		Package: "github.com/labstack/echo/v4",
		ID:      "4",
	}
	GORILLA = ServerModule{
		Name:    "Gorilla",
		Package: "github.com/gorilla/mux",
		ID:      "5",
	}
	IRIS = ServerModule{
		Name:    "Iris",
		Package: "github.com/kataras/iris/v12",
		ID:      "6",
	}
	MUX = ServerModule{
		Name:    "Mux",
		Package: "github.com/gorilla/mux",
		ID:      "7",
	}
	AERO = ServerModule{
		Name:    "Aero",
		Package: "github.com/aerogo/aero",
		ID:      "8",
	}
	FASTHTTP = ServerModule{
		Name:    "FastHTTP",
		Package: "github.com/valyala/fasthttp",
		ID:      "9",
	}
	BEEGO = ServerModule{
		Name:    "Beego",
		Package: "github.com/astaxie/beego",
		ID:      "10",
	}
)

var Servers = []ServerModule{
	STDLIB,
	FIBER,
	GIN,
	CHI,
	ECHO,
	GORILLA,
	IRIS,
	MUX,
	AERO,
	FASTHTTP,
	BEEGO,
}

func GetServerById(id string) ServerModule {
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
		return ServerModule{}
	}
}
