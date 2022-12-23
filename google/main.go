package google

import (
	"context"
	"fmt"
	"io"
	"log"

	iam "google.golang.org/api/iam/v1"
)

func listKeys(w io.Writer, serviceAccountEmail string) ([]*iam.ServiceAccountKey, error) {
	ctx := context.Background()
	service, err := iam.NewService(ctx)
	if err != nil {
		return nil, fmt.Errorf("iam.NewService: %v", err)
	}

	resource := "projects/-/serviceAccounts/" + serviceAccountEmail
	response, err := service.Projects.ServiceAccounts.Keys.List(resource).Do()
	if err != nil {
		return nil, fmt.Errorf("Projects.ServiceAccounts.Keys.List: %v", err)
	}
	for _, key := range response.Keys {
		fmt.Fprintf(w, "Listing key: %v", key.Name)
	}
	return response.Keys, nil
}

func New() {
	ctx := context.Background()

	iamService, err := iam.NewService(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// The resource name of the service account in the following format:
	// `projects/{PROJECT_ID}/serviceAccounts/{	}`.
	// Using `-` as a wildcard for the `PROJECT_ID`, will infer the project from
	// the account. The `ACCOUNT` value can be the `email` address or the
	// `unique_id` of the service account.
	name := "projects/tiinhr/serviceAccounts/b2aa6dfb0bacc513a0be37e498d58325fe57cd45" // TODO: Update placeholder value.

	resp, err := iamService.Projects.ServiceAccounts.Keys.List(name).Context(ctx).Do()
	if err != nil {
		log.Fatal(err)
	}

	// TODO: Change code below to process the `resp` object:
	fmt.Printf("%#v\n", resp)
	// credByBytes, err:=

}
