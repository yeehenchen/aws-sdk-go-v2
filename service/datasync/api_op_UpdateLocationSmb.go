// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates some of the parameters of a previously created location for Server
// Message Block (SMB) file system access. For information about creating an SMB
// location, see Creating a location for SMB
// (https://docs.aws.amazon.com/datasync/latest/userguide/create-smb-location.html).
func (c *Client) UpdateLocationSmb(ctx context.Context, params *UpdateLocationSmbInput, optFns ...func(*Options)) (*UpdateLocationSmbOutput, error) {
	if params == nil {
		params = &UpdateLocationSmbInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateLocationSmb", params, optFns, c.addOperationUpdateLocationSmbMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateLocationSmbOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateLocationSmbInput struct {

	// The Amazon Resource Name (ARN) of the SMB location to update.
	//
	// This member is required.
	LocationArn *string

	// The Amazon Resource Names (ARNs) of agents to use for a Simple Message Block
	// (SMB) location.
	AgentArns []string

	// The name of the Windows domain that the SMB server belongs to.
	Domain *string

	// Represents the mount options that are available for DataSync to access an SMB
	// location.
	MountOptions *types.SmbMountOptions

	// The password of the user who can mount the share has the permissions to access
	// files and folders in the SMB share.
	Password *string

	// The subdirectory in the SMB file system that is used to read data from the SMB
	// source location or write data to the SMB destination. The SMB path should be a
	// path that's exported by the SMB server, or a subdirectory of that path. The path
	// should be such that it can be mounted by other SMB clients in your network.
	// Subdirectory must be specified with forward slashes. For example,
	// /path/to/folder. To transfer all the data in the folder that you specified,
	// DataSync must have permissions to mount the SMB share and to access all the data
	// in that share. To ensure this, do either of the following:
	//
	// * Ensure that the
	// user/password specified belongs to the user who can mount the share and who has
	// the appropriate permissions for all of the files and directories that you want
	// DataSync to access.
	//
	// * Use credentials of a member of the Backup Operators group
	// to mount the share.
	//
	// Doing either of these options enables the agent to access
	// the data. For the agent to access directories, you must also enable all execute
	// access.
	Subdirectory *string

	// The user who can mount the share has the permissions to access files and folders
	// in the SMB share.
	User *string

	noSmithyDocumentSerde
}

type UpdateLocationSmbOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateLocationSmbMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateLocationSmb{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateLocationSmb{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpUpdateLocationSmbValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateLocationSmb(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opUpdateLocationSmb(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datasync",
		OperationName: "UpdateLocationSmb",
	}
}