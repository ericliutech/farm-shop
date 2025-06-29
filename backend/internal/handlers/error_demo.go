// SPDX-License-Identifier: MIT
// Copyright (c) 2025 Eric Liu

package handlers

import (
	"net/http"

	"github.com/ericliutech/farm-shop/backend/pkg/response"
)

func HandleErrorDemo(w http.ResponseWriter, r *http.Request) {
	response.Error(w, http.StatusInternalServerError, response.ErrorJsonResponse{
		Code:    response.CodeErrUnknown,
		Message: response.MessageMap[response.CodeErrUnknown],
		Error: map[string]interface{}{
			"Method": r.Method,
			"URL":    r.URL,
			"Header": r.Header,
			"Host":   r.Host,
		},
	})
}
