package test

import (
	skoopv1 "github.com/mauza/skoop/protos/gen/go/skoop/v1"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"google.golang.org/protobuf/testing/protocmp"
	"testing"
	"time"
)

type SkoopSuite struct {
	suite.Suite
}

func TestSkoopSuite(t *testing.T) {
	suite.Run(t, new(SkoopSuite))
}

func (s *SkoopSuite) SetupSuite() {
	initializeApiClient(5 * time.Second)
	loadTestData(s.T())
}

func assertProtoEqualitySortById(t *testing.T, expected, actual interface{}, opts ...cmp.Option) {
	theOpts := []cmp.Option{
		cmpopts.SortSlices(func(x, y *skoopv1.Hello) bool {
			return *x.Id < *y.Id
		}),
		protocmp.Transform(),
		protocmp.SortRepeated(func(x, y *skoopv1.Hello) bool {
			return *x.Id < *y.Id
		}),
	}
	theOpts = append(theOpts, opts...)
	diff := cmp.Diff(expected, actual, theOpts...)
	require.Empty(t, diff)
}
