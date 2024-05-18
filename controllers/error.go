package controllers

import (
	"net/http"

	"github.com/Vector-ops/rss-aggregator/utils"
)

func HandlerErr(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithError(w, 400, "Something went wrong")
}
