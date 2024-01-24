package model

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/erwinwahyura/go-boilerplate/app/model/constant"
	"github.com/erwinwahyura/go-boilerplate/utils"
)

const DEFAULT_PAGINATION_SIZE = 20

type (
	// BaseResponse is the base response
	BaseResponse struct {
		Code       string      `json:"code"`
		Message    string      `json:"message"`
		Data       interface{} `json:"data"`
		Meta       interface{} `json:"meta"`
		Errors     []string    `json:"errors"`
		ServerTime int64       `json:"server_time"`
	}
)

// MapBaseResponse map response
func MapBaseResponse(w http.ResponseWriter, r *http.Request, message string, data interface{}, meta interface{}, err error) {
	// Check Request ID
	requestID := r.Header.Get(constant.RequestID)
	if requestID != "" {
		bodyByte, _ := json.Marshal(data)
		fmt.Println("[RESPONSE: ", r.URL.String(), "] REQUEST_ID: ", requestID, " BODY:", string(bodyByte))
	}

	statusCode, code := utils.GetStatusCode(err)

	var errors []string
	if err != nil {
		errors = []string{err.Error()}
	}

	// Payload Response
	payload := BaseResponse{
		Code:       code,
		Message:    message,
		Data:       data,
		Errors:     errors,
		ServerTime: utils.TimeNow().Unix(),
		Meta:       meta,
	}

	// Marshal json response
	jsonResponse, _ := json.MarshalIndent(payload, "", "    ")

	// Write Response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonResponse)
}
