# 应该是最快的裤子导入方案了

### 重要

请使用 MySQL8.0 以下，否则 MyISAM 引擎不支持分区，如果用 Innodb 引擎，导入速度和建索引速度将大打折扣，我使用的是 MySQL5.7。

### 下载 mysql shell 工具

mysql shell 下载地址: https://downloads.mysql.com/archives/shell/ 

### 修改配置文件

my.ini 配置文件添加以下两行

```ini
[mysqld]
secure_file_priv=''
local_infile=on
```

同时如果原先已有以下两行，请先注释掉(关闭二进制日志，加快导入速度)

```ini
#binlog_format=mixed
#log-bin=binlog
```

### 建立数据库和表 

```sql
#建立数据库
CREATE DATABASE kuzi;

use kuzi;

#建 weibos 表，分 16 个区
CREATE TABLE `weibos`  (
`phone` BIGINT(20) DEFAULT NULL,`uid` BIGINT(20) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=ASCII PARTITION BY HASH(uid) PARTITIONS 16;

#建 qq 表，分 26 个区
CREATE TABLE `qqs`  (
`qq` BIGINT(20) DEFAULT NULL,`phone` BIGINT(20) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=ASCII PARTITION BY HASH(qq) PARTITIONS 26;
```

### 运行 mysqlsh 导入数据

cmd 命令行进入 mysql shell 的 bin 目录运行  mysqlsh.exe，接着输入如下命令

```
#连接数据库，会提示输入密码
\c root@127.0.0.1:3306

#选择刚刚建的表
\use kuzi

#导入微博数据，把 D:/0/weibo_2019_5e.txt 修改成你的文件路径
util.importTable("D:/0/weibo_2019_5e.txt",{"schema":"kuzi","table":"weibos","fieldsTerminatedBy":"\t","showProgress":true,"threads":8,bytesPerChunk: "1G",maxRate: "2G"});

#上面执行完后，再导入 QQ 数据
util.importTable("D:/0/qq_6.9_8e_update.txt",{"schema":"kuzi","table":"qqs","fieldsTerminatedBy":"----","showProgress":true,"threads":8,bytesPerChunk: "1G",maxRate: "2G"});
```

**该方法导入应该是最快的了，我笔记本机械硬盘导入两个数据，微博用了 5 分钟，QQ 用了 7 分钟**

### 建立索引

```
use kuzi;	
#微博的索引	
ALTER TABLE `kuzi`.`weibos`  ADD INDEX uid (uid);	
ALTER TABLE `kuzi`.`weibos`  ADD INDEX phone (phone); #按需建立	
QQ 的索引	
ALTER TABLE `kuzi`.`qqs`  ADD INDEX qq (qq);	
ALTER TABLE `kuzi`.`qqs`  ADD INDEX phone (phone); #按需建立	
```

这一步就比较久了，微博耗时 45 分钟，QQ 耗时 90 分钟

### 查询

用 go 写了个简单的程序，源码已经放 [github](https://github.com/JuchiaLu/kuzi)，懒得编译的可以直接下载编译好的运行，修改程序配置文件后，cmd 下直接执行即可

```
#数据库类型
db_type=mysql

#数据库连接地址和账户密码
db_url=root:root@tcp(127.0.0.1:3306)/kuzi?charset=utf8mb4&parseTime=True&loc=Local

#http监听端口
listen_port=8080
```

执行程序后浏览器打开，也可以不用界面直接访问接口，返回 Json

QQ：http://127.0.0.1:8080/v1/qq?qq=10001

微博：http://127.0.0.1:8080/v1/weibo?uid=10001

![](https://raw.githubusercontent.com/JuchiaLu/kuzi/main/images/readme1.gif)
### 直接用mysql命令导入 
```
use kuzi

#QQ
LOAD DATA INFILE 'd:qq_6.9_8e_update.txt'  IGNORE INTO TABLE `qqs` FIELDS terminated by '----' lines terminated by '\n'(qq,phone);

#WB
LOAD DATA INFILE 'd:weibo_2019_5e.txt' IGNORE INTO TABLE `weibos`  FIELDS TERMINATED BY '\t' LINES TERMINATED BY '\n'(phone,uid);
```
