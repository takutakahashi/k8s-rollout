package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/takutakahashi/k8s-rollout/pkg/k8s"
	"golang.org/x/exp/slices"
)

func main() {
	e := echo.New()
	e.PUT("/:resource/:namespace/:name", func(c echo.Context) error {
		resource := c.Param("resource")
		if slices.Contains([]string{"deployment", "deploy", "statefulset", "daemonset"}, resource) {
			return c.String(400, "bad input")
		}
		namespace := c.Param("namespace")
		name := c.Param("name")
		out, err := k8s.Rollout(resource, namespace, name)
		if err != nil {
			return c.String(400, err.Error())
		}
		return c.String(http.StatusOK, out)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
