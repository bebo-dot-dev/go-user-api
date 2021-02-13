/*
 * Users API
 *
 * An API for managing system users
 *
 * API version: 2.0.0
 * Contact: joe@jjssoftware.co.uk
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type NewUser struct {

	UserName string `json:"userName"`

	FirstName string `json:"firstName"`

	Surname string `json:"surname"`
}