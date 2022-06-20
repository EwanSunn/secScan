# secScan  
An interactive penetration test scanning tool  
  
## 功能特性  
### 已完成
- [+] 2022/06/08    完成端口扫描与常用服务弱口令爆破功能
- [+] 2022/06/20    完成logger日志记录功能，改进端口扫描的保存结果，结果可供弱口令爆破功能直接读取
### TODO
- 将主机存活探测功能单独抽取出来，与端口扫描等一样，形成一个新的模块
- 添加`gobuster`模块到secScan中
- 对存活端口的服务与`banner`信息进行探测
- 代码结构的优化
  - `task`任务分发的时候能将各个模块的`RunTask`模块化出来，根据传入参数的不同来调用不同的模块任务


  
## 使用指南  
### 端口扫描
`portScan -i 127.0.0.1(or 127.0.0.1/24) -p 22,23-24 -c 1000`

### 弱口令爆破
Usage:
1. `crack -f ip.txt -u ./dict/user.dic -p ./dict/pass.dic [-t 5] [-c 1000]`
2. `crack`  
     (default file: ./dict/ip.txt, default user list: ./dict/user.dic, default pass list: ./dict/pass.dic )

Support Protocol:
- FTP(ip:21)
- SSH(ip:22)
- Mysql(ip:3306)
- Mssql(ip:1433)
- MongoDB(ip:27017)
- Redis(ip:6307)
- Postgres(ip:5432)

ip.txt format:
```
127.0.0.1:22
127.0.0.1:3306
....
127.0.0.1:6379
```


  
## 参考  
- 《白帽子安全开发实战》& 配套代码：https://github.com/netxfly/sec-dev-in-action-src
- https://github.com/desertbit/grumble
- 极客时间专栏 《Go 语言项目开发实战》
