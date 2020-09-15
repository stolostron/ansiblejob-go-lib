/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"strings"

	"github.com/go-logr/logr"
	k8serr "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"

	"sigs.k8s.io/controller-runtime/pkg/client"

	towerv1alpha1 "github.com/open-cluster-management/ansiblejob-go-lib/api/v1alpha1"
)

// AnsibleJobReconciler reconciles a AnsibleJob object
type AnsibleJobReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

const (
	TestAnnotation = "test"
)

// +kubebuilder:rbac:groups=tower.ansible.com,resources=ansiblejobs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=tower.ansible.com,resources=ansiblejobs/status,verbs=get;update;patch

func (r *AnsibleJobReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	logger := r.Log.WithValues("ansiblejob", req.NamespacedName)

	// your logic here
	ins := &towerv1alpha1.AnsibleJob{}
	if err := r.Client.Get(context.TODO(), req.NamespacedName, ins); err != nil {
		// instance is deleted
		if k8serr.IsNotFound(err) {
			return ctrl.Result{}, nil
		}

		return ctrl.Result{}, err
	}

	insAn := ins.GetAnnotations()
	if len(insAn) == 0 || insAn[TestAnnotation] == "" {
		return ctrl.Result{}, nil
	}

	// test marked done, then change the instance status
	if strings.EqualFold(insAn[TestAnnotation], "done") {
		newIns := ins.DeepCopy()
		newIns.Status.AnsibleJobResult.Status = towerv1alpha1.JobScussed
		newIns.Status.Condition.LastTransitionTime = metav1.Now()

		if err := r.Client.Status().Update(context.TODO(), newIns); err != nil {
			logger.Error(err, "failed to update the ansiblejob instance")
			return ctrl.Result{}, nil
		}
	}

	return ctrl.Result{}, nil
}

func (r *AnsibleJobReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&towerv1alpha1.AnsibleJob{}).
		Complete(r)
}
