// SPDX-License-Identifier: AGPL-3.0-only

package parse

import (
	"reflect"

	"github.com/grafana/dskit/grpcclient"
	"github.com/grafana/dskit/kv/memberlist"
	"github.com/weaveworks/common/server"

	"github.com/grafana/phlare/pkg/agent"
	"github.com/grafana/phlare/pkg/distributor"
	"github.com/grafana/phlare/pkg/ingester"
	"github.com/grafana/phlare/pkg/objstore/providers/azure"
	"github.com/grafana/phlare/pkg/objstore/providers/filesystem"
	"github.com/grafana/phlare/pkg/objstore/providers/gcs"
	"github.com/grafana/phlare/pkg/objstore/providers/s3"
	"github.com/grafana/phlare/pkg/objstore/providers/swift"
	"github.com/grafana/phlare/pkg/querier"
)

var (
	// RootBlocks is an ordered list of root blocks. The order is the same order that will
	// follow the markdown generation.
	RootBlocks = []RootBlock{
		{
			Name:       "agent",
			StructType: reflect.TypeOf(agent.Config{}),
			Desc:       "The agent block configures the pull-mode collection of profiles.",
		},
		{
			Name:       "server",
			StructType: reflect.TypeOf(server.Config{}),
			Desc:       "The server block configures the HTTP and gRPC server of the launched service(s).",
		},
		{
			Name:       "distributor",
			StructType: reflect.TypeOf(distributor.Config{}),
			Desc:       "The distributor block configures the distributor.",
		},
		{
			Name:       "ingester",
			StructType: reflect.TypeOf(ingester.Config{}),
			Desc:       "The ingester block configures the ingester.",
		},
		{
			Name:       "querier",
			StructType: reflect.TypeOf(querier.Config{}),
			Desc:       "The querier block configures the querier.",
		},
		{
			Name:       "grpc_client",
			StructType: reflect.TypeOf(grpcclient.Config{}),
			Desc:       "The grpc_client block configures the gRPC client used to communicate between two Mimir components.",
		},
		{
			Name:       "memberlist",
			StructType: reflect.TypeOf(memberlist.KVConfig{}),
			Desc:       "The memberlist block configures the Gossip memberlist.",
		},
		{
			Name:       "s3_storage_backend",
			StructType: reflect.TypeOf(s3.Config{}),
			Desc:       "The s3_backend block configures the connection to Amazon S3 object storage backend.",
		},
		{
			Name:       "gcs_storage_backend",
			StructType: reflect.TypeOf(gcs.Config{}),
			Desc:       "The gcs_backend block configures the connection to Google Cloud Storage object storage backend.",
		},
		{
			Name:       "azure_storage_backend",
			StructType: reflect.TypeOf(azure.Config{}),
			Desc:       "The azure_storage_backend block configures the connection to Azure object storage backend.",
		},
		{
			Name:       "swift_storage_backend",
			StructType: reflect.TypeOf(swift.Config{}),
			Desc:       "The swift_storage_backend block configures the connection to OpenStack Object Storage (Swift) object storage backend.",
		},
		{
			Name:       "filesystem_storage_backend",
			StructType: reflect.TypeOf(filesystem.Config{}),
			Desc:       "The filesystem_storage_backend block configures the usage of local file system as object storage backend.",
		},
	}
)