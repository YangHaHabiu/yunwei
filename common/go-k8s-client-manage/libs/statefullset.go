package libs

import (
	"context"
	"time"

	"fmt"
	"ywadmin-v3/common/go-k8s-client-manage/comm"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

type StatefullsetMsg struct {
	*ClientSet
}

func (t *StatefullsetMsg) Query() error {
	var tips string
	deploymentName := t.ClientSet.ServiceName
	nameSpace := t.ClientSet.NameSpace
	clientset := t.ClientSet.Clientset
	fmt.Println("====================")
	fmt.Println("查询 " + deploymentName + " " + t.ClientSet.Works + "状态：")
	d1, err := clientset.AppsV1().StatefulSets(nameSpace).Get(context.TODO(), deploymentName, metav1.GetOptions{})
	fmt.Println("副本数量：", d1.Status.Replicas)
	if err != nil {
		return err
	}
	success, _, err := comm.GetStatefullsetStatus(*clientset, *d1, nameSpace)
	if success {
		tips = "Running"
	} else {
		tips = "Error"
	}
	fmt.Println(deploymentName + " " + t.ClientSet.Works + "状态：" + tips)
	fmt.Println("====================")
	p, err := comm.PodsGetWithDeploymentNameAndNS(clientset, nameSpace, deploymentName)
	if err != nil {
		return err
	}
	fmt.Println("查询 " + deploymentName + " pods状态：")
	fmt.Println("名称\t\t状态")
	for _, v := range p {
		s, _ := comm.GetPodObjectByName(clientset, v.Name, nameSpace)
		fmt.Println(v.Name, comm.GetPodStatus(s))
	}

	return nil
}

func (t *StatefullsetMsg) Create() error {

	return nil
}
func (t *StatefullsetMsg) Update() error {

	return nil
}

func (t *StatefullsetMsg) Delete() error {

	return nil
}

func (t *StatefullsetMsg) Redeployment() error {
	deploymentName := t.ClientSet.ServiceName
	nameSpace := t.ClientSet.NameSpace
	clientset := t.ClientSet.Clientset
	fmt.Println("开始重新部署" + deploymentName + "服务")
	data := fmt.Sprintf(`{"spec": {"template": {"metadata": {"annotations": {"kubectl.kubernetes.io/restartedAt": "%s"}}}}}`, time.Now().Format("20060102150405"))
	_, err := clientset.AppsV1().StatefulSets(nameSpace).Patch(context.TODO(), deploymentName, types.StrategicMergePatchType, []byte(data), v1.PatchOptions{})
	if err != nil {
		return err
	}
	var tips string
	d1, err := clientset.AppsV1().StatefulSets(nameSpace).Get(context.TODO(), deploymentName, metav1.GetOptions{})
	if err != nil {
		return err
	}
	for {
		success, _, err := comm.GetStatefullsetStatus(*clientset, *d1, nameSpace)
		if err != nil {
			return err
		}
		if success {
			break
		}
		tips += "."
		fmt.Println("重新部署中" + tips)
		time.Sleep(2 * time.Second)
	}
	fmt.Println("重新部署" + deploymentName + "服务成功")
	return nil
}
