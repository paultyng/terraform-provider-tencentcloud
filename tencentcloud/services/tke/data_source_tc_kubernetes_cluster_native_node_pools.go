// Code generated by iacg; DO NOT EDIT.
package tke

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tke2 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tke/v20220501"

	tccommon "github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/common"
	"github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/internal/helper"
)

func DataSourceTencentCloudKubernetesClusterNativeNodePools() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTencentCloudKubernetesClusterNativeNodePoolsRead,
		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "ID of the cluster.",
			},

			"filters": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Query filter conditions: NodePoolsName, Filter according to the node pool name, type: String, required: no. NodePoolsId, Filter according to the node pool ID, type: String, required: no. tags, Filter according to the label key value pairs, type: String, required: no. tag:tag-key, Filter according to the label key value pairs, type: String, required: no.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The attribute name, if there are multiple filters, the relationship between the filters is a logical AND relationship.",
						},
						"values": {
							Type:        schema.TypeSet,
							Required:    true,
							Description: "Attribute values, if there are multiple values in the same filter, the relationship between values under the same filter is a logical OR relationship.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},

			"node_pools": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Node pool list.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cluster_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the cluster.",
						},
						"node_pool_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the node pool.",
						},
						"tags": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Node tags.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"resource_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The resource type bound to the label.",
									},
									"tags": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Tag pair list.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"key": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Tag Key.",
												},
												"value": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Tag Value.",
												},
											},
										},
									},
								},
							},
						},
						"taints": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "node taint.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Key of the taint.",
									},
									"value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Value of the taint.",
									},
									"effect": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Effect of the taint.",
									},
								},
							},
						},
						"deletion_protection": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether to enable deletion protection.",
						},
						"unschedulable": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the node is not schedulable by default.",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Node pool type. Optional value is `Native`.",
						},
						"labels": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Node Labels.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Name in the map table.",
									},
									"value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Value in map table.",
									},
								},
							},
						},
						"life_state": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Node pool status.",
						},
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Creation time.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Node pool name.",
						},
						"native": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Native node pool creation parameters.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"scaling": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Node pool scaling configuration.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"min_replicas": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Minimum number of replicas in node pool.",
												},
												"max_replicas": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Maximum number of replicas in node pool.",
												},
												"create_policy": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Node pool expansion strategy. `ZoneEquality`: multiple availability zones are broken up; `ZonePriority`: the preferred availability zone takes precedence.",
												},
											},
										},
									},
									"subnet_ids": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Subnet list.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"security_group_ids": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Security group list.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"upgrade_settings": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Automatically upgrade configuration.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"auto_upgrade": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Whether to enable automatic upgrade.",
												},
												"upgrade_options": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Operation and maintenance window.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"auto_upgrade_start_time": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Automatic upgrade start time.",
															},
															"duration": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Automatic upgrade duration.",
															},
															"weekly_period": {
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Operation and maintenance date.",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
														},
													},
												},
												"components": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Upgrade items.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"max_unavailable": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "When upgrading, the maximum number of nodes that cannot be upgraded.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"type": {
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Numeric type, 0 is int, 1 is string.",
															},
															"int_val": {
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Integer.",
															},
															"str_val": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "String.",
															},
														},
													},
												},
											},
										},
									},
									"auto_repair": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether to enable self-healing ability.",
									},
									"instance_charge_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Node billing type. `PREPAID` is a yearly and monthly subscription, `POSTPAID_BY_HOUR` is a pay-as-you-go plan. The default is `POSTPAID_BY_HOUR`.",
									},
									"instance_charge_prepaid": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Billing configuration for yearly and monthly models.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"period": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Postpaid billing cycle, unit (month): 1, 2, 3, 4, 5,, 6, 7, 8, 9, 10, 11, 12, 24, 36, 48, 60.",
												},
												"renew_flag": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Prepaid renewal method:\n  - `NOTIFY_AND_AUTO_RENEW`: Notify users of expiration and automatically renew (default).\n  - `NOTIFY_AND_MANUAL_RENEW`: Notify users of expiration, but do not automatically renew.\n  - `DISABLE_NOTIFY_AND_MANUAL_RENEW`: Do not notify users of expiration and do not automatically renew.",
												},
											},
										},
									},
									"system_disk": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "System disk configuration.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"disk_type": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Cloud disk type.",
												},
												"disk_size": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Cloud disk size (G).",
												},
												"auto_format_and_mount": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Whether to automatically format the disk and mount it.",
												},
												"file_system": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "File system.",
												},
												"mount_target": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Mount directory.",
												},
											},
										},
									},
									"key_ids": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Node pool ssh public key id array.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"management": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Node pool management parameter settings.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"nameservers": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Dns configuration.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"hosts": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Hosts configuration.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"kernel_args": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Kernel parameter configuration.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"health_check_policy_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Fault self-healing rule name.",
									},
									"host_name_pattern": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Native node pool hostName pattern string.",
									},
									"kubelet_args": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Kubelet custom parameters.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"lifecycle": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Predefined scripts.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"pre_init": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Custom script before node initialization.",
												},
												"post_init": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Custom script after node initialization.",
												},
											},
										},
									},
									"runtime_root_dir": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Runtime root directory.",
									},
									"enable_autoscaling": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether to enable elastic scaling.",
									},
									"instance_types": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Model list.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"replicas": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Desired number of nodes.",
									},
									"internet_accessible": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Public network bandwidth settings.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"max_bandwidth_out": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Maximum bandwidth output.",
												},
												"charge_type": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Network billing method.",
												},
												"bandwidth_package_id": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Bandwidth package ID.",
												},
											},
										},
									},
									"data_disks": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Native node pool data disk list.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"disk_type": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Cloud disk type.",
												},
												"file_system": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "File system (ext3/ext4/xfs).",
												},
												"disk_size": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Cloud disk size (G).",
												},
												"auto_format_and_mount": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Whether to automatically format the disk and mount it.",
												},
												"disk_partition": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Mount device name or partition name.",
												},
												"mount_target": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Mount directory.",
												},
												"encrypt": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Pass in this parameter to create an encrypted cloud disk. The value is fixed to `ENCRYPT`.",
												},
												"kms_key_id": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Customize the key when purchasing an encrypted disk. When this parameter is passed in, the Encrypt parameter is not empty.",
												},
												"snapshot_id": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Snapshot ID. If passed in, the cloud disk will be created based on this snapshot. The snapshot type must be a data disk snapshot.",
												},
												"throughput_performance": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Cloud disk performance, unit: MB/s. Use this parameter to purchase additional performance for the cloud disk.",
												},
											},
										},
									},
								},
							},
						},
						"annotations": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Node Annotation List.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Name in the map table.",
									},
									"value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Value in the map table.",
									},
								},
							},
						},
					},
				},
			},

			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
		},
	}
}

func dataSourceTencentCloudKubernetesClusterNativeNodePoolsRead(d *schema.ResourceData, meta interface{}) error {
	defer tccommon.LogElapsed("data_source.tencentcloud_kubernetes_cluster_native_node_pools.read")()
	defer tccommon.InconsistentCheck(d, meta)()

	logId := tccommon.GetLogId(nil)
	ctx := tccommon.NewResourceLifeCycleHandleFuncContext(context.Background(), logId, d, meta)

	service := TkeService{client: meta.(tccommon.ProviderMeta).GetAPIV3Conn()}

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("cluster_id"); ok {
		paramMap["ClusterId"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("filters"); ok {
		filtersSet := v.([]interface{})
		tmpSet := make([]*tke2.Filter, 0, len(filtersSet))
		for _, item := range filtersSet {
			filtersMap := item.(map[string]interface{})
			filter := tke2.Filter{}
			if v, ok := filtersMap["name"]; ok {
				filter.Name = helper.String(v.(string))
			}
			if v, ok := filtersMap["values"]; ok {
				valuesSet := v.(*schema.Set).List()
				for i := range valuesSet {
					values := valuesSet[i].(string)
					filter.Values = append(filter.Values, helper.String(values))
				}
			}
			tmpSet = append(tmpSet, &filter)
		}
		paramMap["Filters"] = tmpSet
	}

	var respData []*tke2.NodePool
	err := resource.Retry(tccommon.ReadRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeKubernetesClusterNativeNodePoolsByFilter(ctx, paramMap)
		if e != nil {
			return tccommon.RetryError(e)
		}
		respData = result
		return nil
	})
	if err != nil {
		return err
	}

	ids := make([]string, 0, len(respData))
	nodePoolsList := make([]map[string]interface{}, 0, len(respData))
	if respData != nil {
		for _, nodePools := range respData {
			nodePoolsMap := map[string]interface{}{}

			if nodePools.ClusterId != nil {
				nodePoolsMap["cluster_id"] = nodePools.ClusterId
			}

			var nodePoolId string
			if nodePools.NodePoolId != nil {
				nodePoolsMap["node_pool_id"] = nodePools.NodePoolId
				nodePoolId = *nodePools.NodePoolId
			}

			tagsList := make([]map[string]interface{}, 0, len(nodePools.Tags))
			if nodePools.Tags != nil {
				for _, tags := range nodePools.Tags {
					tagsMap := map[string]interface{}{}

					if tags.ResourceType != nil {
						tagsMap["resource_type"] = tags.ResourceType
					}

					tagsList2 := make([]map[string]interface{}, 0, len(tags.Tags))
					if tags.Tags != nil {
						for _, tags := range tags.Tags {
							tagsMap2 := map[string]interface{}{}

							if tags.Key != nil {
								tagsMap2["key"] = tags.Key
							}

							if tags.Value != nil {
								tagsMap2["value"] = tags.Value
							}

							tagsList2 = append(tagsList2, tagsMap2)
						}

						tagsMap["tags"] = tagsList2
					}
					tagsList = append(tagsList, tagsMap)
				}

				nodePoolsMap["tags"] = tagsList
			}
			taintsList := make([]map[string]interface{}, 0, len(nodePools.Taints))
			if nodePools.Taints != nil {
				for _, taints := range nodePools.Taints {
					taintsMap := map[string]interface{}{}

					if taints.Key != nil {
						taintsMap["key"] = taints.Key
					}

					if taints.Value != nil {
						taintsMap["value"] = taints.Value
					}

					if taints.Effect != nil {
						taintsMap["effect"] = taints.Effect
					}

					taintsList = append(taintsList, taintsMap)
				}

				nodePoolsMap["taints"] = taintsList
			}
			if nodePools.DeletionProtection != nil {
				nodePoolsMap["deletion_protection"] = nodePools.DeletionProtection
			}

			if nodePools.Unschedulable != nil {
				nodePoolsMap["unschedulable"] = nodePools.Unschedulable
			}

			if nodePools.Type != nil {
				nodePoolsMap["type"] = nodePools.Type
			}

			labelsList := make([]map[string]interface{}, 0, len(nodePools.Labels))
			if nodePools.Labels != nil {
				for _, labels := range nodePools.Labels {
					labelsMap := map[string]interface{}{}

					if labels.Name != nil {
						labelsMap["name"] = labels.Name
					}

					if labels.Value != nil {
						labelsMap["value"] = labels.Value
					}

					labelsList = append(labelsList, labelsMap)
				}

				nodePoolsMap["labels"] = labelsList
			}
			if nodePools.LifeState != nil {
				nodePoolsMap["life_state"] = nodePools.LifeState
			}

			if nodePools.CreatedAt != nil {
				nodePoolsMap["created_at"] = nodePools.CreatedAt
			}

			if nodePools.Name != nil {
				nodePoolsMap["name"] = nodePools.Name
			}

			nativeMap := map[string]interface{}{}

			if nodePools.Native != nil {
				scalingMap := map[string]interface{}{}

				if nodePools.Native.Scaling != nil {
					if nodePools.Native.Scaling.MinReplicas != nil {
						scalingMap["min_replicas"] = nodePools.Native.Scaling.MinReplicas
					}

					if nodePools.Native.Scaling.MaxReplicas != nil {
						scalingMap["max_replicas"] = nodePools.Native.Scaling.MaxReplicas
					}

					if nodePools.Native.Scaling.CreatePolicy != nil {
						scalingMap["create_policy"] = nodePools.Native.Scaling.CreatePolicy
					}

					nativeMap["scaling"] = []interface{}{scalingMap}
				}

				if nodePools.Native.SubnetIds != nil {
					nativeMap["subnet_ids"] = nodePools.Native.SubnetIds
				}

				if nodePools.Native.SecurityGroupIds != nil {
					nativeMap["security_group_ids"] = nodePools.Native.SecurityGroupIds
				}

				upgradeSettingsMap := map[string]interface{}{}

				if nodePools.Native.UpgradeSettings != nil {
					if nodePools.Native.UpgradeSettings.AutoUpgrade != nil {
						upgradeSettingsMap["auto_upgrade"] = nodePools.Native.UpgradeSettings.AutoUpgrade
					}

					upgradeOptionsMap := map[string]interface{}{}

					if nodePools.Native.UpgradeSettings.UpgradeOptions != nil {
						if nodePools.Native.UpgradeSettings.UpgradeOptions.AutoUpgradeStartTime != nil {
							upgradeOptionsMap["auto_upgrade_start_time"] = nodePools.Native.UpgradeSettings.UpgradeOptions.AutoUpgradeStartTime
						}

						if nodePools.Native.UpgradeSettings.UpgradeOptions.Duration != nil {
							upgradeOptionsMap["duration"] = nodePools.Native.UpgradeSettings.UpgradeOptions.Duration
						}

						if nodePools.Native.UpgradeSettings.UpgradeOptions.WeeklyPeriod != nil {
							upgradeOptionsMap["weekly_period"] = nodePools.Native.UpgradeSettings.UpgradeOptions.WeeklyPeriod
						}

						upgradeSettingsMap["upgrade_options"] = []interface{}{upgradeOptionsMap}
					}

					if nodePools.Native.UpgradeSettings.Components != nil {
						upgradeSettingsMap["components"] = nodePools.Native.UpgradeSettings.Components
					}

					maxUnavailableMap := map[string]interface{}{}

					if nodePools.Native.UpgradeSettings.MaxUnavailable != nil {
						if nodePools.Native.UpgradeSettings.MaxUnavailable.Type != nil {
							maxUnavailableMap["type"] = nodePools.Native.UpgradeSettings.MaxUnavailable.Type
						}

						if nodePools.Native.UpgradeSettings.MaxUnavailable.IntVal != nil {
							maxUnavailableMap["int_val"] = nodePools.Native.UpgradeSettings.MaxUnavailable.IntVal
						}

						if nodePools.Native.UpgradeSettings.MaxUnavailable.StrVal != nil {
							maxUnavailableMap["str_val"] = nodePools.Native.UpgradeSettings.MaxUnavailable.StrVal
						}

						upgradeSettingsMap["max_unavailable"] = []interface{}{maxUnavailableMap}
					}

					nativeMap["upgrade_settings"] = []interface{}{upgradeSettingsMap}
				}

				if nodePools.Native.AutoRepair != nil {
					nativeMap["auto_repair"] = nodePools.Native.AutoRepair
				}

				if nodePools.Native.InstanceChargeType != nil {
					nativeMap["instance_charge_type"] = nodePools.Native.InstanceChargeType
				}

				instanceChargePrepaidMap := map[string]interface{}{}

				if nodePools.Native.InstanceChargePrepaid != nil {
					if nodePools.Native.InstanceChargePrepaid.Period != nil {
						instanceChargePrepaidMap["period"] = nodePools.Native.InstanceChargePrepaid.Period
					}

					if nodePools.Native.InstanceChargePrepaid.RenewFlag != nil {
						instanceChargePrepaidMap["renew_flag"] = nodePools.Native.InstanceChargePrepaid.RenewFlag
					}

					nativeMap["instance_charge_prepaid"] = []interface{}{instanceChargePrepaidMap}
				}

				systemDiskMap := map[string]interface{}{}

				if nodePools.Native.SystemDisk != nil {
					if nodePools.Native.SystemDisk.DiskType != nil {
						systemDiskMap["disk_type"] = nodePools.Native.SystemDisk.DiskType
					}

					if nodePools.Native.SystemDisk.DiskSize != nil {
						systemDiskMap["disk_size"] = nodePools.Native.SystemDisk.DiskSize
					}

					if nodePools.Native.SystemDisk.AutoFormatAndMount != nil {
						systemDiskMap["auto_format_and_mount"] = nodePools.Native.SystemDisk.AutoFormatAndMount
					}

					if nodePools.Native.SystemDisk.FileSystem != nil {
						systemDiskMap["file_system"] = nodePools.Native.SystemDisk.FileSystem
					}

					if nodePools.Native.SystemDisk.MountTarget != nil {
						systemDiskMap["mount_target"] = nodePools.Native.SystemDisk.MountTarget
					}

					nativeMap["system_disk"] = []interface{}{systemDiskMap}
				}

				if nodePools.Native.KeyIds != nil {
					nativeMap["key_ids"] = nodePools.Native.KeyIds
				}

				managementMap := map[string]interface{}{}

				if nodePools.Native.Management != nil {
					if nodePools.Native.Management.Nameservers != nil {
						managementMap["nameservers"] = nodePools.Native.Management.Nameservers
					}

					if nodePools.Native.Management.Hosts != nil {
						managementMap["hosts"] = nodePools.Native.Management.Hosts
					}

					if nodePools.Native.Management.KernelArgs != nil {
						managementMap["kernel_args"] = nodePools.Native.Management.KernelArgs
					}

					nativeMap["management"] = []interface{}{managementMap}
				}

				if nodePools.Native.HealthCheckPolicyName != nil {
					nativeMap["health_check_policy_name"] = nodePools.Native.HealthCheckPolicyName
				}

				if nodePools.Native.HostNamePattern != nil {
					nativeMap["host_name_pattern"] = nodePools.Native.HostNamePattern
				}

				if nodePools.Native.KubeletArgs != nil {
					nativeMap["kubelet_args"] = nodePools.Native.KubeletArgs
				}

				lifecycleMap := map[string]interface{}{}

				if nodePools.Native.Lifecycle != nil {
					if nodePools.Native.Lifecycle.PreInit != nil {
						lifecycleMap["pre_init"] = nodePools.Native.Lifecycle.PreInit
					}

					if nodePools.Native.Lifecycle.PostInit != nil {
						lifecycleMap["post_init"] = nodePools.Native.Lifecycle.PostInit
					}

					nativeMap["lifecycle"] = []interface{}{lifecycleMap}
				}

				if nodePools.Native.RuntimeRootDir != nil {
					nativeMap["runtime_root_dir"] = nodePools.Native.RuntimeRootDir
				}

				if nodePools.Native.EnableAutoscaling != nil {
					nativeMap["enable_autoscaling"] = nodePools.Native.EnableAutoscaling
				}

				if nodePools.Native.InstanceTypes != nil {
					nativeMap["instance_types"] = nodePools.Native.InstanceTypes
				}

				if nodePools.Native.Replicas != nil {
					nativeMap["replicas"] = nodePools.Native.Replicas
				}

				internetAccessibleMap := map[string]interface{}{}

				if nodePools.Native.InternetAccessible != nil {
					if nodePools.Native.InternetAccessible.MaxBandwidthOut != nil {
						internetAccessibleMap["max_bandwidth_out"] = nodePools.Native.InternetAccessible.MaxBandwidthOut
					}

					if nodePools.Native.InternetAccessible.ChargeType != nil {
						internetAccessibleMap["charge_type"] = nodePools.Native.InternetAccessible.ChargeType
					}

					if nodePools.Native.InternetAccessible.BandwidthPackageId != nil {
						internetAccessibleMap["bandwidth_package_id"] = nodePools.Native.InternetAccessible.BandwidthPackageId
					}

					nativeMap["internet_accessible"] = []interface{}{internetAccessibleMap}
				}

				dataDisksList := make([]map[string]interface{}, 0, len(nodePools.Native.DataDisks))
				if nodePools.Native.DataDisks != nil {
					for _, dataDisks := range nodePools.Native.DataDisks {
						dataDisksMap := map[string]interface{}{}

						if dataDisks.DiskType != nil {
							dataDisksMap["disk_type"] = dataDisks.DiskType
						}

						if dataDisks.FileSystem != nil {
							dataDisksMap["file_system"] = dataDisks.FileSystem
						}

						if dataDisks.DiskSize != nil {
							dataDisksMap["disk_size"] = dataDisks.DiskSize
						}

						if dataDisks.AutoFormatAndMount != nil {
							dataDisksMap["auto_format_and_mount"] = dataDisks.AutoFormatAndMount
						}

						if dataDisks.DiskPartition != nil {
							dataDisksMap["disk_partition"] = dataDisks.DiskPartition
						}

						if dataDisks.MountTarget != nil {
							dataDisksMap["mount_target"] = dataDisks.MountTarget
						}

						if dataDisks.Encrypt != nil {
							dataDisksMap["encrypt"] = dataDisks.Encrypt
						}

						if dataDisks.KmsKeyId != nil {
							dataDisksMap["kms_key_id"] = dataDisks.KmsKeyId
						}

						if dataDisks.SnapshotId != nil {
							dataDisksMap["snapshot_id"] = dataDisks.SnapshotId
						}

						if dataDisks.ThroughputPerformance != nil {
							dataDisksMap["throughput_performance"] = dataDisks.ThroughputPerformance
						}

						dataDisksList = append(dataDisksList, dataDisksMap)
					}

					nativeMap["data_disks"] = dataDisksList
				}
				nodePoolsMap["native"] = []interface{}{nativeMap}
			}

			annotationsList := make([]map[string]interface{}, 0, len(nodePools.Annotations))
			if nodePools.Annotations != nil {
				for _, annotations := range nodePools.Annotations {
					annotationsMap := map[string]interface{}{}

					if annotations.Name != nil {
						annotationsMap["name"] = annotations.Name
					}

					if annotations.Value != nil {
						annotationsMap["value"] = annotations.Value
					}

					annotationsList = append(annotationsList, annotationsMap)
				}

				nodePoolsMap["annotations"] = annotationsList
			}
			ids = append(ids, nodePoolId)
			nodePoolsList = append(nodePoolsList, nodePoolsMap)
		}

		_ = d.Set("node_pools", nodePoolsList)
	}

	d.SetId(helper.DataResourceIdsHash(ids))

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := tccommon.WriteToFile(output.(string), nodePoolsList); e != nil {
			return e
		}
	}

	return nil
}
