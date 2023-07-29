package backend

import (
	"context"
	"fmt"
	"github.com/brunowang/easypaas/pbgen/pod"
	"github.com/brunowang/gframe/gflog"
	"go.uber.org/zap"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"strconv"
)

type K8sOper interface {
	CreateToK8s(context.Context, *pod.PodInfo) error
	DeleteFromK8s(context.Context, *pod.PodInfo) error
	UpdateToK8s(context.Context, *pod.PodInfo) error
}

type K8sOperImpl struct {
	K8sClientSet *kubernetes.Clientset
	deployment   *appsv1.Deployment
}

func NewK8sOper() *K8sOperImpl {
	// use in the k8s cluster
	config, err := rest.InClusterConfig()
	if err != nil {
		gflog.Fatal(context.Background(), "get k8s config failed", zap.Error(err))
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		gflog.Fatal(context.Background(), "init k8s client failed", zap.Error(err))
	}
	return &K8sOperImpl{
		K8sClientSet: clientset,
		deployment:   &appsv1.Deployment{},
	}
}

func (s *K8sOperImpl) CreateToK8s(ctx context.Context, podInfo *pod.PodInfo) (err error) {
	s.SetDeployment(podInfo)
	if _, err = s.K8sClientSet.AppsV1().Deployments(podInfo.PodNamespace).Get(ctx, podInfo.PodName, metav1.GetOptions{}); err != nil {
		if _, err = s.K8sClientSet.AppsV1().Deployments(podInfo.PodNamespace).Create(ctx, s.deployment, metav1.CreateOptions{}); err != nil {
			gflog.Error(ctx, "create to k8s failed", zap.Error(err))
			return err
		}
		gflog.Info(ctx, "create to k8s success", zap.String("pod", podInfo.PodName))
		return nil
	} else {
		gflog.Error(ctx, fmt.Sprintf("pod %s already exists, cannot be created", podInfo.PodName))
		return fmt.Errorf("pod %s already exists, cannot be created", podInfo.PodName)
	}

}

func (s *K8sOperImpl) UpdateToK8s(ctx context.Context, podInfo *pod.PodInfo) (err error) {
	s.SetDeployment(podInfo)
	if _, err = s.K8sClientSet.AppsV1().Deployments(podInfo.PodNamespace).Get(ctx, podInfo.PodName, metav1.GetOptions{}); err != nil {
		gflog.Error(ctx, fmt.Sprintf("pod %s not exists, cannot be updated", podInfo.PodName))
		return fmt.Errorf("pod %s not exists, cannot be update", podInfo.PodName)
	} else {
		if _, err = s.K8sClientSet.AppsV1().Deployments(podInfo.PodNamespace).Update(ctx, s.deployment, metav1.UpdateOptions{}); err != nil {
			gflog.Error(ctx, "update to k8s failed", zap.Error(err))
			return err
		}
		gflog.Info(ctx, "update to k8s success", zap.String("pod", podInfo.PodName))
		return nil
	}

}

func (s *K8sOperImpl) DeleteFromK8s(ctx context.Context, podInfo *pod.PodInfo) (err error) {
	if err = s.K8sClientSet.AppsV1().Deployments(podInfo.PodNamespace).Delete(ctx, podInfo.PodName, metav1.DeleteOptions{}); err != nil {
		gflog.Error(ctx, "delete pod from k8s failed", zap.Error(err))
		return err
	}
	gflog.Info(ctx, "delete pod from k8s success", zap.String("pod", podInfo.PodName))
	return
}

func (s *K8sOperImpl) SetDeployment(podInfo *pod.PodInfo) {
	deployment := &appsv1.Deployment{}
	deployment.TypeMeta = metav1.TypeMeta{
		Kind:       "deployment",
		APIVersion: "v1",
	}
	deployment.ObjectMeta = metav1.ObjectMeta{
		Name:      podInfo.PodName,
		Namespace: podInfo.PodNamespace,
		Labels: map[string]string{
			"app-name": podInfo.PodName,
			"author":   "Robot",
		},
	}
	deployment.Name = podInfo.PodName
	deployment.Spec = appsv1.DeploymentSpec{
		//副本个数
		Replicas: &podInfo.PodReplicas,
		Selector: &metav1.LabelSelector{
			MatchLabels: map[string]string{
				"app-name": podInfo.PodName,
			},
			MatchExpressions: nil,
		},
		Template: corev1.PodTemplateSpec{
			ObjectMeta: metav1.ObjectMeta{
				Labels: map[string]string{
					"app-name": podInfo.PodName,
				},
			},
			Spec: corev1.PodSpec{
				Containers: []corev1.Container{
					{
						Name:            podInfo.PodName,
						Image:           podInfo.PodImage,
						Ports:           s.getContainerPort(podInfo),
						Env:             s.getEnv(podInfo),
						Resources:       s.getResources(podInfo),
						ImagePullPolicy: s.getImagePullPolicy(podInfo),
					},
				},
			},
		},
		Strategy:                appsv1.DeploymentStrategy{},
		MinReadySeconds:         0,
		RevisionHistoryLimit:    nil,
		Paused:                  false,
		ProgressDeadlineSeconds: nil,
	}
	s.deployment = deployment
}

func (s *K8sOperImpl) getContainerPort(podInfo *pod.PodInfo) (containerPort []corev1.ContainerPort) {
	for _, v := range podInfo.PodPort {
		containerPort = append(containerPort, corev1.ContainerPort{
			Name:          "port-" + strconv.FormatInt(int64(v.ContainerPort), 10),
			ContainerPort: v.ContainerPort,
			Protocol:      s.getProtocol(v.Protocol),
		})
	}
	return
}

func (s *K8sOperImpl) getProtocol(protocol string) corev1.Protocol {
	switch protocol {
	case "TCP":
		return "TCP"
	case "UDP":
		return "UDP"
	case "SCTP":
		return "SCTP"
	default:
		return "TCP"
	}
}

func (s *K8sOperImpl) getEnv(podInfo *pod.PodInfo) (envVar []corev1.EnvVar) {
	for _, v := range podInfo.PodEnv {
		envVar = append(envVar, corev1.EnvVar{
			Name:      v.EnvKey,
			Value:     v.EnvValue,
			ValueFrom: nil,
		})
	}
	return
}

func (s *K8sOperImpl) getResources(podInfo *pod.PodInfo) (source corev1.ResourceRequirements) {
	source.Limits = corev1.ResourceList{
		"cpu":    resource.MustParse(strconv.FormatFloat(float64(podInfo.PodCpuMax), 'f', 6, 64)),
		"memory": resource.MustParse(strconv.FormatFloat(float64(podInfo.PodMemoryMax), 'f', 6, 64)),
	}
	source.Requests = corev1.ResourceList{
		"cpu":    resource.MustParse(strconv.FormatFloat(float64(podInfo.PodCpuMax), 'f', 6, 64)),
		"memory": resource.MustParse(strconv.FormatFloat(float64(podInfo.PodMemoryMax), 'f', 6, 64)),
	}
	return
}

func (s *K8sOperImpl) getImagePullPolicy(podInfo *pod.PodInfo) corev1.PullPolicy {
	switch podInfo.PodPullPolicy {
	case "Always":
		return "Always"
	case "Never":
		return "Never"
	case "IfNotPresent":
		return "IfNotPresent"
	default:
		return "Always"
	}
}
