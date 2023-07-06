package tencentLibs

type RequestModel struct {
	SecretId  string
	SecretKey string
	EksCiId   string
}

//解析kubeconfigJson字段

type KubeConfigResponse struct {
	Response struct {
		Kubeconfig string `json:"Kubeconfig"`
		RequestId  string `json:"RequestId"`
	} `json:"Response"`
}
