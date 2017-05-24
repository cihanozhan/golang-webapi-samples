package main
 
import (
    "encoding/json"
    "fmt"
    "net/http"
    "os"
)
 
type API struct {
    Message string "json:message"
}
 
type User struct {
    ID        int    "json:id"
    FirstName string "json:firstname"
    LastName  string "json:lastname"
    Age       int    "json:age"
}
 
func main() {
 
    apiRoot := "/api"
 
    // .../api
    http.HandleFunc(apiRoot, func(w http.ResponseWriter, r *http.Request) {
        message := API{"API Home"}
        output, err := json.Marshal(message)
        checkError(err)
        // w.Header().Set("Content-Type", "application/json")   // Zorunlu değil
        // w.Write([]byte(output))              // Zorunlu değil
        fmt.Fprintf(w, string(output))
    })
 
    // .../api/me
    http.HandleFunc(apiRoot+"/me", func(w http.ResponseWriter, r *http.Request) {
        user := User{3, "Cihan", "Özhan", 28}
        message := user
        output, err := json.Marshal(message)
        checkError(err)
        // w.Header().Set("Content-Type", "application/json")   // Zorunlu değil
        // w.Write([]byte(output))              // Zorunlu değil
        fmt.Fprintf(w, string(output))
    })
 
    // .../api/users
    http.HandleFunc(apiRoot+"/users", func(w http.ResponseWriter, r *http.Request) {
        users := []User{
            User{ID: 1, FirstName: "Murtaza", LastName: "Kılkuyruk", Age: 77},
            User{ID: 2, FirstName: "Hakkı", LastName: "Bulut", Age: 67},
            User{ID: 3, FirstName: "Müslüm", LastName: "Gürses", Age: 80},
        }
        message := users
        output, err := json.Marshal(message)
        checkError(err)
        // w.Header().Set("Content-Type", "application/json")   // Zorunlu değil
        // w.Write([]byte(output))              // Zorunlu değil
        fmt.Fprintf(w, string(output))
    })
 
    // Hizmeti 8080 portu üzerinden yayınla.
    http.ListenAndServe(":8080", nil)
}
 
func checkError(err error) {
    if err != nil {
        fmt.Println("Fatal Error : ", err.Error())
        os.Exit(1)
    }
}
