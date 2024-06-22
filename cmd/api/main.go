package main

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/go-chi/chi" // Flexible and easy package for web development
// 	"github.com/avukadin/goapi/internal/handlers"
// 	log "github.com/sirupsen/logrus" // Logger alased as log
// )
import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"github.com/vipin-jose/goapi/internal/handlers"
)

func main() {
	// Enable reporting of file and line number
	log.SetReportCaller(true)
	fmt.Println("Starting API server...")

	var r *chi.Mux = chi.NewRouter()                // Create a new router
	handlers.Handler(r)                             // Call the handler function
	fmt.Println("Starting Go API Server")           // Print a message
	err := http.ListenAndServe("localhost:8000", r) // Start the server
	if err != nil {
		log.Fatal(err)
	}
}
