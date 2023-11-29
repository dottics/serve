package servemux

import (
	"context"
	"net/http"
	"regexp"
	"slices"
)

// ContextHandlerFunc is a handler function with the addition of the context.
// The context contains all the request URL params:
//
// - see: Params and GetParam
type ContextHandlerFunc func(*context.Context, http.ResponseWriter, *http.Request)

// ContextMiddleware return a stop bool to indicate if the middleware should
// stop and exit. Where true == stop.
type ContextMiddleware func(*context.Context, http.ResponseWriter, *http.Request) (bool, http.HandlerFunc)

type Route struct {
	method       string
	whitelisted  bool
	pattern      string
	patternRegex *regexp.Regexp
	handler      ContextHandlerFunc
	params       []Param
}

type Mux struct {
	http.ServeMux
	middleware map[int]ContextMiddleware
	routes     map[string][]Route
}

// NewMux creates a new multiplexer instance.
//
// Other:
// - should parse the openapi.yaml and keep it in memory, as this will only happen when the server is started.
func NewMux() *Mux {
	m := &Mux{
		middleware: make(map[int]ContextMiddleware),
		routes:     make(map[string][]Route),
	}
	// run any request, as all routes will match the default path of "/"
	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		route := m.findHandler(r)
		if route == nil {
			w.WriteHeader(http.StatusMethodNotAllowed)
			_, _ = w.Write([]byte(`Method not allowed`))
			return // ensure handler function terminates
		}
		con := m.handlerContext(*route, r)
		stop, handler := m.runMiddleware(con, w, r)
		if !stop {
			route.handler(con, w, r)
			return // ensure handler function terminates
		}
		if handler != nil {
			handler(w, r)
			return
		}
		w.WriteHeader(405)
		_, _ = w.Write([]byte(`Method Not Allowed`))
		return // ensure handler function terminates
	})
	return m
}

func (m *Mux) handlerContext(route Route, r *http.Request) *context.Context {
	ctx := context.Background()
	params := ParsePath(r.URL.Path, route.patternRegex, route.params)
	con := context.WithValue(ctx, "params", params)
	con = context.WithValue(ctx, "whitelisted", route.whitelisted)
	return &con
}

func (m *Mux) findHandler(r *http.Request) *Route {
	// if no routes are declared on this method, then end
	if routes := m.routes[r.Method]; len(routes) == 0 {
		return nil
	}
	possibleRoutes := make([]Route, 0)
	for _, route := range m.routes[r.Method] {
		// we want to be able to match two patterns, namely
		// 1. /some/path
		// 2. /some/path/
		if route.patternRegex.MatchString(r.URL.Path) {
			possibleRoutes = append(possibleRoutes, route)
		}
	}

	// call the context handler func
	if len(possibleRoutes) >= 1 {
		// there are multiple routes that match incoming request
		slices.SortFunc(possibleRoutes, func(a, b Route) int {
			return len(a.params) - len(b.params)
		})

		route := possibleRoutes[0]
		return &route
	}

	return nil
}

func (m *Mux) Use(middleware ContextMiddleware) {
	m.middleware[len(m.middleware)+1] = middleware
}

func (m *Mux) runMiddleware(c *context.Context, w http.ResponseWriter, r *http.Request) (bool, http.HandlerFunc) {
	for _, mid := range m.middleware {
		stop, handler := mid(c, w, r)
		if stop {
			return true, handler
		}
	}
	return false, nil
}

// addRoute adds a new route to the list of routes for the multiplexer.
func (m *Mux) addRoute(route Route) {
	// parse the pattern to extract all params, params are defined using the
	// pattern:
	// {variableName}
	patternRegex, params := CompilePathRegex(route.pattern)
	route.patternRegex = regexp.MustCompile(patternRegex)
	route.params = params

	xr := append(m.routes[route.method], route)
	m.routes[route.method] = xr
}

// newRoute does the basic scaffolding of creating a new Route instance.
func newRoute(method, pattern string, handler ContextHandlerFunc) Route {
	r := Route{
		method:       method,
		pattern:      pattern,
		patternRegex: nil,
		handler:      handler,
	}
	return r
}

// Get adds a new GET route handler to the multiplexer.
func (m *Mux) Get(pattern string, handler ContextHandlerFunc) {
	r := newRoute(http.MethodGet, pattern, handler)
	m.addRoute(r)
}

// GetWhitelisted adds a new GET route handler to the multiplexer. The route is
// whitelisted. Thus, should not be authenticated.
func (m *Mux) GetWhitelisted(pattern string, handler ContextHandlerFunc) {
	r := newRoute(http.MethodGet, pattern, handler)
	r.whitelisted = true
	m.addRoute(r)
}

// Post adds a new POST route handler to the multiplexer.
func (m *Mux) Post(pattern string, handler ContextHandlerFunc) {
	r := newRoute(http.MethodPost, pattern, handler)
	m.addRoute(r)
}

// PostWhitelisted adds a new POST route handler to the multiplexer.
func (m *Mux) PostWhitelisted(pattern string, handler ContextHandlerFunc) {
	r := newRoute(http.MethodPost, pattern, handler)
	r.whitelisted = true
	m.addRoute(r)
}

// Put adds a new PUT route handler to the multiplexer.
func (m *Mux) Put(pattern string, handler ContextHandlerFunc) {
	r := newRoute(http.MethodPut, pattern, handler)
	m.addRoute(r)
}

// Patch adds a new PATCH route handler to the multiplexer.
func (m *Mux) Patch(pattern string, handler ContextHandlerFunc) {
	r := newRoute(http.MethodPatch, pattern, handler)
	m.addRoute(r)
}

// Delete adds a new DELETE route handler to the multiplexer.
func (m *Mux) Delete(pattern string, handler ContextHandlerFunc) {
	r := newRoute(http.MethodDelete, pattern, handler)
	m.addRoute(r)
}

// Head adds a new HEAD route handler to the multiplexer.
func (m *Mux) Head(pattern string, handler ContextHandlerFunc) {
	r := newRoute(http.MethodHead, pattern, handler)
	m.addRoute(r)
}

// Options adds a new OPTIONS route handler to the multiplexer.
func (m *Mux) Options(pattern string, handler ContextHandlerFunc) {
	r := newRoute(http.MethodOptions, pattern, handler)
	m.addRoute(r)
}
