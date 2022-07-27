package handler

import (
	"github.com/mrbelka12000/netfix/basic/config"
	"github.com/mrbelka12000/netfix/basic/internal/delivery"
	"github.com/mrbelka12000/netfix/basic/models"
	"github.com/mrbelka12000/netfix/basic/redis"
	"github.com/mrbelka12000/netfix/basic/tools"
	"log"
	"net/http"
)

// Profile get  example
// @Summary get profile
// @Description getting profile by cookie
// @ID profile
// @Tags general
// @Accept  json
// @Produce  json
// @Param session header string true "session"
// @Security ApiKeyAuth
// @Success 200 {object} getGeneral "general"
// @Failure 400,404,405,500
// @Router /profile [get]
func (h *Handler) Profile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cfg := config.GetConf()

	c := r.Header.Get("session")

	ut, err := redis.GetUserType(c)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	g := &models.General{
		ID:   ut.ID,
		UUID: tools.GetRandomString(),
	}
	switch ut.UserType {
	case models.Cmp:
		err = delivery.Publish(tools.MakeJsonString(g), cfg.Kafka.TopicGetCompany)
		if err != nil {
			log.Println("pushlish error: " + err.Error())
			http.Error(w, err.Error(), 500)
			return
		}
		g, err = delivery.Consumer(cfg.Kafka.TopicUserGetResp, g.UUID)
		if err != nil {
			log.Println("consumer error: " + err.Error())
			http.Error(w, "service unavailable", 500)
			return
		}
	case models.Cust:
		err = delivery.Publish(tools.MakeJsonString(g), cfg.Kafka.TopicGetCustomer)
		if err != nil {
			log.Println("pushlish error: " + err.Error())
			http.Error(w, err.Error(), 500)
			return
		}
		g, err = delivery.Consumer(cfg.Kafka.TopicUserGetResp, g.UUID)
		if err != nil {
			log.Println("consumer error: " + err.Error())
			http.Error(w, "service unavailable", 500)
			return
		}
	}
	err = delivery.Publish(tools.MakeJsonString(g), cfg.Kafka.TopicGetWallet)
	if err != nil {
		log.Println("publish error: " + err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	g, err = delivery.Consumer(cfg.Kafka.TopicGetWalletResp, g.UUID)
	if err != nil {
		log.Println("consumer error: " + err.Error())
		http.Error(w, err.Error(), 500)
		return
	}
	g.UUID = ""
	w.Write([]byte(tools.MakeJsonString(g)))
}

type getGeneral struct {
	*models.SwaggerGetGeneral
}
