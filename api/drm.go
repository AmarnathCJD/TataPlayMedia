package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func DRMHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if r.Method == "POST" {
		id = r.FormValue("id")
	}

	if id == "" {
		http.Error(w, `{"error":"ID is empty"}`, http.StatusBadRequest)
		return
	}

	channel, err := MainTV.GetChannel(id)

	fmt.Println(channel)
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}

	drm, err := channel.GetDRMWithoutEnc()
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	enc.SetEscapeHTML(false)

	if err := enc.Encode(drm); err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}
}
