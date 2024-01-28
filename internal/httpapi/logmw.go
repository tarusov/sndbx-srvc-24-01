package httpapi

import (
	"log"
	"net/http"
	"net/http/httputil"
)

type (
	logMW struct {
		next http.Handler
		log  *log.Logger
	}

	logWriter struct {
		w   http.ResponseWriter
		hdr int
		buf []byte
	}
)

// CTOR
func NewLogMW(next http.Handler) *logMW {
	return &logMW{
		next: next,
		log:  log.Default(),
	}
}

// newLogWriter
func newLogWriter(w http.ResponseWriter) *logWriter {
	return &logWriter{
		w:   w,
		hdr: http.StatusOK,
		buf: make([]byte, 0, 1024),
	}
}

// ServeHTTP
func (mw *logMW) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	reqDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		mw.log.Printf("ERROR: dump request: %v\n", err)
	} else {
		mw.log.Printf("REQUEST:\n%s\n", string(reqDump))
	}

	lw := newLogWriter(w)
	mw.next.ServeHTTP(lw, r)

	mw.log.Printf("RESPONSE:\nSTATUS:%d\nBODY:%s\n", lw.hdr, string(lw.buf))
}

// Header
func (lw *logWriter) Header() http.Header {
	return lw.w.Header()
}

// Write
func (lw *logWriter) Write(in []byte) (int, error) {
	lw.buf = append(lw.buf, in...)
	return lw.w.Write(in)
}

// WriteHeader
func (lw *logWriter) WriteHeader(statusCode int) {
	lw.hdr = statusCode
	lw.w.WriteHeader(statusCode)
}
