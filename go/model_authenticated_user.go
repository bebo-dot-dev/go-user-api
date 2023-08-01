/*
 * User API
 *
 * An API for managing system users
 *
 * API version: 1.0.3
 * Contact: joe@jjssoftware.co.uk
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api

type AuthenticatedUser struct {
	SessionId string `json:"sessionId"`

	UserName string `json:"userName"`

	FirstName string `json:"firstName"`

	Surname string `json:"surname"`
}

// AssertAuthenticatedUserRequired checks if the required fields are not zero-ed
func AssertAuthenticatedUserRequired(obj AuthenticatedUser) error {
	elements := map[string]interface{}{
		"sessionId": obj.SessionId,
		"userName":  obj.UserName,
		"firstName": obj.FirstName,
		"surname":   obj.Surname,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseAuthenticatedUserRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of AuthenticatedUser (e.g. [][]AuthenticatedUser), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseAuthenticatedUserRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aAuthenticatedUser, ok := obj.(AuthenticatedUser)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertAuthenticatedUserRequired(aAuthenticatedUser)
	})
}
