// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package globals

type Globals struct {
	// The name of your Cloudinary cloud
	CloudName *string `pathParam:"style=simple,explode=false,name=cloud_name"`
}

func (o *Globals) GetCloudName() *string {
	if o == nil {
		return nil
	}
	return o.CloudName
}
