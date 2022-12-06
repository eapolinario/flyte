/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

// Request to send a notification that a node execution event has occurred.
type AdminNodeExecutionEventRequest struct {
	RequestId string `json:"request_id,omitempty"`
	// Details about the event that occurred.
	Event *EventNodeExecutionEvent `json:"event,omitempty"`
}
