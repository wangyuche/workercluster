package k8s

import (
	"github.com/wangyuche/goutils/log"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func getKubernetesConfig(path string) (*rest.Config, error) {
	if path != "" {
		cfg, err := clientcmd.BuildConfigFromFlags("", path)
		if err != nil {
			return nil, err
		}
		return cfg, nil
	}
	cfg, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func ConnectKubernetes(path string) *clientset.Clientset {
	config, err := getKubernetesConfig(path)
	if err != nil {
		log.Fail(err.Error())
	}
	return clientset.NewForConfigOrDie(config)
}
