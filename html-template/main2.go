package main
import ( "html/template" 
        "net/http" 
        "log" )

func home(w http.ResponseWriter, r *http.Request) {
tmpl, err := template.ParseFiles("html-template/page2.html")

if err != nil {
    log.Fatal(err)
}

tmpl.Execute(w, "Mr.Amir")
}
func main() {
http.HandleFunc("/", home)

log.Println("Server started on :8082")

log.Fatal(http.ListenAndServe(":8082", nil))
}