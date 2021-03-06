package zhttp

// TODO: don't depend on chi.
/*
func TestRealIP(t *testing.T) {
	tests := []struct {
		name       string
		remoteAddr string
		header     http.Header
		want       string
	}{
		{
			"remote-addr",
			"1.1.1.1:42",
			nil,
			"1.1.1.1",
		},
		{
			"x-real-ip",
			"1.1.1.1",
			http.Header{"X-Real-Ip": {"100.100.100.100"}},
			"100.100.100.100",
		},
		{
			"x-real-ip-6",
			"2001:beef::0",
			http.Header{"X-Real-Ip": {"2001:dead::0"}},
			"2001:dead::0",
		},
		{
			"x-forwarded-for",
			"1.1.1.1",
			http.Header{"X-Forwarded-For": {"100.100.100.100"}},
			"100.100.100.100",
		},
		{
			"x-forwarded-for-6",
			"2001:beef::0",
			http.Header{"X-Forwarded-For": {"2001:dead::0"}},
			"2001:dead::0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/", nil)
			req.RemoteAddr = tt.remoteAddr
			req.Header = tt.header

			w := httptest.NewRecorder()
			r := chi.NewRouter()
			r.Use(RealIP)

			realIP := ""
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				realIP = r.RemoteAddr
				w.Write([]byte("Hello World"))
			})
			r.ServeHTTP(w, req)

			if w.Code != 200 {
				t.Fatalf("wrong response code: %d; wanted 200", w.Code)
			}

			if realIP != tt.want {
				t.Fatalf("wrong IP: %q; wanted %q", realIP, tt.want)
			}
		})
	}
}
*/
