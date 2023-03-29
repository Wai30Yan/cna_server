package helpers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

func randomTimestamp(start time.Time, end time.Time) time.Time {
	delta := end.Sub(start)
	offset := time.Duration(rand.Int63n(int64(delta)))
	return start.Add(offset)
}

func JsonToMap(jsonStr string) (map[string]bool, error) {
	var m map[string]bool
	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
        return nil, err
    }
    return m, nil
}

func MapToJson(m map[string]bool) (string, error) {
	jsonBytes, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

func PreflightHandler(w http.ResponseWriter, r *http.Request) {
    headers := w.Header()
    headers.Set("Access-Control-Allow-Origin", "*")
    headers.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
    headers.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    headers.Set("Access-Control-Max-Age", "86400")
    headers.Set("Access-Control-Allow-Credentials", "true")
    w.WriteHeader(http.StatusOK)
}