package httpf

import (
	"net/http"
) // .import

// Server ..
type Server struct {
	Address string
	Router  *http.ServeMux
} // .Httpf

// Start ..
func (hf *Server) Start() error {

	return http.ListenAndServe(hf.Address, hf.Router)

} // .Start
