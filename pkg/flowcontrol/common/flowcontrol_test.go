package common_test

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/fx"
	"google.golang.org/grpc/peer"

	flowcontrolv1 "github.com/fluxninja/aperture/api/gen/proto/go/aperture/flowcontrol/v1"
	"github.com/fluxninja/aperture/pkg/agentinfo"
	"github.com/fluxninja/aperture/pkg/config"
	"github.com/fluxninja/aperture/pkg/entitycache"
	"github.com/fluxninja/aperture/pkg/flowcontrol/common"
	grpcclient "github.com/fluxninja/aperture/pkg/net/grpc"
	"github.com/fluxninja/aperture/pkg/platform"
	"github.com/fluxninja/aperture/pkg/policies/dataplane"
)

var (
	app *fx.App
	svc flowcontrolv1.FlowControlServiceServer
)

var _ = BeforeEach(func() {
	entities := entitycache.NewEntityCache()
	entities.Put(&entitycache.Entity{
		ID:        entitycache.EntityID{},
		Services:  hardCodedServices,
		IPAddress: hardCodedIPAddress,
	})
	app = platform.New(
		config.ModuleConfig{
			MergeConfig: map[string]interface{}{
				"flowcontrol": map[string]interface{}{
					"controller_addr": "",
					"policies_file":   "",
				},
				"sentrywriter": map[string]interface{}{
					"disabled": true,
				},
			},
		}.Module(),
		fx.Provide(agentinfo.ProvideAgentInfo),
		fx.Supply(entities),
		fx.Provide(common.ProvideNopMetrics),
		fx.Provide(common.ProvideHandler),
		fx.Provide(dataplane.ProvideEngineAPI),
		grpcclient.ClientConstructor{Name: "flowcontrol-grpc-client", ConfigKey: "flowcontrol.client.grpc"}.Annotate(),
		fx.Populate(&svc),
	)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	err := app.Start(ctx)
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterEach(func() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_ = app.Stop(ctx)
})

var _ = Describe("FlowControl Check", func() {
	When("client request comes in", func() {
		It("returns decision accepted response", func() {
			ctx := peer.NewContext(context.Background(), newFakeRpcPeer())
			resp, err := svc.Check(ctx, &flowcontrolv1.CheckRequest{})
			Expect(err).NotTo(HaveOccurred())
			Expect((resp.GetDecisionType())).To(Equal(flowcontrolv1.DecisionType_DECISION_TYPE_ACCEPTED))
		})
	})
})

var (
	hardCodedServices  = []string{"service1", "service2"}
	hardCodedIPAddress = "1.2.3.4"
)

type fakeAddr string

func (a fakeAddr) Network() string { return "tcp" }
func (a fakeAddr) String() string  { return string(a) }

func newFakeRpcPeer() *peer.Peer {
	return &peer.Peer{Addr: fakeAddr("1.2.3.4:54321")}
}