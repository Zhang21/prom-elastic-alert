## FAQ

- Q: 是否支持 ES v7 和 v8 两个版本？<br>
A: 支持。go elasticsearch/v7 适配 v7 和 v8 版本，规则文件中统一使用 `v7` 即可。

<br/>

- Q: 是否支持 es https 连接？<br>
A: 支持。

<br/>

- Q: 是否支持一个 ES 集群的多个地址？<br>
A: 支持。

<br/>

- Q: 规则文件是否支持同时连接多个 ES 集群？<br>
A: 不支持。无法同时建立多个集群 client 连接。

<br/>

- Q: 应用依赖有哪些？<br>
A: redis（必须），alertmanager（可选）。

<br/>

- Q: redis 中 key 的过期时间？<br>
A: 默认 1天，配置中可修改。
