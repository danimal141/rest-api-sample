package v1_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/adams-sarah/test2doc/test"
	"github.com/danimal141/rest-api-sample/api/router"
	"github.com/labstack/echo"
)

var server *test.Server

func TestMain(m *testing.M) {
	var err error
	r := router.NewRouter()
	test.RegisterURLVarExtractor(makeURLVarExtractor(r))
	server, err = test.NewServer(r)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Start test
	code := m.Run()

	// Flush to an apib doc file
	server.Finish()

	// Terminate
	os.Exit(code)
}

func makeURLVarExtractor(e *echo.Echo) func(req *http.Request) map[string]string {
	return func(req *http.Request) map[string]string {
		ctx := e.AcquireContext()
		defer e.ReleaseContext(ctx)
		pnames := ctx.ParamNames()
		if len(pnames) == 0 {
			return nil
		}

		paramsMap := make(map[string]string, len(pnames))
		for _, name := range pnames {
			paramsMap[name] = ctx.Param(name)
		}
		return paramsMap
	}
}
