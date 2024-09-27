// @generated by protoc-gen-connect-es v1.5.0
// @generated from file city/event/v1/event_service.proto (package city.event.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { PublishRequest, PublishResponse, PullRequest, PullResponse } from "./event_service_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service city.event.v1.EventService
 */
export declare const EventService: {
  readonly typeName: "city.event.v1.EventService",
  readonly methods: {
    /**
     * 发布事件
     *
     * @generated from rpc city.event.v1.EventService.Publish
     */
    readonly publish: {
      readonly name: "Publish",
      readonly I: typeof PublishRequest,
      readonly O: typeof PublishResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * 从事件中心拉取事件
     *
     * @generated from rpc city.event.v1.EventService.Pull
     */
    readonly pull: {
      readonly name: "Pull",
      readonly I: typeof PullRequest,
      readonly O: typeof PullResponse,
      readonly kind: MethodKind.Unary,
    },
  }
};

