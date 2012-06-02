package main

import (
	"flag"
	"log"
	"net/http"
	"text/template"
)

var (
	httpAddr = flag.String("http", ":8080", "http listen address")

	index = template.Must(template.New("index").Parse(indexHtml))
)

type header struct {
	Key, Value string
}

func main() {
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		var headers = make([]header, 0, len(req.Header))
		for k, _ := range req.Header {
			headers = append(headers, header{k, req.Header.Get(k)})
		}
		index.Execute(w, headers)
	})

	log.Fatal(http.ListenAndServe(*httpAddr, nil))
}

const indexHtml = `
<!doctype html>
<html>
<head>
	<meta charset="UTF-8">
	<title>HTTP Headers</title>
</head>
<body>
	<table>
	{{range .}}
		<tr><td><strong>{{.Key}}</strong</td><td>{{.Value}}</td></tr>
	{{end}}
	</table>
</body>
</html>
`
