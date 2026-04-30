package common

import (
	"context"
	"fmt"

	"github.com/Masterminds/semver"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetOCPVersion(ctx context.Context) (*semver.Version, error) {

	if sharedConfigClient == nil {
		return nil, fmt.Errorf("openshift Client is not initialized yet")
	}

	// Fetch the ClusterVersion object (always named "version")
	cv, err := sharedConfigClient.ConfigV1().ClusterVersions().Get(ctx, "version", metav1.GetOptions{})
	if err != nil {
		// If running on standard Kubernetes, this will return an IsNotFound error.
		// Handle gracefully if your operator supports both vanilla K8s and OCP.
		return nil, err
	}
	versionStr := cv.Status.Desired.Version
	if versionStr == "" {
		return nil, fmt.Errorf("empty OpenShift version in ClusterVersion status")
	}

	v, err := semver.NewVersion(versionStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse OpenShift version %q: %w", versionStr, err)
	}
	return v, nil
}
