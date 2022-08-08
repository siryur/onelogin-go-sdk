/*
OneLogin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.0.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// User struct for User
type User struct {
	Id *int32 `json:"id,omitempty"`
	// A username for the user.
	Username *string `json:"username,omitempty"`
	// A valid email for the user.
	Email *string `json:"email,omitempty"`
	// The user's first name.
	Firstname *string `json:"firstname,omitempty"`
	// The user's last name.
	Lastname *string `json:"lastname,omitempty"`
	// The user's job title.
	Title *string `json:"title,omitempty"`
	// The user's department.
	Department *string `json:"department,omitempty"`
	// The user's company.
	Company *string `json:"company,omitempty"`
	// Free text related to the user.
	Comment *string `json:"comment,omitempty"`
	// The ID of the Group in OneLogin that the user is assigned to.
	GroupId *int32 `json:"group_id,omitempty"`
	// A list of OneLogin Role IDs of the user
	RoleIds []int32 `json:"role_ids,omitempty"`
	// The E.164 format phone number for a user.
	Phone *string `json:"phone,omitempty"`
	State *int32 `json:"state,omitempty"`
	Status *int32 `json:"status,omitempty"`
	// The ID of the OneLogin Directory of the user.
	DirectoryId *int32 `json:"directory_id,omitempty"`
	// The ID of the OneLogin Trusted IDP of the user.
	TrustedIdpId *int32 `json:"trusted_idp_id,omitempty"`
	// The ID of the user's manager in Active Directory.
	ManagerAdId *string `json:"manager_ad_id,omitempty"`
	// The OneLogin User ID for the user's manager.
	ManagerUserId *string `json:"manager_user_id,omitempty"`
	// The user's Active Directory username.
	SamaccountName *string `json:"samaccount_name,omitempty"`
	// The user's directory membership.
	MemberOf *string `json:"member_of,omitempty"`
	// The principle name of the user.
	Userprincipalname *string `json:"userprincipalname,omitempty"`
	// The distinguished name of the user.
	DistinguishedName *string `json:"distinguished_name,omitempty"`
	// The ID of the user in an external directory.
	ExternalId *string `json:"external_id,omitempty"`
	ActivatedAt *string `json:"activated_at,omitempty"`
	LastLogin *string `json:"last_login,omitempty"`
	InvitationSentAt *string `json:"invitation_sent_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	PreferredLocaleCode *string `json:"preferred_locale_code,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	CustomAttributes map[string]interface{} `json:"custom_attributes,omitempty"`
	InvalidLoginAttempts *int32 `json:"invalid_login_attempts,omitempty"`
	LockedUntil *string `json:"locked_until,omitempty"`
	PasswordChangedAt *string `json:"password_changed_at,omitempty"`
	// The password to set for a user.
	Password *string `json:"password,omitempty"`
	// Required if the password is being set.
	PasswordConfirmation *string `json:"password_confirmation,omitempty"`
	// Use this when importing a password that's already hashed. Prepend the salt value to the cleartext password value before SHA-256-encoding it
	PasswordAlgorithm *string `json:"password_algorithm,omitempty"`
	// The salt value used with the password_algorithm.
	Salt *string `json:"salt,omitempty"`
}

// NewUser instantiates a new User object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUser() *User {
	this := User{}
	return &this
}

// NewUserWithDefaults instantiates a new User object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWithDefaults() *User {
	this := User{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *User) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *User) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *User) SetId(v int32) {
	o.Id = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *User) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *User) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *User) SetUsername(v string) {
	o.Username = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *User) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *User) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *User) SetEmail(v string) {
	o.Email = &v
}

// GetFirstname returns the Firstname field value if set, zero value otherwise.
func (o *User) GetFirstname() string {
	if o == nil || o.Firstname == nil {
		var ret string
		return ret
	}
	return *o.Firstname
}

// GetFirstnameOk returns a tuple with the Firstname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetFirstnameOk() (*string, bool) {
	if o == nil || o.Firstname == nil {
		return nil, false
	}
	return o.Firstname, true
}

// HasFirstname returns a boolean if a field has been set.
func (o *User) HasFirstname() bool {
	if o != nil && o.Firstname != nil {
		return true
	}

	return false
}

// SetFirstname gets a reference to the given string and assigns it to the Firstname field.
func (o *User) SetFirstname(v string) {
	o.Firstname = &v
}

// GetLastname returns the Lastname field value if set, zero value otherwise.
func (o *User) GetLastname() string {
	if o == nil || o.Lastname == nil {
		var ret string
		return ret
	}
	return *o.Lastname
}

// GetLastnameOk returns a tuple with the Lastname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetLastnameOk() (*string, bool) {
	if o == nil || o.Lastname == nil {
		return nil, false
	}
	return o.Lastname, true
}

// HasLastname returns a boolean if a field has been set.
func (o *User) HasLastname() bool {
	if o != nil && o.Lastname != nil {
		return true
	}

	return false
}

// SetLastname gets a reference to the given string and assigns it to the Lastname field.
func (o *User) SetLastname(v string) {
	o.Lastname = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *User) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *User) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *User) SetTitle(v string) {
	o.Title = &v
}

// GetDepartment returns the Department field value if set, zero value otherwise.
func (o *User) GetDepartment() string {
	if o == nil || o.Department == nil {
		var ret string
		return ret
	}
	return *o.Department
}

// GetDepartmentOk returns a tuple with the Department field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetDepartmentOk() (*string, bool) {
	if o == nil || o.Department == nil {
		return nil, false
	}
	return o.Department, true
}

// HasDepartment returns a boolean if a field has been set.
func (o *User) HasDepartment() bool {
	if o != nil && o.Department != nil {
		return true
	}

	return false
}

// SetDepartment gets a reference to the given string and assigns it to the Department field.
func (o *User) SetDepartment(v string) {
	o.Department = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *User) GetCompany() string {
	if o == nil || o.Company == nil {
		var ret string
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetCompanyOk() (*string, bool) {
	if o == nil || o.Company == nil {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *User) HasCompany() bool {
	if o != nil && o.Company != nil {
		return true
	}

	return false
}

// SetCompany gets a reference to the given string and assigns it to the Company field.
func (o *User) SetCompany(v string) {
	o.Company = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *User) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *User) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *User) SetComment(v string) {
	o.Comment = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *User) GetGroupId() int32 {
	if o == nil || o.GroupId == nil {
		var ret int32
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetGroupIdOk() (*int32, bool) {
	if o == nil || o.GroupId == nil {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *User) HasGroupId() bool {
	if o != nil && o.GroupId != nil {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given int32 and assigns it to the GroupId field.
func (o *User) SetGroupId(v int32) {
	o.GroupId = &v
}

// GetRoleIds returns the RoleIds field value if set, zero value otherwise.
func (o *User) GetRoleIds() []int32 {
	if o == nil || o.RoleIds == nil {
		var ret []int32
		return ret
	}
	return o.RoleIds
}

// GetRoleIdsOk returns a tuple with the RoleIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetRoleIdsOk() ([]int32, bool) {
	if o == nil || o.RoleIds == nil {
		return nil, false
	}
	return o.RoleIds, true
}

// HasRoleIds returns a boolean if a field has been set.
func (o *User) HasRoleIds() bool {
	if o != nil && o.RoleIds != nil {
		return true
	}

	return false
}

// SetRoleIds gets a reference to the given []int32 and assigns it to the RoleIds field.
func (o *User) SetRoleIds(v []int32) {
	o.RoleIds = v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *User) GetPhone() string {
	if o == nil || o.Phone == nil {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetPhoneOk() (*string, bool) {
	if o == nil || o.Phone == nil {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *User) HasPhone() bool {
	if o != nil && o.Phone != nil {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *User) SetPhone(v string) {
	o.Phone = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *User) GetState() int32 {
	if o == nil || o.State == nil {
		var ret int32
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetStateOk() (*int32, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *User) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given int32 and assigns it to the State field.
func (o *User) SetState(v int32) {
	o.State = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *User) GetStatus() int32 {
	if o == nil || o.Status == nil {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetStatusOk() (*int32, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *User) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *User) SetStatus(v int32) {
	o.Status = &v
}

// GetDirectoryId returns the DirectoryId field value if set, zero value otherwise.
func (o *User) GetDirectoryId() int32 {
	if o == nil || o.DirectoryId == nil {
		var ret int32
		return ret
	}
	return *o.DirectoryId
}

// GetDirectoryIdOk returns a tuple with the DirectoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetDirectoryIdOk() (*int32, bool) {
	if o == nil || o.DirectoryId == nil {
		return nil, false
	}
	return o.DirectoryId, true
}

// HasDirectoryId returns a boolean if a field has been set.
func (o *User) HasDirectoryId() bool {
	if o != nil && o.DirectoryId != nil {
		return true
	}

	return false
}

// SetDirectoryId gets a reference to the given int32 and assigns it to the DirectoryId field.
func (o *User) SetDirectoryId(v int32) {
	o.DirectoryId = &v
}

// GetTrustedIdpId returns the TrustedIdpId field value if set, zero value otherwise.
func (o *User) GetTrustedIdpId() int32 {
	if o == nil || o.TrustedIdpId == nil {
		var ret int32
		return ret
	}
	return *o.TrustedIdpId
}

// GetTrustedIdpIdOk returns a tuple with the TrustedIdpId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetTrustedIdpIdOk() (*int32, bool) {
	if o == nil || o.TrustedIdpId == nil {
		return nil, false
	}
	return o.TrustedIdpId, true
}

// HasTrustedIdpId returns a boolean if a field has been set.
func (o *User) HasTrustedIdpId() bool {
	if o != nil && o.TrustedIdpId != nil {
		return true
	}

	return false
}

// SetTrustedIdpId gets a reference to the given int32 and assigns it to the TrustedIdpId field.
func (o *User) SetTrustedIdpId(v int32) {
	o.TrustedIdpId = &v
}

// GetManagerAdId returns the ManagerAdId field value if set, zero value otherwise.
func (o *User) GetManagerAdId() string {
	if o == nil || o.ManagerAdId == nil {
		var ret string
		return ret
	}
	return *o.ManagerAdId
}

// GetManagerAdIdOk returns a tuple with the ManagerAdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetManagerAdIdOk() (*string, bool) {
	if o == nil || o.ManagerAdId == nil {
		return nil, false
	}
	return o.ManagerAdId, true
}

// HasManagerAdId returns a boolean if a field has been set.
func (o *User) HasManagerAdId() bool {
	if o != nil && o.ManagerAdId != nil {
		return true
	}

	return false
}

// SetManagerAdId gets a reference to the given string and assigns it to the ManagerAdId field.
func (o *User) SetManagerAdId(v string) {
	o.ManagerAdId = &v
}

// GetManagerUserId returns the ManagerUserId field value if set, zero value otherwise.
func (o *User) GetManagerUserId() string {
	if o == nil || o.ManagerUserId == nil {
		var ret string
		return ret
	}
	return *o.ManagerUserId
}

// GetManagerUserIdOk returns a tuple with the ManagerUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetManagerUserIdOk() (*string, bool) {
	if o == nil || o.ManagerUserId == nil {
		return nil, false
	}
	return o.ManagerUserId, true
}

// HasManagerUserId returns a boolean if a field has been set.
func (o *User) HasManagerUserId() bool {
	if o != nil && o.ManagerUserId != nil {
		return true
	}

	return false
}

// SetManagerUserId gets a reference to the given string and assigns it to the ManagerUserId field.
func (o *User) SetManagerUserId(v string) {
	o.ManagerUserId = &v
}

// GetSamaccountName returns the SamaccountName field value if set, zero value otherwise.
func (o *User) GetSamaccountName() string {
	if o == nil || o.SamaccountName == nil {
		var ret string
		return ret
	}
	return *o.SamaccountName
}

// GetSamaccountNameOk returns a tuple with the SamaccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetSamaccountNameOk() (*string, bool) {
	if o == nil || o.SamaccountName == nil {
		return nil, false
	}
	return o.SamaccountName, true
}

// HasSamaccountName returns a boolean if a field has been set.
func (o *User) HasSamaccountName() bool {
	if o != nil && o.SamaccountName != nil {
		return true
	}

	return false
}

// SetSamaccountName gets a reference to the given string and assigns it to the SamaccountName field.
func (o *User) SetSamaccountName(v string) {
	o.SamaccountName = &v
}

// GetMemberOf returns the MemberOf field value if set, zero value otherwise.
func (o *User) GetMemberOf() string {
	if o == nil || o.MemberOf == nil {
		var ret string
		return ret
	}
	return *o.MemberOf
}

// GetMemberOfOk returns a tuple with the MemberOf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetMemberOfOk() (*string, bool) {
	if o == nil || o.MemberOf == nil {
		return nil, false
	}
	return o.MemberOf, true
}

// HasMemberOf returns a boolean if a field has been set.
func (o *User) HasMemberOf() bool {
	if o != nil && o.MemberOf != nil {
		return true
	}

	return false
}

// SetMemberOf gets a reference to the given string and assigns it to the MemberOf field.
func (o *User) SetMemberOf(v string) {
	o.MemberOf = &v
}

// GetUserprincipalname returns the Userprincipalname field value if set, zero value otherwise.
func (o *User) GetUserprincipalname() string {
	if o == nil || o.Userprincipalname == nil {
		var ret string
		return ret
	}
	return *o.Userprincipalname
}

// GetUserprincipalnameOk returns a tuple with the Userprincipalname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetUserprincipalnameOk() (*string, bool) {
	if o == nil || o.Userprincipalname == nil {
		return nil, false
	}
	return o.Userprincipalname, true
}

// HasUserprincipalname returns a boolean if a field has been set.
func (o *User) HasUserprincipalname() bool {
	if o != nil && o.Userprincipalname != nil {
		return true
	}

	return false
}

// SetUserprincipalname gets a reference to the given string and assigns it to the Userprincipalname field.
func (o *User) SetUserprincipalname(v string) {
	o.Userprincipalname = &v
}

// GetDistinguishedName returns the DistinguishedName field value if set, zero value otherwise.
func (o *User) GetDistinguishedName() string {
	if o == nil || o.DistinguishedName == nil {
		var ret string
		return ret
	}
	return *o.DistinguishedName
}

// GetDistinguishedNameOk returns a tuple with the DistinguishedName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetDistinguishedNameOk() (*string, bool) {
	if o == nil || o.DistinguishedName == nil {
		return nil, false
	}
	return o.DistinguishedName, true
}

// HasDistinguishedName returns a boolean if a field has been set.
func (o *User) HasDistinguishedName() bool {
	if o != nil && o.DistinguishedName != nil {
		return true
	}

	return false
}

// SetDistinguishedName gets a reference to the given string and assigns it to the DistinguishedName field.
func (o *User) SetDistinguishedName(v string) {
	o.DistinguishedName = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *User) GetExternalId() string {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetExternalIdOk() (*string, bool) {
	if o == nil || o.ExternalId == nil {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *User) HasExternalId() bool {
	if o != nil && o.ExternalId != nil {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *User) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetActivatedAt returns the ActivatedAt field value if set, zero value otherwise.
func (o *User) GetActivatedAt() string {
	if o == nil || o.ActivatedAt == nil {
		var ret string
		return ret
	}
	return *o.ActivatedAt
}

// GetActivatedAtOk returns a tuple with the ActivatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetActivatedAtOk() (*string, bool) {
	if o == nil || o.ActivatedAt == nil {
		return nil, false
	}
	return o.ActivatedAt, true
}

// HasActivatedAt returns a boolean if a field has been set.
func (o *User) HasActivatedAt() bool {
	if o != nil && o.ActivatedAt != nil {
		return true
	}

	return false
}

// SetActivatedAt gets a reference to the given string and assigns it to the ActivatedAt field.
func (o *User) SetActivatedAt(v string) {
	o.ActivatedAt = &v
}

// GetLastLogin returns the LastLogin field value if set, zero value otherwise.
func (o *User) GetLastLogin() string {
	if o == nil || o.LastLogin == nil {
		var ret string
		return ret
	}
	return *o.LastLogin
}

// GetLastLoginOk returns a tuple with the LastLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetLastLoginOk() (*string, bool) {
	if o == nil || o.LastLogin == nil {
		return nil, false
	}
	return o.LastLogin, true
}

// HasLastLogin returns a boolean if a field has been set.
func (o *User) HasLastLogin() bool {
	if o != nil && o.LastLogin != nil {
		return true
	}

	return false
}

// SetLastLogin gets a reference to the given string and assigns it to the LastLogin field.
func (o *User) SetLastLogin(v string) {
	o.LastLogin = &v
}

// GetInvitationSentAt returns the InvitationSentAt field value if set, zero value otherwise.
func (o *User) GetInvitationSentAt() string {
	if o == nil || o.InvitationSentAt == nil {
		var ret string
		return ret
	}
	return *o.InvitationSentAt
}

// GetInvitationSentAtOk returns a tuple with the InvitationSentAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetInvitationSentAtOk() (*string, bool) {
	if o == nil || o.InvitationSentAt == nil {
		return nil, false
	}
	return o.InvitationSentAt, true
}

// HasInvitationSentAt returns a boolean if a field has been set.
func (o *User) HasInvitationSentAt() bool {
	if o != nil && o.InvitationSentAt != nil {
		return true
	}

	return false
}

// SetInvitationSentAt gets a reference to the given string and assigns it to the InvitationSentAt field.
func (o *User) SetInvitationSentAt(v string) {
	o.InvitationSentAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *User) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *User) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *User) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetPreferredLocaleCode returns the PreferredLocaleCode field value if set, zero value otherwise.
func (o *User) GetPreferredLocaleCode() string {
	if o == nil || o.PreferredLocaleCode == nil {
		var ret string
		return ret
	}
	return *o.PreferredLocaleCode
}

// GetPreferredLocaleCodeOk returns a tuple with the PreferredLocaleCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetPreferredLocaleCodeOk() (*string, bool) {
	if o == nil || o.PreferredLocaleCode == nil {
		return nil, false
	}
	return o.PreferredLocaleCode, true
}

// HasPreferredLocaleCode returns a boolean if a field has been set.
func (o *User) HasPreferredLocaleCode() bool {
	if o != nil && o.PreferredLocaleCode != nil {
		return true
	}

	return false
}

// SetPreferredLocaleCode gets a reference to the given string and assigns it to the PreferredLocaleCode field.
func (o *User) SetPreferredLocaleCode(v string) {
	o.PreferredLocaleCode = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *User) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *User) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *User) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetCustomAttributes returns the CustomAttributes field value if set, zero value otherwise.
func (o *User) GetCustomAttributes() map[string]interface{} {
	if o == nil || o.CustomAttributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomAttributes
}

// GetCustomAttributesOk returns a tuple with the CustomAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetCustomAttributesOk() (map[string]interface{}, bool) {
	if o == nil || o.CustomAttributes == nil {
		return nil, false
	}
	return o.CustomAttributes, true
}

// HasCustomAttributes returns a boolean if a field has been set.
func (o *User) HasCustomAttributes() bool {
	if o != nil && o.CustomAttributes != nil {
		return true
	}

	return false
}

// SetCustomAttributes gets a reference to the given map[string]interface{} and assigns it to the CustomAttributes field.
func (o *User) SetCustomAttributes(v map[string]interface{}) {
	o.CustomAttributes = v
}

// GetInvalidLoginAttempts returns the InvalidLoginAttempts field value if set, zero value otherwise.
func (o *User) GetInvalidLoginAttempts() int32 {
	if o == nil || o.InvalidLoginAttempts == nil {
		var ret int32
		return ret
	}
	return *o.InvalidLoginAttempts
}

// GetInvalidLoginAttemptsOk returns a tuple with the InvalidLoginAttempts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetInvalidLoginAttemptsOk() (*int32, bool) {
	if o == nil || o.InvalidLoginAttempts == nil {
		return nil, false
	}
	return o.InvalidLoginAttempts, true
}

// HasInvalidLoginAttempts returns a boolean if a field has been set.
func (o *User) HasInvalidLoginAttempts() bool {
	if o != nil && o.InvalidLoginAttempts != nil {
		return true
	}

	return false
}

// SetInvalidLoginAttempts gets a reference to the given int32 and assigns it to the InvalidLoginAttempts field.
func (o *User) SetInvalidLoginAttempts(v int32) {
	o.InvalidLoginAttempts = &v
}

// GetLockedUntil returns the LockedUntil field value if set, zero value otherwise.
func (o *User) GetLockedUntil() string {
	if o == nil || o.LockedUntil == nil {
		var ret string
		return ret
	}
	return *o.LockedUntil
}

// GetLockedUntilOk returns a tuple with the LockedUntil field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetLockedUntilOk() (*string, bool) {
	if o == nil || o.LockedUntil == nil {
		return nil, false
	}
	return o.LockedUntil, true
}

// HasLockedUntil returns a boolean if a field has been set.
func (o *User) HasLockedUntil() bool {
	if o != nil && o.LockedUntil != nil {
		return true
	}

	return false
}

// SetLockedUntil gets a reference to the given string and assigns it to the LockedUntil field.
func (o *User) SetLockedUntil(v string) {
	o.LockedUntil = &v
}

// GetPasswordChangedAt returns the PasswordChangedAt field value if set, zero value otherwise.
func (o *User) GetPasswordChangedAt() string {
	if o == nil || o.PasswordChangedAt == nil {
		var ret string
		return ret
	}
	return *o.PasswordChangedAt
}

// GetPasswordChangedAtOk returns a tuple with the PasswordChangedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetPasswordChangedAtOk() (*string, bool) {
	if o == nil || o.PasswordChangedAt == nil {
		return nil, false
	}
	return o.PasswordChangedAt, true
}

// HasPasswordChangedAt returns a boolean if a field has been set.
func (o *User) HasPasswordChangedAt() bool {
	if o != nil && o.PasswordChangedAt != nil {
		return true
	}

	return false
}

// SetPasswordChangedAt gets a reference to the given string and assigns it to the PasswordChangedAt field.
func (o *User) SetPasswordChangedAt(v string) {
	o.PasswordChangedAt = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *User) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *User) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *User) SetPassword(v string) {
	o.Password = &v
}

// GetPasswordConfirmation returns the PasswordConfirmation field value if set, zero value otherwise.
func (o *User) GetPasswordConfirmation() string {
	if o == nil || o.PasswordConfirmation == nil {
		var ret string
		return ret
	}
	return *o.PasswordConfirmation
}

// GetPasswordConfirmationOk returns a tuple with the PasswordConfirmation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetPasswordConfirmationOk() (*string, bool) {
	if o == nil || o.PasswordConfirmation == nil {
		return nil, false
	}
	return o.PasswordConfirmation, true
}

// HasPasswordConfirmation returns a boolean if a field has been set.
func (o *User) HasPasswordConfirmation() bool {
	if o != nil && o.PasswordConfirmation != nil {
		return true
	}

	return false
}

// SetPasswordConfirmation gets a reference to the given string and assigns it to the PasswordConfirmation field.
func (o *User) SetPasswordConfirmation(v string) {
	o.PasswordConfirmation = &v
}

// GetPasswordAlgorithm returns the PasswordAlgorithm field value if set, zero value otherwise.
func (o *User) GetPasswordAlgorithm() string {
	if o == nil || o.PasswordAlgorithm == nil {
		var ret string
		return ret
	}
	return *o.PasswordAlgorithm
}

// GetPasswordAlgorithmOk returns a tuple with the PasswordAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetPasswordAlgorithmOk() (*string, bool) {
	if o == nil || o.PasswordAlgorithm == nil {
		return nil, false
	}
	return o.PasswordAlgorithm, true
}

// HasPasswordAlgorithm returns a boolean if a field has been set.
func (o *User) HasPasswordAlgorithm() bool {
	if o != nil && o.PasswordAlgorithm != nil {
		return true
	}

	return false
}

// SetPasswordAlgorithm gets a reference to the given string and assigns it to the PasswordAlgorithm field.
func (o *User) SetPasswordAlgorithm(v string) {
	o.PasswordAlgorithm = &v
}

// GetSalt returns the Salt field value if set, zero value otherwise.
func (o *User) GetSalt() string {
	if o == nil || o.Salt == nil {
		var ret string
		return ret
	}
	return *o.Salt
}

// GetSaltOk returns a tuple with the Salt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetSaltOk() (*string, bool) {
	if o == nil || o.Salt == nil {
		return nil, false
	}
	return o.Salt, true
}

// HasSalt returns a boolean if a field has been set.
func (o *User) HasSalt() bool {
	if o != nil && o.Salt != nil {
		return true
	}

	return false
}

// SetSalt gets a reference to the given string and assigns it to the Salt field.
func (o *User) SetSalt(v string) {
	o.Salt = &v
}

func (o User) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Firstname != nil {
		toSerialize["firstname"] = o.Firstname
	}
	if o.Lastname != nil {
		toSerialize["lastname"] = o.Lastname
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Department != nil {
		toSerialize["department"] = o.Department
	}
	if o.Company != nil {
		toSerialize["company"] = o.Company
	}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if o.GroupId != nil {
		toSerialize["group_id"] = o.GroupId
	}
	if o.RoleIds != nil {
		toSerialize["role_ids"] = o.RoleIds
	}
	if o.Phone != nil {
		toSerialize["phone"] = o.Phone
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.DirectoryId != nil {
		toSerialize["directory_id"] = o.DirectoryId
	}
	if o.TrustedIdpId != nil {
		toSerialize["trusted_idp_id"] = o.TrustedIdpId
	}
	if o.ManagerAdId != nil {
		toSerialize["manager_ad_id"] = o.ManagerAdId
	}
	if o.ManagerUserId != nil {
		toSerialize["manager_user_id"] = o.ManagerUserId
	}
	if o.SamaccountName != nil {
		toSerialize["samaccount_name"] = o.SamaccountName
	}
	if o.MemberOf != nil {
		toSerialize["member_of"] = o.MemberOf
	}
	if o.Userprincipalname != nil {
		toSerialize["userprincipalname"] = o.Userprincipalname
	}
	if o.DistinguishedName != nil {
		toSerialize["distinguished_name"] = o.DistinguishedName
	}
	if o.ExternalId != nil {
		toSerialize["external_id"] = o.ExternalId
	}
	if o.ActivatedAt != nil {
		toSerialize["activated_at"] = o.ActivatedAt
	}
	if o.LastLogin != nil {
		toSerialize["last_login"] = o.LastLogin
	}
	if o.InvitationSentAt != nil {
		toSerialize["invitation_sent_at"] = o.InvitationSentAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.PreferredLocaleCode != nil {
		toSerialize["preferred_locale_code"] = o.PreferredLocaleCode
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.CustomAttributes != nil {
		toSerialize["custom_attributes"] = o.CustomAttributes
	}
	if o.InvalidLoginAttempts != nil {
		toSerialize["invalid_login_attempts"] = o.InvalidLoginAttempts
	}
	if o.LockedUntil != nil {
		toSerialize["locked_until"] = o.LockedUntil
	}
	if o.PasswordChangedAt != nil {
		toSerialize["password_changed_at"] = o.PasswordChangedAt
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.PasswordConfirmation != nil {
		toSerialize["password_confirmation"] = o.PasswordConfirmation
	}
	if o.PasswordAlgorithm != nil {
		toSerialize["password_algorithm"] = o.PasswordAlgorithm
	}
	if o.Salt != nil {
		toSerialize["salt"] = o.Salt
	}
	return json.Marshal(toSerialize)
}

type NullableUser struct {
	value *User
	isSet bool
}

func (v NullableUser) Get() *User {
	return v.value
}

func (v *NullableUser) Set(val *User) {
	v.value = val
	v.isSet = true
}

func (v NullableUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUser(val *User) *NullableUser {
	return &NullableUser{value: val, isSet: true}
}

func (v NullableUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


