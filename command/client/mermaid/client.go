// Copyright (C) 2017 ScyllaDB

package mermaid

import (
	api "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/scylladb/mermaid/command/client/mermaid/internal/client/operations"
)

// NewClient returns a new mermaid rest API client.
func NewClient(host string) *operations.Client {
	return operations.New(api.New(host, "/api/v1", []string{"http"}), strfmt.Default)
}
