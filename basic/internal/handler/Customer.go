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

// RegisterCustomer example
// @Summary Register new customer
// @Description registration
// @Tags auth
// @ID auth_customer
// @Accept  json
// @Produce  json
// @Param input body customerReg true "registration"
// @Success 200 {object} session
// @Failure 400,404,405,500
// @Router /register/customer [post]
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

	gen, err := delivery.Consumer(cfg.Kafka.TopicAuth, m.UUID)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "service unavailable", 500)
		return
	}

	wallet := &models.Wallet{OwnerID: gen.ID, UUID: m.UUID}

	err = delivery.Publish(tools.MakeJsonString(wallet), cfg.Kafka.TopicWallets)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	_, err = delivery.Consumer(cfg.Kafka.TopicCreateWallet, m.UUID)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	sess := models.Session{ID: gen.ID, Cookie: m.UUID}
	w.Write([]byte(tools.MakeJsonString(sess)))
}

// ApplyForWork example
// @Summary apply for service
// @ID apply for work
// @Tags service
// @Accept  json
// @Produce  json
// @Param input body workAction true "work"
// @Param session header string true "session"
// @Security ApiKeyAuth
// @Success 200 {string} string	"OKEY"
// @Failure 400,404,405,500
// @Router /service/apply [post]
func (h *Handler) ApplyForWork(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	c := r.Header.Get("session")

	ut, err := redis.GetUserType(c)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if ut.UserType != models.Cust {
		log.Println("forbidden for company")
		http.Error(w, "only customer can apply for work", http.StatusForbidden)
		return
	}

	wa := &models.WorkActions{}

	err = json.NewDecoder(r.Body).Decode(&wa)
	if err != nil {
		log.Println("decode error: " + err.Error())
		http.Error(w, err.Error(), 400)
		return
	}

	wa.CustomerID = ut.ID
	status, err := h.srv.GetWorkStatus(wa.WorkID)
	if err != nil {
		log.Println("did not find a work: " + err.Error())

		err = h.srv.ApplyForWork(wa)
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

	err = h.srv.ApplyForWork(wa)
	if err != nil {
		log.Println("apply for work error: " + err.Error())
		http.Error(w, err.Error(), 400)
		return
	}

	log.Println("successfully applied for work")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("OKEY"))
}

// FinishWork example
// @Summary finish work
// @ID finish work
// @Tags service
// @Accept  json
// @Produce  json
// @Param input body workAction true "finished"
// @Param session header string true "session"
// @Security ApiKeyAuth
// @Success 200 {string} string	"OKEY"
// @Failure 400,404,405,500
// @Router /service/finish [post]
func (h *Handler) FinishWork(w http.ResponseWriter, r *http.Request) {

	c := r.Header.Get("session")

	ut, err := redis.GetUserType(c)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if ut.UserType != models.Cust {
		log.Println("forbidden for company")
		http.Error(w, "only customer can apply for work", http.StatusForbidden)
		return
	}

	wa := &models.WorkActions{}
	err = json.NewDecoder(r.Body).Decode(&wa)
	if err != nil {
		log.Println("decode error: " + err.Error())
		http.Error(w, err.Error(), 400)
		return
	}

	wa.CustomerID = ut.ID

	status, err := h.srv.GetWorkStatus(wa.WorkID)
	if err != nil {
		log.Println("get work status error: " + err.Error())
		http.Error(w, err.Error(), 400)
		return
	}

	if status != busy {
		log.Println("the work cannot end because it is not active")
		http.Error(w, "can not finish work", 400)
		return
	}

	err = h.srv.FinishWork(wa)
	if err != nil {
		log.Println("finish work error: " + err.Error())
		http.Error(w, err.Error(), 400)
		return
	}

	w.Write([]byte("finished"))
}

type customerReg struct {
	*models.SwaggerCustomerRegister
}

type workAction struct {
	*models.SwaggerWorkAction
}
