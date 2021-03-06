package pod

import (
	log "github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/caicloud/cyclone/pkg/k8s/clientset"
	"github.com/caicloud/cyclone/pkg/workflow/common"
)

// IsGCPod judges whether a pod is a GC pod by check whether it has
// annotation "cyclone.io/gc".
func IsGCPod(pod *corev1.Pod) bool {
	if pod == nil || pod.Annotations == nil {
		return false
	}

	if _, ok := pod.Annotations[common.GCAnnotationName]; !ok {
		return false
	}

	return true
}

// GCPodUpdated handles GC pod update. If GC pod is terminated, it will be deleted.
func GCPodUpdated(client clientset.Interface, pod *corev1.Pod) {
	if pod.Status.Phase == corev1.PodSucceeded || pod.Status.Phase == corev1.PodFailed {
		if err := client.CoreV1().Pods(pod.Namespace).Delete(pod.Name, &metav1.DeleteOptions{}); err != nil {
			if errors.IsNotFound(err) {
				return
			}
			log.WithField("pod", pod.Name).Warn("Delete GC pod error: ", err)
		}
	}
}
