module fiber-demo

require (
	github.com/envoyproxy/protoc-gen-validate v0.1.0
	github.com/go-kratos/kratos/contrib/registry/consul/v2 v2.0.0-20220215033000-471a2aab794b
	github.com/go-kratos/kratos/v2 v2.2.0
	github.com/gofiber/contrib/fiberzap v0.0.0-20220309073544-9ca8435e0cb7
	github.com/gofiber/fiber/v2 v2.29.0
	github.com/golang/protobuf v1.5.2
	github.com/google/wire v0.5.0
	github.com/hashicorp/consul/api v1.12.0
	github.com/jassue/go-storage v1.0.1
	github.com/jinzhu/copier v0.3.5
	github.com/natefinch/lumberjack v2.0.0+incompatible
	go.uber.org/zap v1.21.0
	google.golang.org/genproto v0.0.0-20220304144024-325a89244dc8
	google.golang.org/grpc v1.44.0
	google.golang.org/protobuf v1.27.1
)

require (
	github.com/aliyun/aliyun-oss-go-sdk v2.2.1+incompatible // indirect
	github.com/andybalholm/brotli v1.0.4 // indirect
	github.com/armon/go-metrics v0.0.0-20180917152333-f0300d1749da // indirect
	github.com/fatih/color v1.9.0 // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/go-playground/form/v4 v4.2.0 // indirect
	github.com/google/go-cmp v0.5.7 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.1 // indirect
	github.com/hashicorp/go-hclog v0.12.0 // indirect
	github.com/hashicorp/go-immutable-radix v1.0.0 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/golang-lru v0.5.0 // indirect
	github.com/hashicorp/serf v0.9.6 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.15.0 // indirect
	github.com/mattn/go-colorable v0.1.6 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/mapstructure v1.1.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/qiniu/go-sdk/v7 v7.11.1 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.34.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	golang.org/x/net v0.0.0-20220225172249-27dd8689420f // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/sys v0.0.0-20220227234510-4e6760a101f9 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/time v0.0.0-20220224211638-0e9765cccd65 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace github.com/go-kratos/kratos/v2 => github.com/Paradox315/kratos/v2 v2.2.9

replace github.com/go-kratos/kratos/contrib/registry/consul/v2 => github.com/Paradox315/kratos/contrib/registry/consul/v2 v2.0.0-20220216150651-8573d06d9606

go 1.17
