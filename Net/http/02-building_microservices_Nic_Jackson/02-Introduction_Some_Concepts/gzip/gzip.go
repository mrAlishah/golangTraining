package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

var name string

/*************************************************************************************************/
/* Handler: helloWorldHandler = Reply message response
/*************************************************************************************************/
type helloWorldResponse struct {
	Message string `json:"message"`
}

func helloWorldHandler(rw http.ResponseWriter, r *http.Request) {
	response := helloWorldResponse{Message: "Hello " + name}

	encoder := json.NewEncoder(rw)
	encoder.Encode(response)
}

/*************************************************************************************************/
/* Main: Start program
/*************************************************************************************************/

func main() {
	port := 8080

	http.Handle("/helloworld",
		NewGzipHandler(http.HandlerFunc(helloWorldHandler)),
	)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

/*************************************************************************************************/
/* Handler: GZipHandler = checking the client has sent the Accept-Encoding header
/*************************************************************************************************/
//NewGzipHandler our handler checks to see if the client has sent the Accept-Encoding header and if so we will
//write the response using the GzipResponseWriter method

type GZipHandler struct {
	next http.Handler
}

// checking the client has sent the Accept-Encoding header
func NewGzipHandler(next http.Handler) http.Handler {
	return &GZipHandler{next}
}

func (h *GZipHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	encodings := r.Header.Get("Accept-Encoding")

	//checking to see if the client has sent the Accept-Encoding header
	if strings.Contains(encodings, "gzip") {
		//we will write the response using the GzipResponseWriter method
		h.serveGzipped(w, r)
	} else if strings.Contains(encodings, "deflate") {
		//TO-do implement
		panic("Deflate not implemented")
	} else {
		//if the client has requested uncompressed content then we only call
		//ServeHttp with the standard ResponseWriter:
		h.servePlain(w, r)
	}
}

// we write the response using the GzipResponseWriter method
func (h *GZipHandler) serveGzipped(w http.ResponseWriter, r *http.Request) {
	//create gzip Writer
	gzw := gzip.NewWriter(w)
	defer gzw.Close()

	w.Header().Set("Content-Encoding", "gzip")
	h.next.ServeHTTP(GzipResponseWriter{gzw, w}, r)
}

// if the client has requested uncompressed content then we only call ServeHttp with the standard ResponseWriter:
func (h *GZipHandler) servePlain(w http.ResponseWriter, r *http.Request) {
	h.next.ServeHTTP(w, r)
}

/*************************************************************************************************/
/* http.ResponseWriter: GzipResponseWriter =  writing a response in a gzipped format
/*************************************************************************************************/
//creating our own(gzip) ResponseWriter that embeds http.ResponseWriter.

type GzipResponseWriter struct {
	gw *gzip.Writer
	http.ResponseWriter
}

func (w GzipResponseWriter) Write(b []byte) (int, error) {
	//type Header map[string][]string
	if _, ok := w.Header()["Content-Type"]; !ok {
		// If content type is not set, infer it from the uncompressed body.
		w.Header().Set("Content-Type", http.DetectContentType(b))
	}
	return w.gw.Write(b)
}

func (w GzipResponseWriter) Flush() {
	w.gw.Flush()
	if fw, ok := w.ResponseWriter.(http.Flusher); ok {
		fw.Flush()
	}
}
