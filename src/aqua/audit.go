package aqua

// Runtime represents a detect/block message
type Runtime struct {
	Action        string `json:"action"`
	Category      string `json:"category"`
	Container     string `json:"container"`
	Containerid   string `json:"containerid"`
	Control       string `json:"control"`
	Host          string `json:"host"`
	Hostgroup     string `json:"hostgroup"`
	Hostid        string `json:"hostid"`
	Hostip        string `json:"hostip"`
	Image         string `json:"image"`
	Imageid       string `json:"imageid"`
	K8SCluster    string `json:"k8s_cluster"`
	Level         string `json:"level"`
	Poddeployment string `json:"poddeployment"`
	Podname       string `json:"podname"`
	Podnamespace  string `json:"podnamespace"`
	Podtype       string `json:"podtype"`
	Reason        string `json:"reason"`
	Result        int    `json:"result"`
	Rule          string `json:"rule"`
	RuleType      string `json:"rule_type"`
	Time          int    `json:"time"`
	VMGroup       string `json:"vm_group"`
	VMID          string `json:"vm_id"`
	VMLocation    string `json:"vm_location"`
	VMName        string `json:"vm_name"`
}
