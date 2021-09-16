/*
Copyright 2021.

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

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	wasmcloudcomv1alpha1 "wasmcloud-k8s-operator/app/api/v1alpha1"
)

// AppReconciler reconciles a App object
type AppReconciler struct {
	client.Client
	Scheme   *runtime.Scheme
	Log      logr.Logger
	Recorder record.EventRecorder
}

//+kubebuilder:rbac:groups=wasmcloud.com.wasmcloud-k8s-operator,resources=apps,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=wasmcloud.com.wasmcloud-k8s-operator,resources=apps/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=wasmcloud.com.wasmcloud-k8s-operator,resources=apps/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the App object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.9.2/pkg/reconcile
func (r *AppReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := r.Log.WithValues("WasmCloud", req.NamespacedName)

	// your logic here
	var application wasmcloudcomv1alpha1.App
	log.Info("payload", "req", req)
	if err := r.Get(ctx, req.NamespacedName, &application); err != nil {
		if errors.IsNotFound(err) {
			log.Info("The resource was probably deleted", "Request", req)
		}
	}
	log.Info("application payload", "application", application)
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *AppReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&wasmcloudcomv1alpha1.App{}).
		Complete(r)
}
