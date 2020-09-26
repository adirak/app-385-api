// Package apis contains an APIs for frontend
package apis

import (
	"net/http"

	"github.com/adirak/app-385-core/apis/webhook"
	"github.com/adirak/app-385-core/util"
)

// Service is function to handle http request and pass data to other function
func Service(w http.ResponseWriter, r *http.Request) {

	// Make ReqtData
	reqt, err := util.GetReqtData(r)

	if err == nil {

		// Call function
		// You can chage the function that you want to call here
		// Call StravaCallBack function
		// ============================

		resp := webhook.StravaCallBack(reqt)

		// ============================

		// Response JSON data to frontend
		util.SetJSONResponse(resp, w)

	} else {

		// Response JSON data to frontend
		util.SetBadRequestJSONResponse(err, w)

	}

}
