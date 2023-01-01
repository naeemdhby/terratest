package types

type ResourceGroups struct {
	Variables struct {
		ResourceGroups struct {
			Value map[string]struct {
				Name         string `json:"name"`
				Region       string `json:"region"`
				Subscription string `json:"subscription"`
			} `json:"value"`
		} `json:"resource_groups"`
	} `json:"variables"`
}

type VirtualMachines struct {
	Variables struct {
		VirtualMachines struct {
			Value map[string]struct {
				Name              string `json:"name"`
				ResourceGroupName string `json:"resource_group_name"`
			} `json:"value"`
		} `json:"virtual_machines"`
	} `json:"variables"`
}
