package instana

import (
	"github.com/gessnerfl/terraform-provider-instana/instana/restapi"
	"github.com/gessnerfl/terraform-provider-instana/utils"
	"github.com/hashicorp/terraform/helper/schema"
)

const (
	//AlertingChannelVictorOpsFieldAPIKey const for the apiKey field of the VictorOps alerting channel
	AlertingChannelVictorOpsFieldAPIKey = "api_key"
	//AlertingChannelVictorOpsFieldRoutingKey const for the routingKey field of the VictorOps alerting channel
	AlertingChannelVictorOpsFieldRoutingKey = "routing_key"
	//ResourceInstanaAlertingChannelVictorOps the name of the terraform-provider-instana resource to manage alerting channels of type VictorOps
	ResourceInstanaAlertingChannelVictorOps = "instana_alerting_channel_victor_ops"
)

//NewAlertingChannelVictorOpsResourceHandle creates the resource handle for Alerting Channels of type Email
func NewAlertingChannelVictorOpsResourceHandle() *ResourceHandle {
	return &ResourceHandle{
		ResourceName: ResourceInstanaAlertingChannelVictorOps,
		Schema: map[string]*schema.Schema{
			AlertingChannelFieldName:     alertingChannelNameSchemaField,
			AlertingChannelFieldFullName: alertingChannelFullNameSchemaField,
			AlertingChannelVictorOpsFieldAPIKey: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The API Key of the VictorOps alerting channel",
			},
			AlertingChannelVictorOpsFieldRoutingKey: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The Routing Key of the VictorOps alerting channel",
			},
		},
		RestResourceFactory:  func(api restapi.InstanaAPI) restapi.RestResource { return api.AlertingChannels() },
		UpdateState:          updateStateForAlertingChannelVictorOps,
		MapStateToDataObject: mapStateToDataObjectForAlertingChannelVictorOps,
	}
}

func updateStateForAlertingChannelVictorOps(d *schema.ResourceData, obj restapi.InstanaDataObject) error {
	alertingChannel := obj.(restapi.AlertingChannel)
	d.Set(AlertingChannelFieldFullName, alertingChannel.Name)
	d.Set(AlertingChannelVictorOpsFieldAPIKey, alertingChannel.APIKey)
	d.Set(AlertingChannelVictorOpsFieldRoutingKey, alertingChannel.RoutingKey)
	d.SetId(alertingChannel.ID)
	return nil
}

func mapStateToDataObjectForAlertingChannelVictorOps(d *schema.ResourceData, formatter utils.ResourceNameFormatter) (restapi.InstanaDataObject, error) {
	name := computeFullAlertingChannelNameString(d, formatter)
	apiKey := d.Get(AlertingChannelVictorOpsFieldAPIKey).(string)
	routingKey := d.Get(AlertingChannelVictorOpsFieldRoutingKey).(string)
	return restapi.AlertingChannel{
		ID:         d.Id(),
		Name:       name,
		Kind:       restapi.VictorOpsChannelType,
		APIKey:     &apiKey,
		RoutingKey: &routingKey,
	}, nil
}
