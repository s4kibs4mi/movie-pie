package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
)

func capture(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		requestDump, err := httputil.DumpRequest(r, true)
		if err != nil {
			fmt.Println("Error in capturing request :", err)
		} else {
			fmt.Println("Captured request")
			fmt.Println(string(requestDump))
		}

		rec := httptest.NewRecorder()

		next.ServeHTTP(rec, r)

		data := rec.Body.Bytes()

		responseDump, err := httputil.DumpResponse(rec.Result(), true)
		if err != nil {
			fmt.Println("Error in capturing response :", err)
		} else {
			fmt.Println("Captured response")
			fmt.Println(string(responseDump))
		}

		for k, v := range rec.Header() {
			w.Header()[k] = v
		}

		w.WriteHeader(rec.Result().StatusCode)
		w.Write(data)
	}
	return http.HandlerFunc(fn)
}
