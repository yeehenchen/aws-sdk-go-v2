// Code generated by smithy-go-codegen DO NOT EDIT.

package synthetics

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Permanently deletes the specified canary. When you delete a canary, resources
// used and created by the canary are not automatically deleted. After you delete a
// canary that you do not intend to use again, you should also delete the
// following:
//
//     * The Lambda functions and layers used by this canary. These
// have the prefix cwsyn-MyCanaryName .
//
//     * The CloudWatch alarms created for
// this canary. These alarms have a name of Synthetics-SharpDrop-Alarm-MyCanaryName
// .
//
//     * Amazon S3 objects and buckets, such as the canary's artifact
// location.
//
//     * IAM roles created for the canary. If they were created in the
// console, these roles have the name
// role/service-role/CloudWatchSyntheticsRole-MyCanaryName .
//
//     * CloudWatch Logs
// log groups created for the canary. These logs groups have the name
// /aws/lambda/cwsyn-MyCanaryName .
//
//     <p>Before you delete a canary, you might
// want to use <code>GetCanary</code> to display the information about this canary.
// Make note of the information returned by this operation so that you can delete
// these resources after you delete the canary.</p>
func (c *Client) DeleteCanary(ctx context.Context, params *DeleteCanaryInput, optFns ...func(*Options)) (*DeleteCanaryOutput, error) {
	stack := middleware.NewStack("DeleteCanary", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteCanaryMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpDeleteCanaryValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteCanary(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "DeleteCanary",
			Err:           err,
		}
	}
	out := result.(*DeleteCanaryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteCanaryInput struct {
	// The name of the canary that you want to delete. To find the names of your
	// canaries, use DescribeCanaries
	// (https://docs.aws.amazon.com/AmazonSynthetics/latest/APIReference/API_DescribeCanaries.html).
	Name *string
}

type DeleteCanaryOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteCanaryMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteCanary{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteCanary{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteCanary(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "synthetics",
		OperationName: "DeleteCanary",
	}
}