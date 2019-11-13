# kubernetes-custom-endpoint-healthcheck

![Travis (.org) branch](https://img.shields.io/travis/HuaJuanJiang/lecter/master?style=plastic)
![GitHub release (latest by date including pre-releases)](https://img.shields.io/github/v/release/HuaJuanJiang/lecter?include_prereleases&style=plastic)
![GitHub last commit (branch)](https://img.shields.io/github/last-commit/HuaJuanJiang/lecter/master?style=plastic)
![](https://img.shields.io/github/license/HuaJuanJiang/lecter?style=plastic)

&emsp;&emsp;以go语言为载体实现对于kubernetes集群当中自建外部访问endpoint资源的健康检查。

- 外部读取配置文件，或者通过参数 **-f** 指定配置文件地址。
- 采用**time.Sleep**进行循环检查，没有使用**time.NewTicker**。

- 脚本建议跑在单独的namespace和pod当中，使用serveraccount实现对ApiServer访问读写。

### 配置文件

#### 脚本配置 my_conf

**interval**  健康检查间隔 

```yaml
interval: 2000
```

**pingtimeout**  ping超时时间

&emsp;&emsp;指的是通过net包对ip port进行4层访问测试的超时时间。

```yaml
pingtimeout: 3000
```



#### k8s访问配置 kubernetes_conf

**url**  访问k8s集群ApiServer地址

​		在pod当中运行，可以填写域名来实现内部访问，示例配置为外部访问。

```yaml
url: https://172.16.0.60:6443
```

**endpointapi**  访问endpoint资源api

&emsp;&emsp;根据k8s版本不同，api的路径上有所不同，需要自己进行配置。

&emsp;&emsp;需要注意的是**myNameSpaces**和**myEndPoints**是后面会进行替换的字段，因此路径自定义但是保证这两个关键字存在。

```yaml
endpointapi: /api/v1/namespaces/myNameSpaces/endpoints/myEndPoints
```

**content_type**  写入时数据包类型

&emsp;&emsp;因为写在代码当中，因此不建议更改，保持**application/yaml**就好。

```yaml
content_type: application/yaml
```

**token_file**  访问k8s集群ApiServer的token

&emsp;&emsp;如果为空，不填：

- linux系统下会默认访问如下路径：`/var/run/secrets/kubernetes.io/serviceaccount/token`

- window系统下会默认访问如下路径：`C:\Users\Administrator\go\my_script\src\lepai.token\token`

&emsp;&emsp;如果为**default**：

- linux和window下，会在脚本执行的相对路径`./`以及`./conf/`进行token文件寻找。

&emsp;&emsp;指定文件路径：

- 例如：C:\Users\Administrator\go\my_script\src\lepai.token\token


&emsp;&emsp;因此建议在pod当中使用挂载的serveraccount文件来实现。也就是`/var/run/secrets/kubernetes.io/serviceaccount/token`

```yaml
token_file: default
```











