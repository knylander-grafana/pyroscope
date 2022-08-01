// Code generated by protoc-gen-connect-go-mux. DO NOT EDIT.
//
// Source: querier/v1/querier.proto

package querierv1connect

import (
	connect_go "github.com/bufbuild/connect-go"
	mux "github.com/gorilla/mux"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

// RegisterQuerierServiceHandler register an HTTP handler to a mux.Router from the service
// implementation.
func RegisterQuerierServiceHandler(mux *mux.Router, svc QuerierServiceHandler, opts ...connect_go.HandlerOption) {
	mux.Handle("/querier.v1.QuerierService/ProfileTypes", connect_go.NewUnaryHandler(
		"/querier.v1.QuerierService/ProfileTypes",
		svc.ProfileTypes,
		opts...,
	))
}
