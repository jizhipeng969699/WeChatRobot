# Apache

## Mesos 资源管理器

# docker

## swarm 资源管理器

# Google k8s介绍

## 发展历史

- google产品
- 前身borg系统，使用golang重写
- 特点：
    1. 轻量级，使用golang，消耗资源少
    2. 开源
    3. 弹性伸缩
    4. IPVS负载均衡 ， 内部实现

## 前世今生

## k8s框架

## k8s 名词含义

> 高可用集群数量最好是 3 5 7 9 个

- Api Server  : 所有服务访问统一入口
- CrontrollerMagnger ： 维持副本期望数目
- Scheduler ： 负责接受任务，选择合适的节点进行分配任务
- ETCD ： 存储k8s集群所有重要的信息，可持久化
- Kubelet ： 直接跟容器引擎交互实现容器的生命周期管理
- Kube-Proxy ： 负责写入规则至IPTABLES ipvs 实现服务映射访问
- Core-DNS : 可以为集群中的SVC创建一个域名ip 的对应关系解析
- Dashboard ： 给k8s集群提供一个 B/S 结构访问体系
- Ingress Controller ： 官方只能实现四层代理，ingress 可以实现七层代理
- Fedetation ： 提供一个可以跨集群中心多k8s的统一管理功能
- Prometheus ： 提供一个k8s的监控能力
- ELK ： 提供k8s 集群日志统一分析介入平台

# 基础概念：

### 什么是pod

- 同一个pod 会公用一个 pause 的网络栈 ， 同一个pod内容器以本地形式互相访问
- pod 之间 通过 Overlay NetWork 进行通讯
- pod 与service 之间 通过 iptables 规则通讯

#### 自主式pod

#### 控制器管理的pod

### 控制器类型

### k8s网络通讯模型

# k8s集群：

# 资源清单：

### 什么是资源

### 资源清单的语法

### 编写pod

### 掌握pod的生命周期

# pod控制器（掌握各种控制器的特点及使用 定义方式）

# 服务发现： （掌握svc原理及构建方式）

### 服务的分类：

- 有状态服务 DBMS
- 无状态服务 LVS APACHE

# 存储 （掌握多种存储类型的特点，并且能够在不同环境中选择合适的存储方案，有自己的见解）

# 调度器： （掌握调度器原理，能够根据要求吧pod定义到想要的节点运行）

# 安全：（掌握集群的认证，鉴权，访问控制，原理 及其 流程）

# HELM ： 类似linux yum （掌握HELM原理，HELM模板定义，HELM部署一些常用插件）

# 运维：

- 修改kubeadm 的证书可用期限为10年或者更长
- 能够构建高可用的k8s集群