package elastigroup_aws_launch_configuration

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	Prefix = "launch_configuration_"
)

const (
	ImageId            commons.FieldName = "image_id"
	IamInstanceProfile commons.FieldName = "iam_instance_profile"
	KeyName            commons.FieldName = "key_name"
	SecurityGroups     commons.FieldName = "security_groups"
	UserData           commons.FieldName = "user_data"
	ShutdownScript     commons.FieldName = "shutdown_script"
	EnableMonitoring   commons.FieldName = "enable_monitoring"
	EbsOptimized       commons.FieldName = "ebs_optimized"
	PlacementTenancy   commons.FieldName = "placement_tenancy"
	CPUCredits         commons.FieldName = "cpu_credits"
	MetadataOptions    commons.FieldName = "metadata_options"
	CPUOptions         commons.FieldName = "cpu_options"

	// - MetadataOptions -----------------------------
	HTTPTokens              commons.FieldName = "http_tokens"
	HTTPPutResponseHopLimit commons.FieldName = "http_put_response_hop_limit"
	// -----------------------------------

	// - CPUOptions -----------------------------
	ThreadsPerCore commons.FieldName = "threads_per_core"
	// -----------------------------------
)

const (
	ResourceTagSpecification commons.FieldName = "resource_tag_specification"
	ShouldTagVolumes         commons.FieldName = "should_tag_volumes"
	ShouldTagSnapshots       commons.FieldName = "should_tag_snapshots"
	ShouldTagENIs            commons.FieldName = "should_tag_enis"
	ShouldTagAMIs            commons.FieldName = "should_tag_amis"
)
