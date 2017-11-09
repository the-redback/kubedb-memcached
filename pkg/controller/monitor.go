package controller

import (
	"fmt"

	tapi "github.com/k8sdb/apimachinery/apis/kubedb/v1alpha1"
	"github.com/k8sdb/apimachinery/pkg/monitor"
)

func (c *Controller) newMonitorController(memcached *tapi.Memcached) (monitor.Monitor, error) {
	monitorSpec := memcached.Spec.Monitor

	if monitorSpec == nil {
		return nil, fmt.Errorf("MonitorSpec not found in %v", memcached.Spec)
	}

	if monitorSpec.Prometheus != nil {
		return monitor.NewPrometheusController(c.Client, c.ApiExtKubeClient, c.promClient, c.opt.OperatorNamespace), nil
	}

	return nil, fmt.Errorf("Monitoring controller not found for %v", monitorSpec)
}

func (c *Controller) addMonitor(memcached *tapi.Memcached) error {
	ctrl, err := c.newMonitorController(memcached)
	if err != nil {
		return err
	}
	return ctrl.AddMonitor(memcached.ObjectMeta, memcached.Spec.Monitor)
}

func (c *Controller) deleteMonitor(memcached *tapi.Memcached) error {
	ctrl, err := c.newMonitorController(memcached)
	if err != nil {
		return err
	}
	return ctrl.DeleteMonitor(memcached.ObjectMeta, memcached.Spec.Monitor)
}

func (c *Controller) updateMonitor(oldMemcached, updatedMemcached *tapi.Memcached) error {
	var err error
	var ctrl monitor.Monitor
	if updatedMemcached.Spec.Monitor == nil {
		ctrl, err = c.newMonitorController(oldMemcached)
	} else {
		ctrl, err = c.newMonitorController(updatedMemcached)
	}
	if err != nil {
		return err
	}
	return ctrl.UpdateMonitor(updatedMemcached.ObjectMeta, oldMemcached.Spec.Monitor, updatedMemcached.Spec.Monitor)
}
