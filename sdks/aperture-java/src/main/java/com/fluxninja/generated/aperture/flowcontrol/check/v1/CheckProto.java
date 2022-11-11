// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: aperture/flowcontrol/check/v1/check.proto

package com.fluxninja.generated.aperture.flowcontrol.check.v1;

public final class CheckProto {
  private CheckProto() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_aperture_flowcontrol_check_v1_CheckRequest_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_aperture_flowcontrol_check_v1_CheckRequest_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_aperture_flowcontrol_check_v1_CheckRequest_LabelsEntry_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_aperture_flowcontrol_check_v1_CheckRequest_LabelsEntry_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_aperture_flowcontrol_check_v1_CheckResponse_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_aperture_flowcontrol_check_v1_CheckResponse_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_aperture_flowcontrol_check_v1_CheckResponse_TelemetryFlowLabelsEntry_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_aperture_flowcontrol_check_v1_CheckResponse_TelemetryFlowLabelsEntry_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_aperture_flowcontrol_check_v1_ControlPointInfo_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_aperture_flowcontrol_check_v1_ControlPointInfo_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_aperture_flowcontrol_check_v1_ClassifierInfo_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_aperture_flowcontrol_check_v1_ClassifierInfo_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_aperture_flowcontrol_check_v1_LimiterDecision_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_aperture_flowcontrol_check_v1_LimiterDecision_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_aperture_flowcontrol_check_v1_LimiterDecision_RateLimiterInfo_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_aperture_flowcontrol_check_v1_LimiterDecision_RateLimiterInfo_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_aperture_flowcontrol_check_v1_LimiterDecision_ConcurrencyLimiterInfo_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_aperture_flowcontrol_check_v1_LimiterDecision_ConcurrencyLimiterInfo_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_aperture_flowcontrol_check_v1_FluxMeterInfo_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_aperture_flowcontrol_check_v1_FluxMeterInfo_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n)aperture/flowcontrol/check/v1/check.pr" +
      "oto\022\035aperture.flowcontrol.check.v1\032\037goog" +
      "le/protobuf/timestamp.proto\"\264\001\n\014CheckReq" +
      "uest\022\030\n\007feature\030\001 \001(\tR\007feature\022O\n\006labels" +
      "\030\002 \003(\01327.aperture.flowcontrol.check.v1.C" +
      "heckRequest.LabelsEntryR\006labels\0329\n\013Label" +
      "sEntry\022\020\n\003key\030\001 \001(\tR\003key\022\024\n\005value\030\002 \001(\tR" +
      "\005value:\0028\001\"\333\010\n\rCheckResponse\0220\n\005start\030\001 " +
      "\001(\0132\032.google.protobuf.TimestampR\005start\022," +
      "\n\003end\030\002 \001(\0132\032.google.protobuf.TimestampR" +
      "\003end\022\032\n\010services\030\004 \003(\tR\010services\022]\n\022cont" +
      "rol_point_info\030\005 \001(\0132/.aperture.flowcont" +
      "rol.check.v1.ControlPointInfoR\020controlPo" +
      "intInfo\022&\n\017flow_label_keys\030\006 \003(\tR\rflowLa" +
      "belKeys\022y\n\025telemetry_flow_labels\030\007 \003(\0132E" +
      ".aperture.flowcontrol.check.v1.CheckResp" +
      "onse.TelemetryFlowLabelsEntryR\023telemetry" +
      "FlowLabels\022^\n\rdecision_type\030\010 \001(\01629.aper" +
      "ture.flowcontrol.check.v1.CheckResponse." +
      "DecisionTypeR\014decisionType\022^\n\rreject_rea" +
      "son\030\t \001(\01629.aperture.flowcontrol.check.v" +
      "1.CheckResponse.RejectReasonR\014rejectReas" +
      "on\022X\n\020classifier_infos\030\n \003(\0132-.aperture." +
      "flowcontrol.check.v1.ClassifierInfoR\017cla" +
      "ssifierInfos\022V\n\020flux_meter_infos\030\013 \003(\0132," +
      ".aperture.flowcontrol.check.v1.FluxMeter" +
      "InfoR\016fluxMeterInfos\022[\n\021limiter_decision" +
      "s\030\014 \003(\0132..aperture.flowcontrol.check.v1." +
      "LimiterDecisionR\020limiterDecisions\032F\n\030Tel" +
      "emetryFlowLabelsEntry\022\020\n\003key\030\001 \001(\tR\003key\022" +
      "\024\n\005value\030\002 \001(\tR\005value:\0028\001\"m\n\014RejectReaso" +
      "n\022\026\n\022REJECT_REASON_NONE\020\000\022\036\n\032REJECT_REAS" +
      "ON_RATE_LIMITED\020\001\022%\n!REJECT_REASON_CONCU" +
      "RRENCY_LIMITED\020\002\"F\n\014DecisionType\022\032\n\026DECI" +
      "SION_TYPE_ACCEPTED\020\000\022\032\n\026DECISION_TYPE_RE" +
      "JECTED\020\001\"\305\001\n\020ControlPointInfo\022H\n\004type\030\001 " +
      "\001(\01624.aperture.flowcontrol.check.v1.Cont" +
      "rolPointInfo.TypeR\004type\022\030\n\007feature\030\002 \001(\t" +
      "R\007feature\"M\n\004Type\022\020\n\014TYPE_UNKNOWN\020\000\022\020\n\014T" +
      "YPE_FEATURE\020\001\022\020\n\014TYPE_INGRESS\020\002\022\017\n\013TYPE_" +
      "EGRESS\020\003\"\212\003\n\016ClassifierInfo\022\037\n\013policy_na" +
      "me\030\001 \001(\tR\npolicyName\022\037\n\013policy_hash\030\002 \001(" +
      "\tR\npolicyHash\022)\n\020classifier_index\030\003 \001(\003R" +
      "\017classifierIndex\022\033\n\tlabel_key\030\004 \001(\tR\010lab" +
      "elKey\022I\n\005error\030\005 \001(\01623.aperture.flowcont" +
      "rol.check.v1.ClassifierInfo.ErrorR\005error" +
      "\"\242\001\n\005Error\022\016\n\nERROR_NONE\020\000\022\025\n\021ERROR_EVAL" +
      "_FAILED\020\001\022\031\n\025ERROR_EMPTY_RESULTSET\020\002\022\035\n\031" +
      "ERROR_AMBIGUOUS_RESULTSET\020\003\022\032\n\026ERROR_MUL" +
      "TI_EXPRESSION\020\004\022\034\n\030ERROR_EXPRESSION_NOT_" +
      "MAP\020\005\"\336\005\n\017LimiterDecision\022\037\n\013policy_name" +
      "\030\001 \001(\tR\npolicyName\022\037\n\013policy_hash\030\002 \001(\tR" +
      "\npolicyHash\022\'\n\017component_index\030\003 \001(\003R\016co" +
      "mponentIndex\022\030\n\007dropped\030\004 \001(\010R\007dropped\022T" +
      "\n\006reason\030\005 \001(\0162<.aperture.flowcontrol.ch" +
      "eck.v1.LimiterDecision.LimiterReasonR\006re" +
      "ason\022l\n\021rate_limiter_info\030\006 \001(\0132>.apertu" +
      "re.flowcontrol.check.v1.LimiterDecision." +
      "RateLimiterInfoH\000R\017rateLimiterInfo\022\201\001\n\030c" +
      "oncurrency_limiter_info\030\007 \001(\0132E.aperture" +
      ".flowcontrol.check.v1.LimiterDecision.Co" +
      "ncurrencyLimiterInfoH\000R\026concurrencyLimit" +
      "erInfo\032_\n\017RateLimiterInfo\022\034\n\tremaining\030\001" +
      " \001(\003R\tremaining\022\030\n\007current\030\002 \001(\003R\007curren" +
      "t\022\024\n\005label\030\003 \001(\tR\005label\032?\n\026ConcurrencyLi" +
      "miterInfo\022%\n\016workload_index\030\001 \001(\tR\rworkl" +
      "oadIndex\"Q\n\rLimiterReason\022\036\n\032LIMITER_REA" +
      "SON_UNSPECIFIED\020\000\022 \n\034LIMITER_REASON_KEY_" +
      "NOT_FOUND\020\001B\t\n\007details\"7\n\rFluxMeterInfo\022" +
      "&\n\017flux_meter_name\030\001 \001(\tR\rfluxMeterName2" +
      "z\n\022FlowControlService\022d\n\005Check\022+.apertur" +
      "e.flowcontrol.check.v1.CheckRequest\032,.ap" +
      "erture.flowcontrol.check.v1.CheckRespons" +
      "e\"\000B\260\002\n5com.fluxninja.generated.aperture" +
      ".flowcontrol.check.v1B\nCheckProtoP\001ZTgit" +
      "hub.com/fluxninja/aperture/api/gen/proto" +
      "/go/aperture/flowcontrol/check/v1;checkv" +
      "1\242\002\003AFC\252\002\035Aperture.Flowcontrol.Check.V1\312" +
      "\002\035Aperture\\Flowcontrol\\Check\\V1\342\002)Apertu" +
      "re\\Flowcontrol\\Check\\V1\\GPBMetadata\352\002 Ap" +
      "erture::Flowcontrol::Check::V1b\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          com.google.protobuf.TimestampProto.getDescriptor(),
        });
    internal_static_aperture_flowcontrol_check_v1_CheckRequest_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_aperture_flowcontrol_check_v1_CheckRequest_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_aperture_flowcontrol_check_v1_CheckRequest_descriptor,
        new java.lang.String[] { "Feature", "Labels", });
    internal_static_aperture_flowcontrol_check_v1_CheckRequest_LabelsEntry_descriptor =
      internal_static_aperture_flowcontrol_check_v1_CheckRequest_descriptor.getNestedTypes().get(0);
    internal_static_aperture_flowcontrol_check_v1_CheckRequest_LabelsEntry_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_aperture_flowcontrol_check_v1_CheckRequest_LabelsEntry_descriptor,
        new java.lang.String[] { "Key", "Value", });
    internal_static_aperture_flowcontrol_check_v1_CheckResponse_descriptor =
      getDescriptor().getMessageTypes().get(1);
    internal_static_aperture_flowcontrol_check_v1_CheckResponse_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_aperture_flowcontrol_check_v1_CheckResponse_descriptor,
        new java.lang.String[] { "Start", "End", "Services", "ControlPointInfo", "FlowLabelKeys", "TelemetryFlowLabels", "DecisionType", "RejectReason", "ClassifierInfos", "FluxMeterInfos", "LimiterDecisions", });
    internal_static_aperture_flowcontrol_check_v1_CheckResponse_TelemetryFlowLabelsEntry_descriptor =
      internal_static_aperture_flowcontrol_check_v1_CheckResponse_descriptor.getNestedTypes().get(0);
    internal_static_aperture_flowcontrol_check_v1_CheckResponse_TelemetryFlowLabelsEntry_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_aperture_flowcontrol_check_v1_CheckResponse_TelemetryFlowLabelsEntry_descriptor,
        new java.lang.String[] { "Key", "Value", });
    internal_static_aperture_flowcontrol_check_v1_ControlPointInfo_descriptor =
      getDescriptor().getMessageTypes().get(2);
    internal_static_aperture_flowcontrol_check_v1_ControlPointInfo_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_aperture_flowcontrol_check_v1_ControlPointInfo_descriptor,
        new java.lang.String[] { "Type", "Feature", });
    internal_static_aperture_flowcontrol_check_v1_ClassifierInfo_descriptor =
      getDescriptor().getMessageTypes().get(3);
    internal_static_aperture_flowcontrol_check_v1_ClassifierInfo_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_aperture_flowcontrol_check_v1_ClassifierInfo_descriptor,
        new java.lang.String[] { "PolicyName", "PolicyHash", "ClassifierIndex", "LabelKey", "Error", });
    internal_static_aperture_flowcontrol_check_v1_LimiterDecision_descriptor =
      getDescriptor().getMessageTypes().get(4);
    internal_static_aperture_flowcontrol_check_v1_LimiterDecision_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_aperture_flowcontrol_check_v1_LimiterDecision_descriptor,
        new java.lang.String[] { "PolicyName", "PolicyHash", "ComponentIndex", "Dropped", "Reason", "RateLimiterInfo", "ConcurrencyLimiterInfo", "Details", });
    internal_static_aperture_flowcontrol_check_v1_LimiterDecision_RateLimiterInfo_descriptor =
      internal_static_aperture_flowcontrol_check_v1_LimiterDecision_descriptor.getNestedTypes().get(0);
    internal_static_aperture_flowcontrol_check_v1_LimiterDecision_RateLimiterInfo_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_aperture_flowcontrol_check_v1_LimiterDecision_RateLimiterInfo_descriptor,
        new java.lang.String[] { "Remaining", "Current", "Label", });
    internal_static_aperture_flowcontrol_check_v1_LimiterDecision_ConcurrencyLimiterInfo_descriptor =
      internal_static_aperture_flowcontrol_check_v1_LimiterDecision_descriptor.getNestedTypes().get(1);
    internal_static_aperture_flowcontrol_check_v1_LimiterDecision_ConcurrencyLimiterInfo_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_aperture_flowcontrol_check_v1_LimiterDecision_ConcurrencyLimiterInfo_descriptor,
        new java.lang.String[] { "WorkloadIndex", });
    internal_static_aperture_flowcontrol_check_v1_FluxMeterInfo_descriptor =
      getDescriptor().getMessageTypes().get(5);
    internal_static_aperture_flowcontrol_check_v1_FluxMeterInfo_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_aperture_flowcontrol_check_v1_FluxMeterInfo_descriptor,
        new java.lang.String[] { "FluxMeterName", });
    com.google.protobuf.TimestampProto.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}