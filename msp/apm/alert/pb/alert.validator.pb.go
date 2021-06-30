// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: alert.proto

package pb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/structpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *QueryAlertRuleRequest) Validate() error {
	return nil
}
func (this *QueryAlertRuleResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *AlertTypeRuleResp) Validate() error {
	for _, item := range this.AlertTypeRules {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("AlertTypeRules", err)
			}
		}
	}
	for _, item := range this.Operators {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Operators", err)
			}
		}
	}
	for _, item := range this.Aggregator {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Aggregator", err)
			}
		}
	}
	for _, item := range this.Silence {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Silence", err)
			}
		}
	}
	return nil
}
func (this *DisplayKey) Validate() error {
	return nil
}
func (this *AlertTypeRule) Validate() error {
	if this.AlertType != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.AlertType); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("AlertType", err)
		}
	}
	for _, item := range this.Rules {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Rules", err)
			}
		}
	}
	return nil
}
func (this *AlertRule) Validate() error {
	if this.AlertIndex != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.AlertIndex); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("AlertIndex", err)
		}
	}
	// Validation of proto3 map<> fields is unsupported.
	for _, item := range this.Functions {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Functions", err)
			}
		}
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *AlertRuleFunction) Validate() error {
	if this.Field != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Field); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Field", err)
		}
	}
	if this.Value != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Value); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Value", err)
		}
	}
	return nil
}
func (this *Operator) Validate() error {
	return nil
}
func (this *NotifySilence) Validate() error {
	if this.Unit != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Unit); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Unit", err)
		}
	}
	return nil
}
func (this *QueryAlertRequest) Validate() error {
	if !(this.PageNo > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("PageNo", fmt.Errorf(`value '%v' must be greater than '0'`, this.PageNo))
	}
	if !(this.PageSize > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("PageSize", fmt.Errorf(`value '%v' must be greater than '0'`, this.PageSize))
	}
	if !(this.PageSize < 101) {
		return github_com_mwitkow_go_proto_validators.FieldError("PageSize", fmt.Errorf(`value '%v' must be less than '101'`, this.PageSize))
	}
	return nil
}
func (this *QueryAlertResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *QueryAlertData) Validate() error {
	for _, item := range this.List {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("List", err)
			}
		}
	}
	return nil
}
func (this *Alert) Validate() error {
	for _, item := range this.Rules {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Rules", err)
			}
		}
	}
	for _, item := range this.Notifies {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Notifies", err)
			}
		}
	}
	// Validation of proto3 map<> fields is unsupported.
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *AlertExpression) Validate() error {
	for _, item := range this.Functions {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Functions", err)
			}
		}
	}
	return nil
}
func (this *AlertExpressionFunction) Validate() error {
	if this.Value != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Value); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Value", err)
		}
	}
	return nil
}
func (this *AlertNotify) Validate() error {
	if this.NotifyGroup != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.NotifyGroup); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("NotifyGroup", err)
		}
	}
	if this.Silence != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Silence); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Silence", err)
		}
	}
	return nil
}
func (this *NotifyGroup) Validate() error {
	for _, item := range this.Targets {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Targets", err)
			}
		}
	}
	if this.CreatedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CreatedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CreatedAt", err)
		}
	}
	return nil
}
func (this *NotifyTarget) Validate() error {
	for _, item := range this.Values {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Values", err)
			}
		}
	}
	return nil
}
func (this *Target) Validate() error {
	return nil
}
func (this *AlertNotifySilence) Validate() error {
	return nil
}
func (this *GetAlertRequest) Validate() error {
	if !(this.Id > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be greater than '0'`, this.Id))
	}
	return nil
}
func (this *GetAlertResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateAlertRequest) Validate() error {
	for _, item := range this.Rules {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Rules", err)
			}
		}
	}
	for _, item := range this.Notifies {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Notifies", err)
			}
		}
	}
	return nil
}
func (this *CreateAlertResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateAlertData) Validate() error {
	return nil
}
func (this *UpdateAlertRequest) Validate() error {
	if !(this.Id > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be greater than '0'`, this.Id))
	}
	for _, item := range this.Rules {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Rules", err)
			}
		}
	}
	for _, item := range this.Notifies {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Notifies", err)
			}
		}
	}
	return nil
}
func (this *UpdateAlertResponse) Validate() error {
	return nil
}
func (this *UpdateAlertEnableRequest) Validate() error {
	if !(this.Id > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be greater than '0'`, this.Id))
	}
	return nil
}
func (this *UpdateAlertEnableResponse) Validate() error {
	return nil
}
func (this *DeleteAlertRequest) Validate() error {
	if !(this.Id > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be greater than '0'`, this.Id))
	}
	return nil
}
func (this *DeleteAlertResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *DeleteAlertData) Validate() error {
	return nil
}
func (this *QueryCustomizeMetricRequest) Validate() error {
	return nil
}
func (this *QueryCustomizeMetricResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CustomizeMetricData) Validate() error {
	for _, item := range this.Metrics {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Metrics", err)
			}
		}
	}
	for _, item := range this.FunctionOperators {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("FunctionOperators", err)
			}
		}
	}
	for _, item := range this.FilterOperators {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("FilterOperators", err)
			}
		}
	}
	for _, item := range this.Aggregator {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Aggregator", err)
			}
		}
	}
	return nil
}
func (this *MetricMeta) Validate() error {
	if this.Name != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Name); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Name", err)
		}
	}
	for _, item := range this.Fields {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Fields", err)
			}
		}
	}
	for _, item := range this.Tags {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Tags", err)
			}
		}
	}
	return nil
}
func (this *FieldMeta) Validate() error {
	if this.Field != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Field); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Field", err)
		}
	}
	return nil
}
func (this *TagMeta) Validate() error {
	if this.Tag != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Tag); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Tag", err)
		}
	}
	return nil
}
func (this *QueryCustomizeNotifyTargetRequest) Validate() error {
	return nil
}
func (this *QueryCustomizeNotifyTargetResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *QueryCustomizeNotifyTargetData) Validate() error {
	for _, item := range this.Targets {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Targets", err)
			}
		}
	}
	return nil
}
func (this *QueryCustomizeAlertsRequest) Validate() error {
	if !(this.PageNo > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("PageNo", fmt.Errorf(`value '%v' must be greater than '0'`, this.PageNo))
	}
	if !(this.PageSize > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("PageSize", fmt.Errorf(`value '%v' must be greater than '0'`, this.PageSize))
	}
	if !(this.PageSize < 101) {
		return github_com_mwitkow_go_proto_validators.FieldError("PageSize", fmt.Errorf(`value '%v' must be less than '101'`, this.PageSize))
	}
	return nil
}
func (this *QueryCustomizeAlertsResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *QueryCustomizeAlertsData) Validate() error {
	for _, item := range this.List {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("List", err)
			}
		}
	}
	return nil
}
func (this *CustomizeAlertOverview) Validate() error {
	return nil
}
func (this *GetCustomizeAlertRequest) Validate() error {
	if !(this.Id > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be greater than '0'`, this.Id))
	}
	return nil
}
func (this *GetCustomizeAlertResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CustomizeAlertDetail) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	for _, item := range this.Rules {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Rules", err)
			}
		}
	}
	for _, item := range this.Notifies {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Notifies", err)
			}
		}
	}
	return nil
}
func (this *CustomizeAlertRule) Validate() error {
	for _, item := range this.Functions {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Functions", err)
			}
		}
	}
	for _, item := range this.Filters {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Filters", err)
			}
		}
	}
	// Validation of proto3 map<> fields is unsupported.
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *CustomizeAlertRuleFunction) Validate() error {
	if this.Value != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Value); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Value", err)
		}
	}
	return nil
}
func (this *CustomizeAlertRuleFilter) Validate() error {
	if this.Value != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Value); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Value", err)
		}
	}
	return nil
}
func (this *CustomizeAlertNotifyTemplates) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *CreateCustomizeAlertRequest) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	for _, item := range this.Rules {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Rules", err)
			}
		}
	}
	for _, item := range this.Notifies {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Notifies", err)
			}
		}
	}
	return nil
}
func (this *CreateCustomizeAlertResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateCustomizeAlertData) Validate() error {
	return nil
}
func (this *UpdateCustomizeAlertRequest) Validate() error {
	if !(this.Id > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be greater than '0'`, this.Id))
	}
	// Validation of proto3 map<> fields is unsupported.
	for _, item := range this.Rules {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Rules", err)
			}
		}
	}
	for _, item := range this.Notifies {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Notifies", err)
			}
		}
	}
	return nil
}
func (this *UpdateCustomizeAlertResponse) Validate() error {
	return nil
}
func (this *UpdateCustomizeAlertEnableRequest) Validate() error {
	if !(this.Id > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be greater than '0'`, this.Id))
	}
	return nil
}
func (this *UpdateCustomizeAlertEnableResponse) Validate() error {
	return nil
}
func (this *DeleteCustomizeAlertRequest) Validate() error {
	if !(this.Id > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be greater than '0'`, this.Id))
	}
	return nil
}
func (this *DeleteCustomizeAlertResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *DeleteCustomizeAlertData) Validate() error {
	return nil
}
func (this *GetAlertRecordAttrsRequest) Validate() error {
	return nil
}
func (this *GetAlertRecordAttrsResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *AlertRecordAttr) Validate() error {
	for _, item := range this.AlertState {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("AlertState", err)
			}
		}
	}
	for _, item := range this.AlertType {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("AlertType", err)
			}
		}
	}
	for _, item := range this.HandleState {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("HandleState", err)
			}
		}
	}
	return nil
}
func (this *GetAlertRecordsRequest) Validate() error {
	return nil
}
func (this *GetAlertRecordsResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *GetAlertRecordsData) Validate() error {
	for _, item := range this.List {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("List", err)
			}
		}
	}
	return nil
}
func (this *AlertRecord) Validate() error {
	return nil
}
func (this *GetAlertRecordRequest) Validate() error {
	return nil
}
func (this *GetAlertRecordResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *GetAlertHistoriesRequest) Validate() error {
	return nil
}
func (this *GetAlertHistoriesResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *AlertHistory) Validate() error {
	return nil
}
func (this *CreateAlertRecordIssueRequest) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *CreateAlertRecordIssueResponse) Validate() error {
	return nil
}
func (this *UpdateAlertRecordIssueRequest) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *UpdateAlertRecordIssueResponse) Validate() error {
	return nil
}
func (this *DashboardPreviewRequest) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	for _, item := range this.Rules {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Rules", err)
			}
		}
	}
	for _, item := range this.Notifies {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Notifies", err)
			}
		}
	}
	return nil
}
func (this *DashboardPreviewResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *View) Validate() error {
	if this.StaticData != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.StaticData); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("StaticData", err)
		}
	}
	if this.Config != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Config); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Config", err)
		}
	}
	if this.Api != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Api); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Api", err)
		}
	}
	if this.Controls != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Controls); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Controls", err)
		}
	}
	return nil
}
func (this *Config) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	if this.DataSourceConfig != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.DataSourceConfig); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("DataSourceConfig", err)
		}
	}
	if this.Option != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Option); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Option", err)
		}
	}
	return nil
}
func (this *API) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	// Validation of proto3 map<> fields is unsupported.
	// Validation of proto3 map<> fields is unsupported.
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
