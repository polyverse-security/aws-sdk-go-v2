// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lambda

type EventSourcePosition string

// Enum values for EventSourcePosition
const (
	EventSourcePositionTrimHorizon EventSourcePosition = "TRIM_HORIZON"
	EventSourcePositionLatest      EventSourcePosition = "LATEST"
	EventSourcePositionAtTimestamp EventSourcePosition = "AT_TIMESTAMP"
)

func (enum EventSourcePosition) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum EventSourcePosition) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type FunctionVersion string

// Enum values for FunctionVersion
const (
	FunctionVersionAll FunctionVersion = "ALL"
)

func (enum FunctionVersion) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum FunctionVersion) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type InvocationType string

// Enum values for InvocationType
const (
	InvocationTypeEvent           InvocationType = "Event"
	InvocationTypeRequestResponse InvocationType = "RequestResponse"
	InvocationTypeDryRun          InvocationType = "DryRun"
)

func (enum InvocationType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum InvocationType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type LogType string

// Enum values for LogType
const (
	LogTypeNone LogType = "None"
	LogTypeTail LogType = "Tail"
)

func (enum LogType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum LogType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Runtime string

// Enum values for Runtime
const (
	RuntimeNodejs       Runtime = "nodejs"
	RuntimeNodejs43     Runtime = "nodejs4.3"
	RuntimeNodejs610    Runtime = "nodejs6.10"
	RuntimeNodejs810    Runtime = "nodejs8.10"
	RuntimeNodejs10X    Runtime = "nodejs10.x"
	RuntimeJava8        Runtime = "java8"
	RuntimePython27     Runtime = "python2.7"
	RuntimePython36     Runtime = "python3.6"
	RuntimePython37     Runtime = "python3.7"
	RuntimeDotnetcore10 Runtime = "dotnetcore1.0"
	RuntimeDotnetcore20 Runtime = "dotnetcore2.0"
	RuntimeDotnetcore21 Runtime = "dotnetcore2.1"
	RuntimeNodejs43Edge Runtime = "nodejs4.3-edge"
	RuntimeGo1X         Runtime = "go1.x"
	RuntimeRuby25       Runtime = "ruby2.5"
	RuntimeProvided     Runtime = "provided"
)

func (enum Runtime) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Runtime) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ThrottleReason string

// Enum values for ThrottleReason
const (
	ThrottleReasonConcurrentInvocationLimitExceeded                 ThrottleReason = "ConcurrentInvocationLimitExceeded"
	ThrottleReasonFunctionInvocationRateLimitExceeded               ThrottleReason = "FunctionInvocationRateLimitExceeded"
	ThrottleReasonReservedFunctionConcurrentInvocationLimitExceeded ThrottleReason = "ReservedFunctionConcurrentInvocationLimitExceeded"
	ThrottleReasonReservedFunctionInvocationRateLimitExceeded       ThrottleReason = "ReservedFunctionInvocationRateLimitExceeded"
	ThrottleReasonCallerRateLimitExceeded                           ThrottleReason = "CallerRateLimitExceeded"
)

func (enum ThrottleReason) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ThrottleReason) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type TracingMode string

// Enum values for TracingMode
const (
	TracingModeActive      TracingMode = "Active"
	TracingModePassThrough TracingMode = "PassThrough"
)

func (enum TracingMode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum TracingMode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
