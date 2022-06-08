# secScan  
An interactive penetration test scanning tool  
  
## 功能特性  
### 已完成
[+] 2022/06/08   完成端口扫描与常用服务弱口令爆破功能
### TODO
- 2022年6月TODO
    - 日志记录logger，可以进行debug，跟踪错误
    - 联动端口扫描与弱口令爆破。端口扫描保存的结果应该匹配弱口令爆破的协议格式
    - 优化代码结构

  
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
