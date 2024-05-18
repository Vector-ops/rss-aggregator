package controllers

import (
	"net/http"

	"github.com/Vector-ops/rss-aggregator/utils"
)

func HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, 200, struct{}{})
}
