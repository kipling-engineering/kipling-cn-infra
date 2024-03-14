module github.com/kipling-engineering/kipling-cn-infra

go 1.16

require (
	github.com/Shopify/sarama v1.20.0
	github.com/Songmu/prompter v0.0.0-20150725163906-b5721e8d5566
	github.com/alicebob/miniredis v2.4.5+incompatible
	github.com/boltdb/bolt v1.3.2-0.20180302180052-fd01fc79c553
	github.com/bshuster-repo/logrus-logstash-hook v1.1.0
	github.com/bsm/sarama-cluster v2.1.15+incompatible
	github.com/evalphobia/logrus_fluent v0.4.0
	github.com/fsnotify/fsnotify v1.4.7
	github.com/ghodss/yaml v1.0.0
	github.com/go-redis/redis v6.14.2+incompatible
	github.com/gocql/gocql v0.0.0-20181030013202-a84ce58083d3
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/gorilla/mux v1.6.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/hashicorp/consul/api v1.12.0
	github.com/hashicorp/consul/sdk v0.8.0
	github.com/howeyc/crc16 v0.0.0-20171223171357-2b2a61e366a6
	github.com/maraino/go-mock v0.0.0-20180321183845-4c74c434cd3a
	github.com/mitchellh/mapstructure v1.1.2
	github.com/namsral/flag v1.7.4-pre
	github.com/onsi/gomega v1.4.3
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.11.1
	github.com/sirupsen/logrus v1.8.1
	github.com/unrolled/render v0.0.0-20180914162206-b9786414de4d
	github.com/willfaught/gockle v0.0.0-20160623235217-4f254e1e0f0a
	go.etcd.io/etcd/api/v3 v3.5.7
	go.etcd.io/etcd/client/v3 v3.5.7
	go.etcd.io/etcd/server/v3 v3.5.7
	go.ligato.io/cn-infra/v2 v2.5.0-alpha.0.20230824082901-356dce1f1754
	golang.org/x/crypto v0.14.0
	golang.org/x/net v0.16.0
	golang.org/x/time v0.0.0-20210220033141-f8bda1e9f3ba
	google.golang.org/grpc v1.41.0
	google.golang.org/grpc/examples v0.0.0-20220120004855-f93e8e673710
	google.golang.org/protobuf v1.27.1
)

require (
	github.com/DataDog/zstd v1.5.5 // indirect
	github.com/alicebob/gopher-json v0.0.0-20230218143504-906a9b012302 // indirect
	github.com/eapache/go-resiliency v1.6.0 // indirect
	github.com/eapache/go-xerial-snappy v0.0.0-20230731223053-c322873962e3 // indirect
	github.com/eapache/queue v1.1.0 // indirect
	github.com/fluent/fluent-logger-golang v1.9.0 // indirect
	github.com/gomodule/redigo v1.9.2 // indirect
	github.com/gorilla/context v1.1.2 // indirect
	github.com/pierrec/lz4 v2.6.1+incompatible // indirect
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
	github.com/tinylib/msgp v1.1.9 // indirect
	github.com/yuin/gopher-lua v1.1.1 // indirect
)

replace go.ligato.io/cn-infra/v2 => ./
