# mongod参数

[TOC]

创建数据库目录

mkdir -p /opt/mongo
```
mongod参数
---
 基本配置
---
-f [--config ] arg    configuration file specifying additional options
--quiet # 安静输出
--port arg  # 指定服务端口号，默认端口27017
--bind_ip arg   # 绑定服务IP，若绑定127.0.0.1，则只能本机访问，不指定默认本地所有IP
--logpath arg   # 指定MongoDB日志文件，注意是指定文件不是目录
--logappend # 使用追加的方式写日志
--pidfilepath arg   # PID File 的完整路径，如果没有设置，则没有PID文件
--keyFile arg   # 集群的私钥的完整路径，只对于Replica Set 架构有效
--unixSocketPrefix arg  # UNIX域套接字替代目录,(默认为 /tmp)
--fork  # 以守护进程的方式运行MongoDB，创建服务器进程
--auth  # 启用验证
--cpu   # 定期显示CPU的CPU利用率和iowait
--dbpath arg    # 指定数据库路径
--diaglog arg   # diaglog选项 0=off 1=W 2=R 3=both 7=W+some reads
--directoryperdb    # 设置每个数据库将被保存在一个单独的目录
--journal   # 启用日志选项，MongoDB的数据操作将会写入到journal文件夹的文件里
--journalOptions arg    # 启用日志诊断选项
--ipv6  # 启用IPv6选项
--jsonp # 允许JSONP形式通过HTTP访问（有安全影响）
--maxConns arg  # 最大同时连接数 默认2000
--noauth    # 不启用验证
--nohttpinterface   # 关闭http接口，默认关闭27018端口访问
--noprealloc    # 禁用数据文件预分配(往往影响性能)
--noscripting   # 禁用脚本引擎
--notablescan   # 不允许表扫描
--nounixsocket  # 禁用Unix套接字监听
--nssize arg (=16)  # 设置信数据库.ns文件大小(MB)
--objcheck  # 在收到客户数据,检查的有效性，
--profile arg   # 档案参数 0=off 1=slow, 2=all
--quota # 限制每个数据库的文件数，设置默认为8
--quotaFiles arg    # number of files allower per db, requires --quota
--repair    # 修复所有数据库run repair on all dbs
--repairpath arg    # 修复库生成的文件的目录,默认为目录名称dbpath
--slowms arg (=100) # value of slow for profile and console log
--smallfiles    # 使用较小的默认文件
--syncdelay arg (=60)   # 数据写入磁盘的时间秒数(0=never,不推荐)
--sysinfo   # 打印一些诊断系统信息
--upgrade   # 如果需要升级数据库
---
 * Replicaton 参数
--fastsync  # 从一个dbpath里启用从库复制服务，该dbpath的数据库是主库的快照，可用于快速启用同步
--autoresync    # 如果从库与主库同步数据差得多，自动重新同步，
--oplogSize arg # 设置oplog的大小(MB)
---
 * 主/从参数
--master    # 主库模式
--slave # 从库模式
--source arg    # 从库 端口号
--only arg  # 指定单一的数据库复制
--slavedelay arg    # 设置从库同步主库的延迟时间
---
 * Replica set(副本集)选项：
---
--replSet arg   # 设置副本集名称
---
 * Sharding(分片)选项
--configsvr # 声明这是一个集群的config服务,默认端口27019，默认目录/data/configdb
--shardsvr  # 声明这是一个集群的分片,默认端口27018
--noMoveParanoia    # 关闭偏执为moveChunk数据保存
```

## mongod.conf

```
fork=true
bind_ip=0.0.0.0
port=27017
dbpath=/opt/mongo
logpath=/opt/mongo/log
directoryperdb=true
logappend=true
```

mongod.conf的几个大块
systemLog:        #日志
```
systemLog:
  verbosity: <int>                #日志级别，默认0,1-5均会包含debug信息
  quiet: <boolean>                #安静，true时mongod将会减少日志的输出量
  traceAllExceptions: <boolean>        #打印异常详细信息
  syslogFacility:  <string>                #指定用于登录时信息到syslog Facility水平，前提是启用syslog
  path:  <string>          #日志路径，默认情况下，MongoDB将覆盖现有的日志文件
  logAppend: <boolean>        #mongod重启后，在现有日志后继续添加日志，否则备份当前日志，然后创建新日志
  logRotate: rename|reopen        #日志轮询，防止一个日志文件特别大。rename重命名日志文件，默认值；reopen使用Linuxrotate特性，关闭并重新打开日志文件，前提为logAppend: true
  destination: <string>        #日志输出目的地，可为file或syslog，若不指定，则会输出到 std out
  timeStampFormat: <string>        #指定日志格式的时间戳，有 ctime, Iso869-utc, iso8691-local
  component:            #为不同的组件指定各自的日志信息级别
      accessControl:
          verbosity: <int>
      command:
          verbosity: <int>
```
storage:          #存储
```
storage:
  dbPath: <string>        #mongodb进程存储数据目录，此配置进队此mongod进程有效，你使用配置文件开启的mongod就可以指定额外的数据目录
  indexBuildRetry:  <boolean>        #当构件索引时mongod意外
  关闭，那么在此启动是否重建索引，默认true
  repairPath: <string>        #在repair期间使用此目录存储临时数据，repair结束后此目录下数据将被删除
  journal:        
      enabled: <boolean>        #journal日志持久存储，journal日志用来数据恢复，通常用于故障恢复，建议开启
      commitIntervalMs: <num>        #mongod日志刷新值，范围1-500毫秒，默认100，不建议修改
  directoryPerDB:  <boolean>        #是否将不同的数据存储在不同的目录中，dbPath子目录
  syncPeriodSecs:  <int>        #fsync操作将数据flush到磁盘的时间间隔，默认为60秒，不建议修改
  engine:  <string>        #存储引擎

  mmapv1:    #mmapv1存储引擎，3.2前默认
      preallocDataFiles:  <boolean>
      nsSize: <int>
      quota:
          enforced: <boolean>
          maxFilesPerDB: <int>
      smallFiles: <boolean>
      journal:
          debugFlags: <int>
          commitIntervalMs: <num>
  wiredTiger:    #WiredTiger存储引擎，3.2后默认
      engineConfig:
          cacheSizeGB: <number>    #最大缓存大小
          journalCompressor: <string>    #日志压缩算法，可选值有 none，snappy(默认)，zlib
          directoryForIndexes: <boolean>    #是否将索引和collections数据分别存储在dbPath单独的目录中
      collectionConfig:
          blockCompressor: <string>    #collection数据压缩算法，可选none, snappy，zlib
      indexConfig:
          prefixCompression: <boolean>    #是否对索引数据使用前缀压缩。对那些经过排序的值存储有很大帮助，可有效减少索引数据的内存使用量。
  inMemory:    #inMemory内存存储引擎，bate版
      engineConfig:
          inMemorySizeGB: <number>
```
processManagement:        #进程管理
```
processManagement:
  fork: <boolean>        #是否以fork模式运行mongod进程，默认情况下，mongod不作为守护进程运行
  pidFilePath: <string>        #将mongod进程ID写入指定文件，如未指定，将不会创建PID文件
```
net:        #网络
```
net:
  prot: <int>    #监听端口，默认27017
  bindIp: <string>    #绑定IP，如果此值是“0.0.0.0”则绑定所有接口
  maxIncomingConnections: <int>    #mongod进程允许的最大连接数，如果此值超过系统配置的连接数阈值，将不会生效(ulimit)
  wireObjectCheck: <boolean>    #当客户端写入数据时，检查数据的有效性（BSON）。如果数据格式不良，update,insert等操作将会被拒绝
  ipv6: <boolean>    #是否支持多实例之间使用ipv6
  unixDomainSocker:    #适用于Unix系统
      enabled: <boolean>   
      pathPrefix: <string>
      filePermissions: <int>
  http:    #
      enabled: <boolean>
      JSONEnabled: <boolean>
      RESTInterfaceEnabled: <boolean>
  ssl:
      sslOnNormalPorts: <boolean>
      mode: <string>
      PEMKeyFile: <string>
      PEMKeyPassword: <string>
      clusterFile: <string>
      clusterPassword: <string>
      CAFile: <string>
      CRLFile: <string>
      allowConnectionsWithoutCertificates: <boolean>
      allowInvalidCertificates: <boolean>
      allowInvalidHostnames: <boolean>
      disabledProtocols: <string>
      FIPSMode: <boolean>
  compression:
      compressors: <string>  
```
security:        #安全
```
security:
  authorization: enabled    #MondoDB认证功能
  keyFile: /path/mongo.key    #MongoDB副本集节点身份验证密钥文件
  clusterAuthMode: <string>    #集群members间的认证模式
  transitionToAuth: <boolean>
   javascriptEnabled:  <boolean>    #是否允许执行JavaScript脚本
   redactClientLogData: <boolean>
   sasl:
      hostName: <string>
      serviceName: <string>
      saslauthdSocketPath: <string>
   enableEncryption: <boolean>
   encryptionCipherMode: <string>
   encryptionKeyFile: <string>
   kmip:
      keyIdentifier: <string>
      rotateMasterKey: <boolean>
      serverName: <string>
      port: <string>
      clientCertificateFile: <string>
      clientCertificatePassword: <string>
      serverCAFile: <string>
   ldap:
      servers: <string>
      bind:
         method: <string>
         saslMechanism: <string>
         queryUser: <string>
         queryPassword: <string>
         useOSDefaults: <boolean>
      transportSecurity: <string>
      timeoutMS: <int>
      userToDNMapping: <string>
      authz:
         queryTemplate: <string>
```
operationProfiling:        #性能分析器
```
operationProfiling:
  slowOpThresholdMs: <int>    #数据库profiler判定一个操作是“慢查询”的时间阈值，单位毫秒。mongod会把慢查询记录到日志中，默认100ms
  mode: <string>    #数据库profiler级别，操作的性能信息将会被写入日志文件中，可选值“off”--关闭profiling，“slowOp”--只包包含慢操作，“all”--记录所有操作
  #数据库profiling会影响性能，建议只在性能调试阶段开启
```
replication:        #主从复制
```
replication:
  oplogSizeMB: <int>    #replication操作日志的最大尺寸，如果太小，secondary将不能通过oplog来同步数据，只能全量同步
  replSetName: <string>    #副本集名称，副本集中所有的mongod实例都必须有相同的名字，Sharding分布式下，不同的sharding应该使用不同的repSetName
  secondaryIndexPrefetch: <string>    #副本集中的secondary，从oplog中应用变更操作之前，将会先把索引加载到内存
  enalbeMajorityReadConcern: <boolean>    #允许readConcern的级别为“majority”
```
sharding:        #分片相关参数
```
sharding:
  clusterRole: <string>    #在sharding集群中，此mongod实例可选的角色。configsvr,默认监听27019端口 和 shardsvr,默认监听27018端口
  archiveMovedChunks: <boolean>    #当chunks因为“负载均衡”而迁移到其他节点时，mongod是否将这些chunks归档，并保存在dbPath/movechunk目录下，mongod不会删除moveChunk下的文件
```
setParameter:        #自定义变量
```
setParameter:
  <parameter1>: <value1>
  <parameter2>: <value2>
  enableLocalhostAuthBypass: false    #栗子
```
auditLog:        #审计相关参数
```
auditLog:
  destination: <string>    #指定审计记录的输出方式，有syslog, console, file
  format: <string>    #输出格式，有JSON 和 BSON
  path: <string>    #如果审计时间输入为文件，那么就需要指定文件完整路径及文件名
  filter: <string>    #过滤器，可限制审计系统记录的操作类型，该选项需要一个表单的查询文档的字符串表示形式
```
snmp:        #
```

```