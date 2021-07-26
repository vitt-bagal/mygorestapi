package main

import (
	"os"

	"github.com/vitt-bagal/mygorestapi/handler"
)

func main() {
	// Set an Environment Variable for suppliers api url
	// We can override this url at runtime
	os.Setenv("FRUIT_SUPPLIER", "https://run.mocky.io/v3/c51441de-5c1a-4dc2-a44e-aab4f619926b")
	os.Setenv("VEG_SUPPLIER", "https://run.mocky.io/v3/4ec58fbc-e9e5-4ace-9ff0-4e893ef9663c")
	os.Setenv("GRAIN_SUPPLIER", "https://run.mocky.io/v3/e6c77e5c-aec9-403f-821b-e14114220148")

	// call to router
	handler.HandleRequests()
}
