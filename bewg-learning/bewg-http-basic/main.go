package main

import (
	"log"
	"net/http"
)

//// basic and primitive approach ⤵️
//// this approach show basic of ListenAndServe method.
//// ListenAndServe required two parameters; address and interface named ServeHTTP
//// passed the struct to the ListenAndServe to apply the interface
//type server struct {
//	addr string
//}
//
//func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	// manual routing
//	if r.Method == http.MethodGet {
//		switch r.URL.Path {
//		case "/":
//			_, err := w.Write([]byte("Index page\n"))
//			if err != nil {
//				fmt.Println(err)
//				return
//			}
//		case "/users":
//			_, err := w.Write([]byte("Users page\n"))
//			if err != nil {
//				fmt.Println(err)
//				return
//			}
//		default:
//			_, err := w.Write([]byte("404 page not found!\n"))
//			if err != nil {
//				fmt.Println(err)
//				return
//			}
//		}
//	}
//}
//
//func main() {
//	srv := &server{
//		addr: ":8080",
//	}
//
//	log.Fatal(http.ListenAndServe(srv.addr, srv))
//}

// one of best practice approach.
// flexible configuration due we can define the routing with more readable approach.
func main() {
	apiData := &api{
		addr: ":8080",
	}

	mux := http.NewServeMux()

	// &http.Server is used to make more flexible on defining a http configuration
	// we can define ReadTimeout of the server using this configuration.
	srv := &http.Server{
		Addr:    apiData.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", apiData.getUserHandler)
	mux.HandleFunc("POST /users", apiData.createUserHandler)

	log.Fatal(srv.ListenAndServe())
}
