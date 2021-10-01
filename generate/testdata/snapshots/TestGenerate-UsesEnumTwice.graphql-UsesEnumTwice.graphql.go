package test

// Code generated by github.com/Khan/genqlient version , DO NOT EDIT.

import (
	"github.com/Khan/genqlient/graphql"
)

// Role is a type a user may have.
type Role string

const (
	// What is a student?
	//
	// A student is primarily a person enrolled in a school or other educational institution and who is under learning with goals of acquiring knowledge, developing professions and achieving employment at desired field. In the broader sense, a student is anyone who applies themselves to the intensive intellectual engagement with some matter necessary to master it as part of some practical affair in which such mastery is basic or decisive.
	//
	// (from [Wikipedia](https://en.wikipedia.org/wiki/Student))
	RoleStudent Role = "STUDENT"
	// Teacher is a teacher, who teaches the students.
	RoleTeacher Role = "TEACHER"
)

// UsesEnumTwiceQueryMeUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type UsesEnumTwiceQueryMeUser struct {
	Roles []Role `json:"roles"`
}

// GetRoles returns UsesEnumTwiceQueryMeUser.Roles, and is useful for accessing the field via an interface.
func (v *UsesEnumTwiceQueryMeUser) GetRoles() []Role { return v.Roles }

// UsesEnumTwiceQueryOtherUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type UsesEnumTwiceQueryOtherUser struct {
	Roles []Role `json:"roles"`
}

// GetRoles returns UsesEnumTwiceQueryOtherUser.Roles, and is useful for accessing the field via an interface.
func (v *UsesEnumTwiceQueryOtherUser) GetRoles() []Role { return v.Roles }

// UsesEnumTwiceQueryResponse is returned by UsesEnumTwiceQuery on success.
type UsesEnumTwiceQueryResponse struct {
	// user looks up a user by some stuff.
	//
	// See UserQueryInput for what stuff is supported.
	// If query is null, returns the current user.
	Me UsesEnumTwiceQueryMeUser `json:"Me"`
	// user looks up a user by some stuff.
	//
	// See UserQueryInput for what stuff is supported.
	// If query is null, returns the current user.
	OtherUser UsesEnumTwiceQueryOtherUser `json:"OtherUser"`
}

// GetMe returns UsesEnumTwiceQueryResponse.Me, and is useful for accessing the field via an interface.
func (v *UsesEnumTwiceQueryResponse) GetMe() UsesEnumTwiceQueryMeUser { return v.Me }

// GetOtherUser returns UsesEnumTwiceQueryResponse.OtherUser, and is useful for accessing the field via an interface.
func (v *UsesEnumTwiceQueryResponse) GetOtherUser() UsesEnumTwiceQueryOtherUser { return v.OtherUser }

func UsesEnumTwiceQuery(
	client graphql.Client,
) (*UsesEnumTwiceQueryResponse, error) {
	var err error

	var retval UsesEnumTwiceQueryResponse
	err = client.MakeRequest(
		nil,
		"UsesEnumTwiceQuery",
		`
query UsesEnumTwiceQuery {
	Me: user {
		roles
	}
	OtherUser: user {
		roles
	}
}
`,
		&retval,
		nil,
	)
	return &retval, err
}

