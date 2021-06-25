// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// The filters applied to Data Store query.
type DatastoreFilter struct {

	// A filter that allows the user to set cutoff dates for records. All Data Stores
	// created after the specified date will be included in the results.
	CreatedAfter *time.Time

	// A filter that allows the user to set cutoff dates for records. All Data Stores
	// created before the specified date will be included in the results.
	CreatedBefore *time.Time

	// Allows the user to filter Data Store results by name.
	DatastoreName *string

	// Allows the user to filter Data Store results by status.
	DatastoreStatus DatastoreStatus
}

// Displays the properties of the Data Store, including the ID, Arn, name, and the
// status of the Data Store.
type DatastoreProperties struct {

	// The Amazon Resource Name used in the creation of the Data Store.
	//
	// This member is required.
	DatastoreArn *string

	// The AWS endpoint for the Data Store. Each Data Store will have it's own endpoint
	// with Data Store ID in the endpoint URL.
	//
	// This member is required.
	DatastoreEndpoint *string

	// The AWS-generated ID number for the Data Store.
	//
	// This member is required.
	DatastoreId *string

	// The status of the Data Store. Possible statuses are 'CREATING', 'ACTIVE',
	// 'DELETING', or 'DELETED'.
	//
	// This member is required.
	DatastoreStatus DatastoreStatus

	// The FHIR version. Only R4 version data is supported.
	//
	// This member is required.
	DatastoreTypeVersion FHIRVersion

	// The time that a Data Store was created.
	CreatedAt *time.Time

	// The user-generated name for the Data Store.
	DatastoreName *string

	// The preloaded data configuration for the Data Store. Only data preloaded from
	// Synthea is supported.
	PreloadDataConfig *PreloadDataConfig
}

// The properties of a FHIR export job, including the ID, ARN, name, and the status
// of the job.
type ExportJobProperties struct {

	// The AWS generated ID for the Data Store from which files are being exported for
	// an export job.
	//
	// This member is required.
	DatastoreId *string

	// The AWS generated ID for an export job.
	//
	// This member is required.
	JobId *string

	// The status of a FHIR export job. Possible statuses are SUBMITTED, IN_PROGRESS,
	// COMPLETED, or FAILED.
	//
	// This member is required.
	JobStatus JobStatus

	// The output data configuration that was supplied when the export job was created.
	//
	// This member is required.
	OutputDataConfig OutputDataConfig

	// The time an export job was initiated.
	//
	// This member is required.
	SubmitTime *time.Time

	// The Amazon Resource Name used during the initiation of the job.
	DataAccessRoleArn *string

	// The time an export job completed.
	EndTime *time.Time

	// The user generated name for an export job.
	JobName *string

	// An explanation of any errors that may have occurred during the export job.
	Message *string
}

// Displays the properties of the import job, including the ID, Arn, Name, and the
// status of the Data Store.
type ImportJobProperties struct {

	// The datastore id used when the Import job was created.
	//
	// This member is required.
	DatastoreId *string

	// The input data configuration that was supplied when the Import job was created.
	//
	// This member is required.
	InputDataConfig InputDataConfig

	// The AWS-generated id number for the Import job.
	//
	// This member is required.
	JobId *string

	// The job status for an Import job. Possible statuses are SUBMITTED, IN_PROGRESS,
	// COMPLETED, FAILED.
	//
	// This member is required.
	JobStatus JobStatus

	// The time that the Import job was submitted for processing.
	//
	// This member is required.
	SubmitTime *time.Time

	// The Amazon Resource Name (ARN) that gives Amazon HealthLake access to your input
	// data.
	DataAccessRoleArn *string

	// The time that the Import job was completed.
	EndTime *time.Time

	// The user-generated name for an Import job.
	JobName *string

	// An explanation of any errors that may have occurred during the FHIR import job.
	Message *string
}

// The input properties for an import job.
//
// The following types satisfy this interface:
//  InputDataConfigMemberS3Uri
type InputDataConfig interface {
	isInputDataConfig()
}

// The S3Uri is the user specified S3 location of the FHIR data to be imported into
// Amazon HealthLake.
type InputDataConfigMemberS3Uri struct {
	Value string
}

func (*InputDataConfigMemberS3Uri) isInputDataConfig() {}

// The output data configuration that was supplied when the export job was created.
//
// The following types satisfy this interface:
//  OutputDataConfigMemberS3Uri
type OutputDataConfig interface {
	isOutputDataConfig()
}

// The S3Uri is the user specified S3 location to which data will be exported from
// a FHIR Data Store.
type OutputDataConfigMemberS3Uri struct {
	Value string
}

func (*OutputDataConfigMemberS3Uri) isOutputDataConfig() {}

// The input properties for the preloaded Data Store. Only data preloaded from
// Synthea is supported.
type PreloadDataConfig struct {

	// The type of preloaded data. Only Synthea preloaded data is supported.
	//
	// This member is required.
	PreloadDataType PreloadDataType
}

// UnknownUnionMember is returned when a union member is returned over the wire,
// but has an unknown tag.
type UnknownUnionMember struct {
	Tag   string
	Value []byte
}

func (*UnknownUnionMember) isInputDataConfig()  {}
func (*UnknownUnionMember) isOutputDataConfig() {}