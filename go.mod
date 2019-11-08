module github.com/HuaJuanJiang/kubernetes-custom-endpoint-healthcheck

go 1.13

require (
	conf v1.0.1
	github.com/golang/protobuf v1.3.2 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	golang.org/x/net v0.0.0-20191002035440-2ec189313ef0 // indirect
	golang.org/x/sys v0.0.0-20191010194322-b09406accb47 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	healthcheck v1.0.1
	k8sapi v1.0.1
	manageyaml v1.0.1
	myredis v1.0.1
	transformation v1.0.1
)

replace conf v1.0.1 => ./src/public/conf

replace deepcopy v1.0.1 => ./src/public/deepcopy

replace mynet v1.0.1 => ./src/public/mynet

replace healthcheck v1.0.1 => ./src/public/healthcheck

replace k8sapi v1.0.1 => ./src/public/k8sapi

replace myredis v1.0.1 => ./src/public/storage

replace manageyaml v1.0.1 => ./src/public/manageyaml

replace transformation v1.0.1 => ./src/public/transformation
