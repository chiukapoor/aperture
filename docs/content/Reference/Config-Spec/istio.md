---
title: Istio
slug: /reference/configuration/istio
description: Aperture Istio page
keywords:
  - istio
  - envoyfilter
---

## Data Collected

The Aperture Agent collects the below data using the Envoy Filter:

| Key                          | Value                                                          | Description                                                                            | Type              |
| ---------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------------------------------- | ----------------- |
| http.method                  | "%REQ(:METHOD)%"                                               | The HTTP method used for the request                                                   | Inbound, Outbound |
| http.target                  | "%REQ(:PATH)%"                                                 | The HTTP path requested by the client                                                  | Inbound, Outbound |
| http.host                    | "%REQ(HOST)%"                                                  | The value of the Host (HTTP/1.1)                                                       | Inbound, Outbound |
| http.user_agent              | "%REQ(USER-AGENT)%"                                            | The user agent string to identify the specific type of software request agent          | Inbound, Outbound |
| http.duration_millis         | "%DURATION%"                                                   | Total duration in milliseconds of the request from the start time to the last byte out | Inbound, Outbound |
| http.request_content_length  | "%BYTES_RECEIVED%"                                             | Body bytes received                                                                    | Inbound, Outbound |
| http.response_content_length | "%BYTES_SENT%"                                                 | Body bytes sent. For WebSocket connection it will also include response header bytes   | Inbound, Outbound |
| http.status_code             | "%RESPONSE_CODE%"                                              | HTTP response code                                                                     | Inbound, Outbound |
| fn.flow                      | "%DYNAMIC_METADATA(envoy.filters.http.ext_authz:fn.flow)%"     | FluxNinja Flow Labels                                                                  | Inbound, Outbound |
| fn.policies                  | "%DYNAMIC_METADATA(envoy.filters.http.ext_authz:fn.policies)%" | FluxNinja Policy details                                                               | Inbound, Outbound |
| control_point                | ingress/egress                                                 | Request Type                                                                           | Inbound, Outbound |
| net.peer.address             | "%UPSTREAM_HOST%"                                              | Upstream host URL                                                                      | Outbound          |
| net.host.address             | "%UPSTREAM_LOCAL_ADDRESS%"                                     | Local address of the upstream connection                                               | Outbound          |
| net.peer.ip                  | "%DOWNSTREAM_REMOTE_ADDRESS_WITHOUT_PORT%"                     | Remote address of the downstream connection, without any port component                | Inbound           |
| net.host.ip                  | "%DOWNSTREAM_LOCAL_ADDRESS_WITHOUT_PORT%"                      | Local address of the downstream connection, without any port component                 | Inbound           |
| net.host.port                | "%DOWNSTREAM_LOCAL_PORT%"                                      | Local port of the downstream connection                                                | Inbound           |

More information about the extracted values can be found on
[this site](https://www.envoyproxy.io/docs/envoy/latest/configuration/observability/access_log/usage#config-access-log).

## Envoy Filter

[Envoy Filter](https://istio.io/latest/docs/reference/config/networking/envoy-filter/)
is used to customize the default configurations generated by the Istio. The
Aperture Agent requires some of the additional details and needs the below
[Config Patches](https://istio.io/latest/docs/reference/config/networking/envoy-filter/#EnvoyFilter-EnvoyConfigObjectPatch)
to be added via the Envoy Filter.

**Note**: In all the below patches, it is presumed that the Aperture Agent is
installed with `agent` as the helm release name and is installed in the
`aperture-system` namespace, which makes the target URL value
`agent-aperture-agent.aperture-system.svc.cluster.local`.

1. The below patch merges the
   [Access Log](https://www.envoyproxy.io/docs/envoy/latest/configuration/observability/access_log/usage#config-access-log)
   config of type
   [Open Telemetry](https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/access_loggers/open_telemetry/v3/logs_service.proto#extensions-access-loggers-open-telemetry-v3-opentelemetryaccesslogconfig)
   with extracted values from the filter, to the
   [HTTP Connection Manager](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_conn_man/http_conn_man#)
   filter for outbound listener, in the Istio sidecar running with the
   application.

   The Open Telemetry config in the below patch has the attributes having
   extracted values which gets forwarded to the Aperture Agent instance using
   gRPC.

   The prepared log has the request method value as log body and `egress` as the
   log name to differentiate between different access logs coming from the same
   Envoy.

   ```yaml
   applyTo: NETWORK_FILTER
   match:
     context: SIDECAR_OUTBOUND
     listener:
       filterChain:
         filter:
           name: "envoy.filters.network.http_connection_manager"
   patch:
     operation: MERGE
     value:
       name: "envoy.filters.network.http_connection_manager"
       typed_config:
         "@type": "type.googleapis.com/envoy.extensions.filters.network.http_connection_manager.v3.HttpConnectionManager"
         access_log:
           - name: envoy.access_loggers.open_telemetry
             typed_config:
               "@type": "type.googleapis.com/envoy.extensions.access_loggers.open_telemetry.v3alpha.OpenTelemetryAccessLogConfig"
               common_config:
                 log_name: egress
                 grpc_service:
                   google_grpc:
                     target_uri: agent-aperture-agent.aperture-system.svc.cluster.local:4317
                     stat_prefix: fn_otlp_access_log
                 transport_api_version: V3
               body:
                 string_value: "%REQ(:METHOD)%"
               attributes:
                 values:
                   - key: http.method
                     value:
                       string_value: "%REQ(:METHOD)%"
                   - key: http.target
                     value:
                       string_value: "%REQ(:PATH)%"
                   - key: http.host
                     value:
                       string_value: "%REQ(HOST)%"
                   - key: http.user_agent
                     value:
                       string_value: "%REQ(USER-AGENT)%"
                   - key: http.duration_millis
                     value:
                       string_value: "%DURATION%"
                   - key: http.request_content_length
                     value:
                       string_value: "%BYTES_RECEIVED%"
                   - key: http.response_content_length
                     value:
                       string_value: "%BYTES_SENT%"
                   - key: http.status_code
                     value:
                       string_value: "%RESPONSE_CODE%"
                   - key: fn.flow
                     value:
                       string_value: "%DYNAMIC_METADATA(envoy.filters.http.ext_authz:fn.flow)%"
                   - key: fn.overwiev
                     value:
                       string_value: "%DYNAMIC_METADATA(envoy.filters.http.ext_authz:fn.overwiev)%"
                   - key: fn.fluxmeters
                     value:
                       string_value: "%DYNAMIC_METADATA(envoy.filters.http.ext_authz:fn.fluxmeters)%"
                   - key: control_point
                     value:
                       string_value: "egress"
                   - key: net.peer.address
                     value:
                       string_value: "%UPSTREAM_HOST%"
                   - key: net.host.address
                     value:
                       string_value: "%UPSTREAM_LOCAL_ADDRESS%"
   ```

2. The below patch also merges the
   [Access Log](https://www.envoyproxy.io/docs/envoy/latest/configuration/observability/access_log/usage#config-access-log)
   config of type
   [Open Telemetry](https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/access_loggers/open_telemetry/v3/logs_service.proto#extensions-access-loggers-open-telemetry-v3-opentelemetryaccesslogconfig)
   to the
   [HTTP Connection Manager](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_conn_man/http_conn_man#)
   filter but for inbound listener in the Istio sidecar running with the
   application.

   The Open Telemetry config in the below patch has the attributes having
   extracted values which gets forwarded to the Aperture Agent instance using
   gRPC.

   The prepared log has the request method value as log body and `ingress` as
   the log name to differentiate between different access logs coming from the
   same Envoy.

   ```yaml
   applyTo: NETWORK_FILTER
   match:
     context: SIDECAR_INBOUND
     listener:
       filterChain:
         filter:
           name: "envoy.filters.network.http_connection_manager"
   patch:
     operation: MERGE
     value:
       name: "envoy.filters.network.http_connection_manager"
       typed_config:
         "@type": "type.googleapis.com/envoy.extensions.filters.network.http_connection_manager.v3.HttpConnectionManager"
         access_log:
           - name: envoy.access_loggers.open_telemetry
             typed_config:
               "@type": "type.googleapis.com/envoy.extensions.access_loggers.open_telemetry.v3alpha.OpenTelemetryAccessLogConfig"
               common_config:
                 log_name: ingress
                 grpc_service:
                   google_grpc:
                     target_uri: agent-aperture-agent.aperture-system.svc.cluster.local:4317
                     stat_prefix: fn_otlp_access_log
                 transport_api_version: V3
               body:
                 string_value: "%REQ(:METHOD)%"
               attributes:
                 values:
                   - key: http.method
                     value:
                       string_value: "%REQ(:METHOD)%"
                   - key: http.target
                     value:
                       string_value: "%REQ(:PATH)%"
                   - key: http.host
                     value:
                       string_value: "%REQ(HOST)%"
                   - key: http.user_agent
                     value:
                       string_value: "%REQ(USER-AGENT)%"
                   - key: http.duration_millis
                     value:
                       string_value: "%DURATION%"
                   - key: http.request_content_length
                     value:
                       string_value: "%BYTES_RECEIVED%"
                   - key: http.response_content_length
                     value:
                       string_value: "%BYTES_SENT%"
                   - key: http.status_code
                     value:
                       string_value: "%RESPONSE_CODE%"
                   - key: fn.flow
                     value:
                       string_value: "%DYNAMIC_METADATA(envoy.filters.http.ext_authz:fn.flow)%"
                   - key: fn.overwiev
                     value:
                       string_value: "%DYNAMIC_METADATA(envoy.filters.http.ext_authz:fn.overwiev)%"
                   - key: fn.fluxmeters
                     value:
                       string_value: "%DYNAMIC_METADATA(envoy.filters.http.ext_authz:fn.fluxmeters)%"
                   - key: control_point
                     value:
                       string_value: "ingress"
                   - key: net.peer.ip
                     value:
                       string_value: "%DOWNSTREAM_REMOTE_ADDRESS_WITHOUT_PORT%"
                   - key: net.host.ip
                     value:
                       string_value: "%DOWNSTREAM_LOCAL_ADDRESS_WITHOUT_PORT%"
                   - key: net.host.port
                     value:
                       string_value: "%DOWNSTREAM_LOCAL_PORT%"
   ```

3. The below patch inserts the
   [External Authorization](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/ext_authz_filter)
   before the `Router` sub-filter of the
   [HTTP Connection Manager](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_conn_man/http_conn_man#)
   filter for inbound listener in the Istio sidecar running with the
   application.

   The External Authorization filter buffers the client request body with
   maximum size of `8192` bytes, and forwards it to the Aperture Agent instance
   using gRPC with a timeout of `0.01s`, having `INBOUND` value for key
   `traffic-direction` metadata included in the streams initiated to the gRPC
   service. The filter will accept client request even if the communication with
   the authorization service has failed, or if the authorization service has
   returned a HTTP 5xx error.

   ```yaml
   applyTo: HTTP_FILTER
   match:
     context: SIDECAR_INBOUND
     listener:
       filterChain:
         filter:
           name: "envoy.filters.network.http_connection_manager"
           subFilter:
             name: "envoy.filters.http.router"
   patch:
     operation: INSERT_BEFORE
     filterClass: AUTHZ
     value:
       name: envoy.filters.http.ext_authz
       typed_config:
         "@type": "type.googleapis.com/envoy.extensions.filters.http.ext_authz.v3.ExtAuthz"
         transport_api_version: V3
         with_request_body:
           max_request_bytes: 8192
           allow_partial_message: true
         failure_mode_allow: true
         grpc_service:
           google_grpc:
             target_uri: agent-aperture-agent.aperture-system.svc.cluster.local:80
             stat_prefix: ext_authz
           timeout: 0.01s
           initial_metadata:
             - key: traffic-direction
               value: INBOUND
   ```

4. The below patch also inserts the
   [External Authorization](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/ext_authz_filter)
   before the `Router` sub-filter of the
   [HTTP Connection Manager](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_conn_man/http_conn_man#)
   filter, but for outbound listener in the Istio sidecar running with the
   application.

   The External Authorization filter buffers the client request body with
   maximum size of `8192` bytes, and forwards it to the Aperture Agent instance
   using gRPC with a timeout of `0.01s`, having `OUTBOUND` value for key
   `traffic-direction` metadata included in the streams initiated to the gRPC
   service. The filter will accept client request even if the communication with
   the authorization service has failed, or if the authorization service has
   returned a HTTP 5xx error.

   ```yaml
   applyTo: HTTP_FILTER
   match:
     context: SIDECAR_OUTBOUND
     listener:
       filterChain:
         filter:
           name: "envoy.filters.network.http_connection_manager"
           subFilter:
             name: "envoy.filters.http.router"
   patch:
     operation: INSERT_BEFORE
     filterClass: AUTHZ
     value:
       name: envoy.filters.http.ext_authz
       typed_config:
         "@type": "type.googleapis.com/envoy.extensions.filters.http.ext_authz.v3.ExtAuthz"
         transport_api_version: V3
         with_request_body:
           max_request_bytes: 8192
           allow_partial_message: true
         failure_mode_allow: true
         grpc_service:
           google_grpc:
             target_uri: agent-aperture-agent.aperture-system.svc.cluster.local:80
             stat_prefix: ext_authz
           timeout: 0.01s
           initial_metadata:
             - key: traffic-direction
               value: OUTBOUND
   ```