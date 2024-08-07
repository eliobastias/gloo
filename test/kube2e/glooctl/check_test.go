package glooctl_test

import (
	"os"

	"github.com/solo-io/gloo/test/kube2e"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/solo-io/go-utils/testutils/exec"
)

var _ = Describe("Check", func() {

	Context("exclude", Ordered, func() {

		BeforeAll(func() {
			// Check that everything is OK
			kube2e.GlooctlCheckEventuallyHealthy(1, testHelper, "90s")
		})

		It("all checks pass with OK status", func() {
			output, err := GlooctlOut("check", "-x", "xds-metrics")
			Expect(err).NotTo(HaveOccurred())

			Expect(output).To(ContainSubstring("Checking Deployments... OK"))
			Expect(output).To(ContainSubstring("Checking Upstreams... OK"))
			Expect(output).To(ContainSubstring("Checking UpstreamGroups... OK"))
			Expect(output).To(ContainSubstring("Checking AuthConfigs... OK"))
			Expect(output).To(ContainSubstring("Checking RateLimitConfigs... OK"))
			Expect(output).To(ContainSubstring("Checking Secrets... OK"))
			Expect(output).To(ContainSubstring("Checking VirtualServices... OK"))
			Expect(output).To(ContainSubstring("Checking Gateways... OK"))
			Expect(output).To(ContainSubstring("Checking Proxies... OK"))
		})

		It("can exclude deployments", func() {
			output, err := GlooctlOut("check", "-x", "xds-metrics,deployments")
			Expect(err).NotTo(HaveOccurred())

			Expect(output).NotTo(ContainSubstring("Checking Deployments..."))
			Expect(output).To(ContainSubstring("Checking Pods... OK"))
			Expect(output).To(ContainSubstring("Checking Upstreams... OK"))
			Expect(output).To(ContainSubstring("Checking UpstreamGroups... OK"))
			Expect(output).To(ContainSubstring("Checking AuthConfigs... OK"))
			Expect(output).To(ContainSubstring("Checking RateLimitConfigs... OK"))
			Expect(output).To(ContainSubstring("Checking Secrets... OK"))
			Expect(output).To(ContainSubstring("Checking VirtualServices... OK"))
			Expect(output).To(ContainSubstring("Checking Gateways... OK"))
			Expect(output).To(ContainSubstring("Checking Proxies... Skipping proxies because deployments were excluded"))
		})

		It("can exclude pods", func() {
			output, err := GlooctlOut("check", "-x", "xds-metrics,pods")
			Expect(err).NotTo(HaveOccurred())

			Expect(output).To(ContainSubstring("Checking Deployments... OK"))
			Expect(output).NotTo(ContainSubstring("Checking Pods..."))
			Expect(output).To(ContainSubstring("Checking Upstreams... OK"))
			Expect(output).To(ContainSubstring("Checking UpstreamGroups... OK"))
			Expect(output).To(ContainSubstring("Checking AuthConfigs... OK"))
			Expect(output).To(ContainSubstring("Checking RateLimitConfigs... OK"))
			Expect(output).To(ContainSubstring("Checking Secrets... OK"))
			Expect(output).To(ContainSubstring("Checking VirtualServices... OK"))
			Expect(output).To(ContainSubstring("Checking Gateways... OK"))
			Expect(output).To(ContainSubstring("Checking Proxies... OK"))
		})

		It("can exclude upstreams", func() {
			output, err := GlooctlOut("check", "-x", "xds-metrics,upstreams,virtual-services")
			Expect(err).NotTo(HaveOccurred())

			Expect(output).To(ContainSubstring("Checking Deployments... OK"))
			Expect(output).To(ContainSubstring("Checking Pods... OK"))
			Expect(output).NotTo(ContainSubstring("Checking Upstreams..."))
			Expect(output).To(ContainSubstring("Checking UpstreamGroups... OK"))
			Expect(output).To(ContainSubstring("Checking AuthConfigs... OK"))
			Expect(output).To(ContainSubstring("Checking RateLimitConfigs... OK"))
			Expect(output).To(ContainSubstring("Checking Secrets... OK"))
			Expect(output).To(ContainSubstring("Checking Gateways... OK"))
			Expect(output).To(ContainSubstring("Checking Proxies... OK"))
		})

		It("can exclude upstreamgroups", func() {
			output, err := GlooctlOut("check", "-x", "xds-metrics,upstreamgroup")
			Expect(err).NotTo(HaveOccurred())

			Expect(output).To(ContainSubstring("Checking Deployments... OK"))
			Expect(output).To(ContainSubstring("Checking Pods... OK"))
			Expect(output).To(ContainSubstring("Checking Upstreams... OK"))
			Expect(output).NotTo(ContainSubstring("Checking UpstreamGroups..."))
			Expect(output).To(ContainSubstring("Checking AuthConfigs... OK"))
			Expect(output).To(ContainSubstring("Checking RateLimitConfigs... OK"))
			Expect(output).To(ContainSubstring("Checking Secrets... OK"))
			Expect(output).To(ContainSubstring("Checking VirtualServices... OK"))
			Expect(output).To(ContainSubstring("Checking Gateways... OK"))
			Expect(output).To(ContainSubstring("Checking Proxies... OK"))
		})

		It("can exclude auth-configs", func() {
			output, err := GlooctlOut("check", "-x", "xds-metrics,auth-configs")
			Expect(err).NotTo(HaveOccurred())

			Expect(output).To(ContainSubstring("Checking Deployments... OK"))
			Expect(output).To(ContainSubstring("Checking Pods... OK"))
			Expect(output).To(ContainSubstring("Checking Upstreams... OK"))
			Expect(output).To(ContainSubstring("Checking UpstreamGroups... OK"))
			Expect(output).NotTo(ContainSubstring("Checking AuthConfigs..."))
			Expect(output).To(ContainSubstring("Checking RateLimitConfigs... OK"))
			Expect(output).To(ContainSubstring("Checking Secrets... OK"))
			Expect(output).To(ContainSubstring("Checking VirtualServices... OK"))
			Expect(output).To(ContainSubstring("Checking Gateways... OK"))
			Expect(output).To(ContainSubstring("Checking Proxies... OK"))
		})

		It("can exclude rate-limit-configs", func() {
			output, err := GlooctlOut("check", "-x", "xds-metrics,rate-limit-configs")
			Expect(err).NotTo(HaveOccurred())

			Expect(output).To(ContainSubstring("Checking Deployments... OK"))
			Expect(output).To(ContainSubstring("Checking Pods... OK"))
			Expect(output).To(ContainSubstring("Checking Upstreams... OK"))
			Expect(output).To(ContainSubstring("Checking UpstreamGroups... OK"))
			Expect(output).To(ContainSubstring("Checking AuthConfigs... OK"))
			Expect(output).NotTo(ContainSubstring("Checking RateLimitConfigs..."))
			Expect(output).To(ContainSubstring("Checking Secrets... OK"))
			Expect(output).To(ContainSubstring("Checking VirtualServices... OK"))
			Expect(output).To(ContainSubstring("Checking Gateways... OK"))
			Expect(output).To(ContainSubstring("Checking Proxies... OK"))
		})

		It("can exclude secrets", func() {
			output, err := GlooctlOut("check", "-x", "xds-metrics,secrets")
			Expect(err).NotTo(HaveOccurred())

			Expect(output).To(ContainSubstring("Checking Deployments... OK"))
			Expect(output).To(ContainSubstring("Checking Pods... OK"))
			Expect(output).To(ContainSubstring("Checking Upstreams... OK"))
			Expect(output).To(ContainSubstring("Checking UpstreamGroups... OK"))
			Expect(output).To(ContainSubstring("Checking AuthConfigs... OK"))
			Expect(output).To(ContainSubstring("Checking RateLimitConfigs... OK"))
			Expect(output).NotTo(ContainSubstring("Checking Secrets..."))
			Expect(output).To(ContainSubstring("Checking VirtualServices... OK"))
			Expect(output).To(ContainSubstring("Checking Gateways... OK"))
			Expect(output).To(ContainSubstring("Checking Proxies... OK"))
		})

		It("can exclude virtual-services", func() {
			output, err := GlooctlOut("check", "-x", "xds-metrics,virtual-services")
			Expect(err).NotTo(HaveOccurred())

			Expect(output).To(ContainSubstring("Checking Deployments... OK"))
			Expect(output).To(ContainSubstring("Checking Pods... OK"))
			Expect(output).To(ContainSubstring("Checking Upstreams... OK"))
			Expect(output).To(ContainSubstring("Checking UpstreamGroups... OK"))
			Expect(output).To(ContainSubstring("Checking AuthConfigs... OK"))
			Expect(output).To(ContainSubstring("Checking RateLimitConfigs... OK"))
			Expect(output).To(ContainSubstring("Checking Secrets... OK"))
			Expect(output).NotTo(ContainSubstring("Checking VirtualServices..."))
			Expect(output).To(ContainSubstring("Checking Gateways... OK"))
			Expect(output).To(ContainSubstring("Checking Proxies... OK"))
		})

		It("can exclude gateways", func() {
			output, err := GlooctlOut("check", "-x", "xds-metrics,gateways")
			Expect(err).NotTo(HaveOccurred())

			Expect(output).To(ContainSubstring("Checking Deployments... OK"))
			Expect(output).To(ContainSubstring("Checking Pods... OK"))
			Expect(output).To(ContainSubstring("Checking Upstreams... OK"))
			Expect(output).To(ContainSubstring("Checking UpstreamGroups... OK"))
			Expect(output).To(ContainSubstring("Checking AuthConfigs... OK"))
			Expect(output).To(ContainSubstring("Checking RateLimitConfigs... OK"))
			Expect(output).To(ContainSubstring("Checking Secrets... OK"))
			Expect(output).To(ContainSubstring("Checking VirtualServices... OK"))
			Expect(output).NotTo(ContainSubstring("Checking Gateways..."))
			Expect(output).To(ContainSubstring("Checking Proxies... OK"))
		})

		It("can exclude proxies", func() {
			output, err := GlooctlOut("check", "-x", "xds-metrics,proxies")
			Expect(err).NotTo(HaveOccurred())

			Expect(output).To(ContainSubstring("Checking Deployments... OK"))
			Expect(output).To(ContainSubstring("Checking Pods... OK"))
			Expect(output).To(ContainSubstring("Checking Upstreams... OK"))
			Expect(output).To(ContainSubstring("Checking UpstreamGroups... OK"))
			Expect(output).To(ContainSubstring("Checking AuthConfigs... OK"))
			Expect(output).To(ContainSubstring("Checking RateLimitConfigs... OK"))
			Expect(output).To(ContainSubstring("Checking Secrets... OK"))
			Expect(output).To(ContainSubstring("Checking VirtualServices... OK"))
			Expect(output).To(ContainSubstring("Checking Gateways... OK"))
			Expect(output).NotTo(ContainSubstring("Checking Proxies..."))
		})

		It("excludes the Kubernetes Gateway resources when it is disabled", func() {
			output, err := GlooctlOut("check", "-x", "xds-metrics")
			Expect(err).NotTo(HaveOccurred())

			Expect(output).To(ContainSubstring("Checking Deployments... OK"))
			Expect(output).To(ContainSubstring("Checking Upstreams... OK"))
			Expect(output).To(ContainSubstring("Checking UpstreamGroups... OK"))
			Expect(output).To(ContainSubstring("Checking AuthConfigs... OK"))
			Expect(output).To(ContainSubstring("Checking RateLimitConfigs... OK"))
			Expect(output).To(ContainSubstring("Checking Secrets... OK"))
			Expect(output).To(ContainSubstring("Checking VirtualServices... OK"))
			Expect(output).To(ContainSubstring("Checking Gateways... OK"))
			Expect(output).To(ContainSubstring("Checking Proxies... OK"))

			Expect(output).NotTo(ContainSubstring("Checking Kubernetes GatewayClasses..."))
			Expect(output).NotTo(ContainSubstring("Checking Kubernetes Gateways..."))
			Expect(output).NotTo(ContainSubstring("Checking Kubernetes HTTPRoutes..."))
		})
	})

	Context("timeouts", Ordered, func() {

		BeforeAll(func() {
			// Check that everything is OK
			kube2e.GlooctlCheckEventuallyHealthy(1, testHelper, "90s")
		})

		It("can set timeouts (too short)", func() {
			values, err := os.CreateTemp("", "*.yaml")
			Expect(err).NotTo(HaveOccurred())
			_, err = values.Write([]byte(`checkTimeoutSeconds: 1ns`))
			Expect(err).NotTo(HaveOccurred())

			_, err = GlooctlOut("check", "-c", values.Name())
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("context deadline exceeded"))
		})

		It("can set timeouts (appropriately)", func() {
			values, err := os.CreateTemp("", "*.yaml")
			Expect(err).NotTo(HaveOccurred())
			_, err = values.Write([]byte(`checkTimeoutSeconds: 300s`))
			Expect(err).NotTo(HaveOccurred())

			_, err = GlooctlOut("check", "-c", values.Name())
			Expect(err).ToNot(HaveOccurred())
		})

	})

	Context("read only", Ordered, func() {

		BeforeAll(func() {
			// Check that everything is OK
			kube2e.GlooctlCheckEventuallyHealthy(1, testHelper, "90s")
		})

		It("can run a read only check", func() {
			output, err := GlooctlOut("check", "--read-only")
			Expect(err).NotTo(HaveOccurred())

			Expect(output).To(ContainSubstring("Checking Deployments... OK"))
			Expect(output).To(ContainSubstring("Checking Pods... OK"))
			Expect(output).To(ContainSubstring("Checking Upstreams... OK"))
			Expect(output).To(ContainSubstring("Checking UpstreamGroups... OK"))
			Expect(output).To(ContainSubstring("Checking AuthConfigs... OK"))
			Expect(output).To(ContainSubstring("Checking RateLimitConfigs... OK"))
			Expect(output).To(ContainSubstring("Checking Secrets... OK"))
			Expect(output).To(ContainSubstring("Checking VirtualServices... OK"))
			Expect(output).To(ContainSubstring("Checking Gateways... OK"))
			Expect(output).To(ContainSubstring("Checking Proxies... OK"))
			Expect(output).To(ContainSubstring("Warning: checking proxies with port forwarding is disabled"))
			Expect(output).To(ContainSubstring("Warning: checking xds with port forwarding is disabled"))
		})

	})

	Context("kube context", Ordered, func() {

		BeforeAll(func() {
			// Check that everything is OK
			kube2e.GlooctlCheckEventuallyHealthy(1, testHelper, "90s")
		})

		It("fails when scanning with invalid kubecontext", func() {
			_, err := GlooctlOut("check", "--kube-context", "not-gloo-context")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("context \"not-gloo-context\" does not exist"))
		})

		It("succeeds with correct kubecontext", func() {
			// The name of the cluster we run these tests in is "kind" which is why this test works
			clusterName, ok := os.LookupEnv("CLUSTER_NAME")
			if !ok {
				clusterName = "kind"
			}
			_, err := GlooctlOut("check", "--kube-context", "kind-"+clusterName)
			Expect(err).NotTo(HaveOccurred())
		})

	})

	Context("namespace", Ordered, func() {

		BeforeAll(func() {
			// Check that everything is OK
			kube2e.GlooctlCheckEventuallyHealthy(1, testHelper, "90s")
		})

		It("connection fails on incorrect namespace check", func() {
			_, err := GlooctlOut("check", "-n", "not-gloo-system")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Could not communicate with kubernetes cluster: namespaces \"not-gloo-system\" not found"))

			_, err = GlooctlOut("check", "-n", "default")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Warning: The provided label selector (gloo) applies to no pods"))

			output, err := GlooctlOut("check", "-p", "not-gloo")
			Expect(err).ToNot(HaveOccurred())
			Expect(output).To(ContainSubstring("Warning: The provided label selector (not-gloo) applies to no pods"))
			Expect(output).To(ContainSubstring("No problems detected."))

			_, err = GlooctlOut("check", "-r", "not-gloo-system")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("No namespaces specified are currently being watched (defaulting to 'gloo-system' namespace)"))
		})

	})

	Context("gateway-proxy replicas", func() {

		BeforeEach(func() {
			// We scale up/down deployments in each test, so we need to be sure we are starting from a healthy point
			kube2e.GlooctlCheckEventuallyHealthy(1, testHelper, "90s")
		})

		When("there are 0 replicas of any gateway-proxy pod", func() {

			It("fails", func() {
				err := exec.RunCommand(testHelper.RootDir, false, "kubectl", "scale", "--replicas=0", "deployment", "gateway-proxy", "-n", "gloo-system")
				Expect(err).ToNot(HaveOccurred())
				err = exec.RunCommand(testHelper.RootDir, false, "kubectl", "scale", "--replicas=0", "deployment", "public-gw", "-n", "gloo-system")
				Expect(err).ToNot(HaveOccurred())

				_, err = GlooctlOut("check", "-x", "xds-metrics")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Gloo installation is incomplete: no active gateway-proxy pods exist in cluster"))

				err = exec.RunCommand(testHelper.RootDir, false, "kubectl", "scale", "--replicas=1", "deployment", "gateway-proxy", "-n", "gloo-system")
				Expect(err).ToNot(HaveOccurred())
				err = exec.RunCommand(testHelper.RootDir, false, "kubectl", "scale", "--replicas=1", "deployment", "public-gw", "-n", "gloo-system")
				Expect(err).ToNot(HaveOccurred())
			})

		})

		When("there are 0 replicas of the gateway-proxy deployment", func() {

			It("warns", func() {
				err := exec.RunCommand(testHelper.RootDir, false, "kubectl", "scale", "--replicas=0", "deployment", "gateway-proxy", "-n", "gloo-system")
				Expect(err).ToNot(HaveOccurred())

				output, err := GlooctlOut("check", "-x", "xds-metrics")
				Expect(err).ToNot(HaveOccurred())
				Expect(output).To(ContainSubstring("Checking Deployments... OK"))
				Expect(output).To(ContainSubstring("Checking Pods... OK"))
				Expect(output).To(ContainSubstring("Checking Upstreams... OK"))
				Expect(output).To(ContainSubstring("Checking UpstreamGroups... OK"))
				Expect(output).To(ContainSubstring("Checking AuthConfigs... OK"))
				Expect(output).To(ContainSubstring("Checking RateLimitConfigs... OK"))
				Expect(output).To(ContainSubstring("Checking VirtualHostOptions... OK"))
				Expect(output).To(ContainSubstring("Checking RouteOptions... OK"))
				Expect(output).To(ContainSubstring("Checking Secrets... OK"))
				Expect(output).To(ContainSubstring("Checking VirtualServices... OK"))
				Expect(output).To(ContainSubstring("Checking Gateways... OK"))
				Expect(output).To(ContainSubstring("Checking Proxies... OK"))
				Expect(output).To(ContainSubstring("Warning: gloo-system:gateway-proxy has zero replicas"))
				Expect(output).To(ContainSubstring("No problems detected."))

				err = exec.RunCommand(testHelper.RootDir, false, "kubectl", "scale", "--replicas=1", "deployment", "gateway-proxy", "-n", "gloo-system")
				Expect(err).ToNot(HaveOccurred())
			})

		})

	})

	Context("error reporting/formatting", func() {

		BeforeEach(func() {
			// We apply resources in these tests, so we need to be sure we are starting from a healthy point
			kube2e.GlooctlCheckEventuallyHealthy(1, testHelper, "90s")
		})

		It("reports multiple errors at one time", func() {
			err := exec.RunCommand(testHelper.RootDir, false, "kubectl", "apply", "-f", testHelper.RootDir+"/test/kube2e/glooctl/reject-me.yaml")
			err = exec.RunCommand(testHelper.RootDir, false, "kubectl", "apply", "-f", testHelper.RootDir+"/test/kube2e/glooctl/reject-me-too.yaml")
			Expect(err).NotTo(HaveOccurred())

			Eventually(func(g Gomega) {
				_, err = GlooctlOut("check")
				g.Expect(err).To(HaveOccurred())
				g.Expect(err.Error()).To(ContainSubstring("* Found rejected virtual service by 'gloo-system': default reject-me-too (Reason: 2 errors occurred:"))
				g.Expect(err.Error()).To(ContainSubstring("* domain conflict: other virtual services that belong to the same Gateway as this one don't specify a domain (and thus default to '*'): [gloo-system.reject-me]"))
				g.Expect(err.Error()).To(ContainSubstring("* VirtualHost Error: DomainsNotUniqueError. Reason: domain * is shared by the following virtual hosts: [default.reject-me-too gloo-system.reject-me]"))

				g.Expect(err.Error()).To(ContainSubstring("* Found rejected virtual service by 'gloo-system': gloo-system reject-me (Reason: 2 errors occurred:"))
				g.Expect(err.Error()).To(ContainSubstring("* domain conflict: other virtual services that belong to the same Gateway as this one don't specify a domain (and thus default to '*'): [default.reject-me-too]"))
				g.Expect(err.Error()).To(ContainSubstring("* VirtualHost Error: DomainsNotUniqueError. Reason: domain * is shared by the following virtual hosts: [default.reject-me-too gloo-system.reject-me]"))
			}, "10s", "2s").Should(Succeed())

			err = exec.RunCommand(testHelper.RootDir, false, "kubectl", "delete", "-n", "gloo-system", "virtualservice", "reject-me")
			Expect(err).NotTo(HaveOccurred())
			err = exec.RunCommand(testHelper.RootDir, false, "kubectl", "delete", "-n", "default", "virtualservice", "reject-me-too")
			Expect(err).NotTo(HaveOccurred())
		})

	})

})
