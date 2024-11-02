// Copyright © 2024 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

package provider

// Section below is generated&owned by "gen/generator.go". //template:begin provider
import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
)

// MerakiProvider defines the provider implementation.
type MerakiProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// MerakiProviderModel describes the provider data model.
type MerakiProviderModel struct {
	ApiKey     types.String `tfsdk:"api_key"`
	BaseUrl    types.String `tfsdk:"base_url"`
	ReqTimeout types.String `tfsdk:"req_timeout"`
	Retries    types.Int64  `tfsdk:"retries"`
}

// MerakiProviderData describes the data maintained by the provider.
type MerakiProviderData struct {
	Client *meraki.Client
}

// Metadata returns the provider type name.
func (p *MerakiProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "meraki"
	resp.Version = p.version
}

func (p *MerakiProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"api_key": schema.StringAttribute{
				MarkdownDescription: "Meraki Dashboard API key. This can also be set as the MERAKI_API_KEY environment variable.",
				Optional:            true,
				Sensitive:           true,
			},
			"base_url": schema.StringAttribute{
				MarkdownDescription: "Base URL to be used. The default value is `https://api.meraki.com/api/v1`. This can also be set as the MERAKI_BASE_URL environment variable.",
				Optional:            true,
			},
			"req_timeout": schema.StringAttribute{
				MarkdownDescription: "Timeout for a single HTTPS request made to REST API before it is retried. This can also be set as the MERAKI_REQ_TIMEOUT environment variable. A string like `\"1s\"` means one second. Defaults to `\"5s\"`.",
				Optional:            true,
			},
			"retries": schema.Int64Attribute{
				MarkdownDescription: "Number of retries for REST API calls. This can also be set as the MERAKI_RETRIES environment variable. Defaults to `3`.",
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 9),
				},
			},
		},
	}
}

func (p *MerakiProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// Retrieve provider data from configuration
	var config MerakiProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// User must provide an API key to the provider
	var apiKey string
	if config.ApiKey.IsUnknown() {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as API key",
		)
		return
	}

	if config.ApiKey.IsNull() {
		apiKey = os.Getenv("MERAKI_API_KEY")
	} else {
		apiKey = config.ApiKey.ValueString()
	}

	if apiKey == "" {
		// Error vs warning - empty value must stop execution
		resp.Diagnostics.AddError(
			"Unable to find API key",
			"API key cannot be an empty string",
		)
		return
	}

	// User can provide a Base URL to the provider
	var baseUrl string
	if config.BaseUrl.IsUnknown() {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as base url",
		)
		return
	}

	if config.BaseUrl.IsNull() {
		baseUrl = os.Getenv("MERAKI_BASE_URL")
		if baseUrl == "" {
			baseUrl = "https://api.meraki.com/api/v1"
		}
	} else {
		baseUrl = config.BaseUrl.ValueString()
	}

	var reqTimeout time.Duration
	if config.ReqTimeout.IsUnknown() {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as req_timeout",
		)
		return
	}

	var reqTimeoutStr string
	if config.ReqTimeout.IsNull() {
		reqTimeoutStr = os.Getenv("MERAKI_REQ_TIMEOUT")
		if reqTimeoutStr == "" {
			reqTimeoutStr = "5s"
		}
	} else {
		reqTimeoutStr = config.ReqTimeout.ValueString()
	}
	reqTimeout, err := time.ParseDuration(reqTimeoutStr)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create client",
			fmt.Sprintf("Cannot parse the req_timeout string: %v", err),
		)
		return
	}

	var retries int64
	if config.Retries.IsUnknown() {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as retries",
		)
		return
	}

	if config.Retries.IsNull() {
		retriesStr := os.Getenv("MERAKI_RETRIES")
		if retriesStr == "" {
			retries = 3
		} else {
			retries, _ = strconv.ParseInt(retriesStr, 0, 64)
		}
	} else {
		retries = config.Retries.ValueInt64()
	}

	tflog.Debug(ctx, fmt.Sprint("Creating a new Meraki client",
		"  base_url=", baseUrl,
		"  req_timeout=", reqTimeout,
		"  retries=", retries,
	))

	// Create a new Meraki client and set it to the provider client
	c, err := meraki.NewClient(apiKey, meraki.BaseUrl(baseUrl), meraki.MaxRetries(int(retries)), meraki.RequestTimeout(reqTimeout))
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create client",
			"Unable to create meraki client:\n\n"+err.Error(),
		)
		return
	}

	data := MerakiProviderData{Client: &c}
	resp.DataSourceData = &data
	resp.ResourceData = &data
}

func (p *MerakiProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewApplianceCellularFirewallRulesResource,
		NewApplianceConnectivityMonitoringDestinationsResource,
		NewApplianceContentFilteringResource,
		NewApplianceFirewallSettingsResource,
		NewApplianceFirewalledServiceResource,
		NewApplianceInboundCellularFirewallRulesResource,
		NewApplianceInboundFirewallRulesResource,
		NewApplianceL3FirewallRulesResource,
		NewApplianceL7FirewallRulesResource,
		NewApplianceNetworkSecurityIntrusionResource,
		NewApplianceOneToManyNATRulesResource,
		NewApplianceOneToOneNATRulesResource,
		NewApplianceOrganizationSecurityIntrusionResource,
		NewAppliancePortResource,
		NewAppliancePortForwardingRulesResource,
		NewAppliancePrefixDelegatedStaticResource,
		NewApplianceRadioSettingsResource,
		NewApplianceRFProfileResource,
		NewApplianceSDWANInternetPoliciesResource,
		NewApplianceSecurityMalwareResource,
		NewApplianceSettingsResource,
		NewApplianceSingleLANResource,
		NewApplianceSiteToSiteVPNResource,
		NewApplianceSSIDResource,
		NewApplianceStaticRouteResource,
		NewApplianceThirdPartyVPNPeersResource,
		NewApplianceTrafficShapingResource,
		NewApplianceTrafficShapingCustomPerformanceClassResource,
		NewApplianceTrafficShapingRulesResource,
		NewApplianceTrafficShapingUplinkBandwidthResource,
		NewApplianceTrafficShapingUplinkSelectionResource,
		NewApplianceTrafficShapingVPNExclusionsResource,
		NewApplianceUplinksSettingsResource,
		NewApplianceVLANResource,
		NewApplianceVLANsSettingsResource,
		NewApplianceVPNBGPResource,
		NewApplianceVPNFirewallRulesResource,
		NewApplianceWarmSpareResource,
		NewDeviceResource,
		NewDeviceCellularSIMsResource,
		NewDeviceManagementInterfaceResource,
		NewNetworkResource,
		NewNetworkClientPolicyResource,
		NewNetworkDeviceClaimResource,
		NewNetworkFloorPlanResource,
		NewNetworkGroupPolicyResource,
		NewNetworkMerakiAuthUserResource,
		NewNetworkMQTTBrokerResource,
		NewNetworkNetflowResource,
		NewNetworkSettingsResource,
		NewNetworkSNMPResource,
		NewNetworkSyslogServersResource,
		NewNetworkTrafficAnalysisResource,
		NewNetworkVLANProfileResource,
		NewNetworkWebhookHTTPServerResource,
		NewNetworkWebhookPayloadTemplateResource,
		NewOrganizationResource,
		NewOrganizationAdaptivePolicyResource,
		NewOrganizationAdaptivePolicyACLResource,
		NewOrganizationAdaptivePolicyGroupResource,
		NewOrganizationAdaptivePolicySettingsResource,
		NewOrganizationAdminResource,
		NewOrganizationAlertsProfileResource,
		NewOrganizationConfigTemplateResource,
		NewOrganizationEarlyAccessFeaturesOptInResource,
		NewOrganizationInventoryClaimResource,
		NewOrganizationLoginSecurityResource,
		NewOrganizationPolicyObjectResource,
		NewOrganizationPolicyObjectGroupResource,
		NewOrganizationSAMLResource,
		NewOrganizationSAMLIdPResource,
		NewOrganizationSAMLRoleResource,
		NewOrganizationSNMPResource,
		NewSwitchAccessControlListsResource,
		NewSwitchAccessPolicyResource,
		NewSwitchAlternateManagementInterfaceResource,
		NewSwitchDHCPServerPolicyResource,
		NewSwitchDHCPServerPolicyARPInspectionTrustedServerResource,
		NewSwitchDSCPToCoSMappingsResource,
		NewSwitchLinkAggregationResource,
		NewSwitchMTUResource,
		NewSwitchOrganizationPortsProfileResource,
		NewSwitchPortResource,
		NewSwitchPortScheduleResource,
		NewSwitchQoSRuleResource,
		NewSwitchQoSRuleOrderResource,
		NewSwitchRoutingInterfaceResource,
		NewSwitchRoutingInterfaceDHCPResource,
		NewSwitchRoutingMulticastResource,
		NewSwitchRoutingMulticastRendezvousPointResource,
		NewSwitchRoutingOSPFResource,
		NewSwitchRoutingStaticRouteResource,
		NewSwitchSettingsResource,
		NewSwitchStackResource,
		NewSwitchStackRoutingInterfaceResource,
		NewSwitchStackRoutingInterfaceDHCPResource,
		NewSwitchStackRoutingStaticRouteResource,
		NewSwitchStormControlResource,
		NewSwitchSTPResource,
		NewSwitchWarmSpareResource,
		NewWirelessAirMarshalRuleResource,
		NewWirelessAirMarshalSettingsResource,
		NewWirelessAlternateManagementInterfaceResource,
		NewWirelessBillingResource,
		NewWirelessDeviceBluetoothSettingsResource,
		NewWirelessDeviceElectronicShelfLabelResource,
		NewWirelessEthernetPortProfileResource,
		NewWirelessEthernetPortProfileAssignResource,
		NewWirelessEthernetPortProfileDefaultResource,
		NewWirelessNetworkBluetoothSettingsResource,
		NewWirelessNetworkElectronicShelfLabelResource,
		NewWirelessRadioSettingsResource,
		NewWirelessRFProfileResource,
		NewWirelessSettingsResource,
		NewWirelessSSIDResource,
		NewWirelessSSIDBonjourForwardingResource,
		NewWirelessSSIDDeviceTypeGroupPoliciesResource,
		NewWirelessSSIDEAPOverrideResource,
		NewWirelessSSIDHotspot20Resource,
		NewWirelessSSIDIdentityPSKResource,
		NewWirelessSSIDL3FirewallRulesResource,
		NewWirelessSSIDL7FirewallRulesResource,
		NewWirelessSSIDSchedulesResource,
		NewWirelessSSIDSplashSettingsResource,
		NewWirelessSSIDTrafficShapingRulesResource,
		NewWirelessSSIDVPNResource,
	}
}

func (p *MerakiProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewApplianceCellularFirewallRulesDataSource,
		NewApplianceConnectivityMonitoringDestinationsDataSource,
		NewApplianceContentFilteringDataSource,
		NewApplianceFirewallSettingsDataSource,
		NewApplianceFirewalledServiceDataSource,
		NewApplianceInboundCellularFirewallRulesDataSource,
		NewApplianceInboundFirewallRulesDataSource,
		NewApplianceL3FirewallRulesDataSource,
		NewApplianceL7FirewallRulesDataSource,
		NewApplianceNetworkSecurityIntrusionDataSource,
		NewApplianceOneToManyNATRulesDataSource,
		NewApplianceOneToOneNATRulesDataSource,
		NewApplianceOrganizationSecurityIntrusionDataSource,
		NewAppliancePortDataSource,
		NewAppliancePortForwardingRulesDataSource,
		NewAppliancePrefixDelegatedStaticDataSource,
		NewAppliancePrefixDelegatedStaticsDataSource,
		NewApplianceRadioSettingsDataSource,
		NewApplianceRFProfileDataSource,
		NewApplianceRFProfilesDataSource,
		NewApplianceSecurityMalwareDataSource,
		NewApplianceSettingsDataSource,
		NewApplianceSingleLANDataSource,
		NewApplianceSiteToSiteVPNDataSource,
		NewApplianceSSIDDataSource,
		NewApplianceStaticRouteDataSource,
		NewApplianceStaticRoutesDataSource,
		NewApplianceThirdPartyVPNPeersDataSource,
		NewApplianceTrafficShapingDataSource,
		NewApplianceTrafficShapingCustomPerformanceClassDataSource,
		NewApplianceTrafficShapingCustomPerformanceClassesDataSource,
		NewApplianceTrafficShapingRulesDataSource,
		NewApplianceTrafficShapingUplinkBandwidthDataSource,
		NewApplianceTrafficShapingUplinkSelectionDataSource,
		NewApplianceTrafficShapingVPNExclusionsDataSource,
		NewApplianceUplinksSettingsDataSource,
		NewApplianceVLANDataSource,
		NewApplianceVLANsDataSource,
		NewApplianceVLANsSettingsDataSource,
		NewApplianceVPNBGPDataSource,
		NewApplianceVPNFirewallRulesDataSource,
		NewApplianceWarmSpareDataSource,
		NewDeviceDataSource,
		NewDeviceCellularSIMsDataSource,
		NewDeviceManagementInterfaceDataSource,
		NewNetworkDataSource,
		NewNetworksDataSource,
		NewNetworkClientPolicyDataSource,
		NewNetworkFloorPlanDataSource,
		NewNetworkFloorPlansDataSource,
		NewNetworkGroupPolicyDataSource,
		NewNetworkGroupPoliciesDataSource,
		NewNetworkMerakiAuthUserDataSource,
		NewNetworkMerakiAuthUsersDataSource,
		NewNetworkMQTTBrokerDataSource,
		NewNetworkMQTTBrokersDataSource,
		NewNetworkNetflowDataSource,
		NewNetworkSettingsDataSource,
		NewNetworkSNMPDataSource,
		NewNetworkSyslogServersDataSource,
		NewNetworkTrafficAnalysisDataSource,
		NewNetworkVLANProfileDataSource,
		NewNetworkVLANProfilesDataSource,
		NewNetworkWebhookHTTPServerDataSource,
		NewNetworkWebhookHTTPServersDataSource,
		NewNetworkWebhookPayloadTemplateDataSource,
		NewNetworkWebhookPayloadTemplatesDataSource,
		NewOrganizationDataSource,
		NewOrganizationsDataSource,
		NewOrganizationAdaptivePolicyDataSource,
		NewOrganizationAdaptivePoliciesDataSource,
		NewOrganizationAdaptivePolicyACLDataSource,
		NewOrganizationAdaptivePolicyACLsDataSource,
		NewOrganizationAdaptivePolicyGroupDataSource,
		NewOrganizationAdaptivePolicyGroupsDataSource,
		NewOrganizationAdaptivePolicySettingsDataSource,
		NewOrganizationAdminDataSource,
		NewOrganizationAdminsDataSource,
		NewOrganizationAlertsProfileDataSource,
		NewOrganizationAlertsProfilesDataSource,
		NewOrganizationConfigTemplateDataSource,
		NewOrganizationConfigTemplatesDataSource,
		NewOrganizationEarlyAccessFeaturesOptInDataSource,
		NewOrganizationEarlyAccessFeaturesOptInsDataSource,
		NewOrganizationLoginSecurityDataSource,
		NewOrganizationPolicyObjectDataSource,
		NewOrganizationPolicyObjectsDataSource,
		NewOrganizationPolicyObjectGroupDataSource,
		NewOrganizationPolicyObjectGroupsDataSource,
		NewOrganizationSAMLDataSource,
		NewOrganizationSAMLIdPDataSource,
		NewOrganizationSAMLIdPsDataSource,
		NewOrganizationSAMLRoleDataSource,
		NewOrganizationSAMLRolesDataSource,
		NewOrganizationSNMPDataSource,
		NewSwitchAccessControlListsDataSource,
		NewSwitchAccessPolicyDataSource,
		NewSwitchAccessPoliciesDataSource,
		NewSwitchAlternateManagementInterfaceDataSource,
		NewSwitchDHCPServerPolicyDataSource,
		NewSwitchDHCPServerPolicyARPInspectionTrustedServerDataSource,
		NewSwitchDHCPServerPolicyARPInspectionTrustedServersDataSource,
		NewSwitchDSCPToCoSMappingsDataSource,
		NewSwitchLinkAggregationDataSource,
		NewSwitchLinkAggregationsDataSource,
		NewSwitchMTUDataSource,
		NewSwitchOrganizationPortsProfileDataSource,
		NewSwitchOrganizationPortsProfilesDataSource,
		NewSwitchPortDataSource,
		NewSwitchPortScheduleDataSource,
		NewSwitchPortSchedulesDataSource,
		NewSwitchQoSRuleDataSource,
		NewSwitchQoSRulesDataSource,
		NewSwitchQoSRuleOrderDataSource,
		NewSwitchRoutingInterfaceDataSource,
		NewSwitchRoutingInterfacesDataSource,
		NewSwitchRoutingInterfaceDHCPDataSource,
		NewSwitchRoutingMulticastDataSource,
		NewSwitchRoutingMulticastRendezvousPointDataSource,
		NewSwitchRoutingMulticastRendezvousPointsDataSource,
		NewSwitchRoutingOSPFDataSource,
		NewSwitchRoutingStaticRouteDataSource,
		NewSwitchRoutingStaticRoutesDataSource,
		NewSwitchSettingsDataSource,
		NewSwitchStackDataSource,
		NewSwitchStacksDataSource,
		NewSwitchStackRoutingInterfaceDataSource,
		NewSwitchStackRoutingInterfacesDataSource,
		NewSwitchStackRoutingInterfaceDHCPDataSource,
		NewSwitchStackRoutingStaticRouteDataSource,
		NewSwitchStackRoutingStaticRoutesDataSource,
		NewSwitchStormControlDataSource,
		NewSwitchSTPDataSource,
		NewSwitchWarmSpareDataSource,
		NewWirelessAirMarshalRuleDataSource,
		NewWirelessAirMarshalSettingsDataSource,
		NewWirelessAlternateManagementInterfaceDataSource,
		NewWirelessBillingDataSource,
		NewWirelessDeviceBluetoothSettingsDataSource,
		NewWirelessDeviceElectronicShelfLabelDataSource,
		NewWirelessEthernetPortProfileDataSource,
		NewWirelessEthernetPortProfilesDataSource,
		NewWirelessEthernetPortProfileDefaultDataSource,
		NewWirelessNetworkBluetoothSettingsDataSource,
		NewWirelessNetworkElectronicShelfLabelDataSource,
		NewWirelessRadioSettingsDataSource,
		NewWirelessRFProfileDataSource,
		NewWirelessRFProfilesDataSource,
		NewWirelessSettingsDataSource,
		NewWirelessSSIDDataSource,
		NewWirelessSSIDBonjourForwardingDataSource,
		NewWirelessSSIDDeviceTypeGroupPoliciesDataSource,
		NewWirelessSSIDEAPOverrideDataSource,
		NewWirelessSSIDHotspot20DataSource,
		NewWirelessSSIDIdentityPSKDataSource,
		NewWirelessSSIDIdentityPSKsDataSource,
		NewWirelessSSIDL3FirewallRulesDataSource,
		NewWirelessSSIDL7FirewallRulesDataSource,
		NewWirelessSSIDSchedulesDataSource,
		NewWirelessSSIDSplashSettingsDataSource,
		NewWirelessSSIDTrafficShapingRulesDataSource,
		NewWirelessSSIDVPNDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &MerakiProvider{
			version: version,
		}
	}
}

// End of section. //template:end provider
