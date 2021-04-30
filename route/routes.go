package route

import (
	"io/ioutil"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mwkosasih/soccer-gateway/util"
	"gopkg.in/go-playground/validator.v9"
	// "go.elastic.co/apm/module/apmechov4"
)

// Route for mapping from json file
type Route struct {
	Path       string   `json:"path"`
	Method     string   `json:"method"`
	Module     string   `json:"module"`
	Tag        string   `json:"tag"`
	Endpoint   string   `json:"endpoint_filter"`
	Middleware []string `json:"middleware"`
}

// Init gateway router
func Init() *echo.Echo {
	routes := loadRoutes("./route/gate/")

	e := echo.New()

	validator := validator.New()
	// Set Bundle MiddleWare
	e.Use(middleware.RequestID())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:  []string{"*"},
		AllowHeaders:  []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, echo.HeaderContentLength, echo.HeaderAcceptEncoding, echo.HeaderAccessControlAllowOrigin, echo.HeaderAccessControlAllowHeaders, echo.HeaderContentDisposition, "X-Request-Id", "device-id", "X-Summary", "X-Account-Number", "X-Business-Name", "client-secret", "X-CSRF-Token"},
		ExposeHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, echo.HeaderContentLength, echo.HeaderAcceptEncoding, echo.HeaderAccessControlAllowOrigin, echo.HeaderAccessControlAllowHeaders, echo.HeaderContentDisposition, "X-Request-Id", "device-id", "X-Summary", "X-Account-Number", "X-Business-Name", "client-secret", "X-CSRF-Token"},
		AllowMethods:  []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	for _, route := range routes {
		e.Add(route.Method, route.Path, endpoint[route.Endpoint].Handle, chainMiddleware(route)...)
	}

	e.Validator = &CustomValidator{Validator: validator}
	return e
}

func loadRoutes(filePath string) []Route {
	var routes []Route
	files, err := ioutil.ReadDir(filePath)
	if err != nil {
		log.Fatalf("Failed to load file: %v", err)
	}
	for _, file := range files {
		byteFile, err := ioutil.ReadFile(filePath + "/" + file.Name())
		if err != nil {
			log.Fatalf("Failed to load file: %v", err)
		}
		var tmp []Route
		if err := util.Json.Unmarshal(byteFile, &tmp); err != nil {
			log.Fatalf("Failed to marshal file: %v", err)
		}
		routes = append(routes, tmp...)
	}

	return routes
}

func chainMiddleware(route Route) []echo.MiddlewareFunc {
	var mwHandlers []echo.MiddlewareFunc
	// init mw for router ,attach router properties
	mwHandlers = append(mwHandlers)
	for _, v := range route.Middleware {
		mwHandlers = append(mwHandlers, middlewareHandler[v])
	}
	return mwHandlers
}

type CustomValidator struct {
	Validator *validator.Validate
}

// Validate struct
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
