///
//  Generated code. Do not modify.
//  source: v1/pacman.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'pacman.pb.dart' as $0;
export 'pacman.pb.dart';

class PacmanServiceClient extends $grpc.Client {
  static final _$add = $grpc.ClientMethod<$0.AddRequest, $0.AddResponse>(
      '/proto.v1.PacmanService/Add',
      ($0.AddRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.AddResponse.fromBuffer(value));
  static final _$search =
      $grpc.ClientMethod<$0.SearchRequest, $0.SearchResponse>(
          '/proto.v1.PacmanService/Search',
          ($0.SearchRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $0.SearchResponse.fromBuffer(value));
  static final _$update =
      $grpc.ClientMethod<$0.UpdateRequest, $0.UpdateResponse>(
          '/proto.v1.PacmanService/Update',
          ($0.UpdateRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $0.UpdateResponse.fromBuffer(value));
  static final _$describe =
      $grpc.ClientMethod<$0.DescribeRequest, $0.DescribeResponse>(
          '/proto.v1.PacmanService/Describe',
          ($0.DescribeRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.DescribeResponse.fromBuffer(value));
  static final _$stats = $grpc.ClientMethod<$0.StatsRequest, $0.StatsResponse>(
      '/proto.v1.PacmanService/Stats',
      ($0.StatsRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.StatsResponse.fromBuffer(value));

  PacmanServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$0.AddResponse> add($0.AddRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$add, request, options: options);
  }

  $grpc.ResponseFuture<$0.SearchResponse> search($0.SearchRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$search, request, options: options);
  }

  $grpc.ResponseFuture<$0.UpdateResponse> update($0.UpdateRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$update, request, options: options);
  }

  $grpc.ResponseFuture<$0.DescribeResponse> describe($0.DescribeRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$describe, request, options: options);
  }

  $grpc.ResponseFuture<$0.StatsResponse> stats($0.StatsRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$stats, request, options: options);
  }
}

abstract class PacmanServiceBase extends $grpc.Service {
  $core.String get $name => 'proto.v1.PacmanService';

  PacmanServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.AddRequest, $0.AddResponse>(
        'Add',
        add_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.AddRequest.fromBuffer(value),
        ($0.AddResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.SearchRequest, $0.SearchResponse>(
        'Search',
        search_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.SearchRequest.fromBuffer(value),
        ($0.SearchResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.UpdateRequest, $0.UpdateResponse>(
        'Update',
        update_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.UpdateRequest.fromBuffer(value),
        ($0.UpdateResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.DescribeRequest, $0.DescribeResponse>(
        'Describe',
        describe_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.DescribeRequest.fromBuffer(value),
        ($0.DescribeResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.StatsRequest, $0.StatsResponse>(
        'Stats',
        stats_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.StatsRequest.fromBuffer(value),
        ($0.StatsResponse value) => value.writeToBuffer()));
  }

  $async.Future<$0.AddResponse> add_Pre(
      $grpc.ServiceCall call, $async.Future<$0.AddRequest> request) async {
    return add(call, await request);
  }

  $async.Future<$0.SearchResponse> search_Pre(
      $grpc.ServiceCall call, $async.Future<$0.SearchRequest> request) async {
    return search(call, await request);
  }

  $async.Future<$0.UpdateResponse> update_Pre(
      $grpc.ServiceCall call, $async.Future<$0.UpdateRequest> request) async {
    return update(call, await request);
  }

  $async.Future<$0.DescribeResponse> describe_Pre(
      $grpc.ServiceCall call, $async.Future<$0.DescribeRequest> request) async {
    return describe(call, await request);
  }

  $async.Future<$0.StatsResponse> stats_Pre(
      $grpc.ServiceCall call, $async.Future<$0.StatsRequest> request) async {
    return stats(call, await request);
  }

  $async.Future<$0.AddResponse> add(
      $grpc.ServiceCall call, $0.AddRequest request);
  $async.Future<$0.SearchResponse> search(
      $grpc.ServiceCall call, $0.SearchRequest request);
  $async.Future<$0.UpdateResponse> update(
      $grpc.ServiceCall call, $0.UpdateRequest request);
  $async.Future<$0.DescribeResponse> describe(
      $grpc.ServiceCall call, $0.DescribeRequest request);
  $async.Future<$0.StatsResponse> stats(
      $grpc.ServiceCall call, $0.StatsRequest request);
}
