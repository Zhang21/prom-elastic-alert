## metrics介绍

prometheus metrics 默认的路径是 `http://prom-elastic-alert:9003/metrics`，默认的指标前缀是 `prom_elastic_alert_`。

指标介绍如下：

| 指标 | 描述 |
| - | - |
| info | 应用信息 |
| link_redis | 连接 redis 的健康状态 |
| op_redis | 操作 redis 命令的次数 |
| query | 规则信息 |
| rule | 规则启用状态 |
| rule_hit | 规则命中状态 |
| webhook_notify | webhook 通知的次数 |

<br/>

### prom告警规则

如果直接使用 metrics 接入 promethues，那么可通过 promql 进行告警，而不通过应用到 alertmanager，这样就可以和告警系统复用一套 prom。

建议日志规则的 `unique_id` 配置为可识别的内容，这样可以在告警通知信息中通过这个 `label` 快速识别告警。

看后续是否需要把日志规则触发的字段（如 message）和内容填充到 `hit` 指标中，便于告警通知信息展示。

```yaml
# 日志关键字告警规则触发
prom_elastic_alert_rule_hit > 0

```
