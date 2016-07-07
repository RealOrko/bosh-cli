package director

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)

type VMInfo struct {
	AgentID string `json:"agent_id"`

	Timestamp time.Time

	JobName   string `json:"job_name"`
	ID        string `json:"id"`
	Index     *int   `json:"index"`
	State     string `json:"job_state"` // e.g. "running"
	Bootstrap bool

	IPs []string `json:"ips"`
	DNS []string `json:"dns"`

	AZ           string `json:"az"`
	VMID         string `json:"vm_cid"`
	VMType       string `json:"vm_type"`
	ResourcePool string `json:"resource_pool"`
	DiskID       string `json:"disk_cid"`

	Processes []VMInfoProcess

	Vitals VMInfoVitals

	ResurrectionPaused bool `json:"resurrection_paused"`
}

type VMInfoProcess struct {
	Name  string
	State string // e.g. "running"

	CPU    VMInfoVitalsCPU `json:"cpu"`
	Mem    VMInfoVitalsMemIntSize
	Uptime VMInfoVitalsUptime
}

type VMInfoVitals struct {
	CPU    VMInfoVitalsCPU `json:"cpu"`
	Mem    VMInfoVitalsMemSize
	Swap   VMInfoVitalsMemSize
	Uptime VMInfoVitalsUptime

	Load []string
	Disk map[string]VMInfoVitalsDiskSize
}

func (v VMInfoVitals) SystemDisk() VMInfoVitalsDiskSize     { return v.Disk["system"] }
func (v VMInfoVitals) EphemeralDisk() VMInfoVitalsDiskSize  { return v.Disk["ephemeral"] }
func (v VMInfoVitals) PersistentDisk() VMInfoVitalsDiskSize { return v.Disk["persistent"] }

type VMInfoVitalsCPU struct {
	Total *float64 // used by VMInfoProcess
	Sys   string
	User  string
	Wait  string
}

type VMInfoVitalsDiskSize struct {
	InodePercent string `json:"inode_percent"`
	Percent      string
}

type VMInfoVitalsMemSize struct {
	KB      string `json:"kb"`
	Percent string
}

type VMInfoVitalsMemIntSize struct {
	KB      *uint64 `json:"kb"`
	Percent *float64
}

type VMInfoVitalsUptime struct {
	Seconds *uint64 `json:"secs"` // e.g. 48307
}

func (i VMInfo) IsRunning() bool {
	if i.State != "running" {
		return false
	}

	for _, p := range i.Processes {
		if !p.IsRunning() {
			return false
		}
	}

	return true
}

func (p VMInfoProcess) IsRunning() bool {
	return p.State == "running"
}

func (d DeploymentImpl) VMInfos() ([]VMInfo, error) {
	infos, err := d.client.DeploymentVMInfos(d.name)
	if err != nil {
		return nil, err
	}

	t := time.Now()

	for _, info := range infos {
		info.Timestamp = t
	}

	return infos, nil
}

func (c Client) DeploymentVMInfos(deploymentName string) ([]VMInfo, error) {
	if len(deploymentName) == 0 {
		return nil, bosherr.Error("Expected non-empty deployment name")
	}

	path := fmt.Sprintf("/deployments/%s/vms?format=full", deploymentName)

	_, resultBytes, err := c.taskClientRequest.GetResult(path)
	if err != nil {
		return nil, bosherr.WrapErrorf(
			err, "Listing deployment '%s' VMs infos", deploymentName)
	}

	var resps []VMInfo

	for _, piece := range strings.Split(string(resultBytes), "\n") {
		if len(piece) == 0 {
			continue
		}

		var resp VMInfo

		err := json.Unmarshal([]byte(piece), &resp)
		if err != nil {
			return nil, bosherr.WrapErrorf(
				err, "Unmarshaling VM info response: '%s'", string(piece))
		}

		resps = append(resps, resp)
	}

	return resps, nil
}