package comm

import (
	"context"
	"fmt"
	"strings"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var (
	Actions     string
	Works       string
	FileConfig  string
	YmalFile    string
	NameSpace   string
	ServiceName string
)

const (
	ContainersReady string = "ContainersReady"
	PodInitialized  string = "Initialized"
	PodReady        string = "Ready"
	PodScheduled    string = "PodScheduled"
)

const (
	ConditionTrue    string = "True"
	ConditionFalse   string = "False"
	ConditionUnknown string = "Unknown"
)

func Help() {
	tips := fmt.Sprintf(`当前k8s操作版本：%s
-a 常规操作：create query update delete redeployment
-f 认证文件：kubeconfig
-w 工作负载：deployment statefullset
-n 命名空间：namespace
-s 服务名：servicename
-y yaml文件：只用于创建和更新部署使用
`, Version)
	fmt.Println(tips)
}

// 根据命名空间和deployment名称，从k8s处获取deployment拥有的pod列表
func PodsGetWithDeploymentNameAndNS(clientset *kubernetes.Clientset, namespace,
	deploymentName string) ([]corev1.Pod, error) {
	var label string
	if deploymentName == "nginx" || deploymentName == "nginx-state" {
		label = "k8s-app=" + deploymentName
	} else {
		label = "app=" + deploymentName
	}
	podList, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{LabelSelector: label})
	if err != nil {
		return nil, err
	}

	return podList.Items, nil
}

//根据pod查状态
func GetPodStatus(pod *corev1.Pod) string {
	for _, cond := range pod.Status.Conditions {
		if string(cond.Type) == ContainersReady {
			if string(cond.Status) != ConditionTrue {
				return "Unavailable"
			}
		} else if string(cond.Type) == PodInitialized && string(cond.Status) != ConditionTrue {
			return "Initializing"
		} else if string(cond.Type) == PodReady {
			if string(cond.Status) != ConditionTrue {
				return "Unavailable"
			}
			for _, containerState := range pod.Status.ContainerStatuses {
				if !containerState.Ready {
					return "Unavailable"
				}
			}
		} else if string(cond.Type) == PodScheduled && string(cond.Status) != ConditionTrue {
			return "Scheduling"
		}
	}
	return string(pod.Status.Phase)
}

//根据pod名称及命名空间返回pod结构体
func GetPodObjectByName(clientset *kubernetes.Clientset, podName, namespace string) (*corev1.Pod, error) {
	return clientset.CoreV1().Pods(namespace).Get(context.TODO(), podName, metav1.GetOptions{})
}

func GetDeploymentStatus(clientset kubernetes.Clientset, deployment appsv1.Deployment, namespace string) (success bool, reasons []string, err error) {
	if namespace == "" {
		namespace = "default"
	}

	k8sDeployment, err := clientset.AppsV1().Deployments(namespace).Get(context.TODO(), deployment.Name, metav1.GetOptions{})
	if err != nil {
		return false, []string{"get deployments status error"}, err
	}

	// 获取pod的状态
	labelSelector := ""
	for key, value := range deployment.Spec.Selector.MatchLabels {
		labelSelector = labelSelector + key + "=" + value + ","
	}
	labelSelector = strings.TrimRight(labelSelector, ",")

	podList, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{LabelSelector: labelSelector})
	if err != nil {
		return false, []string{"get pods status error"}, err
	}

	readyPod := 0
	unavailablePod := 0
	waitingReasons := []string{}
	for _, pod := range podList.Items {
		// 记录等待原因
		for _, containerStatus := range pod.Status.ContainerStatuses {
			if containerStatus.State.Waiting != nil {
				reason := "pod " + pod.Name + ", container " + containerStatus.Name + ", waiting reason: " + containerStatus.State.Waiting.Reason
				waitingReasons = append(waitingReasons, reason)
			}
		}

		podScheduledCondition := GetPodCondition(pod.Status, corev1.PodScheduled)
		initializedCondition := GetPodCondition(pod.Status, corev1.PodInitialized)
		readyCondition := GetPodCondition(pod.Status, corev1.PodReady)
		containersReadyCondition := GetPodCondition(pod.Status, corev1.ContainersReady)

		if pod.Status.Phase == "Running" &&
			podScheduledCondition.Status == "True" &&
			initializedCondition.Status == "True" &&
			readyCondition.Status == "True" &&
			containersReadyCondition.Status == "True" {
			readyPod++
		} else {
			unavailablePod++
		}
	}

	// 根据container状态判定
	if len(waitingReasons) != 0 {
		return false, waitingReasons, nil
	}

	// 根据pod状态判定
	if int32(readyPod) < *(k8sDeployment.Spec.Replicas) ||
		int32(unavailablePod) != 0 {
		return false, []string{"pods not ready!"}, nil
	}

	// deployment进行状态判定

	availableCondition := GetDeploymentCondition(k8sDeployment.Status, appsv1.DeploymentAvailable)
	progressingCondition := GetDeploymentCondition(k8sDeployment.Status, appsv1.DeploymentProgressing)

	if k8sDeployment.Status.UpdatedReplicas != *(k8sDeployment.Spec.Replicas) ||
		k8sDeployment.Status.Replicas != *(k8sDeployment.Spec.Replicas) ||
		k8sDeployment.Status.AvailableReplicas != *(k8sDeployment.Spec.Replicas) ||
		availableCondition.Status != "True" ||
		progressingCondition.Status != "True" {
		return false, []string{"deployments not ready!"}, nil
	}

	if k8sDeployment.Status.ObservedGeneration < k8sDeployment.Generation {
		return false, []string{"observed generation less than generation!"}, nil
	}

	// 发布成功
	return true, []string{}, nil
}

// GetDeploymentCondition returns the condition with the provided type.
func GetDeploymentCondition(status appsv1.DeploymentStatus, condType appsv1.DeploymentConditionType) *appsv1.DeploymentCondition {
	for i := range status.Conditions {
		c := status.Conditions[i]
		if c.Type == condType {
			return &c
		}
	}
	return nil
}
func GetPodCondition(status corev1.PodStatus, condType corev1.PodConditionType) *corev1.PodCondition {
	for i := range status.Conditions {
		c := status.Conditions[i]
		if c.Type == condType {
			return &c
		}
	}
	return nil
}

func GetStatefullsetCondition(status appsv1.StatefulSetStatus, condType appsv1.StatefulSetConditionType) *appsv1.StatefulSetCondition {
	fmt.Println(status.Conditions)
	for i := range status.Conditions {
		c := status.Conditions[i]
		if c.Type == condType {
			return &c
		}
	}
	return nil
}

func GetStatefullsetStatus(clientset kubernetes.Clientset, deployment appsv1.StatefulSet, namespace string) (success bool, reasons []string, err error) {
	if namespace == "" {
		namespace = "default"
	}

	k8sDeployment, err := clientset.AppsV1().StatefulSets(namespace).Get(context.TODO(), deployment.Name, metav1.GetOptions{})
	if err != nil {
		return false, []string{"get deployments status error"}, err
	}

	// 获取pod的状态
	labelSelector := ""
	for key, value := range deployment.Spec.Selector.MatchLabels {
		labelSelector = labelSelector + key + "=" + value + ","
	}
	labelSelector = strings.TrimRight(labelSelector, ",")

	podList, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{LabelSelector: labelSelector})
	if err != nil {
		return false, []string{"get pods status error"}, err
	}

	readyPod := 0
	unavailablePod := 0
	waitingReasons := []string{}
	for _, pod := range podList.Items {
		// 记录等待原因
		for _, containerStatus := range pod.Status.ContainerStatuses {
			if containerStatus.State.Waiting != nil {
				reason := "pod " + pod.Name + ", container " + containerStatus.Name + ", waiting reason: " + containerStatus.State.Waiting.Reason
				waitingReasons = append(waitingReasons, reason)
			}
		}

		podScheduledCondition := GetPodCondition(pod.Status, corev1.PodScheduled)
		initializedCondition := GetPodCondition(pod.Status, corev1.PodInitialized)
		readyCondition := GetPodCondition(pod.Status, corev1.PodReady)
		containersReadyCondition := GetPodCondition(pod.Status, corev1.ContainersReady)

		if pod.Status.Phase == "Running" &&
			podScheduledCondition.Status == "True" &&
			initializedCondition.Status == "True" &&
			readyCondition.Status == "True" &&
			containersReadyCondition.Status == "True" {
			readyPod++
		} else {
			unavailablePod++
		}
	}

	// 根据container状态判定
	if len(waitingReasons) != 0 {
		return false, waitingReasons, nil
	}

	// 根据pod状态判定
	if int32(readyPod) < *(k8sDeployment.Spec.Replicas) ||
		int32(unavailablePod) != 0 {
		return false, []string{"pods not ready!"}, nil
	}
	//fmt.Println(k8sDeployment.Status)
	// deployment进行状态判定
	//GetStatefullsetCondition(k8sDeployment.Status, "1111")

	// availableCondition := GetStatefullsetCondition(k8sDeployment.Status, appsv1.DeploymentAvailable)
	// progressingCondition := GetStatefullsetCondition(k8sDeployment.Status, appsv1.DeploymentProgressing)

	// if k8sDeployment.Status.UpdatedReplicas != *(k8sDeployment.Spec.Replicas) ||
	// 	k8sDeployment.Status.Replicas != *(k8sDeployment.Spec.Replicas) ||
	// 	k8sDeployment.Status.AvailableReplicas != *(k8sDeployment.Spec.Replicas) ||
	// 	availableCondition.Status != "True" ||
	// 	progressingCondition.Status != "True" {
	// 	return false, []string{"deployments not ready!"}, nil
	// }
	if k8sDeployment.Status.UpdatedReplicas != *(k8sDeployment.Spec.Replicas) ||
		k8sDeployment.Status.Replicas != *(k8sDeployment.Spec.Replicas) ||
		k8sDeployment.Status.AvailableReplicas != *(k8sDeployment.Spec.Replicas) {
		return false, []string{"deployments not ready!"}, nil
	}

	if k8sDeployment.Status.ObservedGeneration < k8sDeployment.Generation {
		return false, []string{"observed generation less than generation!"}, nil
	}

	// 发布成功
	return true, []string{}, nil
}
