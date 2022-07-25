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

const busy = true

func (h *Handler) RegisterCustomer(w http.ResponseWriter, r *http.Request) {
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
		log.Println("validate customer error: " + err.Error())
		http.Error(w, "validate customer error", 400)
		return
	}

	err = delivery.Publish(tools.MakeJsonString(m), cfg.Kafka.TopicCustomer)
	if err != nil {
		http.Error(w, "service unavailable", 500)
		return
	}

	err = delivery.Consumer(cfg.Kafka.TopicAuth, m.UUID)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "service unavailable", 500)
		return
	}
	sess := models.Session{Cookie: m.UUID}
	w.Write([]byte(tools.MakeJsonString(sess)))
}

func (h *Handler) ApplyForWork(w http.ResponseWriter, r *http.Request) {
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

	if ut.UserType != models.Cust {
		log.Println("forbidden for company")
		http.Error(w, "only customer can apply for work", http.StatusForbidden)
		return
	}

	ap := &models.ApplyForWork{}

	err = json.NewDecoder(r.Body).Decode(&ap)
	if err != nil {
		log.Println("decode error: " + err.Error())
		http.Error(w, err.Error(), 400)
		return
	}

	ap.CustomerID = ut.ID
	status, err := h.srv.GetWorkStatus(ap.WorkID)
	if err != nil {
		log.Println("did not find a work: " + err.Error())

		err = h.srv.ApplyForWork(ap)
		if err != nil {
			log.Println("apply for work error: " + err.Error())
			http.Error(w, err.Error(), 400)
			return
		}
		log.Println("successfully applied for work")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("OKEY"))
		return
	}

	if status == busy {
		log.Println("can not apply for work")
		http.Error(w, "can not apply for work", 400)
		return
	}

	err = h.srv.ApplyForWork(ap)
	if err != nil {
		log.Println("apply for work error: " + err.Error())
		http.Error(w, err.Error(), 400)
		return
	}

	log.Println("successfully applied for work")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("OKEY"))
}
