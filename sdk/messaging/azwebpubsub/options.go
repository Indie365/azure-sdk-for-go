//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azwebpubsub

// ClientAddConnectionToGroupOptions contains the optional parameters for the Client.AddConnectionToGroup method.
type ClientAddConnectionToGroupOptions struct {
	// placeholder for future optional parameters
}

// ClientAddConnectionsToGroupsOptions contains the optional parameters for the Client.AddConnectionsToGroups method.
type ClientAddConnectionsToGroupsOptions struct {
	// placeholder for future optional parameters
}

// ClientAddUserToGroupOptions contains the optional parameters for the Client.AddUserToGroup method.
type ClientAddUserToGroupOptions struct {
	// placeholder for future optional parameters
}

// ClientCheckPermissionOptions contains the optional parameters for the Client.CheckPermission method.
type ClientCheckPermissionOptions struct {
	// The meaning of the target depends on the specific permission. For joinLeaveGroup and sendToGroup, targetName is a required
	// parameter standing for the group name.
	TargetName *string
}

// ClientCloseAllConnectionsOptions contains the optional parameters for the Client.CloseAllConnections method.
type ClientCloseAllConnectionsOptions struct {
	// Exclude these connectionIds when closing the connections in the hub.
	Excluded []string

	// The reason closing the client connection.
	Reason *string
}

// ClientCloseConnectionOptions contains the optional parameters for the Client.CloseConnection method.
type ClientCloseConnectionOptions struct {
	// The reason closing the client connection.
	Reason *string
}

// ClientCloseGroupConnectionsOptions contains the optional parameters for the Client.CloseGroupConnections method.
type ClientCloseGroupConnectionsOptions struct {
	// Exclude these connectionIds when closing the connections in the group.
	Excluded []string

	// The reason closing the client connection.
	Reason *string
}

// ClientCloseUserConnectionsOptions contains the optional parameters for the Client.CloseUserConnections method.
type ClientCloseUserConnectionsOptions struct {
	// Exclude these connectionIds when closing the connections for the user.
	Excluded []string

	// The reason closing the client connection.
	Reason *string
}

// ClientConnectionExistsOptions contains the optional parameters for the Client.ConnectionExists method.
type ClientConnectionExistsOptions struct {
	// placeholder for future optional parameters
}

// ClientGenerateClientTokenOptions contains the optional parameters for the Client.GenerateClientToken method.
type ClientGenerateClientTokenOptions struct {
	// Groups that the connection will join when it connects.
	Group []string

	// The expire time of the generated token.
	MinutesToExpire *int32

	// Roles that the connection with the generated token will have.
	Role []string

	// User Id.
	UserID *string
}

// ClientGrantPermissionOptions contains the optional parameters for the Client.GrantPermission method.
type ClientGrantPermissionOptions struct {
	// The meaning of the target depends on the specific permission. For joinLeaveGroup and sendToGroup, targetName is a required
	// parameter standing for the group name.
	TargetName *string
}

// ClientGroupExistsOptions contains the optional parameters for the Client.GroupExists method.
type ClientGroupExistsOptions struct {
	// placeholder for future optional parameters
}

// ClientRemoveConnectionFromAllGroupsOptions contains the optional parameters for the Client.RemoveConnectionFromAllGroups
// method.
type ClientRemoveConnectionFromAllGroupsOptions struct {
	// placeholder for future optional parameters
}

// ClientRemoveConnectionFromGroupOptions contains the optional parameters for the Client.RemoveConnectionFromGroup method.
type ClientRemoveConnectionFromGroupOptions struct {
	// placeholder for future optional parameters
}

// ClientRemoveConnectionsFromGroupsOptions contains the optional parameters for the Client.RemoveConnectionsFromGroups method.
type ClientRemoveConnectionsFromGroupsOptions struct {
	// placeholder for future optional parameters
}

// ClientRemoveUserFromAllGroupsOptions contains the optional parameters for the Client.RemoveUserFromAllGroups method.
type ClientRemoveUserFromAllGroupsOptions struct {
	// placeholder for future optional parameters
}

// ClientRemoveUserFromGroupOptions contains the optional parameters for the Client.RemoveUserFromGroup method.
type ClientRemoveUserFromGroupOptions struct {
	// placeholder for future optional parameters
}

// ClientRevokePermissionOptions contains the optional parameters for the Client.RevokePermission method.
type ClientRevokePermissionOptions struct {
	// The meaning of the target depends on the specific permission. For joinLeaveGroup and sendToGroup, targetName is a required
	// parameter standing for the group name.
	TargetName *string
}

// ClientSendToAllOptions contains the optional parameters for the Client.SendToAll method.
type ClientSendToAllOptions struct {
	// Excluded connection Ids.
	Excluded []string

	// Following OData filter syntax to filter out the subscribers receiving the messages.
	Filter *string

	// The time-to-live (TTL) value in seconds for messages sent to the service. 0 is the default value, which means the message
	// never expires. 300 is the maximum value. If this parameter is non-zero,
	// messages that are not consumed by the client within the specified TTL will be dropped by the service. This parameter can
	// help when the client's bandwidth is limited.
	MessageTTLSeconds *int32
}

// ClientSendToConnectionOptions contains the optional parameters for the Client.SendToConnection method.
type ClientSendToConnectionOptions struct {
	// The time-to-live (TTL) value in seconds for messages sent to the service. 0 is the default value, which means the message
	// never expires. 300 is the maximum value. If this parameter is non-zero,
	// messages that are not consumed by the client within the specified TTL will be dropped by the service. This parameter can
	// help when the client's bandwidth is limited.
	MessageTTLSeconds *int32
}

// ClientSendToGroupOptions contains the optional parameters for the Client.SendToGroup method.
type ClientSendToGroupOptions struct {
	// Excluded connection Ids
	Excluded []string

	// Following OData filter syntax to filter out the subscribers receiving the messages.
	Filter *string

	// The time-to-live (TTL) value in seconds for messages sent to the service. 0 is the default value, which means the message
	// never expires. 300 is the maximum value. If this parameter is non-zero,
	// messages that are not consumed by the client within the specified TTL will be dropped by the service. This parameter can
	// help when the client's bandwidth is limited.
	MessageTTLSeconds *int32
}

// ClientSendToUserOptions contains the optional parameters for the Client.SendToUser method.
type ClientSendToUserOptions struct {
	// Following OData filter syntax to filter out the subscribers receiving the messages.
	Filter *string

	// The time-to-live (TTL) value in seconds for messages sent to the service. 0 is the default value, which means the message
	// never expires. 300 is the maximum value. If this parameter is non-zero,
	// messages that are not consumed by the client within the specified TTL will be dropped by the service. This parameter can
	// help when the client's bandwidth is limited.
	MessageTTLSeconds *int32
}

// ClientUserExistsOptions contains the optional parameters for the Client.UserExists method.
type ClientUserExistsOptions struct {
	// placeholder for future optional parameters
}
