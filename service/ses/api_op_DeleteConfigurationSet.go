// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to delete a configuration set. Configuration sets enable
// you to publish email sending events. For information about using configuration
// sets, see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/monitor-sending-activity.html).
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/DeleteConfigurationSetRequest
type DeleteConfigurationSetInput struct {
	_ struct{} `type:"structure"`

	// The name of the configuration set to delete.
	//
	// ConfigurationSetName is a required field
	ConfigurationSetName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteConfigurationSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteConfigurationSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteConfigurationSetInput"}

	if s.ConfigurationSetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigurationSetName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// An empty element returned on a successful request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/DeleteConfigurationSetResponse
type DeleteConfigurationSetOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteConfigurationSetOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteConfigurationSet = "DeleteConfigurationSet"

// DeleteConfigurationSetRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Deletes a configuration set. Configuration sets enable you to publish email
// sending events. For information about using configuration sets, see the Amazon
// SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/monitor-sending-activity.html).
//
// You can execute this operation no more than once per second.
//
//    // Example sending a request using DeleteConfigurationSetRequest.
//    req := client.DeleteConfigurationSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/DeleteConfigurationSet
func (c *Client) DeleteConfigurationSetRequest(input *DeleteConfigurationSetInput) DeleteConfigurationSetRequest {
	op := &aws.Operation{
		Name:       opDeleteConfigurationSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteConfigurationSetInput{}
	}

	req := c.newRequest(op, input, &DeleteConfigurationSetOutput{})
	return DeleteConfigurationSetRequest{Request: req, Input: input, Copy: c.DeleteConfigurationSetRequest}
}

// DeleteConfigurationSetRequest is the request type for the
// DeleteConfigurationSet API operation.
type DeleteConfigurationSetRequest struct {
	*aws.Request
	Input *DeleteConfigurationSetInput
	Copy  func(*DeleteConfigurationSetInput) DeleteConfigurationSetRequest
}

// Send marshals and sends the DeleteConfigurationSet API request.
func (r DeleteConfigurationSetRequest) Send(ctx context.Context) (*DeleteConfigurationSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteConfigurationSetResponse{
		DeleteConfigurationSetOutput: r.Request.Data.(*DeleteConfigurationSetOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteConfigurationSetResponse is the response type for the
// DeleteConfigurationSet API operation.
type DeleteConfigurationSetResponse struct {
	*DeleteConfigurationSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteConfigurationSet request.
func (r *DeleteConfigurationSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}