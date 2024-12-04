package fake

import (
	"fmt"
	"log"

	//cmd "piranhas/platform/pirctl/cmd"
	//pk8s "piranhas/platform/pirctl/pkg/k8s"
	"context"
	"k8s.io/apimachinery/pkg/runtime"
	//"k8s.io/client-go/dynamic"
	dyfake "k8s.io/client-go/dynamic/fake"

	//"k8s.io/client-go/kubernetes/fake"
	"github.com/google/go-cmp/cmp"
	"k8s.io/apimachinery/pkg/api/equality"
	//"reflect"
	"testing"
	"strings"
	//pminio "piranhas/platform/pirctl/pkg/minio"
	//"github.com/spf13/cobra"
	//"k8s.io/apimachinery/pkg/api/errors"
	//"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// func getFakeDynamicClient() dynamic.Interface { // *FakeDynamicClient
// 	log.Printf("getFakeDynamicClient")
// 	scheme := runtime.NewScheme()
// 	dynamicClient := dyfake.NewSimpleDynamicClient(
// 		scheme,
// 		newUnstructured(
// 			"group/version",
// 			"TheKind",
// 			"ns-foo",
// 			"name-foo",
// 			// "piranhas.framework/v1", // apiVersion
// 			// "analyticdeployments",   // kind
// 			// "test-tenant",           // namespace
// 			// "test-ad",               // name
// 		),
// 	)
// 	// adRes := schema.GroupVersionResource{Group: "piranhas.framework", Version: "v1", Resource: "analyticdeployments"}
// 	// dynamicClient.Resource(adRes)
// 	return dynamicClient
// }

func newUnstructured(apiVersion, kind, namespace, name string) *unstructured.Unstructured {
	return &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": apiVersion,
			"kind":       kind,
			"metadata": map[string]interface{}{
				"namespace": namespace,
				"name":      name,
			},
		},
	}
}

// https://github.com/kubernetes/client-go/blob/master/dynamic/fake/simple_test.go
func TestGetExample(t *testing.T) {
	log.Printf("#### TestGetExample\n")
	group := "myGroup"
	version := "v1"
	apiVersion := fmt.Sprintf("%s/%s", group, version)
	namespace := "myNamespace"
	scheme := runtime.NewScheme()
	name := "myName"
	kind := "TheKind"
	resourceStr := strings.ToLower(kind) + "s"

	client := dyfake.NewSimpleDynamicClient(
		scheme,
		newUnstructured(apiVersion, kind, namespace, name),
	)
	get, err := client.Resource(schema.GroupVersionResource{Group: group, Version: version, Resource: resourceStr}).Namespace(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		t.Fatal(err)
	}

	expected := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": apiVersion,
			"kind":       kind,
			"metadata": map[string]interface{}{
				"name":      name,
				"namespace": namespace,
			},
		},
	}
	if !equality.Semantic.DeepEqual(get, expected) {
		t.Fatal(cmp.Diff(expected, get))
	}
}

// func XTestGetFakeDynamicClient(t *testing.T) {
// 	fmt.Printf("#### TestGetFakeDynamicClient\n")
// 	dynamicClient := getFakeDynamicClient()
// 	if dynamicClient == nil {
// 		log.Fatal("getFakeDynamicClient failed")
// 	}
// 	adRes := schema.GroupVersionResource{Group: "piranhas.framework", Version: "v1", Resource: "analyticdeployments"}
// 	AD_NAME := "ad-name"
// 	AD_TENANT := "ad-tenant"
// 	// AD_VERSION := "ALL"
// 	log.Printf("get Analytic deployment, adRes: %v, AD_TENANT: %v, AD_NAME: %v", adRes, AD_TENANT, AD_NAME)
// 	ad := &unstructured.Unstructured{}
// 	if currAd, err := dynamicClient.Resource(adRes).Namespace(AD_TENANT).Get(context.Background(), AD_NAME, metav1.GetOptions{}); errors.IsNotFound(err) {
// 		log.Fatalf("Analytic deployment '%s' does not exist.", AD_NAME)
// 	} else {
// 		ad = currAd
// 	}
// 	if ad == nil {
// 		log.Fatal("dynamicClient.Resource failed")
// 	}
// }
