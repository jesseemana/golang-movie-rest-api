package main

import (
	"fmt"
	"net/http"

	connectdb "github.com/jesseemana/gomongoapi/connect"
	"github.com/jesseemana/gomongoapi/router"
)

var PORT = 3030

func main() {
	connectdb.ConnectDB()

	r := router.Router()

	fmt.Println("Server starting...")
	http.ListenAndServe(":3030", r)
	fmt.Println("server listening on port:", 3030)
}	
