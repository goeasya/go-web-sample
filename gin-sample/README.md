
##  代码编译

编译docker镜像

> sh release.sh

## 填写启动参数

### 脚本方式启动

> 编辑start.sh脚本

```sh

# CONFIG_TYPE is zookeeper, must set:
CONFIG_ZK_NODE=""
CONFIG_ZK_SERVER=""

# CONFIG_TYPE is env, must set
DB_TYPE="DB_TYPE"
DB_COON_INFO="DB_CONN_INFO"
SECRET_KEY="SECRET_KEY"
API_ADDR="API_ADDR"
LOG_LEVEL="LOGLEVEL"
```


2. 修改获取配置信息方式CONFIG_TYPE，目前支持zookeeper和环境变量env。
   - 使用zookeeper，则配置CONFIG_ZK_* 开头的变量信息
   - 使用env，则配置ENV_* 开头的变量信息

zookeeper config示例：

```
# 节点为ginsample

/ginsample=api_addr=0.0.0.0:9999
/ginsample=db_auto_migrate=true
/ginsample=db_conn_info=xxx:xxx@tcp(127.0.0.1:3306)/ginsample
/ginsample=db_type=mysql
/ginsample=loglevel=debug
/ginsample=log_encoder=json
/ginsample=secret_key=xxx
```

> 运行gin-sample
>

```sh
sh start.sh
```

## 检查程序运行状态

```sh
docker logs -f gin-sample
```



# 接口文档

http://IP:PORT/swagger/index.html

更新swagger文档：

```shell
cd cmd/ginsample
swag init -o ../../docs --parseDependency
```

