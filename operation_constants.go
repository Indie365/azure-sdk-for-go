package servicebus

const vendorPrefix = "com.microsoft:"

// Operations
const (
	lockRenewalOperationName   = vendorPrefix + "renew-lock"
	peekMessageOperationID     = vendorPrefix + "peek-message"
	scheduleMessageOperationID = vendorPrefix + "schedule-message"
)

// Field Descriptions
const (
	operationFieldName     = "operation"
	lockTokensFieldName    = "lock-tokens"
	serverTimeoutFieldName = vendorPrefix + "server-timeout"
)
