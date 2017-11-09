package validator

import (
	"fmt"

	tapi "github.com/k8sdb/apimachinery/apis/kubedb/v1alpha1"
	"github.com/k8sdb/apimachinery/pkg/docker"
	amv "github.com/k8sdb/apimachinery/pkg/validator"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// TODO: Change method name. ValidateMemcached -> Validate<--->
func ValidateMemcached(client kubernetes.Interface, memcached *tapi.Memcached) error {
	if memcached.Spec.Version == "" {
		return fmt.Errorf(`Object 'Version' is missing in '%v'`, memcached.Spec)
	}

	// Set Database Image version
	version := memcached.Spec.Version
	// TODO: docker.ImageMemcached should hold correct image name
	if err := docker.CheckDockerImageVersion(docker.ImageMemcached, version); err != nil {
		return fmt.Errorf(`Image %v:%v not found`, docker.ImageMemcached, version)
	}

	if memcached.Spec.Storage != nil {
		var err error
		if err = amv.ValidateStorage(client, memcached.Spec.Storage); err != nil {
			return err
		}
	}

	// ---> Start
	// TODO: Use following if database needs/supports authentication secret
	// otherwise, delete
	databaseSecret := memcached.Spec.DatabaseSecret
	if databaseSecret != nil {
		if _, err := client.CoreV1().Secrets(memcached.Namespace).Get(databaseSecret.SecretName, metav1.GetOptions{}); err != nil {
			return err
		}
	}
	// ---> End

	backupScheduleSpec := memcached.Spec.BackupSchedule
	if backupScheduleSpec != nil {
		if err := amv.ValidateBackupSchedule(client, backupScheduleSpec, memcached.Namespace); err != nil {
			return err
		}
	}

	monitorSpec := memcached.Spec.Monitor
	if monitorSpec != nil {
		if err := amv.ValidateMonitorSpec(monitorSpec); err != nil {
			return err
		}

	}
	return nil
}
