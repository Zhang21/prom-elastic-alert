prom-elastic-alert是一个基于查询Elasticsearch的告警组件. 

如有帮助给个Star⭐鼓励一下~️

<br/>

## 简介

本项目[prom-elastic-alert](https://github.com/dream-mo/prom-elastic-alert)主要是解决Elastic栈中, ELK告警市面上没有更多的日志告警开源组件可供选择.虽然之前有使用过[Elastalert](https://github.com/Yelp/elastalert)项目,
但是该项目已经不维护,并且我们在实际使用的过程中遇到了一些问题:

- 1.组件使用Python编写,性能较差有时候造成告警延迟
- 2.告警收敛、告警聚合、收敛等功能较弱
- 3.组件运行数据不能对接Prometheus监控体系

本项目灵感来自于[Elastalert](https://github.com/Yelp/elastalert)

<br/>

## 特性及优点

- 1.使用Golang编写,跨平台、体积小、性能有足够的优势
- 2.自身不实现告警聚合、收敛、分组等,这是alertmanager的优势所在,没必要自己再造轮子.引入[PrometheusAlert](https://github.com/feiyu563/PrometheusAlert)实现多类型告警
- 3.内置exporter,可以接入Prometheus监控体系,查看当前组件运行状态和规则告警状态等。
- 4.支持Elasticsearch7、Elasticsearch8
- 5.提供现成的Grafana面板json文件

<br/>

## 架构图

![架构图](docs/img/architecture.png)

<br/>

## 告警样例

### 钉钉通知

![钉钉告警图](docs/img/alert.png)

### 告警详情

![告警详情图](docs/img/detail.png)

### Grafana面板

![Grafana面板图](docs/img/grafana.png)

<br/>

## 使用文档

详细文档:  [使用文档](docs/document.md)

<br/>

## metrics介绍

prom metrics介绍: [metrics 介绍](docs/metrics.md)

<br/>

## FAQ

FAQ: [FAQ](docs/faq.md)

## License

prom-elastic-alert is under the Apache 2.0 license. See the [LICENSE](LICENSE) file for details.
