package vault

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/denisaguilar/dummy_automation/restclient"
)

type SysHealthResponse struct {
	Initialized         bool
	Sealed              bool
	Standby             bool
	PerformanceStandby  bool   `json:"performance_standby"`
	ReplicationPerfMode string `json:"replication_perf_mode"`
	ReplicationDrMode   string `json:"replication_dr_mode"`
	ServerTimeUtc       int    `json:"server_time_utc"`
	Version             string
	ClusterName         string `json:"cluster_name"`
	ClusterID           string `json:"cluster_id"`
}

func SysHealth() (*SysHealthResponse, error) {
	resp, err := restclient.Get("http://192.168.31.235:8200/v1/sys/health", nil)
	if err == nil {
		log.Println("Error when calling api")
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error when reading the body")
	}
	s := SysHealthResponse{}
	if err := json.Unmarshal(bodyBytes, &s); err != nil {
		log.Println("Failed to parse response")
	}
	return &s, nil
}
