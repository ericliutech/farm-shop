// SPDX-License-Identifier: MIT
// Copyright (c) 2025 Eric Liu

package handlers

import (
	"net/http"

	"github.com/ericliutech/farm-shop/backend/pkg/response"
)

func HandlePing(w http.ResponseWriter, r *http.Request) {
	response.Ok(w, http.StatusOK, response.OkJsonResponse{
		Code:    response.CodeOk,
		Message: response.MessageMap[response.CodeOk],
		Data:    nil,
	})
}
