package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/placer14/gof-server/internal/database"
	"gorm.io/gorm"
)

func main() {
	db := database.GetDB()
	http.HandleFunc("/ping", service.pingHandler)
	http.HandleFunc("/get_string_value", getStringValue(db))
	http.HandleFunc("/set_flag_value", setFlagValue(db))

	fmt.Println("Server is running on http://localhost:23456")
	// defer db.Close()
	log.Fatal(http.ListenAndServe(":23456", nil))
}

func handlePing(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("pong"))
}

func getStringValue(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		params, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		flag_key := params.Get("flag_key")
		var existing_flag database.FlagKey

		result := db.First(&existing_flag, "key = ?", flag_key)
		if result.RowsAffected == 0 {
			http.Error(w, fmt.Sprintf("flag `%s` not found", flag_key), http.StatusNotFound)
			return
		}
		var current_variation database.FlagKeyStringVariations
		variation_result := db.First(&current_variation, "uuid = ?", existing_flag.ActiveVariation)
		if variation_result.RowsAffected == 0 {
			http.Error(w, fmt.Sprintf("flag `%s` activate variation `%s` does not exist", flag_key, existing_flag.ActiveVariation), http.StatusNotFound)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(current_variation.Value))
	}
}

func setFlagValue(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// if r.Method != http.MethodPost {
		// 	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		// 	return
		// }
		params, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		flag_key := params.Get("flag_key")
		active_variation_uuid := params.Get("active_variation_uuid")

		var existing_flag database.FlagKey

		result := db.First(&existing_flag, "key = ?", flag_key)
		var message string
		now := time.Now()

		if result.RowsAffected == 0 {
			flag_key_uuid := uuid.New()
			variation_one_uuid := uuid.New()
			variation_two_uuid := uuid.New()
			db.Create(database.FlagKeyStringVariations{UUID: variation_one_uuid.String(), FlagKeyUUID: flag_key_uuid.String(), Value: "false", LastUpdated: &now})
			db.Create(database.FlagKeyStringVariations{UUID: variation_two_uuid.String(), FlagKeyUUID: flag_key_uuid.String(), Value: "true", LastUpdated: &now})
			db.Create(database.FlagKey{UUID: flag_key_uuid.String(), Key: flag_key, ActiveVariation: variation_one_uuid.String(), LastUpdated: &now})
			message = fmt.Sprintf("wrote record for flag key %s", flag_key)
		} else {
			var existing_variation database.FlagKeyStringVariations
			result = db.First(&existing_variation, "uuid = ?", active_variation_uuid)
			if result.RowsAffected == 0 {
				http.Error(w, fmt.Sprintf("no variation with uuid %s exists for flag %s", active_variation_uuid, flag_key), http.StatusBadRequest)
				return
			}
			existing_flag.ActiveVariation = active_variation_uuid
			existing_flag.LastUpdated = &now
			db.Save(&existing_flag)
			message = fmt.Sprintf("flag key `%s` updated to use variation %s", flag_key, active_variation_uuid)
		}
		_, _ = w.Write([]byte(message))
	}
}
