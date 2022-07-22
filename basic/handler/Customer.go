package handler

import (
	"encoding/json"
	"github.com/mrbelka12000/netfix/basic/config"
	"github.com/mrbelka12000/netfix/basic/delivery"
	"github.com/mrbelka12000/netfix/basic/models"
	"github.com/mrbelka12000/netfix/basic/tools"
	"net/http"
)

func (h *Handler) RegisterCustomer(w http.ResponseWriter, r *http.Request) {
	cfg := config.GetConf()
	m := &models.General{}

	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	m.UUID = tools.GetRandomString()
	jsonB, _ := json.Marshal(m)

	err = delivery.Publish(string(jsonB), cfg.Kafka.TopicCustomer)
	if err != nil {
		http.Error(w, "service unavailable", 500)
		return
	}

	err = delivery.Consumer(cfg.Kafka.TopicAuth, m.UUID)
	if err != nil {
		http.Error(w, "service unavailable", 500)
		return
	}

	w.Write([]byte("OKEY"))
}
