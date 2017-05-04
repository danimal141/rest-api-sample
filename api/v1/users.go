package v1

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/danimal141/rest-api-sample/models"
	"github.com/labstack/echo"
)

type paginationParams struct {
	Pagination string `query:"pagination"`
}

/*

### Query parameter

key        |value  |description
----------:|------:|----------------------------
pagination |false  |ページネーション機能は未実装なのでfalseが必須

*/
func UsersIndex(c echo.Context) error {
	if err := validatePaginationParams(c); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, models.SampleUsers())
}

func UsersShow(c echo.Context) error {
	users := models.SampleUsers()
	id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return err
	}
	if id > len(users)-1 {
		err := fmt.Errorf("user_id=%d is not found", id)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, users[id])
}

func validatePaginationParams(c echo.Context) error {
	p := new(paginationParams)
	if err := c.Bind(p); err != nil {
		return err
	}
	if p.Pagination != "false" {
		err := errors.New("pagination must be false, because pagination is not supported yet")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
