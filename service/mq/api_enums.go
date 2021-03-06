// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mq

// The status of the broker.
type BrokerState string

// Enum values for BrokerState
const (
	BrokerStateCreationInProgress BrokerState = "CREATION_IN_PROGRESS"
	BrokerStateCreationFailed     BrokerState = "CREATION_FAILED"
	BrokerStateDeletionInProgress BrokerState = "DELETION_IN_PROGRESS"
	BrokerStateRunning            BrokerState = "RUNNING"
	BrokerStateRebootInProgress   BrokerState = "REBOOT_IN_PROGRESS"
)

func (enum BrokerState) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum BrokerState) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

// The type of change pending for the ActiveMQ user.
type ChangeType string

// Enum values for ChangeType
const (
	ChangeTypeCreate ChangeType = "CREATE"
	ChangeTypeUpdate ChangeType = "UPDATE"
	ChangeTypeDelete ChangeType = "DELETE"
)

func (enum ChangeType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ChangeType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type DayOfWeek string

// Enum values for DayOfWeek
const (
	DayOfWeekMonday    DayOfWeek = "MONDAY"
	DayOfWeekTuesday   DayOfWeek = "TUESDAY"
	DayOfWeekWednesday DayOfWeek = "WEDNESDAY"
	DayOfWeekThursday  DayOfWeek = "THURSDAY"
	DayOfWeekFriday    DayOfWeek = "FRIDAY"
	DayOfWeekSaturday  DayOfWeek = "SATURDAY"
	DayOfWeekSunday    DayOfWeek = "SUNDAY"
)

func (enum DayOfWeek) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DayOfWeek) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

// The deployment mode of the broker.
type DeploymentMode string

// Enum values for DeploymentMode
const (
	DeploymentModeSingleInstance       DeploymentMode = "SINGLE_INSTANCE"
	DeploymentModeActiveStandbyMultiAz DeploymentMode = "ACTIVE_STANDBY_MULTI_AZ"
)

func (enum DeploymentMode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DeploymentMode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

// The type of broker engine. Note: Currently, Amazon MQ supports only ActiveMQ.
type EngineType string

// Enum values for EngineType
const (
	EngineTypeActivemq EngineType = "ACTIVEMQ"
)

func (enum EngineType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum EngineType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

// The reason for which the XML elements or attributes were sanitized.
type SanitizationWarningReason string

// Enum values for SanitizationWarningReason
const (
	SanitizationWarningReasonDisallowedElementRemoved     SanitizationWarningReason = "DISALLOWED_ELEMENT_REMOVED"
	SanitizationWarningReasonDisallowedAttributeRemoved   SanitizationWarningReason = "DISALLOWED_ATTRIBUTE_REMOVED"
	SanitizationWarningReasonInvalidAttributeValueRemoved SanitizationWarningReason = "INVALID_ATTRIBUTE_VALUE_REMOVED"
)

func (enum SanitizationWarningReason) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum SanitizationWarningReason) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
