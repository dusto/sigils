package handler

import (
	"sync"

	"github.com/danielgtaylor/huma/v2"
)

type Endpoint interface {
	Register(api huma.API)
}

var regendpointMU sync.RWMutex
var registeredendpoints = make(map[string]Endpoint)

func RegisterEndpoint(path string, endpoint Endpoint) {
	regendpointMU.Lock()
	defer regendpointMU.Unlock()
	if _, ok := registeredendpoints[path]; ok {
		panic("Endpoints: RegisterEndpoint called twice for path " + path)
	}
	registeredendpoints[path] = endpoint
}

func RegisterApi(api huma.API) {
	for _, v := range registeredendpoints {
		v.Register(api)
	}
}
