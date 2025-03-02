package main

import (
	"context"
	"fmt"
	"os"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	mlopsv1 "github.com/nickemma/mlops-starter/api/v1"
)

type ModelDeploymentReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

func (r *ModelDeploymentReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	var deployment mlopsv1.ModelDeployment
	if err := r.Get(ctx, req.NamespacedName, &deployment); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	fmt.Printf("Reconciling ModelDeployment: %s\n", deployment.Name)
	// Add actual reconciliation logic here
	return ctrl.Result{}, nil
}

func main() {
	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme: runtime.NewScheme(),
		Port:   9443,
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err = (&ModelDeploymentReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
