// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type SchemeBasicAuth struct {
	Username string `security:"name=username,env=cloudinary_username"`
	Password string `security:"name=password,env=cloudinary_password"`
}

func (o *SchemeBasicAuth) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

func (o *SchemeBasicAuth) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}
