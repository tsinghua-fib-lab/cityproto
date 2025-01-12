# cityproto

CityProto is a protobuf-based data exchange format for FIBLAB projects.
It is designed to be used in Golang, Python, and JavaScript/TypeScript projects.

## Usage

### Golang

```bash
go mod edit -replace github.com/tsinghua-fib-lab/cityproto/v2=git.fiblab.net/sim/protos/v2@v${major}.${minor}.${patch}
# write code to import "github.com/tsinghua-fib-lab/cityproto/v2/go/city/xxx"
go mod tidy
```

### Python

```bash
pip install pycityproto
```

If you encounter the following error during the installation of `grpcio`:

```
pip install fails with "No such file or directory: 'c++': 'c++'"
```

It means that C++ related dependencies are missing. On the Debian image, execute:

```shell
sudo apt install build-essential
```

On the Alpine image, execute:

```shell
apk add g++
```

### JavaScript/TypeScript

#### NPM

Use the `npm install` command to install the SDK:

```bash
npm install @fiblab/cityproto
```

#### Yarn

Use the `yarn add` command to install the SDK:

```bash
yarn add @fiblab/cityproto
```

## Table of Contents

```{toctree}
:maxdepth: 2

development
docs
```
