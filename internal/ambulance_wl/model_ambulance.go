/*
 * Waiting List Api
 *
 * Ambulance Waiting List management for Web-In-Cloud system
 *
 * API version: 1.0.0
 * Contact: xkello@stuba.sk
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ambulance_wl

type Ambulance struct {

	// Unique identifier of the ambulance
	Id string `json:"id"`

	// Human readable display name of the ambulance
	Name string `json:"name"`

	RoomNumber string `json:"roomNumber"`

	WaitingList []WaitingListEntry `json:"waitingList,omitempty"`

	PredefinedConditions []Condition `json:"predefinedConditions,omitempty"`
}
