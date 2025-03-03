package elastigroup_aws_launch_configuration

import (
	"encoding/base64"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[ImageId] = commons.NewGenericField(
		commons.ElastigroupAWSLaunchConfiguration,
		ImageId,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *string = nil
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.ImageID != nil {
				value = elastigroup.Compute.LaunchSpecification.ImageID
			}
			if err := resourceData.Set(string(ImageId), spotinst.StringValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ImageId), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(ImageId)).(string); ok && v != "" {
				elastigroup.Compute.LaunchSpecification.SetImageId(spotinst.String(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(ImageId)).(string); ok && v != "" {
				elastigroup.Compute.LaunchSpecification.SetImageId(spotinst.String(v))
			}
			return nil
		},
		nil,
	)

	fieldsMap[IamInstanceProfile] = commons.NewGenericField(
		commons.ElastigroupAWSLaunchConfiguration,
		IamInstanceProfile,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value = ""
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.IAMInstanceProfile != nil {

				lc := elastigroup.Compute.LaunchSpecification
				if lc.IAMInstanceProfile.Arn != nil {
					value = spotinst.StringValue(lc.IAMInstanceProfile.Arn)
				} else if lc.IAMInstanceProfile.Name != nil {
					value = spotinst.StringValue(lc.IAMInstanceProfile.Name)
				}
			}
			if err := resourceData.Set(string(IamInstanceProfile), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(IamInstanceProfile), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(IamInstanceProfile)).(string); ok && v != "" {
				iamInstanceProf := &aws.IAMInstanceProfile{}

				if InstanceProfileArnRegex.MatchString(v) {
					iamInstanceProf.SetArn(spotinst.String(v))
				} else {
					iamInstanceProf.SetName(spotinst.String(v))
				}
				elastigroup.Compute.LaunchSpecification.SetIAMInstanceProfile(iamInstanceProf)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(IamInstanceProfile)).(string); ok && v != "" {
				iamInstanceProf := &aws.IAMInstanceProfile{}
				if InstanceProfileArnRegex.MatchString(v) {
					iamInstanceProf.SetArn(spotinst.String(v))
				} else {
					iamInstanceProf.SetName(spotinst.String(v))
				}
				elastigroup.Compute.LaunchSpecification.SetIAMInstanceProfile(iamInstanceProf)
			} else {
				elastigroup.Compute.LaunchSpecification.SetIAMInstanceProfile(nil)
			}
			return nil
		},
		nil,
	)

	fieldsMap[KeyName] = commons.NewGenericField(
		commons.ElastigroupAWSLaunchConfiguration,
		KeyName,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *string = nil
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.KeyPair != nil {
				value = elastigroup.Compute.LaunchSpecification.KeyPair
			}
			if err := resourceData.Set(string(KeyName), spotinst.StringValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(KeyName), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(KeyName)).(string); ok && v != "" {
				elastigroup.Compute.LaunchSpecification.SetKeyPair(spotinst.String(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *string = nil
			if v, ok := resourceData.Get(string(KeyName)).(string); ok && v != "" {
				value = spotinst.String(v)
			}
			elastigroup.Compute.LaunchSpecification.SetKeyPair(value)
			return nil
		},
		nil,
	)

	fieldsMap[SecurityGroups] = commons.NewGenericField(
		commons.ElastigroupAWSLaunchConfiguration,
		SecurityGroups,
		&schema.Schema{
			Type:     schema.TypeList,
			Required: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value []string = nil
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.SecurityGroupIDs != nil {
				value = elastigroup.Compute.LaunchSpecification.SecurityGroupIDs
			}
			if err := resourceData.Set(string(SecurityGroups), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(SecurityGroups), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(SecurityGroups)).([]interface{}); ok {
				ids := make([]string, len(v))
				for i, j := range v {
					ids[i] = j.(string)
				}
				elastigroup.Compute.LaunchSpecification.SetSecurityGroupIDs(ids)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(SecurityGroups)).([]interface{}); ok {
				ids := make([]string, len(v))
				for i, j := range v {
					ids[i] = j.(string)
				}
				elastigroup.Compute.LaunchSpecification.SetSecurityGroupIDs(ids)
			}
			return nil
		},
		nil,
	)

	fieldsMap[UserData] = commons.NewGenericField(
		commons.ElastigroupAWSLaunchConfiguration,
		UserData,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
				// Sometimes the EC2 API responds with the equivalent, empty SHA1 sum
				if (old == "da39a3ee5e6b4b0d3255bfef95601890afd80709" && new == "") ||
					(old == "" && new == "da39a3ee5e6b4b0d3255bfef95601890afd80709") {
					return true
				}
				return false
			},
			StateFunc: Base64StateFunc,
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value = ""
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.UserData != nil {

				userData := elastigroup.Compute.LaunchSpecification.UserData
				userDataValue := spotinst.StringValue(userData)
				if userDataValue != "" {
					if isBase64Encoded(resourceData.Get(string(UserData)).(string)) {
						value = userDataValue
					} else {
						decodedUserData, _ := base64.StdEncoding.DecodeString(userDataValue)
						value = string(decodedUserData)
					}
				}
			}
			if err := resourceData.Set(string(UserData), Base64StateFunc(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(UserData), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(UserData)).(string); ok && v != "" {
				userData := spotinst.String(base64Encode(v))
				elastigroup.Compute.LaunchSpecification.SetUserData(userData)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var userData *string = nil
			if v, ok := resourceData.Get(string(UserData)).(string); ok && v != "" {
				userData = spotinst.String(base64Encode(v))
			}
			elastigroup.Compute.LaunchSpecification.SetUserData(userData)
			return nil
		},
		nil,
	)

	fieldsMap[ShutdownScript] = commons.NewGenericField(
		commons.ElastigroupAWSLaunchConfiguration,
		ShutdownScript,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
				// Sometimes the EC2 API responds with the equivalent, empty SHA1 sum
				if (old == "da39a3ee5e6b4b0d3255bfef95601890afd80709" && new == "") ||
					(old == "" && new == "da39a3ee5e6b4b0d3255bfef95601890afd80709") {
					return true
				}
				return false
			},
			StateFunc: Base64StateFunc,
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value = ""
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.ShutdownScript != nil {

				shutdownScript := elastigroup.Compute.LaunchSpecification.ShutdownScript
				shutdownScriptValue := spotinst.StringValue(shutdownScript)
				if shutdownScriptValue != "" {
					if isBase64Encoded(resourceData.Get(string(ShutdownScript)).(string)) {
						value = shutdownScriptValue
					} else {
						decodedShutdownScript, _ := base64.StdEncoding.DecodeString(shutdownScriptValue)
						value = string(decodedShutdownScript)
					}
				}
			}
			if err := resourceData.Set(string(ShutdownScript), Base64StateFunc(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ShutdownScript), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(ShutdownScript)).(string); ok && v != "" {
				shutdownScript := spotinst.String(base64Encode(v))
				elastigroup.Compute.LaunchSpecification.SetShutdownScript(shutdownScript)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var shutdownScript *string = nil
			if v, ok := resourceData.Get(string(ShutdownScript)).(string); ok && v != "" {
				shutdownScript = spotinst.String(base64Encode(v))
			}
			elastigroup.Compute.LaunchSpecification.SetShutdownScript(shutdownScript)
			return nil
		},
		nil,
	)

	fieldsMap[EnableMonitoring] = commons.NewGenericField(
		commons.ElastigroupAWSLaunchConfiguration,
		EnableMonitoring,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *bool = nil
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.Monitoring != nil {
				value = elastigroup.Compute.LaunchSpecification.Monitoring
			}
			if err := resourceData.Set(string(EnableMonitoring), spotinst.BoolValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(EnableMonitoring), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(EnableMonitoring)).(bool); ok {
				elastigroup.Compute.LaunchSpecification.SetMonitoring(spotinst.Bool(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(EnableMonitoring)).(bool); ok {
				elastigroup.Compute.LaunchSpecification.SetMonitoring(spotinst.Bool(v))
			}
			return nil
		},
		nil,
	)

	fieldsMap[EbsOptimized] = commons.NewGenericField(
		commons.ElastigroupAWSLaunchConfiguration,
		EbsOptimized,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *bool = nil
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.EBSOptimized != nil {
				value = elastigroup.Compute.LaunchSpecification.EBSOptimized
			}
			if err := resourceData.Set(string(EbsOptimized), spotinst.BoolValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(EbsOptimized), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(EbsOptimized)).(bool); ok {
				elastigroup.Compute.LaunchSpecification.SetEBSOptimized(spotinst.Bool(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(EbsOptimized)).(bool); ok {
				elastigroup.Compute.LaunchSpecification.SetEBSOptimized(spotinst.Bool(v))
			}
			return nil
		},
		nil,
	)

	fieldsMap[PlacementTenancy] = commons.NewGenericField(
		commons.ElastigroupAWSLaunchConfiguration,
		PlacementTenancy,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *string = nil
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.Tenancy != nil {
				value = elastigroup.Compute.LaunchSpecification.Tenancy
			}
			if err := resourceData.Set(string(PlacementTenancy), spotinst.StringValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(PlacementTenancy), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(PlacementTenancy)).(string); ok && v != "" {
				tenancy := spotinst.String(v)
				elastigroup.Compute.LaunchSpecification.SetTenancy(tenancy)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var tenancy *string = nil
			if v, ok := resourceData.Get(string(PlacementTenancy)).(string); ok && v != "" {
				tenancy = spotinst.String(v)
			}
			elastigroup.Compute.LaunchSpecification.SetTenancy(tenancy)
			return nil
		},
		nil,
	)

	fieldsMap[CPUCredits] = commons.NewGenericField(
		commons.ElastigroupAWSLaunchConfiguration,
		CPUCredits,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *string = nil
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.CreditSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.CreditSpecification.CPUCredits != nil {
				value = elastigroup.Compute.LaunchSpecification.CreditSpecification.CPUCredits
			}
			if err := resourceData.Set(string(CPUCredits), spotinst.StringValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(CPUCredits), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(CPUCredits)).(string); ok && v != "" {
				if elastigroup.Compute.LaunchSpecification.CreditSpecification == nil {
					elastigroup.Compute.LaunchSpecification.CreditSpecification = &aws.CreditSpecification{}
				}
				credits := spotinst.String(v)
				elastigroup.Compute.LaunchSpecification.CreditSpecification.SetCPUCredits(credits)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			credSpec := &aws.CreditSpecification{}
			if v, ok := resourceData.Get(string(CPUCredits)).(string); ok && v != "" {
				credSpec.SetCPUCredits(spotinst.String(v))
				elastigroup.Compute.LaunchSpecification.SetCreditSpecification(credSpec)
			} else {
				elastigroup.Compute.LaunchSpecification.SetCreditSpecification(nil)
			}
			return nil
		},
		nil,
	)

	fieldsMap[MetadataOptions] = commons.NewGenericField(
		commons.ElastigroupAWSLaunchConfiguration,
		MetadataOptions,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(HTTPTokens): {
						Type:     schema.TypeString,
						Required: true,
					},

					string(HTTPPutResponseHopLimit): {
						Type:     schema.TypeInt,
						Optional: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var result []interface{} = nil
			if elastigroup != nil && elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.MetadataOptions != nil {
				result = flattenMetadataOptions(elastigroup.Compute.LaunchSpecification.MetadataOptions)
			}

			if result != nil {
				if err := resourceData.Set(string(MetadataOptions), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(MetadataOptions), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(MetadataOptions)); ok {
				if metaDataOptions, err := expandMetadataOptions(v); err != nil {
					return err
				} else {
					elastigroup.Compute.LaunchSpecification.SetMetadataOptions(metaDataOptions)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *aws.MetadataOptions = nil
			if v, ok := resourceData.GetOk(string(MetadataOptions)); ok {
				if metaDataOptions, err := expandMetadataOptions(v); err != nil {
					return err
				} else {
					value = metaDataOptions
				}
			}
			elastigroup.Compute.LaunchSpecification.SetMetadataOptions(value)
			return nil
		},
		nil,
	)

	fieldsMap[CPUOptions] = commons.NewGenericField(
		commons.ElastigroupAWSLaunchConfiguration,
		CPUOptions,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(ThreadsPerCore): {
						Type:     schema.TypeInt,
						Required: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var result []interface{} = nil
			if elastigroup != nil && elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.CPUOptions != nil {
				result = flattenCPUOptions(elastigroup.Compute.LaunchSpecification.CPUOptions)
			}

			if result != nil {
				if err := resourceData.Set(string(CPUOptions), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(CPUOptions), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(CPUOptions)); ok {
				if cpuOptions, err := expandCPUOptions(v); err != nil {
					return err
				} else {
					elastigroup.Compute.LaunchSpecification.SetCPUOptions(cpuOptions)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *aws.CPUOptions = nil
			if v, ok := resourceData.GetOk(string(CPUOptions)); ok {
				if cpuOptions, err := expandCPUOptions(v); err != nil {
					return err
				} else {
					value = cpuOptions
				}
			}
			elastigroup.Compute.LaunchSpecification.SetCPUOptions(value)
			return nil
		},
		nil,
	)

	fieldsMap[ResourceTagSpecification] = commons.NewGenericField(
		commons.ElastigroupAWSLaunchConfiguration,
		ResourceTagSpecification,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(ShouldTagVolumes): {
						Type:     schema.TypeBool,
						Optional: true,
					},
					string(ShouldTagAMIs): {
						Type:     schema.TypeBool,
						Optional: true,
					},
					string(ShouldTagENIs): {
						Type:     schema.TypeBool,
						Optional: true,
					},
					string(ShouldTagSnapshots): {
						Type:     schema.TypeBool,
						Optional: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var result []interface{} = nil
			if elastigroup != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.ResourceTagSpecification != nil {
				resourceTagSpecification := elastigroup.Compute.LaunchSpecification.ResourceTagSpecification
				result = flattenResourceTagSpecification(resourceTagSpecification)
			}
			if len(result) > 0 {
				if err := resourceData.Set(string(ResourceTagSpecification), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ResourceTagSpecification), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(ResourceTagSpecification)); ok {
				if v, err := expandResourceTagSpecification(v); err != nil {
					return err
				} else {
					elastigroup.Compute.LaunchSpecification.SetResourceTagSpecification(v)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *aws.ResourceTagSpecification = nil
			if v, ok := resourceData.GetOk(string(ResourceTagSpecification)); ok {
				if resourceTagSpecification, err := expandResourceTagSpecification(v); err != nil {
					return err
				} else {
					value = resourceTagSpecification
				}
			}
			elastigroup.Compute.LaunchSpecification.SetResourceTagSpecification(value)
			return nil
		},
		nil,
	)
}

var InstanceProfileArnRegex = regexp.MustCompile(`arn:aws:iam::\d{12}:instance-profile/?[a-zA-Z_0-9+=,.@\-_/]+`)

func Base64StateFunc(v interface{}) string {
	if isBase64Encoded(v.(string)) {
		return v.(string)
	} else {
		return base64Encode(v.(string))
	}
}

// base64Encode encodes data if the input isn't already encoded using
// base64.StdEncoding.EncodeToString. If the input is already base64 encoded,
// return the original input unchanged.
func base64Encode(data string) string {
	// Check whether the data is already Base64 encoded; don't double-encode
	if isBase64Encoded(data) {
		return data
	}
	// data has not been encoded encode and return
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func isBase64Encoded(data string) bool {
	_, err := base64.StdEncoding.DecodeString(data)
	return err == nil
}

func expandMetadataOptions(data interface{}) (*aws.MetadataOptions, error) {
	metadataOptions := &aws.MetadataOptions{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return metadataOptions, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(HTTPTokens)].(string); ok && v != "" {
		metadataOptions.SetHTTPTokens(spotinst.String(v))
	}
	if v, ok := m[string(HTTPPutResponseHopLimit)].(int); ok && v >= 0 {
		metadataOptions.SetHTTPPutResponseHopLimit(spotinst.Int(v))
	} else {
		metadataOptions.SetHTTPPutResponseHopLimit(nil)
	}

	return metadataOptions, nil
}

func flattenMetadataOptions(metadataOptions *aws.MetadataOptions) []interface{} {
	result := make(map[string]interface{})
	result[string(HTTPTokens)] = spotinst.StringValue(metadataOptions.HTTPTokens)
	result[string(HTTPPutResponseHopLimit)] = spotinst.IntValue(metadataOptions.HTTPPutResponseHopLimit)

	return []interface{}{result}
}

func expandCPUOptions(data interface{}) (*aws.CPUOptions, error) {
	cpuOptions := &aws.CPUOptions{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return cpuOptions, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(ThreadsPerCore)].(int); ok && v >= 0 {
		cpuOptions.SetThreadsPerCore(spotinst.Int(v))
	} else {
		cpuOptions.SetThreadsPerCore(nil)
	}

	return cpuOptions, nil
}

func flattenCPUOptions(cpuOptions *aws.CPUOptions) []interface{} {
	result := make(map[string]interface{})
	result[string(ThreadsPerCore)] = spotinst.IntValue(cpuOptions.ThreadsPerCore)

	return []interface{}{result}
}

func expandResourceTagSpecification(data interface{}) (*aws.ResourceTagSpecification, error) {
	resourceTagSpecification := &aws.ResourceTagSpecification{}
	list := data.([]interface{})

	if list == nil || list[0] == nil {
		return resourceTagSpecification, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(ShouldTagVolumes)].(bool); ok {
		volumes := &aws.Volumes{}
		resourceTagSpecification.SetVolumes(volumes)
		resourceTagSpecification.Volumes.SetShouldTag(spotinst.Bool(v))

	}
	if v, ok := m[string(ShouldTagAMIs)].(bool); ok {
		anis := &aws.AMIs{}
		resourceTagSpecification.SetAMIs(anis)
		resourceTagSpecification.AMIs.SetShouldTag(spotinst.Bool(v))

	}
	if v, ok := m[string(ShouldTagENIs)].(bool); ok {
		enis := &aws.ENIs{}
		resourceTagSpecification.SetENIs(enis)
		resourceTagSpecification.ENIs.SetShouldTag(spotinst.Bool(v))
	}
	if v, ok := m[string(ShouldTagSnapshots)].(bool); ok {
		snapshots := &aws.Snapshots{}
		resourceTagSpecification.SetSnapshots(snapshots)
		resourceTagSpecification.Snapshots.SetShouldTag(spotinst.Bool(v))

	}

	return resourceTagSpecification, nil
}

func flattenResourceTagSpecification(resourceTagSpecification *aws.ResourceTagSpecification) []interface{} {
	result := make(map[string]interface{})
	if resourceTagSpecification != nil && resourceTagSpecification.Snapshots != nil {
		result[string(ShouldTagSnapshots)] = spotinst.BoolValue(resourceTagSpecification.Snapshots.ShouldTag)
	}
	if resourceTagSpecification != nil && resourceTagSpecification.ENIs != nil {
		result[string(ShouldTagENIs)] = spotinst.BoolValue(resourceTagSpecification.ENIs.ShouldTag)
	}
	if resourceTagSpecification != nil && resourceTagSpecification.AMIs != nil {
		result[string(ShouldTagAMIs)] = spotinst.BoolValue(resourceTagSpecification.AMIs.ShouldTag)
	}
	if resourceTagSpecification != nil && resourceTagSpecification.Volumes != nil {
		result[string(ShouldTagVolumes)] = spotinst.BoolValue(resourceTagSpecification.Volumes.ShouldTag)
	}

	return []interface{}{result}
}
