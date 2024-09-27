// @generated by protoc-gen-connect-es v1.5.0
// @generated from file city/sync/v2/sync_service.proto (package city.sync.v2, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { EnterStepSyncRequest, EnterStepSyncResponse, ExitStepSyncRequest, ExitStepSyncResponse, GetURLRequest, GetURLResponse, SetURLRequest, SetURLResponse } from "./sync_service_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service city.sync.v2.SyncService
 */
export const SyncService = {
  typeName: "city.sync.v2.SyncService",
  methods: {
    /**
     * 注册程序URL
     *
     * @generated from rpc city.sync.v2.SyncService.SetURL
     */
    setURL: {
      name: "SetURL",
      I: SetURLRequest,
      O: SetURLResponse,
      kind: MethodKind.Unary,
    },
    /**
     * 获取程序URL
     *
     * @generated from rpc city.sync.v2.SyncService.GetURL
     */
    getURL: {
      name: "GetURL",
      I: GetURLRequest,
      O: GetURLResponse,
      kind: MethodKind.Unary,
    },
    /**
     * 程序完成本步所有操作，进入同步状态。
     * 要求：进入同步状态的程序不再向其他程序发送消息，直到下一步开始。
     *
     * @generated from rpc city.sync.v2.SyncService.EnterStepSync
     */
    enterStepSync: {
      name: "EnterStepSync",
      I: EnterStepSyncRequest,
      O: EnterStepSyncResponse,
      kind: MethodKind.Unary,
    },
    /**
     * 程序完成同步阶段（无通信的安全区域）中必要的处理，如为prepare阶段加锁，可以进入准备阶段（恢复通信）。
     *
     * @generated from rpc city.sync.v2.SyncService.ExitStepSync
     */
    exitStepSync: {
      name: "ExitStepSync",
      I: ExitStepSyncRequest,
      O: ExitStepSyncResponse,
      kind: MethodKind.Unary,
    },
  }
};

