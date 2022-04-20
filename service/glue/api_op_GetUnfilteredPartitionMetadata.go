// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

func (c *Client) GetUnfilteredPartitionMetadata(ctx context.Context, params *GetUnfilteredPartitionMetadataInput, optFns ...func(*Options)) (*GetUnfilteredPartitionMetadataOutput, error) {
	if params == nil {
		params = &GetUnfilteredPartitionMetadataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetUnfilteredPartitionMetadata", params, optFns, c.addOperationGetUnfilteredPartitionMetadataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetUnfilteredPartitionMetadataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetUnfilteredPartitionMetadataInput struct {

	// This member is required.
	CatalogId *string

	// This member is required.
	DatabaseName *string

	// This member is required.
	PartitionValues []string

	// This member is required.
	SupportedPermissionTypes []types.PermissionType

	// This member is required.
	TableName *string

	// A structure containing information for audit.
	AuditContext *types.AuditContext

	noSmithyDocumentSerde
}

type GetUnfilteredPartitionMetadataOutput struct {
	AuthorizedColumns []string

	IsRegisteredWithLakeFormation bool

	// Represents a slice of table data.
	Partition *types.Partition

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetUnfilteredPartitionMetadataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetUnfilteredPartitionMetadata{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetUnfilteredPartitionMetadata{}, middleware.After)
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
	if err = addOpGetUnfilteredPartitionMetadataValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetUnfilteredPartitionMetadata(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetUnfilteredPartitionMetadata(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "GetUnfilteredPartitionMetadata",
	}
}