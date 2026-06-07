package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync/atomic"
	"time"
)

type projetoKorpResponse struct {
	Nome    string `json:"nome"`
	Horario string `json:"horario"`
}

var (
	projetoKorpRequests atomic.Uint64
	totalRequests       atomic.Uint64
	startedAt           = time.Now().UTC()
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /projeto-korp", handleProjetoKorp)
	mux.HandleFunc("GET /healthz", handleHealth)
	mux.HandleFunc("GET /metrics", handleMetrics)

	port := envOrDefault("PORT", "8080")

	log.Printf("servidor iniciado na porta %s", port)
	if err := http.ListenAndServe(":"+port, logMiddleware(mux)); err != nil {
		log.Fatalf("erro ao iniciar servidor: %v", err)
	}
}

func handleProjetoKorp(w http.ResponseWriter, _ *http.Request) {
	projetoKorpRequests.Add(1)

	resp := projetoKorpResponse{
		Nome:    "Projeto Korp",
		Horario: time.Now().UTC().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(resp)
}

func handleHealth(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write([]byte(`{"status":"ok"}` + "\n"))
}

func handleMetrics(w http.ResponseWriter, _ *http.Request) {
	uptime := time.Since(startedAt).Seconds()

	w.Header().Set("Content-Type", "text/plain; version=0.0.4; charset=utf-8")

	fmt.Fprintf(w, `# HELP http_server_projeto_korp_up Disponibilidade do servico.
# TYPE http_server_projeto_korp_up gauge
http_server_projeto_korp_up 1

# HELP http_server_projeto_korp_requests_total Requisicoes no endpoint /projeto-korp.
# TYPE http_server_projeto_korp_requests_total counter
http_server_projeto_korp_requests_total{method="GET",path="/projeto-korp"} %d

# HELP http_server_projeto_korp_total_requests_total Total de requisicoes na aplicacao.
# TYPE http_server_projeto_korp_total_requests_total counter
http_server_projeto_korp_total_requests_total %d

# HELP http_server_projeto_korp_uptime_seconds Tempo de atividade em segundos.
# TYPE http_server_projeto_korp_uptime_seconds gauge
http_server_projeto_korp_uptime_seconds %.0f
`,
		projetoKorpRequests.Load(),
		totalRequests.Load(),
		uptime,
	)
}

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		totalRequests.Add(1)

		next.ServeHTTP(w, r)

		log.Printf("%s %s %s %dms",
			r.Method,
			r.URL.Path,
			r.RemoteAddr,
			time.Since(start).Milliseconds(),
		)
	})
}

func envOrDefault(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
