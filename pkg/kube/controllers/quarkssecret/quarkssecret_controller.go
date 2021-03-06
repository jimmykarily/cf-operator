package quarkssecret

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	crc "sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/source"

	credsgen "code.cloudfoundry.org/cf-operator/pkg/credsgen/in_memory_generator"
	qsv1a1 "code.cloudfoundry.org/cf-operator/pkg/kube/apis/quarkssecret/v1alpha1"
	"code.cloudfoundry.org/quarks-utils/pkg/config"
	"code.cloudfoundry.org/quarks-utils/pkg/ctxlog"
)

// AddQuarksSecret creates a new QuarksSecrets controller to watch for the
// custom resource and reconcile it into k8s secrets.
func AddQuarksSecret(ctx context.Context, config *config.Config, mgr manager.Manager) error {
	ctx = ctxlog.NewContextWithRecorder(ctx, "quarks-secret-reconciler", mgr.GetEventRecorderFor("quarks-secret-recorder"))
	log := ctxlog.ExtractLogger(ctx)
	r := NewQuarksSecretReconciler(ctx, config, mgr, credsgen.NewInMemoryGenerator(log), controllerutil.SetControllerReference)

	// Create a new controller
	c, err := controller.New("quarks-secret-controller", mgr, controller.Options{
		Reconciler:              r,
		MaxConcurrentReconciles: config.MaxQuarksSecretWorkers,
	})
	if err != nil {
		return errors.Wrap(err, "Adding quarks secret controller to manager failed.")
	}

	// Watch for changes to QuarksSecrets
	p := predicate.Funcs{
		CreateFunc: func(e event.CreateEvent) bool {
			o := e.Object.(*qsv1a1.QuarksSecret)
			secrets, err := listSecrets(ctx, mgr.GetClient(), o)
			if err != nil {
				ctxlog.Errorf(ctx, "Failed to list secrets owned by QuarksSecret '%s': %s in quarksSecret reconciler", o.Name, err)
			}
			if len(secrets) == 0 {
				ctxlog.NewPredicateEvent(e.Object).Debug(
					ctx, e.Meta, "qsv1a1.QuarksSecret",
					fmt.Sprintf("Create predicate passed for '%s'", e.Meta.GetName()),
				)
				return true
			}
			return false
		},
		DeleteFunc:  func(e event.DeleteEvent) bool { return false },
		GenericFunc: func(e event.GenericEvent) bool { return false },
		UpdateFunc: func(e event.UpdateEvent) bool {
			n := e.ObjectNew.(*qsv1a1.QuarksSecret)
			if !n.Status.Generated {
				ctxlog.NewPredicateEvent(e.ObjectNew).Debug(
					ctx, e.MetaNew, "qsv1a1.QuarksSecret",
					fmt.Sprintf("Update predicate passed for '%s'", e.MetaNew.GetName()),
				)
				return true
			}
			return false
		},
	}
	err = c.Watch(&source.Kind{Type: &qsv1a1.QuarksSecret{}}, &handler.EnqueueRequestForObject{}, p)
	if err != nil {
		return errors.Wrapf(err, "Watching quarks secrets failed in quarksSecret controller.")
	}

	return nil
}

// listSecrets gets all Secrets owned by the QuarksSecret
func listSecrets(ctx context.Context, client crc.Client, qSecret *qsv1a1.QuarksSecret) ([]corev1.Secret, error) {
	ctxlog.Debug(ctx, "Listing Secrets owned by QuarksSecret '", qSecret.Name, "'.")

	secretLabels := map[string]string{qsv1a1.LabelKind: qsv1a1.GeneratedSecretKind}
	result := []corev1.Secret{}

	allSecrets := &corev1.SecretList{}
	err := client.List(ctx, allSecrets,
		crc.InNamespace(qSecret.Namespace),
		crc.MatchingLabels(secretLabels),
	)
	if err != nil {
		return nil, err
	}

	for _, secret := range allSecrets.Items {
		if metav1.IsControlledBy(&secret, qSecret) {
			result = append(result, secret)
			ctxlog.Debug(ctx, "Found Secret '", secret.Name, "' owned by QuarksSecret '", qSecret.Name, "'.")
		}
	}

	if len(result) == 0 {
		ctxlog.Debug(ctx, "Did not find any Secret owned by QuarksSecret '", qSecret.Name, "'.")
	}

	return result, nil
}
