package web

import (
	"encoding/json"
	"github.com/Devashish08/InterPreter-Compiler/evaluator"
	"github.com/Devashish08/InterPreter-Compiler/lexer"
	"github.com/Devashish08/InterPreter-Compiler/object"
	"github.com/Devashish08/InterPreter-Compiler/parser"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

type ExecuteRequest struct {
	Code string `json:"code"`
}

type ExecuteResponse struct {
	Result string `json:"result"`
	Error  string `json:"error,omitempty"`
}

func StartServer(port string) error {
	r := mux.NewRouter()

	// Serve static files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	// Serve main page
	r.HandleFunc("/", handleIndex).Methods("GET")
	
	// Handle code execution
	r.HandleFunc("/execute", handleExecute).Methods("POST")

	// Log startup message
	log.Printf("Starting web server on http://localhost:%s", port)

	return http.ListenAndServe(":"+port, r)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(filepath.Join("web", "templates", "index.html"))
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func handleExecute(w http.ResponseWriter, r *http.Request) {
	var req ExecuteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Error decoding request: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Create new environment for each execution
	env := object.NewEnvironment()

	// Parse and evaluate the code
	l := lexer.New(req.Code)
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) > 0 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(ExecuteResponse{
			Error: "Parser errors:\n" + formatErrors(p.Errors()),
		})
		return
	}

	result := evaluator.Eval(program, env)
	if result != nil {
		if err, ok := result.(*object.Error); ok {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(ExecuteResponse{
				Error: "Runtime error:\n" + err.Message,
			})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(ExecuteResponse{
			Result: result.Inspect(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ExecuteResponse{
		Result: "null",
	})
}

func formatErrors(errors []string) string {
	result := ""
	for _, err := range errors {
		result += err + "\n"
	}
	return result
}
