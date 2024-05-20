// Copyright (c) 2024 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0
package api

import "flag"

// defines several variables for parameterizing the client command.
var (
	OperationType      string
	Socket             string
	OperationInterface string
	OperationInput     string
)

func Init() {
	flag.StringVar(&OperationType, "operationType", "UnsealSecret", "The operation type to perform")
	flag.StringVar(&Socket, "socket", "/run/confidential-containers/cdh.sock", "The rpc socket path")
	flag.StringVar(&OperationInterface, "optInterface", "UnsealEnv", "The interface to use for the operation")
	flag.StringVar(&OperationInput, "optInput", "", "The input value to use for the operation interface")
}
