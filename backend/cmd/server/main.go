// SPDX-License-Identifier: MIT
// Copyright (c) 2025 Eric Liu

package main

import (
	"github.com/ericliutech/farm-shop/backend/pkg/logger"
)

func main() {
	logger.Init()

	logger.Info("Farm Shop backend ran ok.")
}
