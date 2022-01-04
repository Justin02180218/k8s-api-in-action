package resources

type DeploymentResource struct {
	Namespace   string
	Name        string
	Replicas    int32
	Matchlabels map[string]string
	Labels      map[string]string
	Containers  []Container
}

type Container struct {
	Name  string
	Image string
	Ports []Port
}

type Port struct {
	Name     string
	Port     int32
	NodePort int32
}

type ServiceResource struct {
	Namespace string
	Name      string
	Ports     []Port
	Selector  map[string]string
}
