package controllers

import (
	"context"
	"time"
	wasmcloudcomv1alpha1 "wasmcloud-k8s-operator/app/api/v1alpha1"

	"github.com/oam-dev/kubevela/pkg/oam/util"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	// . "github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

var _ = Describe("Test Create Application", func() {
	const (
		timeout  = time.Second * 10
		duration = time.Second * 10
		interval = time.Millisecond * 250
	)
	var (
		application *wasmcloudcomv1alpha1.App
	)
	BeforeEach(func() {
		application = &wasmcloudcomv1alpha1.App{
			TypeMeta: metav1.TypeMeta{
				Kind:       "App",
				APIVersion: "v1alpha1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      "latticecontroller",
				Namespace: "default",
			},
			Spec: wasmcloudcomv1alpha1.AppSpec{
				Components: []wasmcloudcomv1alpha1.ApplicationComponent{
					{
						Name: "userinfo",
						Type: "actor",
						Properties: util.Object2RawExtension(map[string]interface{}{
							"image": "wasmcloud.azurecr.io/fake:1",
						}),
						Traits: []wasmcloudcomv1alpha1.ApplicationTrait{
							{
								Type: "spreadscaler",
								Properties: util.Object2RawExtension(map[string]interface{}{
									"replicas": 4,
									"spread": []map[string]interface{}{
										{
											"name": "eastcoast",
											"requirements": map[string]interface{}{
												"zone": "us-east-1",
											},
											"weight": 80,
										},
										{
											"name": "westcoast",
											"requirements": map[string]interface{}{
												"zone": "us-west-1",
											},
											"weight": 20,
										},
									},
								}),
							},
							{
								Type: "linkdef",
								Properties: util.Object2RawExtension(map[string]interface{}{
									"target": "webcap",
									"values": map[string]interface{}{
										"port": 8080,
									},
								}),
							},
						},
					},
				},
			},
			Status: wasmcloudcomv1alpha1.AppStatus{},
		}
	})
	AfterEach(func() {})
	Context("Do", func() {
		It("Should create the application", func() {
			ctx := context.Background()
			Expect(k8sClient.Create(ctx, application)).Should(Succeed())
			app := wasmcloudcomv1alpha1.App{}
			Eventually(func() bool {
				err := k8sClient.Get(ctx, types.NamespacedName{
					Namespace: "default",
					Name:      "latticecontroller",
				}, &app)
				if err != nil {
					return false
				}
				return true
			}, timeout, interval).Should(BeTrue())
			Expect(len(app.Spec.Components)).Should(Equal(1))
		})
	})
})
