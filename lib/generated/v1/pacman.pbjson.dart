///
//  Generated code. Do not modify.
//  source: v1/pacman.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use addRequestDescriptor instead')
const AddRequest$json = const {
  '1': 'AddRequest',
  '2': const [
    const {'1': 'packages', '3': 1, '4': 3, '5': 9, '10': 'packages'},
    const {'1': 'token', '3': 2, '4': 1, '5': 9, '10': 'token'},
  ],
};

/// Descriptor for `AddRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List addRequestDescriptor = $convert.base64Decode('CgpBZGRSZXF1ZXN0EhoKCHBhY2thZ2VzGAEgAygJUghwYWNrYWdlcxIUCgV0b2tlbhgCIAEoCVIFdG9rZW4=');
@$core.Deprecated('Use addResponseDescriptor instead')
const AddResponse$json = const {
  '1': 'AddResponse',
};

/// Descriptor for `AddResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List addResponseDescriptor = $convert.base64Decode('CgtBZGRSZXNwb25zZQ==');
@$core.Deprecated('Use searchRequestDescriptor instead')
const SearchRequest$json = const {
  '1': 'SearchRequest',
  '2': const [
    const {'1': 'pattern', '3': 1, '4': 1, '5': 9, '10': 'pattern'},
  ],
};

/// Descriptor for `SearchRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List searchRequestDescriptor = $convert.base64Decode('Cg1TZWFyY2hSZXF1ZXN0EhgKB3BhdHRlcm4YASABKAlSB3BhdHRlcm4=');
@$core.Deprecated('Use searchResponseDescriptor instead')
const SearchResponse$json = const {
  '1': 'SearchResponse',
  '2': const [
    const {'1': 'packages', '3': 1, '4': 3, '5': 9, '10': 'packages'},
  ],
};

/// Descriptor for `SearchResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List searchResponseDescriptor = $convert.base64Decode('Cg5TZWFyY2hSZXNwb25zZRIaCghwYWNrYWdlcxgBIAMoCVIIcGFja2FnZXM=');
@$core.Deprecated('Use updateRequestDescriptor instead')
const UpdateRequest$json = const {
  '1': 'UpdateRequest',
  '2': const [
    const {'1': 'token', '3': 1, '4': 1, '5': 9, '10': 'token'},
  ],
};

/// Descriptor for `UpdateRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List updateRequestDescriptor = $convert.base64Decode('Cg1VcGRhdGVSZXF1ZXN0EhQKBXRva2VuGAEgASgJUgV0b2tlbg==');
@$core.Deprecated('Use updateResponseDescriptor instead')
const UpdateResponse$json = const {
  '1': 'UpdateResponse',
};

/// Descriptor for `UpdateResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List updateResponseDescriptor = $convert.base64Decode('Cg5VcGRhdGVSZXNwb25zZQ==');
@$core.Deprecated('Use describeRequestDescriptor instead')
const DescribeRequest$json = const {
  '1': 'DescribeRequest',
  '2': const [
    const {'1': 'package', '3': 1, '4': 1, '5': 9, '10': 'package'},
  ],
};

/// Descriptor for `DescribeRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List describeRequestDescriptor = $convert.base64Decode('Cg9EZXNjcmliZVJlcXVlc3QSGAoHcGFja2FnZRgBIAEoCVIHcGFja2FnZQ==');
@$core.Deprecated('Use describeResponseDescriptor instead')
const DescribeResponse$json = const {
  '1': 'DescribeResponse',
  '2': const [
    const {'1': 'fields', '3': 1, '4': 3, '5': 11, '6': '.proto.v1.DescribeResponse.FieldsEntry', '10': 'fields'},
  ],
  '3': const [DescribeResponse_FieldsEntry$json],
};

@$core.Deprecated('Use describeResponseDescriptor instead')
const DescribeResponse_FieldsEntry$json = const {
  '1': 'FieldsEntry',
  '2': const [
    const {'1': 'key', '3': 1, '4': 1, '5': 9, '10': 'key'},
    const {'1': 'value', '3': 2, '4': 1, '5': 9, '10': 'value'},
  ],
  '7': const {'7': true},
};

/// Descriptor for `DescribeResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List describeResponseDescriptor = $convert.base64Decode('ChBEZXNjcmliZVJlc3BvbnNlEj4KBmZpZWxkcxgBIAMoCzImLnByb3RvLnYxLkRlc2NyaWJlUmVzcG9uc2UuRmllbGRzRW50cnlSBmZpZWxkcxo5CgtGaWVsZHNFbnRyeRIQCgNrZXkYASABKAlSA2tleRIUCgV2YWx1ZRgCIAEoCVIFdmFsdWU6AjgB');
@$core.Deprecated('Use statsRequestDescriptor instead')
const StatsRequest$json = const {
  '1': 'StatsRequest',
};

/// Descriptor for `StatsRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List statsRequestDescriptor = $convert.base64Decode('CgxTdGF0c1JlcXVlc3Q=');
@$core.Deprecated('Use statsResponseDescriptor instead')
const StatsResponse$json = const {
  '1': 'StatsResponse',
  '2': const [
    const {'1': 'packages_count', '3': 1, '4': 1, '5': 5, '10': 'packagesCount'},
    const {'1': 'outdated_count', '3': 2, '4': 1, '5': 5, '10': 'outdatedCount'},
    const {'1': 'outdated_packages', '3': 3, '4': 3, '5': 11, '6': '.proto.v1.OutdatedPackage', '10': 'outdatedPackages'},
  ],
};

/// Descriptor for `StatsResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List statsResponseDescriptor = $convert.base64Decode('Cg1TdGF0c1Jlc3BvbnNlEiUKDnBhY2thZ2VzX2NvdW50GAEgASgFUg1wYWNrYWdlc0NvdW50EiUKDm91dGRhdGVkX2NvdW50GAIgASgFUg1vdXRkYXRlZENvdW50EkYKEW91dGRhdGVkX3BhY2thZ2VzGAMgAygLMhkucHJvdG8udjEuT3V0ZGF0ZWRQYWNrYWdlUhBvdXRkYXRlZFBhY2thZ2Vz');
@$core.Deprecated('Use outdatedPackageDescriptor instead')
const OutdatedPackage$json = const {
  '1': 'OutdatedPackage',
  '2': const [
    const {'1': 'name', '3': 1, '4': 1, '5': 9, '10': 'name'},
    const {'1': 'current_version', '3': 2, '4': 1, '5': 9, '10': 'currentVersion'},
    const {'1': 'latest_version', '3': 3, '4': 1, '5': 9, '10': 'latestVersion'},
  ],
};

/// Descriptor for `OutdatedPackage`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List outdatedPackageDescriptor = $convert.base64Decode('Cg9PdXRkYXRlZFBhY2thZ2USEgoEbmFtZRgBIAEoCVIEbmFtZRInCg9jdXJyZW50X3ZlcnNpb24YAiABKAlSDmN1cnJlbnRWZXJzaW9uEiUKDmxhdGVzdF92ZXJzaW9uGAMgASgJUg1sYXRlc3RWZXJzaW9u');
@$core.Deprecated('Use loginRequestDescriptor instead')
const LoginRequest$json = const {
  '1': 'LoginRequest',
  '2': const [
    const {'1': 'login', '3': 1, '4': 1, '5': 9, '10': 'login'},
    const {'1': 'password', '3': 2, '4': 1, '5': 9, '10': 'password'},
  ],
};

/// Descriptor for `LoginRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List loginRequestDescriptor = $convert.base64Decode('CgxMb2dpblJlcXVlc3QSFAoFbG9naW4YASABKAlSBWxvZ2luEhoKCHBhc3N3b3JkGAIgASgJUghwYXNzd29yZA==');
@$core.Deprecated('Use loginResponseDescriptor instead')
const LoginResponse$json = const {
  '1': 'LoginResponse',
  '2': const [
    const {'1': 'token', '3': 1, '4': 1, '5': 9, '10': 'token'},
  ],
};

/// Descriptor for `LoginResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List loginResponseDescriptor = $convert.base64Decode('Cg1Mb2dpblJlc3BvbnNlEhQKBXRva2VuGAEgASgJUgV0b2tlbg==');
