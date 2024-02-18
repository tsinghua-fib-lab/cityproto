# Protos

proto 文件仓库。

## 内容

- `city.agent`: 智能体
- `city.cognition`: 认知
- `city.comm`: 通信
- `city.config`: 配置文件基本组件
- `city.elec`: 电网
- `city.event`: 城市事件
- `city.geo`: 坐标体系
- `city.map`: 地图
- `city.routing`: 路径规划
- `city.social`: 社交
- `city.streetview`: 街景
- `city.sync`: 全局同步器
- `city.traffic`: 交通
- `city.traffic_light`: 信控
- `city.trip`: 出行与日程
- `city.wargame`: 兵棋
- `city.water`: 水网

## STYLE

以[buf](https://github.com/bufbuild/buf)作为 linter，按 buf 的要求作为标准。

另外，为避免二义性，所有不在完全相同命名空间下的跨文件引用均需要写出完成包名。

在相同命名空间下的，不要引入包名（主要是针对`xxx.proto`与`xxx_service.proto`）。

参考：

1. <https://developers.google.com/protocol-buffers/docs/style>
2. <https://github.com/uber/prototool/blob/dev/style/README.md>

关于 enum: prefer uber 的 style

1. 独立定义 enum，不嵌入 message
2. enum 的 item 名称加 enum name 作为前缀

## 约定

- 所有浮点数采用 double 精度。
- .proto 中不包含任何`option`，需要指定`option`的目标语言采用[buf managed mode](https://docs.buf.build/tour/use-managed-mode)进行管理。
- 采用聚合性质的 message 作为 mongodb collection 的语义表示，并便于开发 Golang 的数据读取代理。

## BSON兼容性

完全支持通过 Golang struct tag 机制实现 `bson` 到 `pb` 的快速转换，并删去不能在 `bson` 中体现的语义：

- 所有的 `oneof` 由 optional 替代，`oneof` 的功能依靠约定实现。
  - 原因：struct tag 无法正确处理`oneof`，`bson`无法体现`oneof`语义，`oneof`语义实际作用价值不大，目前主要体现在`map_position`与`departure`中。
- 所有的`float`替换为`double`。
  - 原因：`bson`中没有`float32`，`python`没有`float32`，`C++`在 64 位 CPU 上采用`float32`与`float64`计算速度相同。
- 所有的`uint32`替换为`int32`。
  - 原因：`bson`中没有`uint`，`python`没有`uint`，同时简化`C++`与`go`代码编写（统一用`int32`）。
- 枚举在`bson`中存储由字符串替换为整数。
  - struct tag 无法正确处理字符串形式的枚举。
- 丰富`bson`中的 external 数据项，解决关联到原始数据的问题。

## 更新与Golang/Python库发布

完成git仓库修改后，执行以下命令：
```bash
git tag v${major}.${minor}.${patch}  # 例如 v0.1.0，设置版本号
git push --tags  # 推送版本号
```

## Golang库使用

```bash
go get -u git.fiblab.net/sim/protos@v${major}.${minor}.${patch}
```

## Python库使用

```shell
pip install pycityproto --extra-index-url https://__token__:glpat-fq6C-NTr45Z_Te4BV4kC@git.fiblab.net/api/v4/projects/26/packages/pypi/simple
```

安装 grpcio 的过程中，如果出现

```
pip install fails with "No such file or directory: 'c++': 'c++'"
```

代表缺少 C++相关依赖。在 Debian 镜像上，执行：

```shell
sudo apt install build-essential
```

在 alpine 镜像上，执行：

```shell
apk add g++
```

### JS/TS

#### NPM

首先，保证项目已经能够正确访问NPM服务
```bash
npm config set @fiblab:registry https://git.fiblab.net/api/v4/projects/26/packages/npm/
npm config set -- '//git.fiblab.net/api/v4/projects/26/packages/npm/:_authToken' "glpat-sdKgi23Ns3KvfpyrREzy"
```

然后，使用`npm install`命令安装SDK：
```bash
npm install @fiblab/cityproto
```

#### Yarn
首先，保证项目已经能够正确访问NPM服务
```bash
yarn config set @fiblab:registry https://git.fiblab.net/api/v4/projects/26/packages/npm/
# 以下未经测试
yarn config set -- '//git.fiblab.net/api/v4/projects/26/packages/npm/:_authToken' "glpat-sdKgi23Ns3KvfpyrREzy"
```

然后，使用`npm install`命令安装SDK：
```bash
yarn add @fiblab/cityproto
```


