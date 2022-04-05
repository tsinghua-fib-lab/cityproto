# 迁移指南

从 v1 版本迁移到 v2 版本的改动

## v2 版本格式变动

完全支持通过 Golang struct tag 机制实现 `bson` 到 `pb` 的快速转换，并删去不能在 `bson` 中体现的语义：

- 所有的 `oneof` 由 optional 替代，`oneof` 的功能依靠约定实现。
  - 原因：struct tag 无法正确处理`oneof`，`bson`无法体现`oneof`语义，`oneof`语义实际作用价值不大，目前主要体现在`map_position`与`departure`中。
- 所有的`float`替换为`double`。
  - 原因：`bson`中没有`float32`，`python`没有`float32`，`C++`在 64 位 CPU 上采用`float32`与`float64`计算速度相同。
- 所有的`uint32`替换为`int32`。
  - 原因：`bson`中没有`uint`，`python`没有`uint`，同时简化`C++`与`go`代码编写（统一用`int`）。
- 枚举在`bson`中存储由字符串替换为整数。
  - struct tag 无法正确处理字符串形式的枚举。
- 丰富`bson`中的 external 数据项，解决关联到原始数据的问题。

## v2 版本内容变动

- 删去 POI，由 AOI 替代 POI，AOI 包含多个路网连接点。
  - 原因：POI 一般指代“点”，如商场内的某个饭店、小区的大门等，不能很好地反映出人员进出区域（如商场、小区）的情况，因此用 AOI 替代 POI 这一概念，原有的无法归类到 AOI 的 POI（如公交站等）则直接提升为 AOI。
- 添加一些聚合性质的 message 作为 mongodb collection 的语义表示，并便于开发 Golang 的数据读取代理。
- 添加同步器`sync.v1`

## 需要后续确认的变动

- 结果输出方式
- 可视化相关
- 为 v1 添加 deprecated 标记`option deprecated = true;`
- 名称迁移到“水镜 mirage”
