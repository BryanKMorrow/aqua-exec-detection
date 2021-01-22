package aqua

type Policy struct {
	Name               string `json:"name"`
	Author             string `json:"author"`
	Version            string `json:"version"`
	Lastupdate         int    `json:"lastupdate"`
	Enabled            bool   `json:"enabled"`
	Type               string `json:"type"`
	RuntimeType        string `json:"runtime_type"`
	Enforce            bool   `json:"enforce"`
	AllowedExecutables struct {
		Enabled              bool          `json:"enabled"`
		SeparateExecutables  bool          `json:"separate_executables"`
		AllowRootExecutables []interface{} `json:"allow_root_executables"`
		AllowExecutables     []interface{} `json:"allow_executables"`
	} `json:"allowed_executables"`
	ExecutableBlacklist struct {
		Enabled     bool          `json:"enabled"`
		Executables []interface{} `json:"executables"`
	} `json:"executable_blacklist"`
	DriftPrevention struct {
		Enabled       bool `json:"enabled"`
		ExecLockdown  bool `json:"exec_lockdown"`
		ImageLockdown bool `json:"image_lockdown"`
	} `json:"drift_prevention"`
	RestrictedVolumes struct {
		Enabled bool          `json:"enabled"`
		Volumes []interface{} `json:"volumes"`
	} `json:"restricted_volumes"`
	NoNewPrivileges bool `json:"no_new_privileges"`
	BypassScope     struct {
		Enabled bool `json:"enabled"`
		Scope   struct {
			Expression string        `json:"expression"`
			Variables  []interface{} `json:"variables"`
		} `json:"scope"`
	} `json:"bypass_scope"`
	LimitContainerPrivileges struct {
		Enabled               bool `json:"enabled"`
		Privileged            bool `json:"privileged"`
		Netmode               bool `json:"netmode"`
		Pidmode               bool `json:"pidmode"`
		Utsmode               bool `json:"utsmode"`
		Usermode              bool `json:"usermode"`
		Ipcmode               bool `json:"ipcmode"`
		PreventRootUser       bool `json:"prevent_root_user"`
		PreventLowPortBinding bool `json:"prevent_low_port_binding"`
		BlockAddCapabilities  bool `json:"block_add_capabilities"`
		UseHostUser           bool `json:"use_host_user"`
	} `json:"limit_container_privileges"`
	PreventOverrideDefaultConfig struct {
		Enabled         bool `json:"enabled"`
		EnforceSelinux  bool `json:"enforce_selinux"`
		EnforceSeccomp  bool `json:"enforce_seccomp"`
		EnforceApparmor bool `json:"enforce_apparmor"`
	} `json:"prevent_override_default_config"`
	Scope                      Scope    `json:"scope"`
	OnlyRegisteredImages       bool     `json:"only_registered_images"`
	BlockDisallowedImages      bool     `json:"block_disallowed_images"`
	BlockNonCompliantWorkloads bool     `json:"block_non_compliant_workloads"`
	Auditing                   Auditing `json:"auditing"`
	BlacklistedOsUsers         struct {
		Enabled        bool          `json:"enabled"`
		GroupBlackList []interface{} `json:"group_black_list"`
		UserBlackList  []interface{} `json:"user_black_list"`
	} `json:"blacklisted_os_users"`
	WhitelistedOsUsers struct {
		Enabled        bool          `json:"enabled"`
		UserWhiteList  []interface{} `json:"user_white_list"`
		GroupWhiteList []interface{} `json:"group_white_list"`
	} `json:"whitelisted_os_users"`
	DefaultSecurityProfile   string `json:"default_security_profile"`
	EnableForkGuard          bool   `json:"enable_fork_guard"`
	ForkGuardProcessLimit    int    `json:"fork_guard_process_limit"`
	EnableIPReputation       bool   `json:"enable_ip_reputation"`
	EnablePortScanProtection bool   `json:"enable_port_scan_protection"`
	BlockNwUnlinkCont        bool   `json:"block_nw_unlink_cont"`
	FileBlock                struct {
		Enabled           bool          `json:"enabled"`
		FilenameBlockList []interface{} `json:"filename_block_list"`
	} `json:"file_block"`
	PackageBlock struct {
		Enabled           bool          `json:"enabled"`
		PackagesBlackList []interface{} `json:"packages_black_list"`
	} `json:"package_block"`
	LinuxCapabilities struct {
		Enabled                 bool          `json:"enabled"`
		RemoveLinuxCapabilities []interface{} `json:"remove_linux_capabilities"`
	} `json:"linux_capabilities"`
	PortBlock struct {
		Enabled            bool          `json:"enabled"`
		BlockInboundPorts  []interface{} `json:"block_inbound_ports"`
		BlockOutboundPorts []interface{} `json:"block_outbound_ports"`
	} `json:"port_block"`
	Tripwire struct {
		Enabled       bool     `json:"enabled"`
		UserID        string   `json:"user_id"`
		UserPassword  string   `json:"user_password"`
		ApplyOn       []string `json:"apply_on"`
		ServerlessApp string   `json:"serverless_app"`
	} `json:"tripwire"`
	FileIntegrityMonitoring struct {
		Enabled                            bool          `json:"enabled"`
		MonitoredFiles                     []interface{} `json:"monitored_files"`
		ExceptionalMonitoredFiles          []interface{} `json:"exceptional_monitored_files"`
		MonitoredFilesProcesses            []interface{} `json:"monitored_files_processes"`
		ExceptionalMonitoredFilesProcesses []interface{} `json:"exceptional_monitored_files_processes"`
		MonitoredFilesUsers                []interface{} `json:"monitored_files_users"`
		ExceptionalMonitoredFilesUsers     []interface{} `json:"exceptional_monitored_files_users"`
		MonitoredFilesRead                 bool          `json:"monitored_files_read"`
		MonitoredFilesModify               bool          `json:"monitored_files_modify"`
		MonitoredFilesAttributes           bool          `json:"monitored_files_attributes"`
		MonitoredFilesCreate               bool          `json:"monitored_files_create"`
		MonitoredFilesDelete               bool          `json:"monitored_files_delete"`
	} `json:"file_integrity_monitoring"`
	RegistryAccessMonitoring struct {
		Enabled                               bool          `json:"enabled"`
		MonitoredRegistryPaths                []interface{} `json:"monitored_registry_paths"`
		ExceptionalMonitoredRegistryPaths     []interface{} `json:"exceptional_monitored_registry_paths"`
		MonitoredRegistryProcesses            []interface{} `json:"monitored_registry_processes"`
		ExceptionalMonitoredRegistryProcesses []interface{} `json:"exceptional_monitored_registry_processes"`
		MonitoredRegistryUsers                []interface{} `json:"monitored_registry_users"`
		ExceptionalMonitoredRegistryUsers     []interface{} `json:"exceptional_monitored_registry_users"`
		MonitoredRegistryCreate               bool          `json:"monitored_registry_create"`
		MonitoredRegistryRead                 bool          `json:"monitored_registry_read"`
		MonitoredRegistryModify               bool          `json:"monitored_registry_modify"`
		MonitoredRegistryDelete               bool          `json:"monitored_registry_delete"`
		MonitoredRegistryAttributes           bool          `json:"monitored_registry_attributes"`
	} `json:"registry_access_monitoring"`
	ReadonlyRegistry struct {
		Enabled                              bool          `json:"enabled"`
		ReadonlyRegistryPaths                []interface{} `json:"readonly_registry_paths"`
		ExceptionalReadonlyRegistryPaths     []interface{} `json:"exceptional_readonly_registry_paths"`
		ReadonlyRegistryProcesses            []interface{} `json:"readonly_registry_processes"`
		ExceptionalReadonlyRegistryProcesses []interface{} `json:"exceptional_readonly_registry_processes"`
		ReadonlyRegistryUsers                []interface{} `json:"readonly_registry_users"`
		ExceptionalReadonlyRegistryUsers     []interface{} `json:"exceptional_readonly_registry_users"`
	} `json:"readonly_registry"`
	SystemIntegrityProtection struct {
		Enabled                   bool `json:"enabled"`
		AuditSystemtimeChange     bool `json:"audit_systemtime_change"`
		WindowsServicesMonitoring bool `json:"windows_services_monitoring"`
	} `json:"system_integrity_protection"`
	BlockContainerExec bool     `json:"block_container_exec"`
	VulnID             int      `json:"vuln_id"`
	RepoID             int      `json:"repo_id"`
	HeuristicRefID     int      `json:"heuristic_ref_id"`
	ImageID            int      `json:"image_id"`
	IsAutoGenerated    bool     `json:"is_auto_generated"`
	ApplicationScopes  []string `json:"application_scopes"`
	AuditOnFailure     bool     `json:"audit_on_failure"`
	FailCicd           bool     `json:"fail_cicd"`
	BlockFailed        bool     `json:"block_failed"`
	EnforceAfterDays   int      `json:"enforce_after_days"`
	DomainName         string   `json:"domain_name"`
	Domain             string   `json:"domain"`
}

type Variable struct {
	Attribute string `json:"attribute"`
	Value     string `json:"value"`
}

type Scope struct {
	Expression string     `json:"expression"`
	Variables  []Variable `json:"variables"`
}

type Auditing struct {
	Enabled             bool `json:"enabled"`
	AuditAllProcesses   bool `json:"audit_all_processes"`
	AuditProcessCmdline bool `json:"audit_process_cmdline"`
	AuditAllNetwork     bool `json:"audit_all_network"`
	AuditOsUserActivity bool `json:"audit_os_user_activity"`
}
