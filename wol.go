package plugin_wol

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	"wol"
)

// Config the plugin configuration.
type Config struct {
	macAddress	string   `json:"macAddress" yaml:"macAddress"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		macAddress: "",
	}
}

// wol a wol plugin.
type wol struct {
	next     	http.Handler
	name     	string
	macAddress  string
}

// New created a new wol plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	if len(config.macAddress) == 0 {
		return nil, fmt.Errorf("macAddress cannot be empty")
	}

	return &wol{
		name:     	name,
		next:     	next,
		macAddress:	config.Headers,
		}, nil
}

func (a *wol) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	fmt.Printf("%+v\n", rw)
	fmt.Printf("%+v\n", req)

	err := wol.send(a.macAddress)
	if err != nil {
		return
	}

	a.next.ServeHTTP(rw, req)
}
