// Code generated by go-swagger; DO NOT EDIT.

package engines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new engines API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for engines API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	AddOryAccessControlPolicyRoleMembers(params *AddOryAccessControlPolicyRoleMembersParams) (*AddOryAccessControlPolicyRoleMembersOK, error)

	DeleteOryAccessControlPolicy(params *DeleteOryAccessControlPolicyParams) (*DeleteOryAccessControlPolicyNoContent, error)

	DeleteOryAccessControlPolicyRole(params *DeleteOryAccessControlPolicyRoleParams) (*DeleteOryAccessControlPolicyRoleNoContent, error)

	DoOryAccessControlPoliciesAllow(params *DoOryAccessControlPoliciesAllowParams) (*DoOryAccessControlPoliciesAllowOK, error)

	GetOryAccessControlPolicy(params *GetOryAccessControlPolicyParams) (*GetOryAccessControlPolicyOK, error)

	GetOryAccessControlPolicyRole(params *GetOryAccessControlPolicyRoleParams) (*GetOryAccessControlPolicyRoleOK, error)

	ListOryAccessControlPolicies(params *ListOryAccessControlPoliciesParams) (*ListOryAccessControlPoliciesOK, error)

	ListOryAccessControlPolicyRoles(params *ListOryAccessControlPolicyRolesParams) (*ListOryAccessControlPolicyRolesOK, error)

	RemoveOryAccessControlPolicyRoleMembers(params *RemoveOryAccessControlPolicyRoleMembersParams) (*RemoveOryAccessControlPolicyRoleMembersOK, error)

	UpsertOryAccessControlPolicy(params *UpsertOryAccessControlPolicyParams) (*UpsertOryAccessControlPolicyOK, error)

	UpsertOryAccessControlPolicyRole(params *UpsertOryAccessControlPolicyRoleParams) (*UpsertOryAccessControlPolicyRoleOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AddOryAccessControlPolicyRoleMembers adds a member to an o r y access control policy role

  Roles group several subjects into one. Rules can be assigned to ORY Access Control Policy (OACP) by using the Role ID
as subject in the OACP.
*/
func (a *Client) AddOryAccessControlPolicyRoleMembers(params *AddOryAccessControlPolicyRoleMembersParams) (*AddOryAccessControlPolicyRoleMembersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddOryAccessControlPolicyRoleMembersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addOryAccessControlPolicyRoleMembers",
		Method:             "PUT",
		PathPattern:        "/engines/acp/ory/{flavor}/roles/{id}/members",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddOryAccessControlPolicyRoleMembersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddOryAccessControlPolicyRoleMembersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addOryAccessControlPolicyRoleMembers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteOryAccessControlPolicy Delete an ORY Access Control Policy
*/
func (a *Client) DeleteOryAccessControlPolicy(params *DeleteOryAccessControlPolicyParams) (*DeleteOryAccessControlPolicyNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteOryAccessControlPolicyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteOryAccessControlPolicy",
		Method:             "DELETE",
		PathPattern:        "/engines/acp/ory/{flavor}/policies/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteOryAccessControlPolicyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteOryAccessControlPolicyNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteOryAccessControlPolicy: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteOryAccessControlPolicyRole deletes an o r y access control policy role

  Roles group several subjects into one. Rules can be assigned to ORY Access Control Policy (OACP) by using the Role ID
as subject in the OACP.
*/
func (a *Client) DeleteOryAccessControlPolicyRole(params *DeleteOryAccessControlPolicyRoleParams) (*DeleteOryAccessControlPolicyRoleNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteOryAccessControlPolicyRoleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteOryAccessControlPolicyRole",
		Method:             "DELETE",
		PathPattern:        "/engines/acp/ory/{flavor}/roles/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteOryAccessControlPolicyRoleReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteOryAccessControlPolicyRoleNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteOryAccessControlPolicyRole: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DoOryAccessControlPoliciesAllow checks if a request is allowed

  Use this endpoint to check if a request is allowed or not. If the request is allowed, a 200 response with
`{"allowed":"true"}` will be sent. If the request is denied, a 403 response with `{"allowed":"false"}` will
be sent instead.
*/
func (a *Client) DoOryAccessControlPoliciesAllow(params *DoOryAccessControlPoliciesAllowParams) (*DoOryAccessControlPoliciesAllowOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDoOryAccessControlPoliciesAllowParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "doOryAccessControlPoliciesAllow",
		Method:             "POST",
		PathPattern:        "/engines/acp/ory/{flavor}/allowed",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DoOryAccessControlPoliciesAllowReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DoOryAccessControlPoliciesAllowOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for doOryAccessControlPoliciesAllow: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOryAccessControlPolicy Get an ORY Access Control Policy
*/
func (a *Client) GetOryAccessControlPolicy(params *GetOryAccessControlPolicyParams) (*GetOryAccessControlPolicyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOryAccessControlPolicyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getOryAccessControlPolicy",
		Method:             "GET",
		PathPattern:        "/engines/acp/ory/{flavor}/policies/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetOryAccessControlPolicyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOryAccessControlPolicyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getOryAccessControlPolicy: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOryAccessControlPolicyRole gets an o r y access control policy role

  Roles group several subjects into one. Rules can be assigned to ORY Access Control Policy (OACP) by using the Role ID
as subject in the OACP.
*/
func (a *Client) GetOryAccessControlPolicyRole(params *GetOryAccessControlPolicyRoleParams) (*GetOryAccessControlPolicyRoleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOryAccessControlPolicyRoleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getOryAccessControlPolicyRole",
		Method:             "GET",
		PathPattern:        "/engines/acp/ory/{flavor}/roles/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetOryAccessControlPolicyRoleReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOryAccessControlPolicyRoleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getOryAccessControlPolicyRole: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListOryAccessControlPolicies List ORY Access Control Policies
*/
func (a *Client) ListOryAccessControlPolicies(params *ListOryAccessControlPoliciesParams) (*ListOryAccessControlPoliciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListOryAccessControlPoliciesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listOryAccessControlPolicies",
		Method:             "GET",
		PathPattern:        "/engines/acp/ory/{flavor}/policies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListOryAccessControlPoliciesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListOryAccessControlPoliciesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listOryAccessControlPolicies: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListOryAccessControlPolicyRoles lists o r y access control policy roles

  Roles group several subjects into one. Rules can be assigned to ORY Access Control Policy (OACP) by using the Role ID
as subject in the OACP.
*/
func (a *Client) ListOryAccessControlPolicyRoles(params *ListOryAccessControlPolicyRolesParams) (*ListOryAccessControlPolicyRolesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListOryAccessControlPolicyRolesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listOryAccessControlPolicyRoles",
		Method:             "GET",
		PathPattern:        "/engines/acp/ory/{flavor}/roles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListOryAccessControlPolicyRolesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListOryAccessControlPolicyRolesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listOryAccessControlPolicyRoles: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RemoveOryAccessControlPolicyRoleMembers removes a member from an o r y access control policy role

  Roles group several subjects into one. Rules can be assigned to ORY Access Control Policy (OACP) by using the Role ID
as subject in the OACP.
*/
func (a *Client) RemoveOryAccessControlPolicyRoleMembers(params *RemoveOryAccessControlPolicyRoleMembersParams) (*RemoveOryAccessControlPolicyRoleMembersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveOryAccessControlPolicyRoleMembersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "removeOryAccessControlPolicyRoleMembers",
		Method:             "DELETE",
		PathPattern:        "/engines/acp/ory/{flavor}/roles/{id}/members/{member}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RemoveOryAccessControlPolicyRoleMembersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RemoveOryAccessControlPolicyRoleMembersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for removeOryAccessControlPolicyRoleMembers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpsertOryAccessControlPolicy Upsert an ORY Access Control Policy
*/
func (a *Client) UpsertOryAccessControlPolicy(params *UpsertOryAccessControlPolicyParams) (*UpsertOryAccessControlPolicyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpsertOryAccessControlPolicyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "upsertOryAccessControlPolicy",
		Method:             "PUT",
		PathPattern:        "/engines/acp/ory/{flavor}/policies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpsertOryAccessControlPolicyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpsertOryAccessControlPolicyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for upsertOryAccessControlPolicy: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpsertOryAccessControlPolicyRole upserts an o r y access control policy role

  Roles group several subjects into one. Rules can be assigned to ORY Access Control Policy (OACP) by using the Role ID
as subject in the OACP.
*/
func (a *Client) UpsertOryAccessControlPolicyRole(params *UpsertOryAccessControlPolicyRoleParams) (*UpsertOryAccessControlPolicyRoleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpsertOryAccessControlPolicyRoleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "upsertOryAccessControlPolicyRole",
		Method:             "PUT",
		PathPattern:        "/engines/acp/ory/{flavor}/roles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpsertOryAccessControlPolicyRoleReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpsertOryAccessControlPolicyRoleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for upsertOryAccessControlPolicyRole: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
