/*
Copyright 2020 The Tekton Authors

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/tektoncd/operator/pkg/apis/operator/v1alpha1"
	operatorv1alpha1 "github.com/tektoncd/operator/pkg/client/clientset/versioned/typed/operator/v1alpha1"
	gentype "k8s.io/client-go/gentype"
)

// fakeTektonDashboards implements TektonDashboardInterface
type fakeTektonDashboards struct {
	*gentype.FakeClientWithList[*v1alpha1.TektonDashboard, *v1alpha1.TektonDashboardList]
	Fake *FakeOperatorV1alpha1
}

func newFakeTektonDashboards(fake *FakeOperatorV1alpha1) operatorv1alpha1.TektonDashboardInterface {
	return &fakeTektonDashboards{
		gentype.NewFakeClientWithList[*v1alpha1.TektonDashboard, *v1alpha1.TektonDashboardList](
			fake.Fake,
			"",
			v1alpha1.SchemeGroupVersion.WithResource("tektondashboards"),
			v1alpha1.SchemeGroupVersion.WithKind("TektonDashboard"),
			func() *v1alpha1.TektonDashboard { return &v1alpha1.TektonDashboard{} },
			func() *v1alpha1.TektonDashboardList { return &v1alpha1.TektonDashboardList{} },
			func(dst, src *v1alpha1.TektonDashboardList) { dst.ListMeta = src.ListMeta },
			func(list *v1alpha1.TektonDashboardList) []*v1alpha1.TektonDashboard {
				return gentype.ToPointerSlice(list.Items)
			},
			func(list *v1alpha1.TektonDashboardList, items []*v1alpha1.TektonDashboard) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
