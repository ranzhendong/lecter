# kubernetes-custom-endpoint-healthcheck

![Travis (.org) branch](https://img.shields.io/travis/ranzhendong/lecter/master?style=plastic)
![GitHub release (latest by date including pre-releases)](https://img.shields.io/github/v/release/ranzhendong/lecter?include_prereleases&style=plastic)
![GitHub last commit (branch)](https://img.shields.io/github/last-commit/ranzhendong/lecter/master?style=plastic)
![](https://img.shields.io/github/license/ranzhendong/lecter?style=plastic)

&emsp;&emsp;以go语言为载体实现对于kubernetes集群当中自建外部访问endpoint资源的健康检查。

- 外部读取配置文件，或者通过参数 **-f** 指定配置文件地址。
- 采用**time.Sleep**进行循环检查，没有使用**time.NewTicker**。
- 脚本建议跑在单独的namespace和pod当中，使用serveraccount实现对ApiServer访问读写。



- [kubernetes-custom-endpoint-healthcheck](#kubernetes-custom-endpoint-healthcheck)
  - [配置](#配置)
    - [脚本配置 my_conf](#脚本配置my_conf)
      - [interval健康检查间隔](#interval-健康检查间隔 )
      - [pingtimeout-ping超时时间](#pingtimeout-ping超时时间 )
    - [k8s配置kubernetes_conf](#k8s配置kubernetes_conf)
      - [url访问k8s集群ApiServer地址](#url-访问k8s集群ApiServer地址 )
      - [endpointapi访问endpoint资源api ](#endpointapi-访问endpoint资源api )
      - [content_type写入时数据包类型](#content_type-写入时数据包类型 )
      - [token_file访问k8s集群ApiServer的token](#token_file-访问k8s集群ApiServer的token )
    - [存储配置storage_conf](#存储配置storage_conf)
      - [redis-redis存储配置](#redis-redis存储配置 )
      - [ipport-redis IP端口配置](#ipport-redisIP端口配置 )
      - [password-redis访问auth](#password-redis访问auth)
    - [service配置service_info](#service配置service_info)
      - [redis-示例名称](#redis-示例名称 )
      - [namespace-endpoint命名空间](#namespace-endpoint命名空间 )
      - [name-endpoint名称](#name-endpoint名称 )
    - [endpoint健康检查配置service_healthcheck_list](#endpoint健康检查配置service_healthcheck_list )
      - [redis-示例名称](#redis-示例名称 )
      - [port-端口](#port-端口 )
      - [iplist-ip列表](#iplist-ip列表 )
    - [endpoint.yaml配置endpoint_template](#endpoint.yaml配置endpoint_template )



## 配置

&emsp;&emsp;通过[config.yaml](https://github.com/ranzhendong/lecter/blob/master/src/config.yaml)进行配置文件管理，配置文件位置建议和脚本同级或者在同级目录下的config文件夹下。



### 脚本配置my_conf

&emsp;&emsp;基础脚本运行配置。



#### interval-健康检查间隔 

```yaml
interval: 2000
```



#### pingtimeout-ping超时时间

&emsp;&emsp;指的是通过net包对ip port进行4层访问测试的超时时间。

```yaml
pingtimeout: 3000
```



### k8s配置kubernetes_conf

​		通过这个配置实现对k8sApiServer基本读写配置。



#### url-访问k8s集群ApiServer地址

&emsp;&emsp;在pod当中运行，可以填写域名来实现内部访问，示例配置为外部访问。

```yaml
url: https://172.16.0.60:6443
```



#### endpointapi-访问endpoint资源api

&emsp;&emsp;根据k8s版本不同，api的路径上有所不同，需要自己进行配置。

&emsp;&emsp;需要注意的是**myNameSpaces**和**myEndPoints**是后面会进行替换的字段，因此路径自定义但是保证这两个关键字存在。

```yaml
endpointapi: /api/v1/namespaces/myNameSpaces/endpoints/myEndPoints
```



#### content_type-写入时数据包类型

&emsp;&emsp;因为写在代码当中，因此不建议更改，保持**application/yaml**就好。

```yaml
content_type: application/yaml
```



#### token_file-访问k8s集群ApiServer的token

&emsp;&emsp;如果为空，不填：

- linux系统下会默认访问如下路径：`/var/run/secrets/kubernetes.io/serviceaccount/token`

- window系统下会默认访问如下路径：`C:\Users\Administrator\go\my_script\src\lepai.token\token`

&emsp;&emsp;如果为**default**：

- linux和window下，会在脚本执行的相对路径`./`以及`./conf/`进行token文件寻找。

&emsp;&emsp;指定文件路径：

- 例如：C:\Users\Administrator\go\my_script\src\lepai.token\token


&emsp;&emsp;**因此建议在pod当中使用挂载的serveraccount文件来实现。也就是`/var/run/secrets/kubernetes.io/serviceaccount/token`。**

```yaml
token_file: default
```



### 存储配置storage_conf

&emsp;&emsp;对健康检查信息进行存储，方便迭代。

&emsp;&emsp;后期计划加入ectd存储选项，保证多个相同脚本同时运行，分布式锁的实现。



#### redis-redis存储配置

```yaml
  redis:
    ipport: 172.16.0.61:6379
    password:
```



#### ipport-redisIP端口配置

```yaml
ipport: 172.16.0.61:6379
```



#### password-redis访问auth

&emsp;&emsp;如果没有密码，保持空就可以

```yaml
password: 123
```



### service配置service_info

&emsp;&emsp;这部分配置主要是为了保证endpoint对应的service信息一致，也就是需要通过`/api/v1/namespaces/myNameSpaces/endpoints/myEndPoints`来对某个endpoint资源实现访问，示例当中配置了redis集群，eqmx集群以及ceph集群。

&emsp;&emsp;最后为了保证可用性，可以自行通过命令行访问**k8sApiServer**的endpoint接口：`/api/v1/namespaces/myNameSpaces/endpoints/myEndPoints`来验证，注意需要对**myNameSpaces**和**myEndPoints**进行替换。



#### redis-示例名称

&emsp;&emsp;可自定义，但是需要和下面的保持一致。

```yaml
service_info:
  redis:
  emqx:
  ceph:
```



#### namespace-endpoint命名空间

```yaml
namespace: default
```



#### **name**-endpoint名称

```yaml
name: redis
```



### endpoint健康检查配置service_healthcheck_list

&emsp;&emsp;这部分是配置是告诉脚本这些ip地址和port端口是正常的，如果检测到从k8sApiServer取出的数据和它不一致，那么就可以确定是有ip地址对应的节点down。

#### redis-示例名称

&emsp;&emsp;可自定义，注意需要保持和上面的service_info当中的名称保持一致。

```yaml
service_healthcheck_list:
  redis:
  emqx:
  ceph:
```



#### port-端口

&emsp;&emsp;指定endpoint访问端口其实也就是健康检查端口。

```yaml
port: 38080
```



#### iplist-ip列表

&emsp;&emsp;指定一个服务的多个节点ip信息。

&emsp;&emsp;需要通过name进行区分，方便后面的钉钉报警组件（还未开发）

```yaml
  - iplist:
    - name: node-01
      ip: 172.16.0.61
    - name: node-02
      ip: 172.16.0.62
    - name: node-03
      ip: 172.16.0.63
```



### endpoint.yaml配置endpoint_template

&emsp;&emsp;模板文件，需要通过它将主要的**subsets**字段进行替换，对通过健康检查的iplist进行更新，然后写入到k8s，实现覆盖，保证service访问。

&emsp;&emsp;这部分除非k8s版本不一致导致需要更改**apiVersion**字段之外，其他的不建议更改，如果因为k8s导致**subsets**字段以及结构不一致的，则不建议使用该脚本，因为脚本对于字符的处理是基于下面的模板进行的，也可以进行自己需求的二次开发。

```yaml
endpoint_template:
  kind: Endpoints
  apiVersion: v1
  metadata:
    name: ceph
    namespace: default
  subsets:
    - addresses:
        - ip: 10.0.0.71
        - ip: 10.0.0.72
      ports:
        - port: 8080
          protocol: TCP
```



## 使用

&emsp;&emsp;使用方式在前面有叙述

### 编译

```shell
$ go build ./src/...
```



### 运行

```powershell
main.exe

2019/11/25 17:58:34 main.go:94: 检查次数:[ 1 ]
2019/11/25 17:58:36 main.go:94: 检查次数:[ 2 ]
2019/11/25 17:58:38 main.go:94: 检查次数:[ 3 ]
2019/11/25 17:58:39 healthcheck.go:55: {Name:redis} {Addresses:[map[ip:172.16.0.61]] != SuccessMapList: [map[ip:172.16.0.61] map[ip:172.16.0.62]]}
2019/11/25 17:58:40 main.go:94: 检查次数:[ 4 ]
2019/11/25 17:58:41 healthcheck.go:52: {Name:redis} {Addresses:[map[ip:172.16.0.61] map[ip:172.16.0.62]] == SuccessMapList: [map[ip:172.16.0.61] map[ip:172.16.0.62]]}
redis: nil
2019/11/25 17:58:41 healthcheck.go:41: RedisListGet CAN'T GET
2019/11/25 17:58:41 healthcheck.go:52: {Name:ceph} {Addresses:[map[ip:172.16.0.61] map[ip:172.16.0.62]] == SuccessMapList: [map[ip:172.16.0.61] map[ip:172.16.0.62]]}
2019/11/25 17:58:42 main.go:94: 检查次数:[ 5 ]
2019/11/25 17:58:43 healthcheck.go:52: {Name:redis} {Addresses:[map[ip:172.16.0.61] map[ip:172.16.0.62]] == SuccessMapList: [map[ip:172.16.0.61] map[ip:172.16.0.62]]}
redis: nil
2019/11/25 17:58:43 healthcheck.go:41: RedisListGet CAN'T GET
2019/11/25 17:58:43 healthcheck.go:52: {Name:ceph} {Addresses:[map[ip:172.16.0.61] map[ip:172.16.0.62]] == SuccessMapList: [map[ip:172.16.0.61] map[ip:172.16.0.62]]}
```



