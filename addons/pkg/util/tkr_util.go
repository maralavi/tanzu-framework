// Copyright 2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package util

import (
	"context"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/vmware-tanzu/tanzu-framework/addons/pkg/constants"
	runtanzuv1alpha1 "github.com/vmware-tanzu/tanzu-framework/apis/run/v1alpha1"
	runtanzuv1alpha3 "github.com/vmware-tanzu/tanzu-framework/apis/run/v1alpha3"
)

// GetTKRByName gets TKR object given a TKR name and a label to help determine version of the TKR object
func GetTKRByName(ctx context.Context, c client.Client, tkrName, label string) (client.Object, error) {
	tkrV1Alpha1 := &runtanzuv1alpha1.TanzuKubernetesRelease{}
	tkrV1Alpha3 := &runtanzuv1alpha3.TanzuKubernetesRelease{}

	if tkrName == "" {
		return nil, nil
	}

	tkrNamespaceName := client.ObjectKey{
		Name: tkrName,
	}

	switch label {
	case constants.TKRLabel:
		if err := c.Get(ctx, tkrNamespaceName, tkrV1Alpha1); err != nil {
			if apierrors.IsNotFound(err) {
				return nil, nil
			}
			return nil, err
		}
		return tkrV1Alpha1, nil

	case constants.TKRLabelClassyClusters:
		if err := c.Get(ctx, tkrNamespaceName, tkrV1Alpha3); err != nil {
			if apierrors.IsNotFound(err) {
				return nil, nil
			}
			return nil, err
		}
		return tkrV1Alpha3, nil
	}

	return nil, nil
}
