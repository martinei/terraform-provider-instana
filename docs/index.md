# instana Provider

Terraform provider implementation of the Instana Web REST API. The provider can be used to configure different
assents in Instana. The provider is aligned with the REST API and links to the endpoint is provided for each 
resource. 

**NOTE:** Starting with version 0.6.0 Terraform version 0.12.x or later is required.

## Supported Resources:

* Application Settings
  * Application Configuration - `instana_application_config`
* Event Settings
  * Custom Event Specification
    * Entity Verification Rule - `instana_custom_event_spec_entity_verification_rule`
    * System Rule - `instana_custom_event_spec_system_rule`
    * Threshold Rule - `instana_custom_event_spec_threshold_rule`
  * Alerting Channels
    * Email - `instana_alerting_channel_email`
    * Google Chat - `instana_alerting_channel_google_chat`
    * Office 365 - `instana_alerting_channel_office_365`
    * OpsGenie - `instana_alerting_channel_ops_genie`
    * Pager Duty - `instana_alerting_channel_pager_duty`
    * Slack - `instana_alerting_channel_slack`
    * Splunk - `instana_alerting_channel_splunk`
    * VictorOps - `instana_alerting_channel_victor_ops`
    * Webhook - `instana_alerting_channel_webhook`
* Settings
  * User Roles - `instana_user_role`

## Example Usage

```hcl
provider "instana" {
  api_token = "secure-api-token"  
  endpoint = "<tenant>-<org>.instana.io"
  default_name_prefix = ""
  default_name_suffix = "(TF managed)"
}
```

## Argument Reference

* `api_token` - Required - The API token which is created in the Settings area of Instana for remote access through 
the REST API. You have to make sure that you assign the proper permissions for this token to configure the desired 
resources with this provider. E.g. when User Roles should be provisioned by terraform using this provider implementation 
then the permission 'Access role configuration' must be activated
* `endpoint` - Required - The endpoint of the instana backend. For SaaS the endpoint URL has the pattern 
`<tenant>-<organization>.instana.io`. For onPremise installation the endpoint URL depends on your local setup.
* `default_name_prefix` - Optional - string will be added in front the resource UI name or label by default
(not supported by all resources). For existing resources the string will only be added when the name/label is changed.
* default_name_suffix - `Optional` - Default value " (TF managed)" - string will be appended to the resource UI name or 
label by default (not supported by all resources). For existing resources the string will only be appended when the 
name/label is changed.

