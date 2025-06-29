// SPDX-License-Identifier: MIT
// Copyright (c) 2025 Eric Liu

package main

import (
	"net/http"

	"github.com/ericliutech/farm-shop/backend/internal/handlers"
	"github.com/ericliutech/farm-shop/backend/pkg/logger"
	"github.com/go-chi/chi/v5"
)

func main() {
	logger.Init()

	r := chi.NewRouter()

	r.Get("/ping", handlers.HandlePing)
	r.Get("/error-demo", handlers.HandleErrorDemo)

	logger.Info("Starting Farm Shop backend...")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Info("Farm Shop backend has stopped.")
}
