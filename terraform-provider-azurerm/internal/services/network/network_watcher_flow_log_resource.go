package network

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonschema"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/location"
	"github.com/hashicorp/terraform-provider-azurerm/helpers/azure"
	"github.com/hashicorp/terraform-provider-azurerm/helpers/tf"
	"github.com/hashicorp/terraform-provider-azurerm/internal/clients"
	"github.com/hashicorp/terraform-provider-azurerm/internal/locks"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/network/migration"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/network/parse"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/network/validate"
	storageValidate "github.com/hashicorp/terraform-provider-azurerm/internal/services/storage/validate"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tags"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/validation"
	"github.com/hashicorp/terraform-provider-azurerm/internal/timeouts"
	"github.com/hashicorp/terraform-provider-azurerm/utils"
	"github.com/tombuildsstuff/kermit/sdk/network/2022-07-01/network"
)

func resourceNetworkWatcherFlowLog() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		Create: resourceNetworkWatcherFlowLogCreateUpdate,
		Read:   resourceNetworkWatcherFlowLogRead,
		Update: resourceNetworkWatcherFlowLogCreateUpdate,
		Delete: resourceNetworkWatcherFlowLogDelete,

		SchemaVersion: 1,
		StateUpgraders: pluginsdk.StateUpgrades(map[int]pluginsdk.StateUpgrade{
			0: migration.NetworkWatcherFlowLogV0ToV1{},
		}),

		Importer: pluginsdk.ImporterValidatingResourceId(func(id string) error {
			_, err := parse.FlowLogID(id)
			return err
		}),

		Timeouts: &pluginsdk.ResourceTimeout{
			Create: pluginsdk.DefaultTimeout(30 * time.Minute),
			Read:   pluginsdk.DefaultTimeout(5 * time.Minute),
			Update: pluginsdk.DefaultTimeout(30 * time.Minute),
			Delete: pluginsdk.DefaultTimeout(30 * time.Minute),
		},

		Schema: map[string]*pluginsdk.Schema{
			"network_watcher_name": {
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.NoZeroValues,
			},

			"resource_group_name": commonschema.ResourceGroupName(),

			//lintignore: S013
			"name": {
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.NetworkWatcherFlowLogName,
			},

			"network_security_group_id": {
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.NetworkSecurityGroupID,
			},

			"storage_account_id": {
				Type:         pluginsdk.TypeString,
				Required:     true,
				ValidateFunc: storageValidate.StorageAccountID,
			},

			"enabled": {
				Type:     pluginsdk.TypeBool,
				Required: true,
			},

			"retention_policy": {
				Type:     pluginsdk.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"enabled": {
							Type:             pluginsdk.TypeBool,
							Required:         true,
							DiffSuppressFunc: azureRMSuppressFlowLogRetentionPolicyEnabledDiff,
						},

						"days": {
							Type:             pluginsdk.TypeInt,
							Required:         true,
							DiffSuppressFunc: azureRMSuppressFlowLogRetentionPolicyDaysDiff,
						},
					},
				},
			},

			"traffic_analytics": {
				Type:     pluginsdk.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"enabled": {
							Type:     pluginsdk.TypeBool,
							Required: true,
						},

						"workspace_id": {
							Type:         pluginsdk.TypeString,
							Required:     true,
							ValidateFunc: validation.IsUUID,
						},

						"workspace_region": {
							Type:             pluginsdk.TypeString,
							Required:         true,
							StateFunc:        location.StateFunc,
							DiffSuppressFunc: location.DiffSuppressFunc,
						},

						"workspace_resource_id": {
							Type:         pluginsdk.TypeString,
							Required:     true,
							ValidateFunc: azure.ValidateResourceIDOrEmpty, // nolint: staticcheck
						},

						"interval_in_minutes": {
							Type:         pluginsdk.TypeInt,
							Optional:     true,
							ValidateFunc: validation.IntInSlice([]int{10, 60}),
							Default:      60,
						},
					},
				},
			},

			"version": {
				Type:         pluginsdk.TypeInt,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.IntBetween(1, 2),
			},

			"location": {
				Type:             pluginsdk.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				ValidateFunc:     location.EnhancedValidate,
				StateFunc:        location.StateFunc,
				DiffSuppressFunc: location.DiffSuppressFunc,
			},

			"tags": tags.Schema(),
		},
	}
}

func azureRMSuppressFlowLogRetentionPolicyEnabledDiff(_, old, _ string, d *pluginsdk.ResourceData) bool {
	// Ignore if flow log is disabled as the returned flow log configuration
	// returns default value `false` which may differ from config
	return old != "" && !d.Get("enabled").(bool)
}

func azureRMSuppressFlowLogRetentionPolicyDaysDiff(_, old, _ string, d *pluginsdk.ResourceData) bool {
	// Ignore if flow log is disabled as the returned flow log configuration
	// returns default value `0` which may differ from config
	return old != "" && !d.Get("enabled").(bool)
}

func resourceNetworkWatcherFlowLogCreateUpdate(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Network.FlowLogsClient
	subscriptionId := meta.(*clients.Client).Account.SubscriptionId
	ctx, cancel := timeouts.ForCreateUpdate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	resourceGroupName := d.Get("resource_group_name").(string)
	networkWatcherName := d.Get("network_watcher_name").(string)
	name := d.Get("name").(string)
	id := parse.NewFlowLogID(subscriptionId, resourceGroupName, networkWatcherName, name)

	networkSecurityGroupID := d.Get("network_security_group_id").(string)
	nsgId, _ := parse.NetworkSecurityGroupID(networkSecurityGroupID)

	if d.IsNewResource() {
		// For newly created resources, the "name" is required, it is set as Optional and Computed is merely for the existing ones for the sake of backward compatibility.
		if name == "" {
			return fmt.Errorf("`name` is required for Network Watcher Flow Log")
		}

		existing, err := client.Get(ctx, id.ResourceGroup, id.NetworkWatcherName, id.Name)
		if err != nil {
			if !utils.ResponseWasNotFound(existing.Response) {
				return fmt.Errorf("failed checking for presence of existing %s: %+v", id, err)
			}
		}

		if !utils.ResponseWasNotFound(existing.Response) {
			return tf.ImportAsExistsError("azurerm_network_watcher_flow_log", id.ID())
		}
	}

	locks.ByID(nsgId.ID())
	defer locks.UnlockByID(nsgId.ID())

	loc := d.Get("location").(string)
	if loc == "" {
		// Get the containing network watcher in order to reuse its location if the "location" is not specified.
		watcherClient := meta.(*clients.Client).Network.WatcherClient
		resp, err := watcherClient.Get(ctx, id.ResourceGroup, id.NetworkWatcherName)
		if err != nil {
			return fmt.Errorf("retrieving %s: %v", parse.NewNetworkWatcherID(id.SubscriptionId, id.ResourceGroup, id.NetworkWatcherName).ID(), err)
		}
		if resp.Location != nil {
			loc = *resp.Location
		}
	}

	parameters := network.FlowLog{
		Location: utils.String(location.Normalize(loc)),
		FlowLogPropertiesFormat: &network.FlowLogPropertiesFormat{
			TargetResourceID: utils.String(nsgId.ID()),
			StorageID:        utils.String(d.Get("storage_account_id").(string)),
			Enabled:          utils.Bool(d.Get("enabled").(bool)),
			RetentionPolicy:  expandAzureRmNetworkWatcherFlowLogRetentionPolicy(d.Get("retention_policy").([]interface{})),
		},
		Tags: tags.Expand(d.Get("tags").(map[string]interface{})),
	}

	if _, ok := d.GetOk("traffic_analytics"); ok {
		parameters.FlowAnalyticsConfiguration = expandAzureRmNetworkWatcherFlowLogTrafficAnalytics(d)
	}

	if version, ok := d.GetOk("version"); ok {
		format := &network.FlowLogFormatParameters{
			Version: utils.Int32(int32(version.(int))),
		}

		parameters.Format = format
	}

	future, err := client.CreateOrUpdate(ctx, id.ResourceGroup, id.NetworkWatcherName, id.Name, parameters)
	if err != nil {
		return fmt.Errorf("creating %q: %+v", id, err)
	}

	if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
		return fmt.Errorf("waiting for completion of creating %q: %+v", id, err)
	}

	d.SetId(id.ID())

	return resourceNetworkWatcherFlowLogRead(d, meta)
}

func resourceNetworkWatcherFlowLogRead(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Network.FlowLogsClient
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.FlowLogID(d.Id())
	if err != nil {
		return err
	}

	// Get current flow log status
	resp, err := client.Get(ctx, id.ResourceGroup, id.NetworkWatcherName, id.Name)
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			log.Printf("[DEBUG] %s was not found - removing from state!", id)
			d.SetId("")
			return nil
		}

		return fmt.Errorf("retrieving %q: %+v", id, err)
	}

	d.Set("name", id.Name)
	d.Set("network_watcher_name", id.NetworkWatcherName)
	d.Set("resource_group_name", id.ResourceGroup)
	d.Set("location", location.NormalizeNilable(resp.Location))

	if prop := resp.FlowLogPropertiesFormat; prop != nil {
		if err := d.Set("traffic_analytics", flattenAzureRmNetworkWatcherFlowLogTrafficAnalytics(prop.FlowAnalyticsConfiguration)); err != nil {
			return fmt.Errorf("setting `traffic_analytics`: %+v", err)
		}

		d.Set("enabled", prop.Enabled)

		version := 0
		if format := prop.Format; format != nil {
			version = int(*format.Version)
		}
		d.Set("version", version)

		// Azure API returns "" when flow log is disabled
		// Don't overwrite to prevent storage account ID diff when that is the case
		if prop.StorageID != nil && *prop.StorageID != "" {
			d.Set("storage_account_id", prop.StorageID)
		}

		networkSecurityGroupId := ""
		if nsgIdRaw := prop.TargetResourceID; nsgIdRaw != nil {
			nsgId, err := parse.NetworkSecurityGroupIDInsensitively(*nsgIdRaw)
			if err != nil {
				return err
			}
			networkSecurityGroupId = nsgId.ID()
		}
		d.Set("network_security_group_id", networkSecurityGroupId)

		if err := d.Set("retention_policy", flattenAzureRmNetworkWatcherFlowLogRetentionPolicy(prop.RetentionPolicy)); err != nil {
			return fmt.Errorf("setting `retention_policy`: %+v", err)
		}
	}

	return tags.FlattenAndSet(d, resp.Tags)
}

func resourceNetworkWatcherFlowLogDelete(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Network.FlowLogsClient
	ctx, cancel := timeouts.ForDelete(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.FlowLogID(d.Id())
	if err != nil {
		return err
	}

	resp, err := client.Get(ctx, id.ResourceGroup, id.NetworkWatcherName, id.Name)
	if err != nil {
		return fmt.Errorf("retrieving %s: %+v", id, err)
	}
	if resp.FlowLogPropertiesFormat == nil || resp.FlowLogPropertiesFormat.TargetResourceID == nil {
		return fmt.Errorf("retreiving %s: `properties` or `properties.TargetResourceID` was nil", id)
	}

	networkSecurityGroupId, err := parse.NetworkSecurityGroupIDInsensitively(*resp.FlowLogPropertiesFormat.TargetResourceID)
	if err != nil {
		return fmt.Errorf("parsing %q as a Network Security Group ID: %+v", *resp.FlowLogPropertiesFormat.TargetResourceID, err)
	}

	locks.ByID(networkSecurityGroupId.ID())
	defer locks.UnlockByID(networkSecurityGroupId.ID())

	future, err := client.Delete(ctx, id.ResourceGroup, id.NetworkWatcherName, id.Name)
	if err != nil {
		return fmt.Errorf("deleting %s: %v", id, err)
	}

	if err := future.WaitForCompletionRef(ctx, client.Client); err != nil {
		return fmt.Errorf("waiting for deletion of %s: %v", id, err)
	}

	return nil
}

func expandAzureRmNetworkWatcherFlowLogRetentionPolicy(input []interface{}) *network.RetentionPolicyParameters {
	if len(input) < 1 || input[0] == nil {
		return nil
	}

	v := input[0].(map[string]interface{})
	enabled := v["enabled"].(bool)
	days := v["days"].(int)

	return &network.RetentionPolicyParameters{
		Enabled: utils.Bool(enabled),
		Days:    utils.Int32(int32(days)),
	}
}

func flattenAzureRmNetworkWatcherFlowLogRetentionPolicy(input *network.RetentionPolicyParameters) []interface{} {
	output := make([]interface{}, 0)

	if input != nil {
		enabled := false
		if input.Enabled != nil {
			enabled = *input.Enabled
		}
		days := 0
		if input.Days != nil {
			days = int(*input.Days)
		}
		output = append(output, map[string]interface{}{
			"days":    days,
			"enabled": enabled,
		})
	}

	return output
}

func flattenAzureRmNetworkWatcherFlowLogTrafficAnalytics(input *network.TrafficAnalyticsProperties) []interface{} {
	output := make([]interface{}, 0)
	if input != nil {
		if cfg := input.NetworkWatcherFlowAnalyticsConfiguration; cfg != nil {
			enabled := false
			if cfg.Enabled != nil {
				enabled = *cfg.Enabled
			}
			workspaceId := ""
			if cfg.WorkspaceID != nil {
				workspaceId = *cfg.WorkspaceID
			}
			workspaceRegion := ""
			if cfg.WorkspaceRegion != nil {
				workspaceRegion = *cfg.WorkspaceRegion
			}
			workspaceResourceId := ""
			if cfg.WorkspaceResourceID != nil {
				workspaceResourceId = *cfg.WorkspaceResourceID
			}

			intervalInMinutes := 0
			if cfg.TrafficAnalyticsInterval != nil {
				intervalInMinutes = int(*cfg.TrafficAnalyticsInterval)
			}
			output = append(output, map[string]interface{}{
				"enabled":               enabled,
				"interval_in_minutes":   intervalInMinutes,
				"workspace_id":          workspaceId,
				"workspace_region":      workspaceRegion,
				"workspace_resource_id": workspaceResourceId,
			})
		}
	}

	return output
}

func expandAzureRmNetworkWatcherFlowLogTrafficAnalytics(d *pluginsdk.ResourceData) *network.TrafficAnalyticsProperties {
	vs := d.Get("traffic_analytics").([]interface{})

	v := vs[0].(map[string]interface{})
	enabled := v["enabled"].(bool)
	workspaceID := v["workspace_id"].(string)
	workspaceRegion := v["workspace_region"].(string)
	workspaceResourceID := v["workspace_resource_id"].(string)
	interval := v["interval_in_minutes"].(int)

	return &network.TrafficAnalyticsProperties{
		NetworkWatcherFlowAnalyticsConfiguration: &network.TrafficAnalyticsConfigurationProperties{
			Enabled:                  utils.Bool(enabled),
			WorkspaceID:              utils.String(workspaceID),
			WorkspaceRegion:          utils.String(workspaceRegion),
			WorkspaceResourceID:      utils.String(workspaceResourceID),
			TrafficAnalyticsInterval: utils.Int32(int32(interval)),
		},
	}
}