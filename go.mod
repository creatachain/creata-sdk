module github.com/creatachain/creata-sdk

go 1.16

replace google.golang.org/grpc => google.golang.org/grpc v1.33.2

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

require (
	github.com/99designs/keyring v1.1.6
	github.com/armon/go-metrics v0.3.10
	github.com/bgentry/speakeasy v0.1.0
	github.com/btcsuite/btcd v0.21.0-beta
	github.com/btcsuite/btcutil v1.0.2
	github.com/confio/ics23/go v0.6.6
	github.com/creatachain/augusteum v0.34.8
	github.com/creatachain/btcd v0.1.1
	github.com/creatachain/btcutil v0.0.0-20210413104051-16e75e19e838
	github.com/creatachain/creata-proto v0.3.1
	github.com/creatachain/crypto v0.0.0-20210408061910-5e21359f3593
	github.com/creatachain/go-amino v0.16.0
	github.com/creatachain/go-bip39 v1.0.0
	github.com/creatachain/iavl v0.15.3
	github.com/creatachain/ledger-creata-go v0.11.1
	github.com/creatachain/tm-db v0.6.4
	github.com/gogo/gateway v1.1.0
	github.com/gogo/protobuf v1.3.2
	github.com/golang/mock v1.6.0
	github.com/golang/protobuf v1.5.2
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/hashicorp/golang-lru v0.5.4
	github.com/magiconair/properties v1.8.5
	github.com/mattn/go-isatty v0.0.14
	github.com/otiai10/copy v1.7.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.12.1
	github.com/prometheus/common v0.32.1
	github.com/rakyll/statik v0.1.7
	github.com/rs/zerolog v1.26.1
	github.com/spf13/cast v1.4.1
	github.com/spf13/cobra v1.3.0
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.10.1
	github.com/stretchr/testify v1.7.0
	golang.org/x/crypto v0.0.0-20220131195533-30dcbda58838
	google.golang.org/genproto v0.0.0-20220204002441-d6cc3cc0770e
	google.golang.org/grpc v1.43.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/yaml.v2 v2.4.0
)
