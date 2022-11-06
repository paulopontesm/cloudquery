// Code generated by codegen; DO NOT EDIT.

package appstream

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/appstream"
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildAppstreamImageBuildersMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockAppstreamClient(ctrl)
	object := types.ImageBuilder{}
	err := faker.FakeObject(&object)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeImageBuilders(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&appstream.DescribeImageBuildersOutput{
			ImageBuilders: []types.ImageBuilder{object},
		}, nil)

	return client.Services{
		Appstream: m,
	}
}

func TestAppstreamImageBuilders(t *testing.T) {
	client.AwsMockTestHelper(t, ImageBuilders(), buildAppstreamImageBuildersMock, client.TestOptions{})
}
