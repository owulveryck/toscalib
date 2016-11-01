package toscalib

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

// Resolver defines a function spec that the Parser will use to resolve
// remote Imports.
type Resolver func(string) ([]byte, error)

// DefaultResolver provides a basic implementation for retrieving imports that reference
// remote locations. The file will be downloaded over HTTP(s) and the contents are returned.
func defaultResolver(location string) ([]byte, error) {
	var r []byte

	u, err := url.Parse(location)
	if err != nil {
		return r, err
	}

	switch u.Scheme {
	case "http", "https":
		var res *http.Response
		res, err = http.Get(u.String())
		if err != nil {
			return r, err
		}
		defer res.Body.Close()
		r, err = ioutil.ReadAll(res.Body)
		if err != nil {
			return r, err
		}
		return r, nil
	default:
		r, err = ioutil.ReadFile(location)
		if err != nil {
			return r, err
		}
		return r, nil
	}

}
