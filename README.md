# Protos

proto文件仓库。

## 内容

- `wolong.agent`: 通用，智能体
- `wolong.geo`: 通用，坐标体系
- `wolong.map`: 通用，地图
- `wolong.routing`: 技术，路径规划
- `wolong.sim`: 技术，模拟器
- `wolong.trip`: 通用，出行与日程

## STYLE

以[buf](https://github.com/bufbuild/buf)作为linter，按buf的要求作为标准。

参考：
1. https://developers.google.com/protocol-buffers/docs/style
2. https://github.com/uber/prototool/blob/dev/style/README.md

关于enum: prefer uber的style
1. 独立定义enum，不嵌入message
2. enum的item名称加enum name作为前缀

## 约定

- 只有涉及到空间位置表示的数据（xy坐标、经纬度坐标）提供double精度。
- .proto中不包含任何`option`，需要指定`option`的目标语言采用[buf managed mode](https://docs.buf.build/tour/use-managed-mode)进行管理。

## 待解决的问题

- 代码内的TODO项
