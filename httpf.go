package httpf

import (
	"net/http"
) // .import

// Httpf  ..
type Httpf struct {
	Address string
	Router  *http.ServeMux
} // .Httpf

// StartServer ..
func (hf *Httpf) StartServer() error {

	return http.ListenAndServe(hf.Address, hf.Router)

} // .StartServer
