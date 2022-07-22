package handler

import (
	"encoding/json"
	"github.com/mrbelka12000/netfix/basic/models"
	"github.com/mrbelka12000/netfix/basic/redis"
	"net/http"
)

func (h *Handler) CreateService(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	c := r.Header.Get("session")
	ut := &models.Role{}
	jsonB, err := redis.GetValue(c)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = json.Unmarshal([]byte(jsonB), &ut)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if ut.UserType != models.Cmp {
		http.Error(w, "only companies can create service", http.StatusForbidden)
		return
	}

	w.Write([]byte(jsonB))
}
