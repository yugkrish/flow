// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: model.proto

/*
Package model is a generated protocol buffer package.

It is generated from these files:
	model.proto

It has these top-level messages:
	CompletionResult
	BlobDatum
	HTTPHeader
	HTTPReqDatum
	HTTPRespDatum
	EmptyDatum
	StageRefDatum
	ErrorDatum
	StateDatum
	Datum
	AddStageRequest
	CompleteStageExternallyRequest
	AddCompletedValueStageRequest
	AddDelayStageRequest
	AddInvokeFunctionStageRequest
	AddStageResponse
	CommitGraphRequest
	GraphRequestProcessedResponse
	CompleteDelayStageRequest
	CompleteStageExternallyResponse
	DeactivateGraphRequest
	CreateGraphRequest
	CreateGraphResponse
	FaasInvocationResponse
	GetGraphStateRequest
	GetGraphStateResponse
	ListGraphsRequest
	StreamLifecycle
	StreamGraph
	StreamRequest
	ListGraphResponse
	ListGraphsResponse
	AwaitStageResultRequest
	AwaitStageResultResponse
	InvalidGraphOperation
	InvalidStageOperation
	InvokeFunctionRequest
	InvokeStageRequest
	GraphEvent
	GraphCreatedEvent
	DelayScheduledEvent
	GraphTerminatingEvent
	GraphCompletedEvent
	GraphCommittedEvent
	StageAddedEvent
	StageCompletedEvent
	StageComposedEvent
	FaasInvocationStartedEvent
	FaasInvocationCompletedEvent
	RuntimeInvokeStageRequest
	RuntimeInvokeStageResponse
*/
package model

import regexp "regexp"
import fmt "fmt"
import go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import math "math"
import _ "github.com/golang/protobuf/ptypes/timestamp"
import _ "github.com/mwitkow/go-proto-validators"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *CompletionResult) Validate() error {
	if nil == this.Datum {
		return go_proto_validators.FieldError("Datum", fmt.Errorf("message must exist"))
	}
	if this.Datum != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Datum); err != nil {
			return go_proto_validators.FieldError("Datum", err)
		}
	}
	return nil
}
func (this *BlobDatum) Validate() error {
	if this.BlobId == "" {
		return go_proto_validators.FieldError("BlobId", fmt.Errorf(`value '%v' must not be an empty string`, this.BlobId))
	}
	if this.ContentType == "" {
		return go_proto_validators.FieldError("ContentType", fmt.Errorf(`value '%v' must not be an empty string`, this.ContentType))
	}
	return nil
}
func (this *HTTPHeader) Validate() error {
	if this.Key == "" {
		return go_proto_validators.FieldError("Key", fmt.Errorf(`value '%v' must not be an empty string`, this.Key))
	}
	if this.Value == "" {
		return go_proto_validators.FieldError("Value", fmt.Errorf(`value '%v' must not be an empty string`, this.Value))
	}
	return nil
}
func (this *HTTPReqDatum) Validate() error {
	if this.Body != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Body); err != nil {
			return go_proto_validators.FieldError("Body", err)
		}
	}
	for _, item := range this.Headers {
		if item != nil {
			if err := go_proto_validators.CallValidatorIfExists(item); err != nil {
				return go_proto_validators.FieldError("Headers", err)
			}
		}
	}
	return nil
}
func (this *HTTPRespDatum) Validate() error {
	if this.Body != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Body); err != nil {
			return go_proto_validators.FieldError("Body", err)
		}
	}
	for _, item := range this.Headers {
		if item != nil {
			if err := go_proto_validators.CallValidatorIfExists(item); err != nil {
				return go_proto_validators.FieldError("Headers", err)
			}
		}
	}
	if !(this.StatusCode > 0) {
		return go_proto_validators.FieldError("StatusCode", fmt.Errorf(`value '%v' must be greater than '0'`, this.StatusCode))
	}
	return nil
}
func (this *EmptyDatum) Validate() error {
	return nil
}
func (this *StageRefDatum) Validate() error {
	if this.StageId == "" {
		return go_proto_validators.FieldError("StageId", fmt.Errorf(`value '%v' must not be an empty string`, this.StageId))
	}
	return nil
}
func (this *ErrorDatum) Validate() error {
	return nil
}
func (this *StateDatum) Validate() error {
	return nil
}
func (this *Datum) Validate() error {
	if oneOfNester, ok := this.GetVal().(*Datum_Empty); ok {
		if oneOfNester.Empty != nil {
			if err := go_proto_validators.CallValidatorIfExists(oneOfNester.Empty); err != nil {
				return go_proto_validators.FieldError("Empty", err)
			}
		}
	}
	if oneOfNester, ok := this.GetVal().(*Datum_Blob); ok {
		if oneOfNester.Blob != nil {
			if err := go_proto_validators.CallValidatorIfExists(oneOfNester.Blob); err != nil {
				return go_proto_validators.FieldError("Blob", err)
			}
		}
	}
	if oneOfNester, ok := this.GetVal().(*Datum_Error); ok {
		if oneOfNester.Error != nil {
			if err := go_proto_validators.CallValidatorIfExists(oneOfNester.Error); err != nil {
				return go_proto_validators.FieldError("Error", err)
			}
		}
	}
	if oneOfNester, ok := this.GetVal().(*Datum_StageRef); ok {
		if oneOfNester.StageRef != nil {
			if err := go_proto_validators.CallValidatorIfExists(oneOfNester.StageRef); err != nil {
				return go_proto_validators.FieldError("StageRef", err)
			}
		}
	}
	if oneOfNester, ok := this.GetVal().(*Datum_HttpReq); ok {
		if oneOfNester.HttpReq != nil {
			if err := go_proto_validators.CallValidatorIfExists(oneOfNester.HttpReq); err != nil {
				return go_proto_validators.FieldError("HttpReq", err)
			}
		}
	}
	if oneOfNester, ok := this.GetVal().(*Datum_HttpResp); ok {
		if oneOfNester.HttpResp != nil {
			if err := go_proto_validators.CallValidatorIfExists(oneOfNester.HttpResp); err != nil {
				return go_proto_validators.FieldError("HttpResp", err)
			}
		}
	}
	if oneOfNester, ok := this.GetVal().(*Datum_State); ok {
		if oneOfNester.State != nil {
			if err := go_proto_validators.CallValidatorIfExists(oneOfNester.State); err != nil {
				return go_proto_validators.FieldError("State", err)
			}
		}
	}
	return nil
}
func (this *AddStageRequest) Validate() error {
	if this.FlowId == "" {
		return go_proto_validators.FieldError("FlowId", fmt.Errorf(`value '%v' must not be an empty string`, this.FlowId))
	}
	if this.Closure != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Closure); err != nil {
			return go_proto_validators.FieldError("Closure", err)
		}
	}
	return nil
}
func (this *CompleteStageExternallyRequest) Validate() error {
	if this.FlowId == "" {
		return go_proto_validators.FieldError("FlowId", fmt.Errorf(`value '%v' must not be an empty string`, this.FlowId))
	}
	if this.StageId == "" {
		return go_proto_validators.FieldError("StageId", fmt.Errorf(`value '%v' must not be an empty string`, this.StageId))
	}
	if nil == this.Value {
		return go_proto_validators.FieldError("Value", fmt.Errorf("message must exist"))
	}
	if this.Value != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Value); err != nil {
			return go_proto_validators.FieldError("Value", err)
		}
	}
	return nil
}
func (this *AddCompletedValueStageRequest) Validate() error {
	if this.FlowId == "" {
		return go_proto_validators.FieldError("FlowId", fmt.Errorf(`value '%v' must not be an empty string`, this.FlowId))
	}
	if nil == this.Value {
		return go_proto_validators.FieldError("Value", fmt.Errorf("message must exist"))
	}
	if this.Value != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Value); err != nil {
			return go_proto_validators.FieldError("Value", err)
		}
	}
	return nil
}
func (this *AddDelayStageRequest) Validate() error {
	if this.FlowId == "" {
		return go_proto_validators.FieldError("FlowId", fmt.Errorf(`value '%v' must not be an empty string`, this.FlowId))
	}
	if !(this.DelayMs > -1) {
		return go_proto_validators.FieldError("DelayMs", fmt.Errorf(`value '%v' must be greater than '-1'`, this.DelayMs))
	}
	return nil
}

var _regex_AddInvokeFunctionStageRequest_FunctionId = regexp.MustCompile("^(\\.|[a-zA-Z0-9_\\-]{1,255})(/[a-zA-Z0-9_/\\-]{0,255})(\\?.*)?$")

func (this *AddInvokeFunctionStageRequest) Validate() error {
	if this.FlowId == "" {
		return go_proto_validators.FieldError("FlowId", fmt.Errorf(`value '%v' must not be an empty string`, this.FlowId))
	}
	if !_regex_AddInvokeFunctionStageRequest_FunctionId.MatchString(this.FunctionId) {
		return go_proto_validators.FieldError("FunctionId", fmt.Errorf(`value '%v' must be a string conforming to regex "^(\\.|[a-zA-Z0-9_\\-]{1,255})(/[a-zA-Z0-9_/\\-]{0,255})(\\?.*)?$"`, this.FunctionId))
	}
	if nil == this.Arg {
		return go_proto_validators.FieldError("Arg", fmt.Errorf("message must exist"))
	}
	if this.Arg != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Arg); err != nil {
			return go_proto_validators.FieldError("Arg", err)
		}
	}
	return nil
}
func (this *AddStageResponse) Validate() error {
	return nil
}
func (this *CommitGraphRequest) Validate() error {
	if this.FlowId == "" {
		return go_proto_validators.FieldError("FlowId", fmt.Errorf(`value '%v' must not be an empty string`, this.FlowId))
	}
	return nil
}
func (this *GraphRequestProcessedResponse) Validate() error {
	return nil
}
func (this *CompleteDelayStageRequest) Validate() error {
	if this.FlowId == "" {
		return go_proto_validators.FieldError("FlowId", fmt.Errorf(`value '%v' must not be an empty string`, this.FlowId))
	}
	if this.StageId == "" {
		return go_proto_validators.FieldError("StageId", fmt.Errorf(`value '%v' must not be an empty string`, this.StageId))
	}
	if this.Result != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Result); err != nil {
			return go_proto_validators.FieldError("Result", err)
		}
	}
	return nil
}
func (this *CompleteStageExternallyResponse) Validate() error {
	return nil
}
func (this *DeactivateGraphRequest) Validate() error {
	if this.FlowId == "" {
		return go_proto_validators.FieldError("FlowId", fmt.Errorf(`value '%v' must not be an empty string`, this.FlowId))
	}
	return nil
}

var _regex_CreateGraphRequest_FunctionId = regexp.MustCompile("^([a-zA-Z0-9_\\-]{1,255})(/[a-zA-Z0-9_/\\-]{0,255})(\\?.*)?$")

func (this *CreateGraphRequest) Validate() error {
	if !_regex_CreateGraphRequest_FunctionId.MatchString(this.FunctionId) {
		return go_proto_validators.FieldError("FunctionId", fmt.Errorf(`function id must be an absolute function of form "app/route"`))
	}
	return nil
}
func (this *CreateGraphResponse) Validate() error {
	return nil
}
func (this *FaasInvocationResponse) Validate() error {
	if this.Result != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Result); err != nil {
			return go_proto_validators.FieldError("Result", err)
		}
	}
	return nil
}
func (this *GetGraphStateRequest) Validate() error {
	if this.FlowId == "" {
		return go_proto_validators.FieldError("FlowId", fmt.Errorf(`value '%v' must not be an empty string`, this.FlowId))
	}
	return nil
}
func (this *GetGraphStateResponse) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *GetGraphStateResponse_StageRepresentation) Validate() error {
	return nil
}
func (this *ListGraphsRequest) Validate() error {
	return nil
}
func (this *StreamLifecycle) Validate() error {
	return nil
}
func (this *StreamGraph) Validate() error {
	return nil
}
func (this *StreamRequest) Validate() error {
	if oneOfNester, ok := this.GetQuery().(*StreamRequest_Lifecycle); ok {
		if oneOfNester.Lifecycle != nil {
			if err := go_proto_validators.CallValidatorIfExists(oneOfNester.Lifecycle); err != nil {
				return go_proto_validators.FieldError("Lifecycle", err)
			}
		}
	}
	if oneOfNester, ok := this.GetQuery().(*StreamRequest_Graph); ok {
		if oneOfNester.Graph != nil {
			if err := go_proto_validators.CallValidatorIfExists(oneOfNester.Graph); err != nil {
				return go_proto_validators.FieldError("Graph", err)
			}
		}
	}
	return nil
}
func (this *ListGraphResponse) Validate() error {
	return nil
}
func (this *ListGraphsResponse) Validate() error {
	for _, item := range this.Graphs {
		if item != nil {
			if err := go_proto_validators.CallValidatorIfExists(item); err != nil {
				return go_proto_validators.FieldError("Graphs", err)
			}
		}
	}
	return nil
}
func (this *AwaitStageResultRequest) Validate() error {
	if this.FlowId == "" {
		return go_proto_validators.FieldError("FlowId", fmt.Errorf(`value '%v' must not be an empty string`, this.FlowId))
	}
	if this.StageId == "" {
		return go_proto_validators.FieldError("StageId", fmt.Errorf(`value '%v' must not be an empty string`, this.StageId))
	}
	return nil
}
func (this *AwaitStageResultResponse) Validate() error {
	if this.Result != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Result); err != nil {
			return go_proto_validators.FieldError("Result", err)
		}
	}
	return nil
}
func (this *InvalidGraphOperation) Validate() error {
	return nil
}
func (this *InvalidStageOperation) Validate() error {
	return nil
}

var _regex_InvokeFunctionRequest_FunctionId = regexp.MustCompile("^(\\.|[a-zA-Z0-9_\\-]{1,255})(/[a-zA-Z0-9_/\\-]{0,255})(\\?.*)?$")

func (this *InvokeFunctionRequest) Validate() error {
	if this.FlowId == "" {
		return go_proto_validators.FieldError("FlowId", fmt.Errorf(`value '%v' must not be an empty string`, this.FlowId))
	}
	if this.StageId == "" {
		return go_proto_validators.FieldError("StageId", fmt.Errorf(`value '%v' must not be an empty string`, this.StageId))
	}
	if !_regex_InvokeFunctionRequest_FunctionId.MatchString(this.FunctionId) {
		return go_proto_validators.FieldError("FunctionId", fmt.Errorf(`invalid function id`))
	}
	if nil == this.Arg {
		return go_proto_validators.FieldError("Arg", fmt.Errorf("message must exist"))
	}
	if this.Arg != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Arg); err != nil {
			return go_proto_validators.FieldError("Arg", err)
		}
	}
	return nil
}

var _regex_InvokeStageRequest_FunctionId = regexp.MustCompile("^(\\.|[a-zA-Z0-9_\\-]{1,255})(/[a-zA-Z0-9_/\\-]{0,255})(\\?.*)?$")

func (this *InvokeStageRequest) Validate() error {
	if this.FlowId == "" {
		return go_proto_validators.FieldError("FlowId", fmt.Errorf(`value '%v' must not be an empty string`, this.FlowId))
	}
	if this.StageId == "" {
		return go_proto_validators.FieldError("StageId", fmt.Errorf(`value '%v' must not be an empty string`, this.StageId))
	}
	if !_regex_InvokeStageRequest_FunctionId.MatchString(this.FunctionId) {
		return go_proto_validators.FieldError("FunctionId", fmt.Errorf(`value '%v' must be a string conforming to regex "^(\\.|[a-zA-Z0-9_\\-]{1,255})(/[a-zA-Z0-9_/\\-]{0,255})(\\?.*)?$"`, this.FunctionId))
	}
	for _, item := range this.Args {
		if item != nil {
			if err := go_proto_validators.CallValidatorIfExists(item); err != nil {
				return go_proto_validators.FieldError("Args", err)
			}
		}
	}
	if nil == this.Closure {
		return go_proto_validators.FieldError("Closure", fmt.Errorf("message must exist"))
	}
	if this.Closure != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Closure); err != nil {
			return go_proto_validators.FieldError("Closure", err)
		}
	}
	return nil
}
func (this *GraphEvent) Validate() error {
	if oneOfNester, ok := this.GetVal().(*GraphEvent_GraphCreated); ok {
		if oneOfNester.GraphCreated != nil {
			if err := go_proto_validators.CallValidatorIfExists(oneOfNester.GraphCreated); err != nil {
				return go_proto_validators.FieldError("GraphCreated", err)
			}
		}
	}
	if oneOfNester, ok := this.GetVal().(*GraphEvent_GraphTermianting); ok {
		if oneOfNester.GraphTermianting != nil {
			if err := go_proto_validators.CallValidatorIfExists(oneOfNester.GraphTermianting); err != nil {
				return go_proto_validators.FieldError("GraphTermianting", err)
			}
		}
	}
	if oneOfNester, ok := this.GetVal().(*GraphEvent_GraphCompleted); ok {
		if oneOfNester.GraphCompleted != nil {
			if err := go_proto_validators.CallValidatorIfExists(oneOfNester.GraphCompleted); err != nil {
				return go_proto_validators.FieldError("GraphCompleted", err)
			}
		}
	}
	if oneOfNester, ok := this.GetVal().(*GraphEvent_DelayScheduled); ok {
		if oneOfNester.DelayScheduled != nil {
			if err := go_proto_validators.CallValidatorIfExists(oneOfNester.DelayScheduled); err != nil {
				return go_proto_validators.FieldError("DelayScheduled", err)
			}
		}
	}
	if oneOfNester, ok := this.GetVal().(*GraphEvent_StageAdded); ok {
		if oneOfNester.StageAdded != nil {
			if err := go_proto_validators.CallValidatorIfExists(oneOfNester.StageAdded); err != nil {
				return go_proto_validators.FieldError("StageAdded", err)
			}
		}
	}
	if oneOfNester, ok := this.GetVal().(*GraphEvent_StageCompleted); ok {
		if oneOfNester.StageCompleted != nil {
			if err := go_proto_validators.CallValidatorIfExists(oneOfNester.StageCompleted); err != nil {
				return go_proto_validators.FieldError("StageCompleted", err)
			}
		}
	}
	if oneOfNester, ok := this.GetVal().(*GraphEvent_StageComposed); ok {
		if oneOfNester.StageComposed != nil {
			if err := go_proto_validators.CallValidatorIfExists(oneOfNester.StageComposed); err != nil {
				return go_proto_validators.FieldError("StageComposed", err)
			}
		}
	}
	if oneOfNester, ok := this.GetVal().(*GraphEvent_FaasInvocationStarted); ok {
		if oneOfNester.FaasInvocationStarted != nil {
			if err := go_proto_validators.CallValidatorIfExists(oneOfNester.FaasInvocationStarted); err != nil {
				return go_proto_validators.FieldError("FaasInvocationStarted", err)
			}
		}
	}
	if oneOfNester, ok := this.GetVal().(*GraphEvent_FaasInvocationCompleted); ok {
		if oneOfNester.FaasInvocationCompleted != nil {
			if err := go_proto_validators.CallValidatorIfExists(oneOfNester.FaasInvocationCompleted); err != nil {
				return go_proto_validators.FieldError("FaasInvocationCompleted", err)
			}
		}
	}
	return nil
}
func (this *GraphCreatedEvent) Validate() error {
	if this.Ts != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Ts); err != nil {
			return go_proto_validators.FieldError("Ts", err)
		}
	}
	return nil
}
func (this *DelayScheduledEvent) Validate() error {
	if this.Ts != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Ts); err != nil {
			return go_proto_validators.FieldError("Ts", err)
		}
	}
	return nil
}
func (this *GraphTerminatingEvent) Validate() error {
	if this.Ts != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Ts); err != nil {
			return go_proto_validators.FieldError("Ts", err)
		}
	}
	return nil
}
func (this *GraphCompletedEvent) Validate() error {
	if this.Ts != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Ts); err != nil {
			return go_proto_validators.FieldError("Ts", err)
		}
	}
	return nil
}
func (this *GraphCommittedEvent) Validate() error {
	if this.Ts != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Ts); err != nil {
			return go_proto_validators.FieldError("Ts", err)
		}
	}
	return nil
}
func (this *StageAddedEvent) Validate() error {
	if this.Closure != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Closure); err != nil {
			return go_proto_validators.FieldError("Closure", err)
		}
	}
	if this.Ts != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Ts); err != nil {
			return go_proto_validators.FieldError("Ts", err)
		}
	}
	return nil
}
func (this *StageCompletedEvent) Validate() error {
	if this.Result != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Result); err != nil {
			return go_proto_validators.FieldError("Result", err)
		}
	}
	if this.Ts != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Ts); err != nil {
			return go_proto_validators.FieldError("Ts", err)
		}
	}
	return nil
}
func (this *StageComposedEvent) Validate() error {
	if this.Ts != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Ts); err != nil {
			return go_proto_validators.FieldError("Ts", err)
		}
	}
	return nil
}
func (this *FaasInvocationStartedEvent) Validate() error {
	if this.Ts != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Ts); err != nil {
			return go_proto_validators.FieldError("Ts", err)
		}
	}
	return nil
}
func (this *FaasInvocationCompletedEvent) Validate() error {
	if this.Result != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Result); err != nil {
			return go_proto_validators.FieldError("Result", err)
		}
	}
	if this.Ts != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Ts); err != nil {
			return go_proto_validators.FieldError("Ts", err)
		}
	}
	return nil
}
func (this *RuntimeInvokeStageRequest) Validate() error {
	if this.FlowId == "" {
		return go_proto_validators.FieldError("FlowId", fmt.Errorf(`value '%v' must not be an empty string`, this.FlowId))
	}
	if this.StageId == "" {
		return go_proto_validators.FieldError("StageId", fmt.Errorf(`value '%v' must not be an empty string`, this.StageId))
	}
	for _, item := range this.Args {
		if item != nil {
			if err := go_proto_validators.CallValidatorIfExists(item); err != nil {
				return go_proto_validators.FieldError("Args", err)
			}
		}
	}
	if nil == this.Closure {
		return go_proto_validators.FieldError("Closure", fmt.Errorf("message must exist"))
	}
	if this.Closure != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Closure); err != nil {
			return go_proto_validators.FieldError("Closure", err)
		}
	}
	return nil
}
func (this *RuntimeInvokeStageResponse) Validate() error {
	if this.Result != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Result); err != nil {
			return go_proto_validators.FieldError("Result", err)
		}
	}
	return nil
}
