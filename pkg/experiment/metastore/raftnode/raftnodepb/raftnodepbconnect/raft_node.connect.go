// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: experiment/metastore/raftnode/raftnodepb/raft_node.proto

package raftnodepbconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	raftnodepb "github.com/grafana/pyroscope/pkg/experiment/metastore/raftnode/raftnodepb"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// RaftNodeServiceName is the fully-qualified name of the RaftNodeService service.
	RaftNodeServiceName = "raft_node.RaftNodeService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// RaftNodeServiceReadIndexProcedure is the fully-qualified name of the RaftNodeService's ReadIndex
	// RPC.
	RaftNodeServiceReadIndexProcedure = "/raft_node.RaftNodeService/ReadIndex"
	// RaftNodeServiceNodeInfoProcedure is the fully-qualified name of the RaftNodeService's NodeInfo
	// RPC.
	RaftNodeServiceNodeInfoProcedure = "/raft_node.RaftNodeService/NodeInfo"
	// RaftNodeServiceRemoveNodeProcedure is the fully-qualified name of the RaftNodeService's
	// RemoveNode RPC.
	RaftNodeServiceRemoveNodeProcedure = "/raft_node.RaftNodeService/RemoveNode"
	// RaftNodeServiceAddNodeProcedure is the fully-qualified name of the RaftNodeService's AddNode RPC.
	RaftNodeServiceAddNodeProcedure = "/raft_node.RaftNodeService/AddNode"
	// RaftNodeServiceDemoteLeaderProcedure is the fully-qualified name of the RaftNodeService's
	// DemoteLeader RPC.
	RaftNodeServiceDemoteLeaderProcedure = "/raft_node.RaftNodeService/DemoteLeader"
	// RaftNodeServicePromoteToLeaderProcedure is the fully-qualified name of the RaftNodeService's
	// PromoteToLeader RPC.
	RaftNodeServicePromoteToLeaderProcedure = "/raft_node.RaftNodeService/PromoteToLeader"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	raftNodeServiceServiceDescriptor               = raftnodepb.File_experiment_metastore_raftnode_raftnodepb_raft_node_proto.Services().ByName("RaftNodeService")
	raftNodeServiceReadIndexMethodDescriptor       = raftNodeServiceServiceDescriptor.Methods().ByName("ReadIndex")
	raftNodeServiceNodeInfoMethodDescriptor        = raftNodeServiceServiceDescriptor.Methods().ByName("NodeInfo")
	raftNodeServiceRemoveNodeMethodDescriptor      = raftNodeServiceServiceDescriptor.Methods().ByName("RemoveNode")
	raftNodeServiceAddNodeMethodDescriptor         = raftNodeServiceServiceDescriptor.Methods().ByName("AddNode")
	raftNodeServiceDemoteLeaderMethodDescriptor    = raftNodeServiceServiceDescriptor.Methods().ByName("DemoteLeader")
	raftNodeServicePromoteToLeaderMethodDescriptor = raftNodeServiceServiceDescriptor.Methods().ByName("PromoteToLeader")
)

// RaftNodeServiceClient is a client for the raft_node.RaftNodeService service.
type RaftNodeServiceClient interface {
	ReadIndex(context.Context, *connect.Request[raftnodepb.ReadIndexRequest]) (*connect.Response[raftnodepb.ReadIndexResponse], error)
	NodeInfo(context.Context, *connect.Request[raftnodepb.NodeInfoRequest]) (*connect.Response[raftnodepb.NodeInfoResponse], error)
	RemoveNode(context.Context, *connect.Request[raftnodepb.RemoveNodeRequest]) (*connect.Response[raftnodepb.RemoveNodeResponse], error)
	AddNode(context.Context, *connect.Request[raftnodepb.AddNodeRequest]) (*connect.Response[raftnodepb.AddNodeResponse], error)
	DemoteLeader(context.Context, *connect.Request[raftnodepb.DemoteLeaderRequest]) (*connect.Response[raftnodepb.DemoteLeaderResponse], error)
	PromoteToLeader(context.Context, *connect.Request[raftnodepb.PromoteToLeaderRequest]) (*connect.Response[raftnodepb.PromoteToLeaderResponse], error)
}

// NewRaftNodeServiceClient constructs a client for the raft_node.RaftNodeService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewRaftNodeServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) RaftNodeServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &raftNodeServiceClient{
		readIndex: connect.NewClient[raftnodepb.ReadIndexRequest, raftnodepb.ReadIndexResponse](
			httpClient,
			baseURL+RaftNodeServiceReadIndexProcedure,
			connect.WithSchema(raftNodeServiceReadIndexMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		nodeInfo: connect.NewClient[raftnodepb.NodeInfoRequest, raftnodepb.NodeInfoResponse](
			httpClient,
			baseURL+RaftNodeServiceNodeInfoProcedure,
			connect.WithSchema(raftNodeServiceNodeInfoMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		removeNode: connect.NewClient[raftnodepb.RemoveNodeRequest, raftnodepb.RemoveNodeResponse](
			httpClient,
			baseURL+RaftNodeServiceRemoveNodeProcedure,
			connect.WithSchema(raftNodeServiceRemoveNodeMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		addNode: connect.NewClient[raftnodepb.AddNodeRequest, raftnodepb.AddNodeResponse](
			httpClient,
			baseURL+RaftNodeServiceAddNodeProcedure,
			connect.WithSchema(raftNodeServiceAddNodeMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		demoteLeader: connect.NewClient[raftnodepb.DemoteLeaderRequest, raftnodepb.DemoteLeaderResponse](
			httpClient,
			baseURL+RaftNodeServiceDemoteLeaderProcedure,
			connect.WithSchema(raftNodeServiceDemoteLeaderMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		promoteToLeader: connect.NewClient[raftnodepb.PromoteToLeaderRequest, raftnodepb.PromoteToLeaderResponse](
			httpClient,
			baseURL+RaftNodeServicePromoteToLeaderProcedure,
			connect.WithSchema(raftNodeServicePromoteToLeaderMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// raftNodeServiceClient implements RaftNodeServiceClient.
type raftNodeServiceClient struct {
	readIndex       *connect.Client[raftnodepb.ReadIndexRequest, raftnodepb.ReadIndexResponse]
	nodeInfo        *connect.Client[raftnodepb.NodeInfoRequest, raftnodepb.NodeInfoResponse]
	removeNode      *connect.Client[raftnodepb.RemoveNodeRequest, raftnodepb.RemoveNodeResponse]
	addNode         *connect.Client[raftnodepb.AddNodeRequest, raftnodepb.AddNodeResponse]
	demoteLeader    *connect.Client[raftnodepb.DemoteLeaderRequest, raftnodepb.DemoteLeaderResponse]
	promoteToLeader *connect.Client[raftnodepb.PromoteToLeaderRequest, raftnodepb.PromoteToLeaderResponse]
}

// ReadIndex calls raft_node.RaftNodeService.ReadIndex.
func (c *raftNodeServiceClient) ReadIndex(ctx context.Context, req *connect.Request[raftnodepb.ReadIndexRequest]) (*connect.Response[raftnodepb.ReadIndexResponse], error) {
	return c.readIndex.CallUnary(ctx, req)
}

// NodeInfo calls raft_node.RaftNodeService.NodeInfo.
func (c *raftNodeServiceClient) NodeInfo(ctx context.Context, req *connect.Request[raftnodepb.NodeInfoRequest]) (*connect.Response[raftnodepb.NodeInfoResponse], error) {
	return c.nodeInfo.CallUnary(ctx, req)
}

// RemoveNode calls raft_node.RaftNodeService.RemoveNode.
func (c *raftNodeServiceClient) RemoveNode(ctx context.Context, req *connect.Request[raftnodepb.RemoveNodeRequest]) (*connect.Response[raftnodepb.RemoveNodeResponse], error) {
	return c.removeNode.CallUnary(ctx, req)
}

// AddNode calls raft_node.RaftNodeService.AddNode.
func (c *raftNodeServiceClient) AddNode(ctx context.Context, req *connect.Request[raftnodepb.AddNodeRequest]) (*connect.Response[raftnodepb.AddNodeResponse], error) {
	return c.addNode.CallUnary(ctx, req)
}

// DemoteLeader calls raft_node.RaftNodeService.DemoteLeader.
func (c *raftNodeServiceClient) DemoteLeader(ctx context.Context, req *connect.Request[raftnodepb.DemoteLeaderRequest]) (*connect.Response[raftnodepb.DemoteLeaderResponse], error) {
	return c.demoteLeader.CallUnary(ctx, req)
}

// PromoteToLeader calls raft_node.RaftNodeService.PromoteToLeader.
func (c *raftNodeServiceClient) PromoteToLeader(ctx context.Context, req *connect.Request[raftnodepb.PromoteToLeaderRequest]) (*connect.Response[raftnodepb.PromoteToLeaderResponse], error) {
	return c.promoteToLeader.CallUnary(ctx, req)
}

// RaftNodeServiceHandler is an implementation of the raft_node.RaftNodeService service.
type RaftNodeServiceHandler interface {
	ReadIndex(context.Context, *connect.Request[raftnodepb.ReadIndexRequest]) (*connect.Response[raftnodepb.ReadIndexResponse], error)
	NodeInfo(context.Context, *connect.Request[raftnodepb.NodeInfoRequest]) (*connect.Response[raftnodepb.NodeInfoResponse], error)
	RemoveNode(context.Context, *connect.Request[raftnodepb.RemoveNodeRequest]) (*connect.Response[raftnodepb.RemoveNodeResponse], error)
	AddNode(context.Context, *connect.Request[raftnodepb.AddNodeRequest]) (*connect.Response[raftnodepb.AddNodeResponse], error)
	DemoteLeader(context.Context, *connect.Request[raftnodepb.DemoteLeaderRequest]) (*connect.Response[raftnodepb.DemoteLeaderResponse], error)
	PromoteToLeader(context.Context, *connect.Request[raftnodepb.PromoteToLeaderRequest]) (*connect.Response[raftnodepb.PromoteToLeaderResponse], error)
}

// NewRaftNodeServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewRaftNodeServiceHandler(svc RaftNodeServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	raftNodeServiceReadIndexHandler := connect.NewUnaryHandler(
		RaftNodeServiceReadIndexProcedure,
		svc.ReadIndex,
		connect.WithSchema(raftNodeServiceReadIndexMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	raftNodeServiceNodeInfoHandler := connect.NewUnaryHandler(
		RaftNodeServiceNodeInfoProcedure,
		svc.NodeInfo,
		connect.WithSchema(raftNodeServiceNodeInfoMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	raftNodeServiceRemoveNodeHandler := connect.NewUnaryHandler(
		RaftNodeServiceRemoveNodeProcedure,
		svc.RemoveNode,
		connect.WithSchema(raftNodeServiceRemoveNodeMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	raftNodeServiceAddNodeHandler := connect.NewUnaryHandler(
		RaftNodeServiceAddNodeProcedure,
		svc.AddNode,
		connect.WithSchema(raftNodeServiceAddNodeMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	raftNodeServiceDemoteLeaderHandler := connect.NewUnaryHandler(
		RaftNodeServiceDemoteLeaderProcedure,
		svc.DemoteLeader,
		connect.WithSchema(raftNodeServiceDemoteLeaderMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	raftNodeServicePromoteToLeaderHandler := connect.NewUnaryHandler(
		RaftNodeServicePromoteToLeaderProcedure,
		svc.PromoteToLeader,
		connect.WithSchema(raftNodeServicePromoteToLeaderMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/raft_node.RaftNodeService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case RaftNodeServiceReadIndexProcedure:
			raftNodeServiceReadIndexHandler.ServeHTTP(w, r)
		case RaftNodeServiceNodeInfoProcedure:
			raftNodeServiceNodeInfoHandler.ServeHTTP(w, r)
		case RaftNodeServiceRemoveNodeProcedure:
			raftNodeServiceRemoveNodeHandler.ServeHTTP(w, r)
		case RaftNodeServiceAddNodeProcedure:
			raftNodeServiceAddNodeHandler.ServeHTTP(w, r)
		case RaftNodeServiceDemoteLeaderProcedure:
			raftNodeServiceDemoteLeaderHandler.ServeHTTP(w, r)
		case RaftNodeServicePromoteToLeaderProcedure:
			raftNodeServicePromoteToLeaderHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedRaftNodeServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedRaftNodeServiceHandler struct{}

func (UnimplementedRaftNodeServiceHandler) ReadIndex(context.Context, *connect.Request[raftnodepb.ReadIndexRequest]) (*connect.Response[raftnodepb.ReadIndexResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("raft_node.RaftNodeService.ReadIndex is not implemented"))
}

func (UnimplementedRaftNodeServiceHandler) NodeInfo(context.Context, *connect.Request[raftnodepb.NodeInfoRequest]) (*connect.Response[raftnodepb.NodeInfoResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("raft_node.RaftNodeService.NodeInfo is not implemented"))
}

func (UnimplementedRaftNodeServiceHandler) RemoveNode(context.Context, *connect.Request[raftnodepb.RemoveNodeRequest]) (*connect.Response[raftnodepb.RemoveNodeResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("raft_node.RaftNodeService.RemoveNode is not implemented"))
}

func (UnimplementedRaftNodeServiceHandler) AddNode(context.Context, *connect.Request[raftnodepb.AddNodeRequest]) (*connect.Response[raftnodepb.AddNodeResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("raft_node.RaftNodeService.AddNode is not implemented"))
}

func (UnimplementedRaftNodeServiceHandler) DemoteLeader(context.Context, *connect.Request[raftnodepb.DemoteLeaderRequest]) (*connect.Response[raftnodepb.DemoteLeaderResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("raft_node.RaftNodeService.DemoteLeader is not implemented"))
}

func (UnimplementedRaftNodeServiceHandler) PromoteToLeader(context.Context, *connect.Request[raftnodepb.PromoteToLeaderRequest]) (*connect.Response[raftnodepb.PromoteToLeaderResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("raft_node.RaftNodeService.PromoteToLeader is not implemented"))
}
