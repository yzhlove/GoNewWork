package pressure

import "net/http"

const Number = 10000

func NewHttpServer() {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Ok."))
	})
	http.ListenAndServe(":1234", nil)

}
