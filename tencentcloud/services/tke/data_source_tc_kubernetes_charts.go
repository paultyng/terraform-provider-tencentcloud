// Code generated by iacg; DO NOT EDIT.
package tke

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tkev20180525 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tke/v20180525"

	tccommon "github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/common"
	"github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/internal/helper"
)

func DataSourceTencentCloudKubernetesCharts() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTencentCloudKubernetesChartsRead,
		Schema: map[string]*schema.Schema{
			"kind": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Kind of app chart. Available values: `log`, `scheduler`, `network`, `storage`, `monitor`, `dns`, `image`, `other`, `invisible`.",
			},

			"arch": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Operation system app supported. Available values: `arm32`, `arm64`, `amd64`.",
			},

			"cluster_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Cluster type. Available values: `tke`, `eks`.",
			},

			"chart_list": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "App chart list.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of chart.",
						},
						"label": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "Label of chart.",
						},
						"latest_version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Chart latest version.",
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

func dataSourceTencentCloudKubernetesChartsRead(d *schema.ResourceData, meta interface{}) error {
	defer tccommon.LogElapsed("data_source.tencentcloud_kubernetes_charts.read")()
	defer tccommon.InconsistentCheck(d, meta)()

	logId := tccommon.GetLogId(nil)
	ctx := tccommon.NewResourceLifeCycleHandleFuncContext(context.Background(), logId, d, meta)

	service := TkeService{client: meta.(tccommon.ProviderMeta).GetAPIV3Conn()}

	var (
		kind        string
		arch        string
		clusterType string
	)
	if v, ok := d.GetOk("kind"); ok {
		kind = v.(string)
	}
	if v, ok := d.GetOk("arch"); ok {
		arch = v.(string)
	}
	if v, ok := d.GetOk("cluster_type"); ok {
		clusterType = v.(string)
	}
	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("kind"); ok {
		paramMap["Kind"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("arch"); ok {
		paramMap["Arch"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("cluster_type"); ok {
		paramMap["ClusterType"] = helper.String(v.(string))
	}

	var respData []*tkev20180525.AppChart
	err := resource.Retry(tccommon.ReadRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeKubernetesChartsByFilter(ctx, paramMap)
		if e != nil {
			return tccommon.RetryError(e)
		}
		respData = result
		return nil
	})
	if err != nil {
		return err
	}

	appChartsList := make([]map[string]interface{}, 0, len(respData))
	if respData != nil {
		for _, appCharts := range respData {
			appChartsMap := map[string]interface{}{}

			if appCharts.Name != nil {
				appChartsMap["name"] = appCharts.Name
			}

			if appCharts.LatestVersion != nil {
				appChartsMap["latest_version"] = appCharts.LatestVersion
			}

			if appCharts.Label != nil {
				tmpMap := make(map[string]interface{})
				if err := json.Unmarshal([]byte(*appCharts.Label), &tmpMap); err != nil {
					return err
				}
				appChartsMap["label"] = tmpMap
			}

			appChartsList = append(appChartsList, appChartsMap)
		}

		_ = d.Set("chart_list", appChartsList)
	}

	d.SetId(strings.Join([]string{kind, arch, clusterType}, tccommon.FILED_SP))

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := tccommon.WriteToFile(output.(string), appChartsList); e != nil {
			return e
		}
	}

	return nil
}
