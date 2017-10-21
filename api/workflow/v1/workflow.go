package v1

import (
	"encoding/json"
	"fmt"
	"hash/fnv"

	"k8s.io/apimachinery/pkg/runtime"
)

func (wf *Workflow) Completed() bool {
	status := wf.Status.Nodes[wf.NodeID(wf.ObjectMeta.Name)].Status
	return status == WorkflowStatusSuccess ||
		status == WorkflowStatusFailed ||
		status == WorkflowStatusCanceled
}

func (wf *Workflow) DeepCopyObject() runtime.Object {
	wfBytes, err := json.Marshal(wf)
	if err != nil {
		panic(err)
	}
	var copy Workflow
	err = json.Unmarshal(wfBytes, &copy)
	if err != nil {
		panic(err)
	}
	return &copy
}

func (wfl *WorkflowList) DeepCopyObject() runtime.Object {
	wflBytes, err := json.Marshal(wfl)
	if err != nil {
		panic(err)
	}
	var copy WorkflowList
	err = json.Unmarshal(wflBytes, &copy)
	if err != nil {
		panic(err)
	}
	return &copy
}

func (wf *Workflow) GetTemplate(name string) *Template {
	for _, t := range wf.Spec.Templates {
		if t.Name == name {
			return &t
		}
	}
	return nil
}

// NodeID creates a deterministic node ID based on a node name
func (wf *Workflow) NodeID(name string) string {
	if name == wf.ObjectMeta.Name {
		return wf.ObjectMeta.Name
	}
	h := fnv.New32a()
	h.Write([]byte(name))
	return fmt.Sprintf("%s-%v", wf.ObjectMeta.Name, h.Sum32())
}