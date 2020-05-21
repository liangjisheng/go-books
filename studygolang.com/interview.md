# interview

## Go相关

* map、channel、slice的底层实现
* sync.map、sync.pool、sync.Once的原理
* GC的过程、写屏障的含义及作用
* GMP模型，触发Goroutine切换的原因有哪些？for死循环会怎么样？全局goroutine里面存储什么？
* interface的底层实现，怎么判空？
* reflect的使用
* 逃逸分析
* context的使用
* go 性能问题的定位过程（pprof的使用）
* 协程池的使用

## Mysql相关

* 索引的分类
* 为什么选择B+树实现索引？一般深度为多少？b+树和红黑树的区别?
* 聚簇索引和非聚簇索引的区别
* 创建索引后，查询读取I/O的次数
* 索引的最左前缀原则
* mysql数据的索引优化以及失效
* 从学生表中查询每个班的分数的前3名
* mysql的隔离级别？处理什么问题的（脏读、幻读、不可重复读）
* mysql的主从复制过程？
* mysql的大表优化方式

## Redis相关

* redis的数据类型以及日常的应用
* redis的发布/订阅的原理
* zset的底层实现
* 数据缓存过期策略
* redis的部署模式
* redis为什么速度比较快
* reids的大key、热key的处理
* 如何实现分布式锁的
* 持久化策略及其对比
* 缓存雪崩、缓存击穿、缓存穿透

## Etcd相关

* etcd是什么？有什么优势
* raft选主逻辑
* 日志复制
* 脑裂问题
* etcd的watch机制
* etcd如何实现配置下发和服务发现
* etcd对于偶数机器的集群的选主处理
* 选主实现逻辑

## Prometheus相关

* 简介
* 数据存储原理
* 数据类型

## Grpc相关

* 相较于restful的优势
* 数据交互方式
* 限流（通过流模式传输时，发送方数据量过大，会发生什么？）
* protobuf和json的对比
* grpc负载均衡的实现

## Linux相关

* awk
* poll、epoll、select
* I/O模型

## 网络协议相关

* http2的优势
* https的建连过程（7次握手）
* icmp协议的原理
* tcp三次握手、四次挥手
* tcp 拥塞策略
* tcp的time_wait状态和colse_wait状态
* 如何解决tcp的粘包问题
* quic协议是什么
* 如何理解网络模型
* http的状态码含义

## 智力题

* 25匹马，每次只能比赛5组，最快几次找到前3名

* 宝石问题（3个盒子，其中2个宝石，一个石头；先随机选取一个，然后剔除剩余两个中的宝石；第三次选择，选择哪个为宝石的概率大？）

## 系统设计

* RPC的设计
* 架构设计分单系统，每秒3000订单有效期15分钟，50W司机进行抢单操作，如果一直没有抢单，则订单失效
* 字符串hash算法的实现
* 敏感词过滤
* 设计一个高可用的稳定的并发模型处理HTTP请求

## 其他

* 一致性hash算法
* 微服务概述
* 什么是死锁，如何避免
* 限流策略

## 算法

* 如何原地交换两个数
* 岛屿问题
* 数组中重复的数据  
* 1到n乱序排列的数据，少了其中一个，找出这个数
* 二叉树的右视图
* LRU缓存机制  （考虑并发访问）
* 高并发的生产者消费者模式
* 通过中序遍历序列和先序序列恢复二叉树
* 爬楼梯问题
* 单链表逆序
* 单向链表排序
* string1 = 1234dsafaserewr，string2 = 23aefasdfwer，求string3 = string1 + string2
* 二叉树节点的公共祖先
* 二叉树的最大深度
* 二叉树的中序遍历和层次遍历
* 寻找两个升序数组的第K大值
* 最长回文子串长度
* 最短回文串
* 合并两个有序链表  
* 全排列
* 接雨水
* 盛最多水的容器  
* Pow(x, n)  

## 海量数据处理问题（面试官很喜欢问）

* hash
* 字典树
* bitmap
* 布隆过滤器
* MapReduce
* 桶

## article

[index](https://mp.weixin.qq.com/s?__biz=MzI4Njg5MDA5NA==&mid=2247484480&idx=1&sn=757cdf8f07dc9ae9b79fb28b94ecf5d6&chksm=ebd74541dca0cc57f0b3fce8c5038c8512bccbf09e6c363f38a1e4f924d67ac2564b4fd03744&token=620000779&lang=zh_CN&scene=21#wechat_redirect)
[index](https://mp.weixin.qq.com/s?__biz=Mzg2NzA4MTkxNQ==&mid=2247486101&idx=1&sn=980f6dfb7643a9ff4f5a661d4a496046&chksm=ce404141f937c85750232523583435e97f3965a3761fa327e5d79e2b720dfced1a1dfc731d3b&token=1321503479&lang=zh_CN#rd)
[http](https://segmentfault.com/a/1190000013271378)
[AVL](https://mp.weixin.qq.com/s/dYP5-fM22BgM3viWg4V44A)
[RBTree](https://mp.weixin.qq.com/s/p_fEMMNjlnPbbwY9dDQMAQ)
[RBTree](https://mp.weixin.qq.com/s/-8JFh5iLr88XA4AJ9mMf6g)
[graph](https://mp.weixin.qq.com/s/4JEHZWanGtsQHYrZ0MDq7Q)
[topN](https://mp.weixin.qq.com/s?__biz=MzIzMTE1ODkyNQ==&mid=2649410393&idx=1&sn=e2aae1e16baede316922c53256a10c5f&chksm=f0b60ebbc7c187ad4e48b140daff2c80e5007076e99b26427fef50dd7f19ffed5181ba45baed&scene=21#wechat_redirect)
[bloom](https://www.cnblogs.com/CodeBear/p/10911177.html)
[bloom](https://www.jianshu.com/p/2104d11ee0a2)
[bloom](https://www.cnblogs.com/liyulong1982/p/6013002.html)
[bloom](https://github.com/ArashPartow/bloom)
