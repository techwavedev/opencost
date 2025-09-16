package prom

import (
	"fmt"
	"github.com/opencost/opencost/core/pkg/log"
	"github.com/opencost/opencost/core/pkg/util"
	"github.com/prometheus/common/model"
)

func (pds *PrometheusMetricsQuerier) decodeContainerMetricsResult(qr interface{}) (ContainerMetricsResult, error) {
	data, ok := qr.(model.Vector)
	if !ok {
		return nil, fmt.Errorf("failed to cast query result to model.Vector")
	}

	results := make(ContainerMetricsResult)

	for _, s := range data {
		metric := s.Metric

		// Old prometheus-adapter code created a pod_name label, but the new one is just pod.
		// We will support both for now.
		podName, ok := metric["pod_name"]
		if !ok {
			podName = metric["pod"]
		}

		ns, ok := metric["namespace"]
		if !ok {
			log.Warnf("Prometheus Result Missing Namespace: %+v", metric)
		}

		// containerName is the name of the container requested by the user, e.g. "cost-model"
		containerName, ok := metric["container_name"]
		if !ok {
			containerName = metric["container"]
			if !ok {
				log.Warnf("Prometheus Result Missing Container Name: %+v", metric)
			}
		}

		// node is the name of the node that the container is running on, e.g. "gke-cluster-1-default-pool-6a-2579"
		node, ok := metric["node"]
		if !ok {
			log.Debugf("Prometheus Result Missing Node: %+v", metric)
		}

		clusterID, ok := metric[model.LabelName(pds.promConfig.ClusterLabel)]
		if !ok {
			clusterID = ""
		}

		key := fmt.Sprintf("%s/%s/%s/%s/%s", clusterID, ns, podName, containerName, node)
		results[key] = append(results[key], &util.Vector{
			Timestamp: float64(s.Timestamp),
			Value:     float64(s.Value),
		})
	}

	return results, nil
}
