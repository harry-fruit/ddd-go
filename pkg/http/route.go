package restapi

import "regexp"

type Route struct {
	Method      string
	Path        string
	RouteParam  *string
	HandlerName string
}

func NewRoute(method string, path string, handlerName string) *Route {

	route := &Route{
		Method:      method,
		Path:        path,
		HandlerName: handlerName,
	}

	routerParam := route.getRouterParam()
	if routerParam != nil {
		route.RouteParam = routerParam
	}

	return route
}

func (r *Route) getRouterParam() *string {
	re := regexp.MustCompile(`\{([^}]+)\}`)
	match := re.FindStringSubmatch(r.Path)

	if len(match) == 0 {
		return nil
	}

	return &match[1]
}
