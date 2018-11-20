package main

// - contenttypeheader
// Set the Content-Type header
//
// w.Header().Set("Content-Type", "text/plain; charset=utf-8")
//
//
// - customserver
// Configured Go server
//
// srv := &http.Server{
//    Addr:         GoLabServiceAddr,
//    ReadTimeout:  5 * time.Second,
//    WriteTimeout: 10 * time.Second,
//    IdleTimeout:  120 * time.Second,
//    TLSConfig:    tlsConfig,
//    Handler:      mux,
// }
//
//
// - customtlsconfig
// Custom TLS configuration
//
// // See https://blog.cloudflare.com/exposing-go-on-the-internet/ for details
// // about these settings
// tlsConfig := &tls.Config{
//    // Causes servers to use Go's default cipher suite preferences,
//    // which are tuned to avoid attacks. Does nothing on clients.
//    PreferServerCipherSuites: true,
//    // Only use curves which have assembly implementations
//    CurvePreferences: []tls.CurveID{
//        tls.CurveP256,
//        tls.X25519, // Go 1.8 only
//    },
//
//    MinVersion: tls.VersionTLS12,
//    CipherSuites: []uint16{
//        tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
//        tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
//        tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
//        tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
//        tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
//        tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
//    },
// }
//
//
// - envvars
// Environment Variables
//
// var (
// 	GoLabCertFile    = os.Getenv("GoLab_CERT_FILE")
// 	GoLabKeyFile     = os.Getenv("GoLab_KEY_FILE")
// 	GoLabServiceAddr = os.Getenv("GoLab_SERVICE_ADDR")
// )
//
//
// - sqlxconnect
// sqlx.Connect
//
// db, err := sqlx.Open("postgres", "postgres://postgres:postgres@10.0.75.1:5432/goland?sslmode=disable")
// if err != nil {
//    log.Fatalln(err)
// }
//
// err = db.Ping()
// if err != nil {
//    log.Fatalln(err)
// }
//
//
// - testhomehandler
// Test the Home Handler
//
// func TestHandlers_Handler(t *testing.T) {
// 	tests := []struct {
// 		name           string
// 		in             *http.Request
// 		out            *httptest.ResponseRecorder
// 		expectedStatus int
// 		expectedBody   string
// 	}{
// 		{
// 			name:           "good",
// 			in:             httptest.NewRequest("GET", "/", nil),
// 			out:            httptest.NewRecorder(),
// 			expectedStatus: http.StatusOK,
// 			expectedBody:   message,
// 		},
// 	}
//
// 	for _, test := range tests {
// 		test := test
// 		t.Run(test.name, func(t *testing.T) {
// 			h := New(nil)
// 			h.Handler(test.out, test.in)
// 			if test.out.Code != test.expectedStatus {
// 				t.Logf("expected: %d\ngot: %d\n", test.expectedStatus, test.out.Code)
// 				t.Fail()
// 			}
//
// 			body := test.out.Body.String()
// 			if body != test.expectedBody {
// 				t.Logf("expected: %s\ngot: %s\n", test.expectedBody, body)
// 				t.Fail()
// 			}
// 		})
// 	}
// }
