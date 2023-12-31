package migration

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-provider-azurerm/utils"
)

func TestFeatureResourceV0ToV1(t *testing.T) {
	testData := []struct {
		name     string
		input    map[string]interface{}
		expected *string
	}{
		{
			name: "old id (normal)",
			input: map[string]interface{}{
				"id": "/subscriptions/12345678-1234-5678-1234-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.AppConfiguration/configurationStores/appConf1/AppConfigurationFeature/keyName/Label/labelName",
			},
			expected: utils.String("https://appConf1.azconfig.io/kv/.appconfig.featureflag%2FkeyName?label=labelName"),
		},
		{
			name: "old id (complicated)",
			input: map[string]interface{}{
				"id": "/subscriptions/12345678-1234-5678-1234-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.AppConfiguration/configurationStores/appConf1/AppConfigurationFeature/key:name/test/Label/test:label/name",
			},
			expected: utils.String("https://appConf1.azconfig.io/kv/.appconfig.featureflag%2Fkey:name%2Ftest?label=test%3Alabel%2Fname"),
		},
		{
			name: "old id (no label)",
			input: map[string]interface{}{
				"id": "/subscriptions/12345678-1234-5678-1234-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.AppConfiguration/configurationStores/appConf1/AppConfigurationFeature/keyName/Label/%00",
			},
			expected: utils.String("https://appConf1.azconfig.io/kv/.appconfig.featureflag%2FkeyName?label="),
		},
		{
			name: "old id (\000 label)",
			input: map[string]interface{}{
				"id": "/subscriptions/12345678-1234-5678-1234-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.AppConfiguration/configurationStores/appConf1/AppConfigurationFeature/keyName/Label/\000",
			},
			expected: utils.String("https://appConf1.azconfig.io/kv/.appconfig.featureflag%2FkeyName?label="),
		},
		{
			name: "old id (empty label)",
			input: map[string]interface{}{
				"id": "/subscriptions/12345678-1234-5678-1234-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.AppConfiguration/configurationStores/appConf1/AppConfigurationFeature/keyName/Label/",
			},
			expected: utils.String("https://appConf1.azconfig.io/kv/.appconfig.featureflag%2FkeyName?label="),
		},
	}
	for _, test := range testData {
		t.Logf("Testing %q...", test.name)
		result, err := FeatureResourceV0ToV1{}.UpgradeFunc()(context.TODO(), test.input, nil)
		if err != nil && test.expected == nil {
			continue
		} else {
			if err == nil && test.expected == nil {
				t.Fatalf("Expected an error but didn't get one")
			} else if err != nil && test.expected != nil {
				t.Fatalf("Expected no error but got: %+v", err)
			}
		}

		actualId := result["id"].(string)
		if *test.expected != actualId {
			t.Fatalf("expected %q but got %q", *test.expected, actualId)
		}
	}
}
