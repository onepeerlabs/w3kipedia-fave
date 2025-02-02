/*
 * fave
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.0-prealpha
 * Contact: sabyasachi@datafund.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Response for the Nearest documents request
type NearestDocumentsResponse struct {
	// The actual list of Objects.
	Documents []Document `json:"documents,omitempty"`
	// Name of the collection as URI relative to the schema URL.
	Name string `json:"name,omitempty"`
}
