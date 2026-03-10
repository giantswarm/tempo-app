package mc

import (
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/giantswarm/apptest-framework/v3/pkg/state"
	"github.com/giantswarm/apptest-framework/v3/pkg/suite"
	"github.com/giantswarm/clustertest/v3/pkg/logger"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/types"
)

const (
	isUpgrade        = false
	installNamespace = "tempo"
)

func TestMC(t *testing.T) {
	suite.New().
		WithInstallNamespace(installNamespace).
		WithIsUpgrade(isUpgrade).
		WithValuesFile("./values.yaml").
		AfterClusterReady(func() {
			It("should connect to the management cluster", func() {
				Expect(state.GetFramework().MC().CheckConnection()).To(Succeed())
			})
		}).
		Tests(func() {
			// Write path
			It("should have tempo-distributor deployment ready on the MC", func() {
				mcClient := state.GetFramework().MC()
				Eventually(func() bool {
					logger.Log("Checking tempo-distributor deployment")
					var dep appsv1.Deployment
					if err := mcClient.Get(state.GetContext(), types.NamespacedName{Namespace: installNamespace, Name: "tempo-distributor"}, &dep); err != nil {
						return false
					}
					return dep.Status.ReadyReplicas == *dep.Spec.Replicas
				}).WithPolling(5 * time.Second).WithTimeout(10 * time.Minute).Should(BeTrue())
			})

			It("should have tempo-ingester statefulset ready on the MC", func() {
				mcClient := state.GetFramework().MC()
				Eventually(func() bool {
					logger.Log("Checking tempo-ingester statefulset")
					var sts appsv1.StatefulSet
					if err := mcClient.Get(state.GetContext(), types.NamespacedName{Namespace: installNamespace, Name: "tempo-ingester"}, &sts); err != nil {
						return false
					}
					return sts.Status.ReadyReplicas == *sts.Spec.Replicas
				}).WithPolling(5 * time.Second).WithTimeout(10 * time.Minute).Should(BeTrue())
			})

			// Read path
			It("should have tempo-querier deployment ready on the MC", func() {
				mcClient := state.GetFramework().MC()
				Eventually(func() bool {
					logger.Log("Checking tempo-querier deployment")
					var dep appsv1.Deployment
					if err := mcClient.Get(state.GetContext(), types.NamespacedName{Namespace: installNamespace, Name: "tempo-querier"}, &dep); err != nil {
						return false
					}
					return dep.Status.ReadyReplicas == *dep.Spec.Replicas
				}).WithPolling(5 * time.Second).WithTimeout(10 * time.Minute).Should(BeTrue())
			})

			It("should have tempo-query-frontend deployment ready on the MC", func() {
				mcClient := state.GetFramework().MC()
				Eventually(func() bool {
					logger.Log("Checking tempo-query-frontend deployment")
					var dep appsv1.Deployment
					if err := mcClient.Get(state.GetContext(), types.NamespacedName{Namespace: installNamespace, Name: "tempo-query-frontend"}, &dep); err != nil {
						return false
					}
					return dep.Status.ReadyReplicas == *dep.Spec.Replicas
				}).WithPolling(5 * time.Second).WithTimeout(10 * time.Minute).Should(BeTrue())
			})

			// Block compaction
			It("should have tempo-compactor deployment ready on the MC", func() {
				mcClient := state.GetFramework().MC()
				Eventually(func() bool {
					logger.Log("Checking tempo-compactor deployment")
					var dep appsv1.Deployment
					if err := mcClient.Get(state.GetContext(), types.NamespacedName{Namespace: installNamespace, Name: "tempo-compactor"}, &dep); err != nil {
						return false
					}
					return dep.Status.ReadyReplicas == *dep.Spec.Replicas
				}).WithPolling(5 * time.Second).WithTimeout(10 * time.Minute).Should(BeTrue())
			})

			// Entry point
			It("should have tempo-gateway deployment ready on the MC", func() {
				mcClient := state.GetFramework().MC()
				Eventually(func() bool {
					logger.Log("Checking tempo-gateway deployment")
					var dep appsv1.Deployment
					if err := mcClient.Get(state.GetContext(), types.NamespacedName{Namespace: installNamespace, Name: "tempo-gateway"}, &dep); err != nil {
						return false
					}
					return dep.Status.ReadyReplicas == *dep.Spec.Replicas
				}).WithPolling(5 * time.Second).WithTimeout(10 * time.Minute).Should(BeTrue())
			})
		}).
		Run(t, "Tempo MC test")
}
