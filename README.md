# k8s-api-in-action
Kubernetes API 是基于HTTP协议的restful编程接口，API支持POST、PUT、PATCH、DELETE 和 GET等方法对资源进行增、删、改、查等操作。 Kubernetes API 还通过 "watch" 机制支持高效的资源变更通知， 从而允许其他组件对资源的状态进行高效的缓存和同步。


在使用 Kubernetes REST API 编写应用程序时，我们不需要自己实现http rest api接口的请求及传输数据的编解码，可以使用Kubernetes官方提供的客户端库来实现。  本系列就是使用kubernetes官方开源的客户端client-go调用Api Server的接口来部署微服务。


本系列分为以下几篇文章：

一，[kubernetes原生api部署微服务-构建微服务](https://mp.weixin.qq.com/s?__biz=Mzg5MjA1ODYzNg==&mid=2247484399&idx=1&sn=70e07602e526f053409f75bd4d53e62a&chksm=cfc2ae8cf8b5279adb28f20635d60418e054b968966f450ed358641b111571257f2079a2c53e&token=1595797558&lang=zh_CN#rd)  

使用gin框架搭建简单的博客系统，包含2个简单的微服务。

二，kubernetes原生api部署微服务-创建Namespace、Deployment、Service

创建Namespace，Deplement，service，测试扩缩容及滚动升级

三，kubernetes原生api部署微服务-创建Configmap

创建configmap来存储微服务的配置信息

四，kubernetes原生api部署微服务-创建Ingress

创建ingress，在kubernetes集群外部访问微服务。

五，kubernetes原生api部署微服务-监听Pod

创建informer来监听pod的创建、修改及删除。


### 代码详解，请关注微信公众号：coding到灯火阑珊

![Image](https://github.com/Justin02180218/distribute-election-bully/blob/master/qrcode_for_gh_8a5b7b90c100_258.jpg)
