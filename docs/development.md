# Development

There are key concerns when developing the protocol buffer library:

## Style Guide

We use [buf](https://github.com/bufbuild/buf) to lint our proto files.

To avoid ambiguity, all cross-file references that are not in the same namespace must be written with the full package name.

Within the same namespace, do not include the package name (mainly for `xxx.proto` and `xxx_service.proto`).

References:
1. [Google Protocol Buffers Style Guide](https://developers.google.com/protocol-buffers/docs/style)
2. [Uber Prototool Style Guide](https://github.com/uber/prototool/blob/dev/style/README.md)

For enums, we prefer the Uber style:
1. Define enums independently, not embedded in messages.
2. Prefix enum item names with the enum name.

## Conventions

- All floating-point numbers use double precision first.
- Do not include any `option` in the `.proto` file. Use [buf managed mode](https://docs.buf.build/tour/use-managed-mode) to manage the options for the target language.
- Use aggregate messages to represent the semantics of MongoDB collections and facilitate the development of Golang data reading agents.

## BSON Compatibility

We fully support fast conversion from `bson` to `pb` using Golang struct tags, and remove semantics that cannot be reflected in `bson`:
- All `oneof` are replaced by `optional`, and the functionality of `oneof` is implemented by convention.
  - Reason: struct tags cannot correctly handle `oneof`, `bson` cannot reflect the semantics of `oneof`, the actual value of `oneof` semantics is not great, and it is mainly reflected in `map_position` and `departure`.
- All `float` are replaced by `double`.
  - Reason: `bson` does not have `float32`, `python` does not have `float32`, and `C++` uses `float32` and `float64` at the same speed on 64-bit CPUs.
- All `uint32` are replaced by `int32`.
  - Reason: `bson` does not have `uint`, `python` does not have `uint`, and `C++` and `go` code writing is simplified (unified with `int32`).
- Enums are stored as integers in `bson`.
  - struct tags cannot correctly handle string-form enums.
- Enrich the external data items in `bson` to solve the problem of associating with the original data.

## How to Publish

To publish a new version of the protocol buffer library, follow these steps:
```bash
git tag v${major}.${minor}.${patch}  # For example, v0.1.0, set the version number
git push --tags  # Push the version number
```
