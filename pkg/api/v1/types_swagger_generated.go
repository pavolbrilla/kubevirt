// Automatically generated by swagger-doc. DO NOT EDIT!

package v1

func (VirtualMachine) SwaggerDoc() map[string]string {
	return map[string]string{
		"":       "VirtualMachine is *the* VM Definition. It represents a virtual machine in the runtime environment of kubernetes.",
		"spec":   "VM Spec contains the VM specification.",
		"status": "Status is the high level overview of how the VM is doing. It contains information available to controllers and users.",
	}
}

func (VirtualMachineList) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "VirtualMachineList is a list of VirtualMachines",
	}
}

func (VirtualMachineSpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                              "VirtualMachineSpec is a description of a VirtualMachine.",
		"domain":                        "Specification of the desired behavior of the VirtualMachine on the host.",
		"nodeSelector":                  "NodeSelector is a selector which must be true for the vm to fit on a node.\nSelector which must match a node's labels for the vm to be scheduled on that node.\nMore info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/\n+optional",
		"affinity":                      "If affinity is specifies, obey all the affinity rules",
		"terminationGracePeriodSeconds": "Grace period observed after signalling a VM to stop after which the VM is force terminated.",
		"volumes":                       "List of volumes that can be mounted by disks belonging to the vm.",
		"hostname":                      "Specifies the hostname of the vm\nIf not specified, the hostname will be set to the name of the vm, if dhcp or cloud-init is configured properly.\n+optional",
		"subdomain":                     "If specified, the fully qualified vm hostname will be \"<hostname>.<subdomain>.<pod namespace>.svc.<cluster domain>\".\nIf not specified, the vm will not have a domainname at all. The DNS entry will resolve to the vm,\nno matter if the vm itself can pick up a hostname.\n+optional",
	}
}

func (Affinity) SwaggerDoc() map[string]string {
	return map[string]string{
		"":             "Affinity groups all the affinity rules related to a VM",
		"nodeAffinity": "Node affinity support",
	}
}

func (VirtualMachineStatus) SwaggerDoc() map[string]string {
	return map[string]string{
		"":           "VirtualMachineStatus represents information about the status of a VM. Status may trail the actual\nstate of a system.",
		"nodeName":   "NodeName is the name where the VM is currently running.",
		"conditions": "Conditions are specific points in VM's pod runtime.",
		"phase":      "Phase is the status of the VM in kubernetes world. It is not the VM status, but partially correlates to it.",
		"interfaces": "Interfaces represent the details of available network interfaces.",
	}
}

func (VirtualMachineCondition) SwaggerDoc() map[string]string {
	return map[string]string{}
}

func (VirtualMachineNetworkInterface) SwaggerDoc() map[string]string {
	return map[string]string{
		"ipAddress": "IP address of a Virtual Machine interface",
		"mac":       "Hardware address of a Virtual Machine interface",
	}
}

func (VMSelector) SwaggerDoc() map[string]string {
	return map[string]string{
		"name": "Name of the VM to migrate",
	}
}

func (VirtualMachineReplicaSet) SwaggerDoc() map[string]string {
	return map[string]string{
		"":       "VM is *the* VM Definition. It represents a virtual machine in the runtime environment of kubernetes.",
		"spec":   "VM Spec contains the VM specification.",
		"status": "Status is the high level overview of how the VM is doing. It contains information available to controllers and users.",
	}
}

func (VirtualMachineReplicaSetList) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "VMList is a list of VMs",
	}
}

func (VMReplicaSetSpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"replicas": "Number of desired pods. This is a pointer to distinguish between explicit\nzero and not specified. Defaults to 1.\n+optional",
		"selector": "Label selector for pods. Existing ReplicaSets whose pods are\nselected by this will be the ones affected by this deployment.",
		"template": "Template describes the pods that will be created.",
		"paused":   "Indicates that the replica set is paused.\n+optional",
	}
}

func (VMReplicaSetStatus) SwaggerDoc() map[string]string {
	return map[string]string{
		"replicas":      "Total number of non-terminated pods targeted by this deployment (their labels match the selector).\n+optional",
		"readyReplicas": "The number of ready replicas for this replica set.\n+optional",
	}
}

func (VMReplicaSetCondition) SwaggerDoc() map[string]string {
	return map[string]string{}
}

func (VMTemplateSpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"spec": "VM Spec contains the VM specification.",
	}
}

func (VirtualMachinePreset) SwaggerDoc() map[string]string {
	return map[string]string{
		"spec": "VM Spec contains the VM specification.",
	}
}

func (VirtualMachinePresetList) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "VirtualMachinePresetList is a list of VirtualMachinePresets",
	}
}

func (VirtualMachinePresetSpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"selector": "Selector is a label query over a set of VMs.\nRequired.",
		"domain":   "Domain is the same object type as contained in VirtualMachineSpec",
	}
}

func (OfflineVirtualMachine) SwaggerDoc() map[string]string {
	return map[string]string{
		"":       "OfflineVirtualMachine handles the VirtualMachines that are not running\nor are in a stopped state\nThe OfflineVirtualMachine contains the template to create the\nVirtualMachine. It also mirrors the running state of the created\nVirtualMachine in its status.",
		"spec":   "Spec contains the specification of VirtualMachine created",
		"status": "Status holds the current state of the controller and brief information\nabout its associated VirtualMachine",
	}
}

func (OfflineVirtualMachineList) SwaggerDoc() map[string]string {
	return map[string]string{
		"":      "OfflineVirtualMachineList is a list of offlinevirtualmachines",
		"items": "Items is a list of OfflineVirtualMachines",
	}
}

func (OfflineVirtualMachineSpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"":         "OfflineVirtualMachineSpec describes how the proper OfflineVirtualMachine\nshould look like",
		"running":  "Running controls whether the associatied VirtualMachine is created or not",
		"template": "Template is the direct specification of VirtualMachine",
	}
}

func (OfflineVirtualMachineStatus) SwaggerDoc() map[string]string {
	return map[string]string{
		"":           "OfflineVirtualMachineStatus represents the status returned by the\ncontroller to describe how the OfflineVirtualMachine is doing",
		"created":    "Created indicates if the virtual machine is created in the cluster",
		"ready":      "Ready indicates if the virtual machine is running and ready",
		"conditions": "Hold the state information of the OfflineVirtualMachine and its VirtualMachine",
	}
}

func (OfflineVirtualMachineCondition) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "OfflineVirtualMachineCondition represents the state of OfflineVirtualMachine",
	}
}
