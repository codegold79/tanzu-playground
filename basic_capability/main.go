package main

import (
	"fmt"

	capdiscovery "github.com/vmware-tanzu/tanzu-framework/pkg/v1/sdk/capabilities/discovery"
	"github.com/vmware-tanzu/tanzu-framework/pkg/v1/tkg/kappclient"
)

func main() {
	cfg, err := kappclient.GetKubeConfig("")
	if err != nil {
		fmt.Printf("get kubeconfig: %v\n", err)
	}

	clusterQueryClient, err := capdiscovery.NewClusterQueryClientForConfig(cfg)
	if err != nil {
		fmt.Printf("create new cluster query builder instance: %v", err)
	}

	apiGroup := capdiscovery.
		Group("resource-empty-string", "autoscaling").
		WithVersions("v1").
		WithResource("")
	fmt.Printf("Query API Group:\n %#v\n", apiGroup)

	query := clusterQueryClient.Query(apiGroup)
	_, err = query.Execute()
	if err != nil {
		fmt.Printf("execute query: %v\n", err)
	}

	results := query.Results()
	for _, v := range results {
		fmt.Printf("Result: %#v\n", v)
	}

	// Output:
	// Query API Group:
	//  &discovery.QueryGVR{name:"resource-empty-string", group:"autoscaling", resource:"", versions:[]string{"v1"}, unmatchedGVRs:[]string(nil)}
	//  Result: &discovery.QueryResult{Found:true, NotFoundReason:""}
}
