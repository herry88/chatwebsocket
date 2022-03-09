package main 
import ( 
  "log" 
  "net/http" 
) 
func main() { 
  http.HandleFunc("/", func(w http.ResponseWriter, r  *http.Request) { 
    w.Write([]byte("\n" +
	"<html>\n" +
	"<head>\n" +
	"<title>Hello World</title>\n" +
	"</head>\n" +
	"<body>\n" +
	"<h1>Hello World, Golang Saya Berhasil Menampilkan Hello World</h1>\n" +
	"</body>\n" +
	"</htmL>\n" ))
  }) 
  // start the web server 
  if err := http.ListenAndServe(":8080", nil); err != nil { 
    log.Fatal("ListenAndServe:", err) 
  } 
} 