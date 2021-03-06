// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package migrationhubiface provides an interface to enable mocking the AWS Migration Hub service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package migrationhubiface

import (
	"github.com/aws/aws-sdk-go-v2/service/migrationhub"
)

// ClientAPI provides an interface to enable mocking the
// migrationhub.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Migration Hub.
//    func myFunc(svc migrationhubiface.ClientAPI) bool {
//        // Make svc.AssociateCreatedArtifact request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := migrationhub.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        migrationhubiface.ClientPI
//    }
//    func (m *mockClientClient) AssociateCreatedArtifact(input *migrationhub.AssociateCreatedArtifactInput) (*migrationhub.AssociateCreatedArtifactOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockClientClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type ClientAPI interface {
	AssociateCreatedArtifactRequest(*migrationhub.AssociateCreatedArtifactInput) migrationhub.AssociateCreatedArtifactRequest

	AssociateDiscoveredResourceRequest(*migrationhub.AssociateDiscoveredResourceInput) migrationhub.AssociateDiscoveredResourceRequest

	CreateProgressUpdateStreamRequest(*migrationhub.CreateProgressUpdateStreamInput) migrationhub.CreateProgressUpdateStreamRequest

	DeleteProgressUpdateStreamRequest(*migrationhub.DeleteProgressUpdateStreamInput) migrationhub.DeleteProgressUpdateStreamRequest

	DescribeApplicationStateRequest(*migrationhub.DescribeApplicationStateInput) migrationhub.DescribeApplicationStateRequest

	DescribeMigrationTaskRequest(*migrationhub.DescribeMigrationTaskInput) migrationhub.DescribeMigrationTaskRequest

	DisassociateCreatedArtifactRequest(*migrationhub.DisassociateCreatedArtifactInput) migrationhub.DisassociateCreatedArtifactRequest

	DisassociateDiscoveredResourceRequest(*migrationhub.DisassociateDiscoveredResourceInput) migrationhub.DisassociateDiscoveredResourceRequest

	ImportMigrationTaskRequest(*migrationhub.ImportMigrationTaskInput) migrationhub.ImportMigrationTaskRequest

	ListCreatedArtifactsRequest(*migrationhub.ListCreatedArtifactsInput) migrationhub.ListCreatedArtifactsRequest

	ListDiscoveredResourcesRequest(*migrationhub.ListDiscoveredResourcesInput) migrationhub.ListDiscoveredResourcesRequest

	ListMigrationTasksRequest(*migrationhub.ListMigrationTasksInput) migrationhub.ListMigrationTasksRequest

	ListProgressUpdateStreamsRequest(*migrationhub.ListProgressUpdateStreamsInput) migrationhub.ListProgressUpdateStreamsRequest

	NotifyApplicationStateRequest(*migrationhub.NotifyApplicationStateInput) migrationhub.NotifyApplicationStateRequest

	NotifyMigrationTaskStateRequest(*migrationhub.NotifyMigrationTaskStateInput) migrationhub.NotifyMigrationTaskStateRequest

	PutResourceAttributesRequest(*migrationhub.PutResourceAttributesInput) migrationhub.PutResourceAttributesRequest
}

var _ ClientAPI = (*migrationhub.Client)(nil)
