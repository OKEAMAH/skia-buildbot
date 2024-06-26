// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.21.12
// source: sheriff_config.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AnomalyConfig_Step int32

const (
	// Step detection algorithm.
	AnomalyConfig_ORIGINAL_STEP AnomalyConfig_Step = 0
	// Step detection using absolute magnitude as threshold.
	AnomalyConfig_ABSOLUTE_STEP AnomalyConfig_Step = 1
	// Step detection using a constant as threshold.
	AnomalyConfig_CONST_STEP AnomalyConfig_Step = 2
	// Step detection that checks if step size is greater than some
	// percentage of the mean of the first half of the trace.
	AnomalyConfig_PERCENT_STEP AnomalyConfig_Step = 3
	// CohenStep uses Cohen's d method to detect a change.
	// https://en.wikipedia.org/wiki/Effect_size#Cohen's_d
	AnomalyConfig_COHEN_STEP AnomalyConfig_Step = 4
	// MannWhitneyU uses the Mann-Whitney U test to detect a change.
	// https://en.wikipedia.org/wiki/Mann%E2%80%93Whitney_U_test
	AnomalyConfig_MANN_WHITNEY_U AnomalyConfig_Step = 5
)

// Enum value maps for AnomalyConfig_Step.
var (
	AnomalyConfig_Step_name = map[int32]string{
		0: "ORIGINAL_STEP",
		1: "ABSOLUTE_STEP",
		2: "CONST_STEP",
		3: "PERCENT_STEP",
		4: "COHEN_STEP",
		5: "MANN_WHITNEY_U",
	}
	AnomalyConfig_Step_value = map[string]int32{
		"ORIGINAL_STEP":  0,
		"ABSOLUTE_STEP":  1,
		"CONST_STEP":     2,
		"PERCENT_STEP":   3,
		"COHEN_STEP":     4,
		"MANN_WHITNEY_U": 5,
	}
)

func (x AnomalyConfig_Step) Enum() *AnomalyConfig_Step {
	p := new(AnomalyConfig_Step)
	*p = x
	return p
}

func (x AnomalyConfig_Step) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AnomalyConfig_Step) Descriptor() protoreflect.EnumDescriptor {
	return file_sheriff_config_proto_enumTypes[0].Descriptor()
}

func (AnomalyConfig_Step) Type() protoreflect.EnumType {
	return &file_sheriff_config_proto_enumTypes[0]
}

func (x AnomalyConfig_Step) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AnomalyConfig_Step.Descriptor instead.
func (AnomalyConfig_Step) EnumDescriptor() ([]byte, []int) {
	return file_sheriff_config_proto_rawDescGZIP(), []int{2, 0}
}

// What actions should be taken for detected anomalies.
// - NOACTION: Only show anomaly in UI. Don't triage or bisect.
// - TRIAGE: File Buganizer issue for anomalies found. Don't bisect.
// - BISECT: Triage and bisect anomaly groups.
type AnomalyConfig_Action int32

const (
	AnomalyConfig_NOACTION AnomalyConfig_Action = 0
	AnomalyConfig_TRIAGE   AnomalyConfig_Action = 1
	AnomalyConfig_BISECT   AnomalyConfig_Action = 2
)

// Enum value maps for AnomalyConfig_Action.
var (
	AnomalyConfig_Action_name = map[int32]string{
		0: "NOACTION",
		1: "TRIAGE",
		2: "BISECT",
	}
	AnomalyConfig_Action_value = map[string]int32{
		"NOACTION": 0,
		"TRIAGE":   1,
		"BISECT":   2,
	}
)

func (x AnomalyConfig_Action) Enum() *AnomalyConfig_Action {
	p := new(AnomalyConfig_Action)
	*p = x
	return p
}

func (x AnomalyConfig_Action) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AnomalyConfig_Action) Descriptor() protoreflect.EnumDescriptor {
	return file_sheriff_config_proto_enumTypes[1].Descriptor()
}

func (AnomalyConfig_Action) Type() protoreflect.EnumType {
	return &file_sheriff_config_proto_enumTypes[1]
}

func (x AnomalyConfig_Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AnomalyConfig_Action.Descriptor instead.
func (AnomalyConfig_Action) EnumDescriptor() ([]byte, []int) {
	return file_sheriff_config_proto_rawDescGZIP(), []int{2, 1}
}

type Subscription_Priority int32

const (
	// If unspecified, default is P2.
	Subscription_P_UNSPECIFIED Subscription_Priority = 0
	Subscription_P0            Subscription_Priority = 1
	Subscription_P1            Subscription_Priority = 2
	Subscription_P2            Subscription_Priority = 3
	Subscription_P3            Subscription_Priority = 4
	Subscription_P4            Subscription_Priority = 5
)

// Enum value maps for Subscription_Priority.
var (
	Subscription_Priority_name = map[int32]string{
		0: "P_UNSPECIFIED",
		1: "P0",
		2: "P1",
		3: "P2",
		4: "P3",
		5: "P4",
	}
	Subscription_Priority_value = map[string]int32{
		"P_UNSPECIFIED": 0,
		"P0":            1,
		"P1":            2,
		"P2":            3,
		"P3":            4,
		"P4":            5,
	}
)

func (x Subscription_Priority) Enum() *Subscription_Priority {
	p := new(Subscription_Priority)
	*p = x
	return p
}

func (x Subscription_Priority) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Subscription_Priority) Descriptor() protoreflect.EnumDescriptor {
	return file_sheriff_config_proto_enumTypes[2].Descriptor()
}

func (Subscription_Priority) Type() protoreflect.EnumType {
	return &file_sheriff_config_proto_enumTypes[2]
}

func (x Subscription_Priority) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Subscription_Priority.Descriptor instead.
func (Subscription_Priority) EnumDescriptor() ([]byte, []int) {
	return file_sheriff_config_proto_rawDescGZIP(), []int{3, 0}
}

type Subscription_Severity int32

const (
	// If unspecified, default is S2.
	Subscription_S_UNSPECIFIED Subscription_Severity = 0
	Subscription_S0            Subscription_Severity = 1
	Subscription_S1            Subscription_Severity = 2
	Subscription_S2            Subscription_Severity = 3
	Subscription_S3            Subscription_Severity = 4
	Subscription_S4            Subscription_Severity = 5
)

// Enum value maps for Subscription_Severity.
var (
	Subscription_Severity_name = map[int32]string{
		0: "S_UNSPECIFIED",
		1: "S0",
		2: "S1",
		3: "S2",
		4: "S3",
		5: "S4",
	}
	Subscription_Severity_value = map[string]int32{
		"S_UNSPECIFIED": 0,
		"S0":            1,
		"S1":            2,
		"S2":            3,
		"S3":            4,
		"S4":            5,
	}
)

func (x Subscription_Severity) Enum() *Subscription_Severity {
	p := new(Subscription_Severity)
	*p = x
	return p
}

func (x Subscription_Severity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Subscription_Severity) Descriptor() protoreflect.EnumDescriptor {
	return file_sheriff_config_proto_enumTypes[3].Descriptor()
}

func (Subscription_Severity) Type() protoreflect.EnumType {
	return &file_sheriff_config_proto_enumTypes[3]
}

func (x Subscription_Severity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Subscription_Severity.Descriptor instead.
func (Subscription_Severity) EnumDescriptor() ([]byte, []int) {
	return file_sheriff_config_proto_rawDescGZIP(), []int{3, 1}
}

// A Pattern message defines regular expression patterns for capturing
// a group of metrics. A metric is uniquely identified by the
// combination of all the keys specified within a Pattern.
// To specify that a value is a Regex, a "~" must be added at the beginning
// of the string.
type Pattern struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Main      string `protobuf:"bytes,1,opt,name=main,proto3" json:"main,omitempty"`           // e.g. "ChromiumPerf", "Chromium*"
	Bot       string `protobuf:"bytes,2,opt,name=bot,proto3" json:"bot,omitempty"`             // e.g. "linux-perf", "~lacros-.*"
	Benchmark string `protobuf:"bytes,3,opt,name=benchmark,proto3" json:"benchmark,omitempty"` // e.g. "Speedometer2"
	Test      string `protobuf:"bytes,4,opt,name=test,proto3" json:"test,omitempty"`           // e.g. "speedometer2"
	Subtest1  string `protobuf:"bytes,5,opt,name=subtest1,proto3" json:"subtest1,omitempty"`
	Subtest2  string `protobuf:"bytes,6,opt,name=subtest2,proto3" json:"subtest2,omitempty"`
	Subtest3  string `protobuf:"bytes,7,opt,name=subtest3,proto3" json:"subtest3,omitempty"`
}

func (x *Pattern) Reset() {
	*x = Pattern{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sheriff_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pattern) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pattern) ProtoMessage() {}

func (x *Pattern) ProtoReflect() protoreflect.Message {
	mi := &file_sheriff_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pattern.ProtoReflect.Descriptor instead.
func (*Pattern) Descriptor() ([]byte, []int) {
	return file_sheriff_config_proto_rawDescGZIP(), []int{0}
}

func (x *Pattern) GetMain() string {
	if x != nil {
		return x.Main
	}
	return ""
}

func (x *Pattern) GetBot() string {
	if x != nil {
		return x.Bot
	}
	return ""
}

func (x *Pattern) GetBenchmark() string {
	if x != nil {
		return x.Benchmark
	}
	return ""
}

func (x *Pattern) GetTest() string {
	if x != nil {
		return x.Test
	}
	return ""
}

func (x *Pattern) GetSubtest1() string {
	if x != nil {
		return x.Subtest1
	}
	return ""
}

func (x *Pattern) GetSubtest2() string {
	if x != nil {
		return x.Subtest2
	}
	return ""
}

func (x *Pattern) GetSubtest3() string {
	if x != nil {
		return x.Subtest3
	}
	return ""
}

// We can use patterns to specify which metrics we want to include or exclude.
//
// For matching, if a Pattern field is not specified, the default is wildcard "*",
// meaning match to any value.
// For excluding, only filter on specified Pattern fields. Exclude patterns are
// only allowed to have one field specified.
//
// Consider the example below:
//
//	{
//	  match: [
//	    {main:"ChromiumPerf",bot:"~lacros-.*-perf",benchmark:"Speedometer2"},
//	    {main:"ChromiumPerf",benchmark:"Jetstream2"},
//	  ],
//	  exclude: [
//	    {bot:"lacros-eve-perf"},
//	    {bot:"lacros-x86-perf"},
//	  ]
//	}
//
// In SQL grammar, this would translate to:
// ...
// SELECT * FROM Metrics
// WHERE
// (main='ChromiumPerf' AND bot REGEXP 'lacros-.*-perf' AND benchmark='Speedometer2'
// AND bot!='lacros-eve-perf' AND bot!='lacros-x86-perf')
// OR
// (main='ChromiumPerf' AND benchmark='Jetstream')
// AND bot!='lacros-eve-perf' AND bot!='lacros-x86-perf')
type Rules struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Match   []*Pattern `protobuf:"bytes,1,rep,name=match,proto3" json:"match,omitempty"`
	Exclude []*Pattern `protobuf:"bytes,2,rep,name=exclude,proto3" json:"exclude,omitempty"`
}

func (x *Rules) Reset() {
	*x = Rules{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sheriff_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rules) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rules) ProtoMessage() {}

func (x *Rules) ProtoReflect() protoreflect.Message {
	mi := &file_sheriff_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rules.ProtoReflect.Descriptor instead.
func (*Rules) Descriptor() ([]byte, []int) {
	return file_sheriff_config_proto_rawDescGZIP(), []int{1}
}

func (x *Rules) GetMatch() []*Pattern {
	if x != nil {
		return x.Match
	}
	return nil
}

func (x *Rules) GetExclude() []*Pattern {
	if x != nil {
		return x.Exclude
	}
	return nil
}

// An AnomalyConfig defines the bounds for which a change in a matching metric
// can be considered "anomalous". For metrics that are matched, we apply the
// anomaly config to determine whether we should create an alert.
//
// The configuration settings defined for an anomaly configuration override
// defaults that are defined by the anomaly detection algorithm.
//
// TODO(eduardoyap): Figure out default values and document them here.
type AnomalyConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Anomaly detection algorithm.
	Step AnomalyConfig_Step `protobuf:"varint,1,opt,name=step,proto3,enum=sheriff_config.v1.AnomalyConfig_Step" json:"step,omitempty"`
	// How many commits to each side of a commit to consider when looking for a step.
	Radius int32 `protobuf:"varint,2,opt,name=radius,proto3" json:"radius,omitempty"`
	// The threshold value beyond which values become interesting
	// (indicates a real regression). Range of this value depends on algorithm used.
	Threshold float32              `protobuf:"fixed32,3,opt,name=threshold,proto3" json:"threshold,omitempty"`
	Action    AnomalyConfig_Action `protobuf:"varint,4,opt,name=action,proto3,enum=sheriff_config.v1.AnomalyConfig_Action" json:"action,omitempty"`
	// Which metrics should be captured by this AnomalyConfig.
	Rules *Rules `protobuf:"bytes,5,opt,name=rules,proto3" json:"rules,omitempty"`
}

func (x *AnomalyConfig) Reset() {
	*x = AnomalyConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sheriff_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnomalyConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnomalyConfig) ProtoMessage() {}

func (x *AnomalyConfig) ProtoReflect() protoreflect.Message {
	mi := &file_sheriff_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnomalyConfig.ProtoReflect.Descriptor instead.
func (*AnomalyConfig) Descriptor() ([]byte, []int) {
	return file_sheriff_config_proto_rawDescGZIP(), []int{2}
}

func (x *AnomalyConfig) GetStep() AnomalyConfig_Step {
	if x != nil {
		return x.Step
	}
	return AnomalyConfig_ORIGINAL_STEP
}

func (x *AnomalyConfig) GetRadius() int32 {
	if x != nil {
		return x.Radius
	}
	return 0
}

func (x *AnomalyConfig) GetThreshold() float32 {
	if x != nil {
		return x.Threshold
	}
	return 0
}

func (x *AnomalyConfig) GetAction() AnomalyConfig_Action {
	if x != nil {
		return x.Action
	}
	return AnomalyConfig_NOACTION
}

func (x *AnomalyConfig) GetRules() *Rules {
	if x != nil {
		return x.Rules
	}
	return nil
}

// A Subscription describes a configuration through which we determine:
//   - A set of metrics a group of users are interested in alert monitoring.
//     These anomalies are grouped together into anomaly groups if they
//     overlap.
//   - Anomaly detection settings.
//   - Alerting settings.
type Subscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A name is a free-form name for human readability purposes. Also
	// serves as a unique key for the subscription and should be unique from
	// all other subscription names.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The contact email address who owns this subscription. This is a required
	// field. There must be an owner we can contact for each subscription.
	ContactEmail string `protobuf:"bytes,2,opt,name=contact_email,json=contactEmail,proto3" json:"contact_email,omitempty"`
	// A list of labels applied to the Buganizer issues associated with
	// this subscription.
	BugLabels []string `protobuf:"bytes,3,rep,name=bug_labels,json=bugLabels,proto3" json:"bug_labels,omitempty"`
	// A list of Hotlist labels applied to the Buganizer issues associated with
	// this subscription.
	HotlistLabels []string `protobuf:"bytes,4,rep,name=hotlist_labels,json=hotlistLabels,proto3" json:"hotlist_labels,omitempty"`
	// A Buganizer component in which to file issues for this subscription.
	BugComponent string `protobuf:"bytes,6,opt,name=bug_component,json=bugComponent,proto3" json:"bug_component,omitempty"`
	// Priority to set in Buganizer issue. Default is P2.
	BugPriority Subscription_Priority `protobuf:"varint,9,opt,name=bug_priority,json=bugPriority,proto3,enum=sheriff_config.v1.Subscription_Priority" json:"bug_priority,omitempty"`
	// Severity to set in Buganizer issue. Default is S2.
	BugSeverity Subscription_Severity `protobuf:"varint,10,opt,name=bug_severity,json=bugSeverity,proto3,enum=sheriff_config.v1.Subscription_Severity" json:"bug_severity,omitempty"`
	// A list of e-mails to add to Buganizer issue CC list.
	BugCcEmails []string `protobuf:"bytes,7,rep,name=bug_cc_emails,json=bugCcEmails,proto3" json:"bug_cc_emails,omitempty"`
	// Here we specify the subset of metrics we are interested in and what anomaly
	// detection algorithms to apply. This field can be repeated so that
	// different algorithms can be applied depending on the metrics captured.
	//
	// Anomaly configs in the same subscription should not have
	// overlapping metrics.
	AnomalyConfigs []*AnomalyConfig `protobuf:"bytes,8,rep,name=anomaly_configs,json=anomalyConfigs,proto3" json:"anomaly_configs,omitempty"`
}

func (x *Subscription) Reset() {
	*x = Subscription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sheriff_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subscription) ProtoMessage() {}

func (x *Subscription) ProtoReflect() protoreflect.Message {
	mi := &file_sheriff_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subscription.ProtoReflect.Descriptor instead.
func (*Subscription) Descriptor() ([]byte, []int) {
	return file_sheriff_config_proto_rawDescGZIP(), []int{3}
}

func (x *Subscription) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Subscription) GetContactEmail() string {
	if x != nil {
		return x.ContactEmail
	}
	return ""
}

func (x *Subscription) GetBugLabels() []string {
	if x != nil {
		return x.BugLabels
	}
	return nil
}

func (x *Subscription) GetHotlistLabels() []string {
	if x != nil {
		return x.HotlistLabels
	}
	return nil
}

func (x *Subscription) GetBugComponent() string {
	if x != nil {
		return x.BugComponent
	}
	return ""
}

func (x *Subscription) GetBugPriority() Subscription_Priority {
	if x != nil {
		return x.BugPriority
	}
	return Subscription_P_UNSPECIFIED
}

func (x *Subscription) GetBugSeverity() Subscription_Severity {
	if x != nil {
		return x.BugSeverity
	}
	return Subscription_S_UNSPECIFIED
}

func (x *Subscription) GetBugCcEmails() []string {
	if x != nil {
		return x.BugCcEmails
	}
	return nil
}

func (x *Subscription) GetAnomalyConfigs() []*AnomalyConfig {
	if x != nil {
		return x.AnomalyConfigs
	}
	return nil
}

// A SheriffConfig lists the subscriptions for a Skia Perf instance.
// Subscriptions may only capture metrics which are uploaded to the Skia Perf
// instance specified.
type SheriffConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subscriptions []*Subscription `protobuf:"bytes,1,rep,name=subscriptions,proto3" json:"subscriptions,omitempty"`
}

func (x *SheriffConfig) Reset() {
	*x = SheriffConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sheriff_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SheriffConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SheriffConfig) ProtoMessage() {}

func (x *SheriffConfig) ProtoReflect() protoreflect.Message {
	mi := &file_sheriff_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SheriffConfig.ProtoReflect.Descriptor instead.
func (*SheriffConfig) Descriptor() ([]byte, []int) {
	return file_sheriff_config_proto_rawDescGZIP(), []int{4}
}

func (x *SheriffConfig) GetSubscriptions() []*Subscription {
	if x != nil {
		return x.Subscriptions
	}
	return nil
}

var File_sheriff_config_proto protoreflect.FileDescriptor

var file_sheriff_config_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x68, 0x65, 0x72, 0x69, 0x66, 0x66, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x73, 0x68, 0x65, 0x72, 0x69, 0x66, 0x66, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x22, 0xb5, 0x01, 0x0a, 0x07, 0x50, 0x61,
	0x74, 0x74, 0x65, 0x72, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x6f, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x62, 0x6f, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x62,
	0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x73,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x73, 0x75, 0x62, 0x74, 0x65, 0x73, 0x74, 0x31, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x73, 0x75, 0x62, 0x74, 0x65, 0x73, 0x74, 0x31, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x75, 0x62,
	0x74, 0x65, 0x73, 0x74, 0x32, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x75, 0x62,
	0x74, 0x65, 0x73, 0x74, 0x32, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x74, 0x65, 0x73, 0x74,
	0x33, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x75, 0x62, 0x74, 0x65, 0x73, 0x74,
	0x33, 0x22, 0x6f, 0x0a, 0x05, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x05, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x68, 0x65, 0x72,
	0x69, 0x66, 0x66, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61,
	0x74, 0x74, 0x65, 0x72, 0x6e, 0x52, 0x05, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x34, 0x0a, 0x07,
	0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x73, 0x68, 0x65, 0x72, 0x69, 0x66, 0x66, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x52, 0x07, 0x65, 0x78, 0x63, 0x6c, 0x75,
	0x64, 0x65, 0x22, 0x95, 0x03, 0x0a, 0x0d, 0x41, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x39, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x25, 0x2e, 0x73, 0x68, 0x65, 0x72, 0x69, 0x66, 0x66, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x53, 0x74, 0x65, 0x70, 0x52, 0x04, 0x73, 0x74, 0x65, 0x70, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x68, 0x72, 0x65, 0x73,
	0x68, 0x6f, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x74, 0x68, 0x72, 0x65,
	0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x3f, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x73, 0x68, 0x65, 0x72, 0x69, 0x66, 0x66, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x6f, 0x6d, 0x61, 0x6c,
	0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x68, 0x65, 0x72, 0x69, 0x66, 0x66, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x52,
	0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x22, 0x72, 0x0a, 0x04, 0x53, 0x74, 0x65, 0x70, 0x12, 0x11,
	0x0a, 0x0d, 0x4f, 0x52, 0x49, 0x47, 0x49, 0x4e, 0x41, 0x4c, 0x5f, 0x53, 0x54, 0x45, 0x50, 0x10,
	0x00, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x42, 0x53, 0x4f, 0x4c, 0x55, 0x54, 0x45, 0x5f, 0x53, 0x54,
	0x45, 0x50, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x4f, 0x4e, 0x53, 0x54, 0x5f, 0x53, 0x54,
	0x45, 0x50, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x45, 0x52, 0x43, 0x45, 0x4e, 0x54, 0x5f,
	0x53, 0x54, 0x45, 0x50, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x4f, 0x48, 0x45, 0x4e, 0x5f,
	0x53, 0x54, 0x45, 0x50, 0x10, 0x04, 0x12, 0x12, 0x0a, 0x0e, 0x4d, 0x41, 0x4e, 0x4e, 0x5f, 0x57,
	0x48, 0x49, 0x54, 0x4e, 0x45, 0x59, 0x5f, 0x55, 0x10, 0x05, 0x22, 0x2e, 0x0a, 0x06, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x0a, 0x08, 0x4e, 0x4f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e,
	0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x54, 0x52, 0x49, 0x41, 0x47, 0x45, 0x10, 0x01, 0x12, 0x0a,
	0x0a, 0x06, 0x42, 0x49, 0x53, 0x45, 0x43, 0x54, 0x10, 0x02, 0x22, 0xc9, 0x04, 0x0a, 0x0c, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x75, 0x67, 0x5f, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x62, 0x75, 0x67, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x68, 0x6f, 0x74, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x68, 0x6f, 0x74,
	0x6c, 0x69, 0x73, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x75,
	0x67, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x62, 0x75, 0x67, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12,
	0x4b, 0x0a, 0x0c, 0x62, 0x75, 0x67, 0x5f, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x73, 0x68, 0x65, 0x72, 0x69, 0x66, 0x66, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x52,
	0x0b, 0x62, 0x75, 0x67, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x4b, 0x0a, 0x0c,
	0x62, 0x75, 0x67, 0x5f, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x28, 0x2e, 0x73, 0x68, 0x65, 0x72, 0x69, 0x66, 0x66, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x52, 0x0b, 0x62, 0x75,
	0x67, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x22, 0x0a, 0x0d, 0x62, 0x75, 0x67,
	0x5f, 0x63, 0x63, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0b, 0x62, 0x75, 0x67, 0x43, 0x63, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x49, 0x0a,
	0x0f, 0x61, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73,
	0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x68, 0x65, 0x72, 0x69, 0x66, 0x66,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x6f, 0x6d, 0x61,
	0x6c, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0e, 0x61, 0x6e, 0x6f, 0x6d, 0x61, 0x6c,
	0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x22, 0x45, 0x0a, 0x08, 0x50, 0x72, 0x69, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x12, 0x11, 0x0a, 0x0d, 0x50, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x50, 0x30, 0x10, 0x01, 0x12,
	0x06, 0x0a, 0x02, 0x50, 0x31, 0x10, 0x02, 0x12, 0x06, 0x0a, 0x02, 0x50, 0x32, 0x10, 0x03, 0x12,
	0x06, 0x0a, 0x02, 0x50, 0x33, 0x10, 0x04, 0x12, 0x06, 0x0a, 0x02, 0x50, 0x34, 0x10, 0x05, 0x22,
	0x45, 0x0a, 0x08, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x11, 0x0a, 0x0d, 0x53,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x06,
	0x0a, 0x02, 0x53, 0x30, 0x10, 0x01, 0x12, 0x06, 0x0a, 0x02, 0x53, 0x31, 0x10, 0x02, 0x12, 0x06,
	0x0a, 0x02, 0x53, 0x32, 0x10, 0x03, 0x12, 0x06, 0x0a, 0x02, 0x53, 0x33, 0x10, 0x04, 0x12, 0x06,
	0x0a, 0x02, 0x53, 0x34, 0x10, 0x05, 0x22, 0x56, 0x0a, 0x0d, 0x53, 0x68, 0x65, 0x72, 0x69, 0x66,
	0x66, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x45, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x73, 0x68, 0x65, 0x72, 0x69, 0x66, 0x66, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x32,
	0x5a, 0x30, 0x67, 0x6f, 0x2e, 0x73, 0x6b, 0x69, 0x61, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x2f, 0x70, 0x65, 0x72, 0x66, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x68, 0x65, 0x72,
	0x69, 0x66, 0x66, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sheriff_config_proto_rawDescOnce sync.Once
	file_sheriff_config_proto_rawDescData = file_sheriff_config_proto_rawDesc
)

func file_sheriff_config_proto_rawDescGZIP() []byte {
	file_sheriff_config_proto_rawDescOnce.Do(func() {
		file_sheriff_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_sheriff_config_proto_rawDescData)
	})
	return file_sheriff_config_proto_rawDescData
}

var file_sheriff_config_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_sheriff_config_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_sheriff_config_proto_goTypes = []interface{}{
	(AnomalyConfig_Step)(0),    // 0: sheriff_config.v1.AnomalyConfig.Step
	(AnomalyConfig_Action)(0),  // 1: sheriff_config.v1.AnomalyConfig.Action
	(Subscription_Priority)(0), // 2: sheriff_config.v1.Subscription.Priority
	(Subscription_Severity)(0), // 3: sheriff_config.v1.Subscription.Severity
	(*Pattern)(nil),            // 4: sheriff_config.v1.Pattern
	(*Rules)(nil),              // 5: sheriff_config.v1.Rules
	(*AnomalyConfig)(nil),      // 6: sheriff_config.v1.AnomalyConfig
	(*Subscription)(nil),       // 7: sheriff_config.v1.Subscription
	(*SheriffConfig)(nil),      // 8: sheriff_config.v1.SheriffConfig
}
var file_sheriff_config_proto_depIdxs = []int32{
	4, // 0: sheriff_config.v1.Rules.match:type_name -> sheriff_config.v1.Pattern
	4, // 1: sheriff_config.v1.Rules.exclude:type_name -> sheriff_config.v1.Pattern
	0, // 2: sheriff_config.v1.AnomalyConfig.step:type_name -> sheriff_config.v1.AnomalyConfig.Step
	1, // 3: sheriff_config.v1.AnomalyConfig.action:type_name -> sheriff_config.v1.AnomalyConfig.Action
	5, // 4: sheriff_config.v1.AnomalyConfig.rules:type_name -> sheriff_config.v1.Rules
	2, // 5: sheriff_config.v1.Subscription.bug_priority:type_name -> sheriff_config.v1.Subscription.Priority
	3, // 6: sheriff_config.v1.Subscription.bug_severity:type_name -> sheriff_config.v1.Subscription.Severity
	6, // 7: sheriff_config.v1.Subscription.anomaly_configs:type_name -> sheriff_config.v1.AnomalyConfig
	7, // 8: sheriff_config.v1.SheriffConfig.subscriptions:type_name -> sheriff_config.v1.Subscription
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_sheriff_config_proto_init() }
func file_sheriff_config_proto_init() {
	if File_sheriff_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sheriff_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pattern); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sheriff_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rules); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sheriff_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnomalyConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sheriff_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subscription); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sheriff_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SheriffConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sheriff_config_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sheriff_config_proto_goTypes,
		DependencyIndexes: file_sheriff_config_proto_depIdxs,
		EnumInfos:         file_sheriff_config_proto_enumTypes,
		MessageInfos:      file_sheriff_config_proto_msgTypes,
	}.Build()
	File_sheriff_config_proto = out.File
	file_sheriff_config_proto_rawDesc = nil
	file_sheriff_config_proto_goTypes = nil
	file_sheriff_config_proto_depIdxs = nil
}
