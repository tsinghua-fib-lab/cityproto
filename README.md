# Protos

proto 文件仓库。

## 内容

- `wolong.agent`: 智能体
- `wolong.comm`: 通信
- `wolong.config`: 配置文件基本组件
- `wolong.event`: 城市事件
- `wolong.elec`: 电网
- `wolong.geo`: 坐标体系
- `wolong.map`: 地图
- `wolong.routing`: 路径规划
- `wolong.sim`: 模拟器
- `wolong.sync`: 全局同步器
- `wolong.thing`: 物体
- `wolong.traffic_light`: 信控
- `wolong.trip`: 出行与日程

## STYLE

以[buf](https://github.com/bufbuild/buf)作为 linter，按 buf 的要求作为标准。

另外，为避免二义性，所有不在完全相同命名空间下的跨文件引用均需要写出完成包名。

在相同命名空间下的，不要引入包名（主要是针对`xxx.proto`与`xxx_service.proto`）。

参考：

1. https://developers.google.com/protocol-buffers/docs/style
2. https://github.com/uber/prototool/blob/dev/style/README.md

关于 enum: prefer uber 的 style

1. 独立定义 enum，不嵌入 message
2. enum 的 item 名称加 enum name 作为前缀

## 约定

- 所有浮点数提供 double 精度。
- .proto 中不包含任何`option`，需要指定`option`的目标语言采用[buf managed mode](https://docs.buf.build/tour/use-managed-mode)进行管理。
