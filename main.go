package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/hupe1980/go-huggingface"
)

type SummarizationRequest struct {
	Text string `json:"text"`
}

type SummarizationResponse struct {
	Summary string `json:"summary"`
}

func summarizeHandler(w http.ResponseWriter, r *http.Request) {
	var req SummarizationRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ic := huggingface.NewInferenceClient(os.Getenv("HUGGINGFACE_TOKEN"))

	summary, err := ic.Summarization(context.Background(), &huggingface.SummarizationRequest{
		Inputs: []string{req.Text},
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := SummarizationResponse{
		Summary: summary[0].SummaryText,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/summarize", summarizeHandler)
	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
