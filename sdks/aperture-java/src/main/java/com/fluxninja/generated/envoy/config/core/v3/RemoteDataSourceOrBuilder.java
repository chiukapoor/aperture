// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: envoy/config/core/v3/base.proto

package com.fluxninja.generated.envoy.config.core.v3;

public interface RemoteDataSourceOrBuilder extends
    // @@protoc_insertion_point(interface_extends:envoy.config.core.v3.RemoteDataSource)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <pre>
   * The HTTP URI to fetch the remote data.
   * </pre>
   *
   * <code>.envoy.config.core.v3.HttpUri http_uri = 1 [json_name = "httpUri", (.validate.rules) = { ... }</code>
   * @return Whether the httpUri field is set.
   */
  boolean hasHttpUri();
  /**
   * <pre>
   * The HTTP URI to fetch the remote data.
   * </pre>
   *
   * <code>.envoy.config.core.v3.HttpUri http_uri = 1 [json_name = "httpUri", (.validate.rules) = { ... }</code>
   * @return The httpUri.
   */
  com.fluxninja.generated.envoy.config.core.v3.HttpUri getHttpUri();
  /**
   * <pre>
   * The HTTP URI to fetch the remote data.
   * </pre>
   *
   * <code>.envoy.config.core.v3.HttpUri http_uri = 1 [json_name = "httpUri", (.validate.rules) = { ... }</code>
   */
  com.fluxninja.generated.envoy.config.core.v3.HttpUriOrBuilder getHttpUriOrBuilder();

  /**
   * <pre>
   * SHA256 string for verifying data.
   * </pre>
   *
   * <code>string sha256 = 2 [json_name = "sha256", (.validate.rules) = { ... }</code>
   * @return The sha256.
   */
  java.lang.String getSha256();
  /**
   * <pre>
   * SHA256 string for verifying data.
   * </pre>
   *
   * <code>string sha256 = 2 [json_name = "sha256", (.validate.rules) = { ... }</code>
   * @return The bytes for sha256.
   */
  com.google.protobuf.ByteString
      getSha256Bytes();

  /**
   * <pre>
   * Retry policy for fetching remote data.
   * </pre>
   *
   * <code>.envoy.config.core.v3.RetryPolicy retry_policy = 3 [json_name = "retryPolicy"];</code>
   * @return Whether the retryPolicy field is set.
   */
  boolean hasRetryPolicy();
  /**
   * <pre>
   * Retry policy for fetching remote data.
   * </pre>
   *
   * <code>.envoy.config.core.v3.RetryPolicy retry_policy = 3 [json_name = "retryPolicy"];</code>
   * @return The retryPolicy.
   */
  com.fluxninja.generated.envoy.config.core.v3.RetryPolicy getRetryPolicy();
  /**
   * <pre>
   * Retry policy for fetching remote data.
   * </pre>
   *
   * <code>.envoy.config.core.v3.RetryPolicy retry_policy = 3 [json_name = "retryPolicy"];</code>
   */
  com.fluxninja.generated.envoy.config.core.v3.RetryPolicyOrBuilder getRetryPolicyOrBuilder();
}
