# 分布式BOT

## 由来
就是写， 这是我写的最恶心心的项目了

agent安装地方有网络限制 没法入网

导致 backend 和 agnet的通讯全部都是异步的

但是 web管理端和backend的通讯是同步的

这个转换恶心死我了BUS总线

![](./s1.png)

协议实现
- KCP
- DNS
- ICMP 
