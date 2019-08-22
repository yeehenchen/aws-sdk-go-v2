// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codecommit

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/BatchGetCommitsInput
type BatchGetCommitsInput struct {
	_ struct{} `type:"structure"`

	// The full commit IDs of the commits to get information about.
	//
	// You must supply the full SHAs of each commit. You cannot use shortened SHAs.
	//
	// CommitIds is a required field
	CommitIds []string `locationName:"commitIds" type:"list" required:"true"`

	// The name of the repository that contains the commits.
	//
	// RepositoryName is a required field
	RepositoryName *string `locationName:"repositoryName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s BatchGetCommitsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchGetCommitsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchGetCommitsInput"}

	if s.CommitIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("CommitIds"))
	}

	if s.RepositoryName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RepositoryName"))
	}
	if s.RepositoryName != nil && len(*s.RepositoryName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RepositoryName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/BatchGetCommitsOutput
type BatchGetCommitsOutput struct {
	_ struct{} `type:"structure"`

	// An array of commit data type objects, each of which contains information
	// about a specified commit.
	Commits []Commit `locationName:"commits" type:"list"`

	// Returns any commit IDs for which information could not be found. For example,
	// if one of the commit IDs was a shortened SHA or that commit was not found
	// in the specified repository, the ID will return an error object with additional
	// information.
	Errors []BatchGetCommitsError `locationName:"errors" type:"list"`
}

// String returns the string representation
func (s BatchGetCommitsOutput) String() string {
	return awsutil.Prettify(s)
}

const opBatchGetCommits = "BatchGetCommits"

// BatchGetCommitsRequest returns a request value for making API operation for
// AWS CodeCommit.
//
// Returns information about the contents of one or more commits in a repository.
//
//    // Example sending a request using BatchGetCommitsRequest.
//    req := client.BatchGetCommitsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/BatchGetCommits
func (c *Client) BatchGetCommitsRequest(input *BatchGetCommitsInput) BatchGetCommitsRequest {
	op := &aws.Operation{
		Name:       opBatchGetCommits,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &BatchGetCommitsInput{}
	}

	req := c.newRequest(op, input, &BatchGetCommitsOutput{})
	return BatchGetCommitsRequest{Request: req, Input: input, Copy: c.BatchGetCommitsRequest}
}

// BatchGetCommitsRequest is the request type for the
// BatchGetCommits API operation.
type BatchGetCommitsRequest struct {
	*aws.Request
	Input *BatchGetCommitsInput
	Copy  func(*BatchGetCommitsInput) BatchGetCommitsRequest
}

// Send marshals and sends the BatchGetCommits API request.
func (r BatchGetCommitsRequest) Send(ctx context.Context) (*BatchGetCommitsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchGetCommitsResponse{
		BatchGetCommitsOutput: r.Request.Data.(*BatchGetCommitsOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchGetCommitsResponse is the response type for the
// BatchGetCommits API operation.
type BatchGetCommitsResponse struct {
	*BatchGetCommitsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchGetCommits request.
func (r *BatchGetCommitsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}