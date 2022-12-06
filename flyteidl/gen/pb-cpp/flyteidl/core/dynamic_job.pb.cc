// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: flyteidl/core/dynamic_job.proto

#include "flyteidl/core/dynamic_job.pb.h"

#include <algorithm>

#include <google/protobuf/stubs/common.h>
#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/extension_set.h>
#include <google/protobuf/wire_format_lite_inl.h>
#include <google/protobuf/descriptor.h>
#include <google/protobuf/generated_message_reflection.h>
#include <google/protobuf/reflection_ops.h>
#include <google/protobuf/wire_format.h>
// @@protoc_insertion_point(includes)
#include <google/protobuf/port_def.inc>

extern PROTOBUF_INTERNAL_EXPORT_flyteidl_2fcore_2fliterals_2eproto ::google::protobuf::internal::SCCInfo<1> scc_info_Binding_flyteidl_2fcore_2fliterals_2eproto;
extern PROTOBUF_INTERNAL_EXPORT_flyteidl_2fcore_2ftasks_2eproto ::google::protobuf::internal::SCCInfo<9> scc_info_TaskTemplate_flyteidl_2fcore_2ftasks_2eproto;
extern PROTOBUF_INTERNAL_EXPORT_flyteidl_2fcore_2fworkflow_2eproto ::google::protobuf::internal::SCCInfo<6> scc_info_WorkflowTemplate_flyteidl_2fcore_2fworkflow_2eproto;
extern PROTOBUF_INTERNAL_EXPORT_flyteidl_2fcore_2fworkflow_2eproto ::google::protobuf::internal::SCCInfo<8> scc_info_BranchNode_flyteidl_2fcore_2fworkflow_2eproto;
namespace flyteidl {
namespace core {
class DynamicJobSpecDefaultTypeInternal {
 public:
  ::google::protobuf::internal::ExplicitlyConstructed<DynamicJobSpec> _instance;
} _DynamicJobSpec_default_instance_;
}  // namespace core
}  // namespace flyteidl
static void InitDefaultsDynamicJobSpec_flyteidl_2fcore_2fdynamic_5fjob_2eproto() {
  GOOGLE_PROTOBUF_VERIFY_VERSION;

  {
    void* ptr = &::flyteidl::core::_DynamicJobSpec_default_instance_;
    new (ptr) ::flyteidl::core::DynamicJobSpec();
    ::google::protobuf::internal::OnShutdownDestroyMessage(ptr);
  }
  ::flyteidl::core::DynamicJobSpec::InitAsDefaultInstance();
}

::google::protobuf::internal::SCCInfo<4> scc_info_DynamicJobSpec_flyteidl_2fcore_2fdynamic_5fjob_2eproto =
    {{ATOMIC_VAR_INIT(::google::protobuf::internal::SCCInfoBase::kUninitialized), 4, InitDefaultsDynamicJobSpec_flyteidl_2fcore_2fdynamic_5fjob_2eproto}, {
      &scc_info_BranchNode_flyteidl_2fcore_2fworkflow_2eproto.base,
      &scc_info_Binding_flyteidl_2fcore_2fliterals_2eproto.base,
      &scc_info_TaskTemplate_flyteidl_2fcore_2ftasks_2eproto.base,
      &scc_info_WorkflowTemplate_flyteidl_2fcore_2fworkflow_2eproto.base,}};

void InitDefaults_flyteidl_2fcore_2fdynamic_5fjob_2eproto() {
  ::google::protobuf::internal::InitSCC(&scc_info_DynamicJobSpec_flyteidl_2fcore_2fdynamic_5fjob_2eproto.base);
}

::google::protobuf::Metadata file_level_metadata_flyteidl_2fcore_2fdynamic_5fjob_2eproto[1];
constexpr ::google::protobuf::EnumDescriptor const** file_level_enum_descriptors_flyteidl_2fcore_2fdynamic_5fjob_2eproto = nullptr;
constexpr ::google::protobuf::ServiceDescriptor const** file_level_service_descriptors_flyteidl_2fcore_2fdynamic_5fjob_2eproto = nullptr;

const ::google::protobuf::uint32 TableStruct_flyteidl_2fcore_2fdynamic_5fjob_2eproto::offsets[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
  ~0u,  // no _has_bits_
  PROTOBUF_FIELD_OFFSET(::flyteidl::core::DynamicJobSpec, _internal_metadata_),
  ~0u,  // no _extensions_
  ~0u,  // no _oneof_case_
  ~0u,  // no _weak_field_map_
  PROTOBUF_FIELD_OFFSET(::flyteidl::core::DynamicJobSpec, nodes_),
  PROTOBUF_FIELD_OFFSET(::flyteidl::core::DynamicJobSpec, min_successes_),
  PROTOBUF_FIELD_OFFSET(::flyteidl::core::DynamicJobSpec, outputs_),
  PROTOBUF_FIELD_OFFSET(::flyteidl::core::DynamicJobSpec, tasks_),
  PROTOBUF_FIELD_OFFSET(::flyteidl::core::DynamicJobSpec, subworkflows_),
};
static const ::google::protobuf::internal::MigrationSchema schemas[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
  { 0, -1, sizeof(::flyteidl::core::DynamicJobSpec)},
};

static ::google::protobuf::Message const * const file_default_instances[] = {
  reinterpret_cast<const ::google::protobuf::Message*>(&::flyteidl::core::_DynamicJobSpec_default_instance_),
};

::google::protobuf::internal::AssignDescriptorsTable assign_descriptors_table_flyteidl_2fcore_2fdynamic_5fjob_2eproto = {
  {}, AddDescriptors_flyteidl_2fcore_2fdynamic_5fjob_2eproto, "flyteidl/core/dynamic_job.proto", schemas,
  file_default_instances, TableStruct_flyteidl_2fcore_2fdynamic_5fjob_2eproto::offsets,
  file_level_metadata_flyteidl_2fcore_2fdynamic_5fjob_2eproto, 1, file_level_enum_descriptors_flyteidl_2fcore_2fdynamic_5fjob_2eproto, file_level_service_descriptors_flyteidl_2fcore_2fdynamic_5fjob_2eproto,
};

const char descriptor_table_protodef_flyteidl_2fcore_2fdynamic_5fjob_2eproto[] =
  "\n\037flyteidl/core/dynamic_job.proto\022\rflyte"
  "idl.core\032\031flyteidl/core/tasks.proto\032\034fly"
  "teidl/core/workflow.proto\032\034flyteidl/core"
  "/literals.proto\"\327\001\n\016DynamicJobSpec\022\"\n\005no"
  "des\030\001 \003(\0132\023.flyteidl.core.Node\022\025\n\rmin_su"
  "ccesses\030\002 \001(\003\022\'\n\007outputs\030\003 \003(\0132\026.flyteid"
  "l.core.Binding\022*\n\005tasks\030\004 \003(\0132\033.flyteidl"
  ".core.TaskTemplate\0225\n\014subworkflows\030\005 \003(\013"
  "2\037.flyteidl.core.WorkflowTemplateB6Z4git"
  "hub.com/flyteorg/flyteidl/gen/pb-go/flyt"
  "eidl/coreb\006proto3"
  ;
::google::protobuf::internal::DescriptorTable descriptor_table_flyteidl_2fcore_2fdynamic_5fjob_2eproto = {
  false, InitDefaults_flyteidl_2fcore_2fdynamic_5fjob_2eproto, 
  descriptor_table_protodef_flyteidl_2fcore_2fdynamic_5fjob_2eproto,
  "flyteidl/core/dynamic_job.proto", &assign_descriptors_table_flyteidl_2fcore_2fdynamic_5fjob_2eproto, 417,
};

void AddDescriptors_flyteidl_2fcore_2fdynamic_5fjob_2eproto() {
  static constexpr ::google::protobuf::internal::InitFunc deps[3] =
  {
    ::AddDescriptors_flyteidl_2fcore_2ftasks_2eproto,
    ::AddDescriptors_flyteidl_2fcore_2fworkflow_2eproto,
    ::AddDescriptors_flyteidl_2fcore_2fliterals_2eproto,
  };
 ::google::protobuf::internal::AddDescriptors(&descriptor_table_flyteidl_2fcore_2fdynamic_5fjob_2eproto, deps, 3);
}

// Force running AddDescriptors() at dynamic initialization time.
static bool dynamic_init_dummy_flyteidl_2fcore_2fdynamic_5fjob_2eproto = []() { AddDescriptors_flyteidl_2fcore_2fdynamic_5fjob_2eproto(); return true; }();
namespace flyteidl {
namespace core {

// ===================================================================

void DynamicJobSpec::InitAsDefaultInstance() {
}
class DynamicJobSpec::HasBitSetters {
 public:
};

void DynamicJobSpec::clear_nodes() {
  nodes_.Clear();
}
void DynamicJobSpec::clear_outputs() {
  outputs_.Clear();
}
void DynamicJobSpec::clear_tasks() {
  tasks_.Clear();
}
void DynamicJobSpec::clear_subworkflows() {
  subworkflows_.Clear();
}
#if !defined(_MSC_VER) || _MSC_VER >= 1900
const int DynamicJobSpec::kNodesFieldNumber;
const int DynamicJobSpec::kMinSuccessesFieldNumber;
const int DynamicJobSpec::kOutputsFieldNumber;
const int DynamicJobSpec::kTasksFieldNumber;
const int DynamicJobSpec::kSubworkflowsFieldNumber;
#endif  // !defined(_MSC_VER) || _MSC_VER >= 1900

DynamicJobSpec::DynamicJobSpec()
  : ::google::protobuf::Message(), _internal_metadata_(nullptr) {
  SharedCtor();
  // @@protoc_insertion_point(constructor:flyteidl.core.DynamicJobSpec)
}
DynamicJobSpec::DynamicJobSpec(const DynamicJobSpec& from)
  : ::google::protobuf::Message(),
      _internal_metadata_(nullptr),
      nodes_(from.nodes_),
      outputs_(from.outputs_),
      tasks_(from.tasks_),
      subworkflows_(from.subworkflows_) {
  _internal_metadata_.MergeFrom(from._internal_metadata_);
  min_successes_ = from.min_successes_;
  // @@protoc_insertion_point(copy_constructor:flyteidl.core.DynamicJobSpec)
}

void DynamicJobSpec::SharedCtor() {
  ::google::protobuf::internal::InitSCC(
      &scc_info_DynamicJobSpec_flyteidl_2fcore_2fdynamic_5fjob_2eproto.base);
  min_successes_ = PROTOBUF_LONGLONG(0);
}

DynamicJobSpec::~DynamicJobSpec() {
  // @@protoc_insertion_point(destructor:flyteidl.core.DynamicJobSpec)
  SharedDtor();
}

void DynamicJobSpec::SharedDtor() {
}

void DynamicJobSpec::SetCachedSize(int size) const {
  _cached_size_.Set(size);
}
const DynamicJobSpec& DynamicJobSpec::default_instance() {
  ::google::protobuf::internal::InitSCC(&::scc_info_DynamicJobSpec_flyteidl_2fcore_2fdynamic_5fjob_2eproto.base);
  return *internal_default_instance();
}


void DynamicJobSpec::Clear() {
// @@protoc_insertion_point(message_clear_start:flyteidl.core.DynamicJobSpec)
  ::google::protobuf::uint32 cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  nodes_.Clear();
  outputs_.Clear();
  tasks_.Clear();
  subworkflows_.Clear();
  min_successes_ = PROTOBUF_LONGLONG(0);
  _internal_metadata_.Clear();
}

#if GOOGLE_PROTOBUF_ENABLE_EXPERIMENTAL_PARSER
const char* DynamicJobSpec::_InternalParse(const char* begin, const char* end, void* object,
                  ::google::protobuf::internal::ParseContext* ctx) {
  auto msg = static_cast<DynamicJobSpec*>(object);
  ::google::protobuf::int32 size; (void)size;
  int depth; (void)depth;
  ::google::protobuf::uint32 tag;
  ::google::protobuf::internal::ParseFunc parser_till_end; (void)parser_till_end;
  auto ptr = begin;
  while (ptr < end) {
    ptr = ::google::protobuf::io::Parse32(ptr, &tag);
    GOOGLE_PROTOBUF_PARSER_ASSERT(ptr);
    switch (tag >> 3) {
      // repeated .flyteidl.core.Node nodes = 1;
      case 1: {
        if (static_cast<::google::protobuf::uint8>(tag) != 10) goto handle_unusual;
        do {
          ptr = ::google::protobuf::io::ReadSize(ptr, &size);
          GOOGLE_PROTOBUF_PARSER_ASSERT(ptr);
          parser_till_end = ::flyteidl::core::Node::_InternalParse;
          object = msg->add_nodes();
          if (size > end - ptr) goto len_delim_till_end;
          ptr += size;
          GOOGLE_PROTOBUF_PARSER_ASSERT(ctx->ParseExactRange(
              {parser_till_end, object}, ptr - size, ptr));
          if (ptr >= end) break;
        } while ((::google::protobuf::io::UnalignedLoad<::google::protobuf::uint64>(ptr) & 255) == 10 && (ptr += 1));
        break;
      }
      // int64 min_successes = 2;
      case 2: {
        if (static_cast<::google::protobuf::uint8>(tag) != 16) goto handle_unusual;
        msg->set_min_successes(::google::protobuf::internal::ReadVarint(&ptr));
        GOOGLE_PROTOBUF_PARSER_ASSERT(ptr);
        break;
      }
      // repeated .flyteidl.core.Binding outputs = 3;
      case 3: {
        if (static_cast<::google::protobuf::uint8>(tag) != 26) goto handle_unusual;
        do {
          ptr = ::google::protobuf::io::ReadSize(ptr, &size);
          GOOGLE_PROTOBUF_PARSER_ASSERT(ptr);
          parser_till_end = ::flyteidl::core::Binding::_InternalParse;
          object = msg->add_outputs();
          if (size > end - ptr) goto len_delim_till_end;
          ptr += size;
          GOOGLE_PROTOBUF_PARSER_ASSERT(ctx->ParseExactRange(
              {parser_till_end, object}, ptr - size, ptr));
          if (ptr >= end) break;
        } while ((::google::protobuf::io::UnalignedLoad<::google::protobuf::uint64>(ptr) & 255) == 26 && (ptr += 1));
        break;
      }
      // repeated .flyteidl.core.TaskTemplate tasks = 4;
      case 4: {
        if (static_cast<::google::protobuf::uint8>(tag) != 34) goto handle_unusual;
        do {
          ptr = ::google::protobuf::io::ReadSize(ptr, &size);
          GOOGLE_PROTOBUF_PARSER_ASSERT(ptr);
          parser_till_end = ::flyteidl::core::TaskTemplate::_InternalParse;
          object = msg->add_tasks();
          if (size > end - ptr) goto len_delim_till_end;
          ptr += size;
          GOOGLE_PROTOBUF_PARSER_ASSERT(ctx->ParseExactRange(
              {parser_till_end, object}, ptr - size, ptr));
          if (ptr >= end) break;
        } while ((::google::protobuf::io::UnalignedLoad<::google::protobuf::uint64>(ptr) & 255) == 34 && (ptr += 1));
        break;
      }
      // repeated .flyteidl.core.WorkflowTemplate subworkflows = 5;
      case 5: {
        if (static_cast<::google::protobuf::uint8>(tag) != 42) goto handle_unusual;
        do {
          ptr = ::google::protobuf::io::ReadSize(ptr, &size);
          GOOGLE_PROTOBUF_PARSER_ASSERT(ptr);
          parser_till_end = ::flyteidl::core::WorkflowTemplate::_InternalParse;
          object = msg->add_subworkflows();
          if (size > end - ptr) goto len_delim_till_end;
          ptr += size;
          GOOGLE_PROTOBUF_PARSER_ASSERT(ctx->ParseExactRange(
              {parser_till_end, object}, ptr - size, ptr));
          if (ptr >= end) break;
        } while ((::google::protobuf::io::UnalignedLoad<::google::protobuf::uint64>(ptr) & 255) == 42 && (ptr += 1));
        break;
      }
      default: {
      handle_unusual:
        if ((tag & 7) == 4 || tag == 0) {
          ctx->EndGroup(tag);
          return ptr;
        }
        auto res = UnknownFieldParse(tag, {_InternalParse, msg},
          ptr, end, msg->_internal_metadata_.mutable_unknown_fields(), ctx);
        ptr = res.first;
        GOOGLE_PROTOBUF_PARSER_ASSERT(ptr != nullptr);
        if (res.second) return ptr;
      }
    }  // switch
  }  // while
  return ptr;
len_delim_till_end:
  return ctx->StoreAndTailCall(ptr, end, {_InternalParse, msg},
                               {parser_till_end, object}, size);
}
#else  // GOOGLE_PROTOBUF_ENABLE_EXPERIMENTAL_PARSER
bool DynamicJobSpec::MergePartialFromCodedStream(
    ::google::protobuf::io::CodedInputStream* input) {
#define DO_(EXPRESSION) if (!PROTOBUF_PREDICT_TRUE(EXPRESSION)) goto failure
  ::google::protobuf::uint32 tag;
  // @@protoc_insertion_point(parse_start:flyteidl.core.DynamicJobSpec)
  for (;;) {
    ::std::pair<::google::protobuf::uint32, bool> p = input->ReadTagWithCutoffNoLastTag(127u);
    tag = p.first;
    if (!p.second) goto handle_unusual;
    switch (::google::protobuf::internal::WireFormatLite::GetTagFieldNumber(tag)) {
      // repeated .flyteidl.core.Node nodes = 1;
      case 1: {
        if (static_cast< ::google::protobuf::uint8>(tag) == (10 & 0xFF)) {
          DO_(::google::protobuf::internal::WireFormatLite::ReadMessage(
                input, add_nodes()));
        } else {
          goto handle_unusual;
        }
        break;
      }

      // int64 min_successes = 2;
      case 2: {
        if (static_cast< ::google::protobuf::uint8>(tag) == (16 & 0xFF)) {

          DO_((::google::protobuf::internal::WireFormatLite::ReadPrimitive<
                   ::google::protobuf::int64, ::google::protobuf::internal::WireFormatLite::TYPE_INT64>(
                 input, &min_successes_)));
        } else {
          goto handle_unusual;
        }
        break;
      }

      // repeated .flyteidl.core.Binding outputs = 3;
      case 3: {
        if (static_cast< ::google::protobuf::uint8>(tag) == (26 & 0xFF)) {
          DO_(::google::protobuf::internal::WireFormatLite::ReadMessage(
                input, add_outputs()));
        } else {
          goto handle_unusual;
        }
        break;
      }

      // repeated .flyteidl.core.TaskTemplate tasks = 4;
      case 4: {
        if (static_cast< ::google::protobuf::uint8>(tag) == (34 & 0xFF)) {
          DO_(::google::protobuf::internal::WireFormatLite::ReadMessage(
                input, add_tasks()));
        } else {
          goto handle_unusual;
        }
        break;
      }

      // repeated .flyteidl.core.WorkflowTemplate subworkflows = 5;
      case 5: {
        if (static_cast< ::google::protobuf::uint8>(tag) == (42 & 0xFF)) {
          DO_(::google::protobuf::internal::WireFormatLite::ReadMessage(
                input, add_subworkflows()));
        } else {
          goto handle_unusual;
        }
        break;
      }

      default: {
      handle_unusual:
        if (tag == 0) {
          goto success;
        }
        DO_(::google::protobuf::internal::WireFormat::SkipField(
              input, tag, _internal_metadata_.mutable_unknown_fields()));
        break;
      }
    }
  }
success:
  // @@protoc_insertion_point(parse_success:flyteidl.core.DynamicJobSpec)
  return true;
failure:
  // @@protoc_insertion_point(parse_failure:flyteidl.core.DynamicJobSpec)
  return false;
#undef DO_
}
#endif  // GOOGLE_PROTOBUF_ENABLE_EXPERIMENTAL_PARSER

void DynamicJobSpec::SerializeWithCachedSizes(
    ::google::protobuf::io::CodedOutputStream* output) const {
  // @@protoc_insertion_point(serialize_start:flyteidl.core.DynamicJobSpec)
  ::google::protobuf::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  // repeated .flyteidl.core.Node nodes = 1;
  for (unsigned int i = 0,
      n = static_cast<unsigned int>(this->nodes_size()); i < n; i++) {
    ::google::protobuf::internal::WireFormatLite::WriteMessageMaybeToArray(
      1,
      this->nodes(static_cast<int>(i)),
      output);
  }

  // int64 min_successes = 2;
  if (this->min_successes() != 0) {
    ::google::protobuf::internal::WireFormatLite::WriteInt64(2, this->min_successes(), output);
  }

  // repeated .flyteidl.core.Binding outputs = 3;
  for (unsigned int i = 0,
      n = static_cast<unsigned int>(this->outputs_size()); i < n; i++) {
    ::google::protobuf::internal::WireFormatLite::WriteMessageMaybeToArray(
      3,
      this->outputs(static_cast<int>(i)),
      output);
  }

  // repeated .flyteidl.core.TaskTemplate tasks = 4;
  for (unsigned int i = 0,
      n = static_cast<unsigned int>(this->tasks_size()); i < n; i++) {
    ::google::protobuf::internal::WireFormatLite::WriteMessageMaybeToArray(
      4,
      this->tasks(static_cast<int>(i)),
      output);
  }

  // repeated .flyteidl.core.WorkflowTemplate subworkflows = 5;
  for (unsigned int i = 0,
      n = static_cast<unsigned int>(this->subworkflows_size()); i < n; i++) {
    ::google::protobuf::internal::WireFormatLite::WriteMessageMaybeToArray(
      5,
      this->subworkflows(static_cast<int>(i)),
      output);
  }

  if (_internal_metadata_.have_unknown_fields()) {
    ::google::protobuf::internal::WireFormat::SerializeUnknownFields(
        _internal_metadata_.unknown_fields(), output);
  }
  // @@protoc_insertion_point(serialize_end:flyteidl.core.DynamicJobSpec)
}

::google::protobuf::uint8* DynamicJobSpec::InternalSerializeWithCachedSizesToArray(
    ::google::protobuf::uint8* target) const {
  // @@protoc_insertion_point(serialize_to_array_start:flyteidl.core.DynamicJobSpec)
  ::google::protobuf::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  // repeated .flyteidl.core.Node nodes = 1;
  for (unsigned int i = 0,
      n = static_cast<unsigned int>(this->nodes_size()); i < n; i++) {
    target = ::google::protobuf::internal::WireFormatLite::
      InternalWriteMessageToArray(
        1, this->nodes(static_cast<int>(i)), target);
  }

  // int64 min_successes = 2;
  if (this->min_successes() != 0) {
    target = ::google::protobuf::internal::WireFormatLite::WriteInt64ToArray(2, this->min_successes(), target);
  }

  // repeated .flyteidl.core.Binding outputs = 3;
  for (unsigned int i = 0,
      n = static_cast<unsigned int>(this->outputs_size()); i < n; i++) {
    target = ::google::protobuf::internal::WireFormatLite::
      InternalWriteMessageToArray(
        3, this->outputs(static_cast<int>(i)), target);
  }

  // repeated .flyteidl.core.TaskTemplate tasks = 4;
  for (unsigned int i = 0,
      n = static_cast<unsigned int>(this->tasks_size()); i < n; i++) {
    target = ::google::protobuf::internal::WireFormatLite::
      InternalWriteMessageToArray(
        4, this->tasks(static_cast<int>(i)), target);
  }

  // repeated .flyteidl.core.WorkflowTemplate subworkflows = 5;
  for (unsigned int i = 0,
      n = static_cast<unsigned int>(this->subworkflows_size()); i < n; i++) {
    target = ::google::protobuf::internal::WireFormatLite::
      InternalWriteMessageToArray(
        5, this->subworkflows(static_cast<int>(i)), target);
  }

  if (_internal_metadata_.have_unknown_fields()) {
    target = ::google::protobuf::internal::WireFormat::SerializeUnknownFieldsToArray(
        _internal_metadata_.unknown_fields(), target);
  }
  // @@protoc_insertion_point(serialize_to_array_end:flyteidl.core.DynamicJobSpec)
  return target;
}

size_t DynamicJobSpec::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:flyteidl.core.DynamicJobSpec)
  size_t total_size = 0;

  if (_internal_metadata_.have_unknown_fields()) {
    total_size +=
      ::google::protobuf::internal::WireFormat::ComputeUnknownFieldsSize(
        _internal_metadata_.unknown_fields());
  }
  ::google::protobuf::uint32 cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  // repeated .flyteidl.core.Node nodes = 1;
  {
    unsigned int count = static_cast<unsigned int>(this->nodes_size());
    total_size += 1UL * count;
    for (unsigned int i = 0; i < count; i++) {
      total_size +=
        ::google::protobuf::internal::WireFormatLite::MessageSize(
          this->nodes(static_cast<int>(i)));
    }
  }

  // repeated .flyteidl.core.Binding outputs = 3;
  {
    unsigned int count = static_cast<unsigned int>(this->outputs_size());
    total_size += 1UL * count;
    for (unsigned int i = 0; i < count; i++) {
      total_size +=
        ::google::protobuf::internal::WireFormatLite::MessageSize(
          this->outputs(static_cast<int>(i)));
    }
  }

  // repeated .flyteidl.core.TaskTemplate tasks = 4;
  {
    unsigned int count = static_cast<unsigned int>(this->tasks_size());
    total_size += 1UL * count;
    for (unsigned int i = 0; i < count; i++) {
      total_size +=
        ::google::protobuf::internal::WireFormatLite::MessageSize(
          this->tasks(static_cast<int>(i)));
    }
  }

  // repeated .flyteidl.core.WorkflowTemplate subworkflows = 5;
  {
    unsigned int count = static_cast<unsigned int>(this->subworkflows_size());
    total_size += 1UL * count;
    for (unsigned int i = 0; i < count; i++) {
      total_size +=
        ::google::protobuf::internal::WireFormatLite::MessageSize(
          this->subworkflows(static_cast<int>(i)));
    }
  }

  // int64 min_successes = 2;
  if (this->min_successes() != 0) {
    total_size += 1 +
      ::google::protobuf::internal::WireFormatLite::Int64Size(
        this->min_successes());
  }

  int cached_size = ::google::protobuf::internal::ToCachedSize(total_size);
  SetCachedSize(cached_size);
  return total_size;
}

void DynamicJobSpec::MergeFrom(const ::google::protobuf::Message& from) {
// @@protoc_insertion_point(generalized_merge_from_start:flyteidl.core.DynamicJobSpec)
  GOOGLE_DCHECK_NE(&from, this);
  const DynamicJobSpec* source =
      ::google::protobuf::DynamicCastToGenerated<DynamicJobSpec>(
          &from);
  if (source == nullptr) {
  // @@protoc_insertion_point(generalized_merge_from_cast_fail:flyteidl.core.DynamicJobSpec)
    ::google::protobuf::internal::ReflectionOps::Merge(from, this);
  } else {
  // @@protoc_insertion_point(generalized_merge_from_cast_success:flyteidl.core.DynamicJobSpec)
    MergeFrom(*source);
  }
}

void DynamicJobSpec::MergeFrom(const DynamicJobSpec& from) {
// @@protoc_insertion_point(class_specific_merge_from_start:flyteidl.core.DynamicJobSpec)
  GOOGLE_DCHECK_NE(&from, this);
  _internal_metadata_.MergeFrom(from._internal_metadata_);
  ::google::protobuf::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  nodes_.MergeFrom(from.nodes_);
  outputs_.MergeFrom(from.outputs_);
  tasks_.MergeFrom(from.tasks_);
  subworkflows_.MergeFrom(from.subworkflows_);
  if (from.min_successes() != 0) {
    set_min_successes(from.min_successes());
  }
}

void DynamicJobSpec::CopyFrom(const ::google::protobuf::Message& from) {
// @@protoc_insertion_point(generalized_copy_from_start:flyteidl.core.DynamicJobSpec)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

void DynamicJobSpec::CopyFrom(const DynamicJobSpec& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:flyteidl.core.DynamicJobSpec)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool DynamicJobSpec::IsInitialized() const {
  return true;
}

void DynamicJobSpec::Swap(DynamicJobSpec* other) {
  if (other == this) return;
  InternalSwap(other);
}
void DynamicJobSpec::InternalSwap(DynamicJobSpec* other) {
  using std::swap;
  _internal_metadata_.Swap(&other->_internal_metadata_);
  CastToBase(&nodes_)->InternalSwap(CastToBase(&other->nodes_));
  CastToBase(&outputs_)->InternalSwap(CastToBase(&other->outputs_));
  CastToBase(&tasks_)->InternalSwap(CastToBase(&other->tasks_));
  CastToBase(&subworkflows_)->InternalSwap(CastToBase(&other->subworkflows_));
  swap(min_successes_, other->min_successes_);
}

::google::protobuf::Metadata DynamicJobSpec::GetMetadata() const {
  ::google::protobuf::internal::AssignDescriptors(&::assign_descriptors_table_flyteidl_2fcore_2fdynamic_5fjob_2eproto);
  return ::file_level_metadata_flyteidl_2fcore_2fdynamic_5fjob_2eproto[kIndexInFileMessages];
}


// @@protoc_insertion_point(namespace_scope)
}  // namespace core
}  // namespace flyteidl
namespace google {
namespace protobuf {
template<> PROTOBUF_NOINLINE ::flyteidl::core::DynamicJobSpec* Arena::CreateMaybeMessage< ::flyteidl::core::DynamicJobSpec >(Arena* arena) {
  return Arena::CreateInternal< ::flyteidl::core::DynamicJobSpec >(arena);
}
}  // namespace protobuf
}  // namespace google

// @@protoc_insertion_point(global_scope)
#include <google/protobuf/port_undef.inc>
