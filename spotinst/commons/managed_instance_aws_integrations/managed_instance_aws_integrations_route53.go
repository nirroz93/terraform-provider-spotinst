package managed_instance_aws_integrations

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/managedinstance/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func SetupRoute53(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[IntegrationRoute53] = commons.NewGenericField(
		commons.ManagedInstanceAWSIntegrations,
		IntegrationRoute53,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(Domains): {
						Type:     schema.TypeSet,
						Required: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(HostedZoneId): {
									Type:     schema.TypeString,
									Required: true,
								},

								string(SpotinstAcctID): {
									Type:     schema.TypeString,
									Optional: true,
								},

								string(RecordSetType): {
									Type:     schema.TypeString,
									Optional: true,
								},

								string(RecordSets): {
									Type:     schema.TypeSet,
									Required: true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											string(UsePublicIP): {
												Type:     schema.TypeBool,
												Optional: true,
											},

											string(UsePublicDNS): {
												Type:     schema.TypeBool,
												Optional: true,
											},

											string(Route53Name): {
												Type:     schema.TypeString,
												Required: true,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			if v, ok := resourceData.GetOk(string(IntegrationRoute53)); ok {
				if integration, err := expandAWSManagedInstanceRoute53Integration(v); err != nil {
					return err
				} else {
					managedInstance.Integration.SetRoute53(integration)
				}
			}
			return nil
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			var value *aws.Route53Integration = nil

			if v, ok := resourceData.GetOk(string(IntegrationRoute53)); ok {
				if integration, err := expandAWSManagedInstanceRoute53Integration(v); err != nil {
					return err
				} else {
					value = integration
				}
			}
			managedInstance.Integration.SetRoute53(value)
			return nil
		},

		nil,
	)

}

func expandAWSManagedInstanceRoute53Integration(data interface{}) (*aws.Route53Integration, error) {
	integration := &aws.Route53Integration{}
	list := data.([]interface{})

	if list != nil && list[0] != nil {
		m := list[0].(map[string]interface{})

		if v, ok := m[string(Domains)]; ok {
			domains, err := expandAWSManagedInstanceRoute53IntegrationDomains(v)

			if err != nil {
				return nil, err
			}
			integration.SetDomains(domains)
		}
	}
	return integration, nil
}

func expandAWSManagedInstanceRoute53IntegrationDomains(data interface{}) ([]*aws.Domain, error) {
	list := data.(*schema.Set).List()
	domains := make([]*aws.Domain, 0, len(list))

	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}

		domain := new(aws.Domain)

		if v, ok := attr[string(HostedZoneId)].(string); ok && v != "" {
			domain.SetHostedZoneId(spotinst.String(v))
		}

		if v, ok := attr[string(SpotinstAcctID)].(string); ok && v != "" {
			domain.SetSpotinstAccountId(spotinst.String(v))
		}

		if v, ok := attr[string(RecordSetType)].(string); ok && v != "" {
			domain.SetRecordSetType(spotinst.String(v))
		}

		if r, ok := attr[string(RecordSets)]; ok {
			if recordSets, err := expandAWSManagedInstanceRoute53IntegrationDomainsRecordSets(r); err != nil {
				return nil, err
			} else {
				domain.SetRecordSets(recordSets)
			}
		}

		domains = append(domains, domain)
	}

	return domains, nil
}

func expandAWSManagedInstanceRoute53IntegrationDomainsRecordSets(data interface{}) ([]*aws.RecordSet, error) {
	list := data.(*schema.Set).List()
	recordSets := make([]*aws.RecordSet, 0, len(list))

	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}

		recordSet := new(aws.RecordSet)

		if v, ok := attr[string(Route53Name)].(string); ok {
			recordSet.SetName(spotinst.String(v))
		}

		if v, ok := attr[string(UsePublicIP)].(bool); ok && v {
			recordSet.SetUsePublicIP(spotinst.Bool(v))
		}

		if v, ok := attr[string(UsePublicDNS)].(bool); ok && v {
			recordSet.SetUsePublicDNS(spotinst.Bool(v))
		}

		recordSets = append(recordSets, recordSet)
	}

	return recordSets, nil
}
