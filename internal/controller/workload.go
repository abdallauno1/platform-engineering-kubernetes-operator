package controller

import "reflect"

type Workload struct {
	Name      string
	Namespace string
	Image     string
	Replicas  int
	Env       map[string]string
	Labels    map[string]string
}

func (w Workload) Key() string {
	return w.Namespace + "/" + w.Name
}

func (w Workload) Equal(other Workload) bool {
	return w.Name == other.Name &&
		w.Namespace == other.Namespace &&
		w.Image == other.Image &&
		w.Replicas == other.Replicas &&
		reflect.DeepEqual(w.Env, other.Env) &&
		reflect.DeepEqual(w.Labels, other.Labels)
}
