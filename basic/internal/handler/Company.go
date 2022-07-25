package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mrbelka12000/netfix/basic/config"
	"github.com/mrbelka12000/netfix/basic/internal/delivery"
	"github.com/mrbelka12000/netfix/basic/models"
	"github.com/mrbelka12000/netfix/basic/redis"
	"github.com/mrbelka12000/netfix/basic/tools"
)

func (h *Handler) RegisterCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cfg := config.GetConf()
	m := &models.General{}

	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	m.UUID = tools.GetRandomString()

	err = m.Validate()
	if err != nil {
		log.Println("validate company error: " + err.Error())
		http.Error(w, err.Error(), 400)
		return
	}

	err = delivery.Publish(tools.MakeJsonString(m), cfg.Kafka.TopicCompany)
	if err != nil {
		http.Error(w, "service unavailable", 500)
		return
	}

	err = delivery.Consumer(cfg.Kafka.TopicAuth, m.UUID)
	if err != nil {
		http.Error(w, "service unavailable", 500)
		return
	}

	sess := models.Session{Cookie: m.UUID}
	w.Write([]byte(tools.MakeJsonString(sess)))
}

func (h *Handler) CreateService(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	c := r.Header.Get("session")
	ut := &models.Role{}
	jsonB, err := redis.GetValue(c)
	if err != nil {
		log.Println("no value in redis: " + err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	err = json.Unmarshal([]byte(jsonB), &ut)
	if err != nil {
		log.Println("unmarshall error: " + err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	if ut.UserType != models.Cmp {
		log.Println("forbidden customer company")
		http.Error(w, "only companies can create service", http.StatusForbidden)
		return
	}

	cw := &models.Work{}
	err = json.NewDecoder(r.Body).Decode(&cw)
	if err != nil {
		log.Println("decode error: " + err.Error())
		http.Error(w, err.Error(), 400)
		return
	}

	err = cw.Validate()
	if err != nil {
		log.Println("validate create work error: " + err.Error())
		http.Error(w, "validate create work error", 400)
		return
	}

	exists := h.srv.IsExists(cw.WorkField)
	if !exists {
		log.Println("unknown work field: " + cw.WorkField)
		http.Error(w, "unknown work field", 400)
		return
	}

	cw.CompanyID = ut.ID

	err = h.srv.CreateWork(cw)
	if err != nil {
		log.Println("create work error: " + err.Error())
		http.Error(w, err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("OKEY"))
}
