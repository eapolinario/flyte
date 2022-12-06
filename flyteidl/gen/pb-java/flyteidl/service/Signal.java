// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: flyteidl/service/signal.proto

package flyteidl.service;

public final class Signal {
  private Signal() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\035flyteidl/service/signal.proto\022\020flyteid" +
      "l.service\032\034google/api/annotations.proto\032" +
      "\033flyteidl/admin/signal.proto\032,protoc-gen" +
      "-swagger/options/annotations.proto2\347\005\n\rS" +
      "ignalService\022\220\001\n\021GetOrCreateSignal\022(.fly" +
      "teidl.admin.SignalGetOrCreateRequest\032\026.f" +
      "lyteidl.admin.Signal\"9\222A6\0324Retrieve a si" +
      "gnal, creating it if it does not exist.\022" +
      "\216\002\n\013ListSignals\022!.flyteidl.admin.SignalL" +
      "istRequest\032\032.flyteidl.admin.SignalList\"\277" +
      "\001\202\323\344\223\002m\022k/api/v1/signals/{workflow_execu" +
      "tion_id.project}/{workflow_execution_id." +
      "domain}/{workflow_execution_id.name}\222AI\032" +
      "GFetch existing signal definitions match" +
      "ing the input signal id filters.\022\261\002\n\tSet" +
      "Signal\022 .flyteidl.admin.SignalSetRequest" +
      "\032!.flyteidl.admin.SignalSetResponse\"\336\001\202\323" +
      "\344\223\002\024\"\017/api/v1/signals:\001*\222A\300\001\032\023Set a sign" +
      "al value.JB\n\003400\022;\n9Returned for bad req" +
      "uest that may have failed validation.Je\n" +
      "\003409\022^\n\\Returned for a request that refe" +
      "rences an identical entity that has alre" +
      "ady been registered.B9Z7github.com/flyte" +
      "org/flyteidl/gen/pb-go/flyteidl/serviceb" +
      "\006proto3"
    };
    com.google.protobuf.Descriptors.FileDescriptor.InternalDescriptorAssigner assigner =
        new com.google.protobuf.Descriptors.FileDescriptor.    InternalDescriptorAssigner() {
          public com.google.protobuf.ExtensionRegistry assignDescriptors(
              com.google.protobuf.Descriptors.FileDescriptor root) {
            descriptor = root;
            return null;
          }
        };
    com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          com.google.api.AnnotationsProto.getDescriptor(),
          flyteidl.admin.SignalOuterClass.getDescriptor(),
          grpc.gateway.protoc_gen_swagger.options.Annotations.getDescriptor(),
        }, assigner);
    com.google.protobuf.ExtensionRegistry registry =
        com.google.protobuf.ExtensionRegistry.newInstance();
    registry.add(com.google.api.AnnotationsProto.http);
    registry.add(grpc.gateway.protoc_gen_swagger.options.Annotations.openapiv2Operation);
    com.google.protobuf.Descriptors.FileDescriptor
        .internalUpdateFileDescriptor(descriptor, registry);
    com.google.api.AnnotationsProto.getDescriptor();
    flyteidl.admin.SignalOuterClass.getDescriptor();
    grpc.gateway.protoc_gen_swagger.options.Annotations.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
