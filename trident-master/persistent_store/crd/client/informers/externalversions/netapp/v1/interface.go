// Copyright 2023 NetApp, Inc. All Rights Reserved.

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "github.com/netapp/trident/persistent_store/crd/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// TridentActionMirrorUpdates returns a TridentActionMirrorUpdateInformer.
	TridentActionMirrorUpdates() TridentActionMirrorUpdateInformer
	// TridentActionSnapshotRestores returns a TridentActionSnapshotRestoreInformer.
	TridentActionSnapshotRestores() TridentActionSnapshotRestoreInformer
	// TridentBackends returns a TridentBackendInformer.
	TridentBackends() TridentBackendInformer
	// TridentBackendConfigs returns a TridentBackendConfigInformer.
	TridentBackendConfigs() TridentBackendConfigInformer
	// TridentMirrorRelationships returns a TridentMirrorRelationshipInformer.
	TridentMirrorRelationships() TridentMirrorRelationshipInformer
	// TridentNodes returns a TridentNodeInformer.
	TridentNodes() TridentNodeInformer
	// TridentSnapshots returns a TridentSnapshotInformer.
	TridentSnapshots() TridentSnapshotInformer
	// TridentSnapshotInfos returns a TridentSnapshotInfoInformer.
	TridentSnapshotInfos() TridentSnapshotInfoInformer
	// TridentStorageClasses returns a TridentStorageClassInformer.
	TridentStorageClasses() TridentStorageClassInformer
	// TridentTransactions returns a TridentTransactionInformer.
	TridentTransactions() TridentTransactionInformer
	// TridentVersions returns a TridentVersionInformer.
	TridentVersions() TridentVersionInformer
	// TridentVolumes returns a TridentVolumeInformer.
	TridentVolumes() TridentVolumeInformer
	// TridentVolumePublications returns a TridentVolumePublicationInformer.
	TridentVolumePublications() TridentVolumePublicationInformer
	// TridentVolumeReferences returns a TridentVolumeReferenceInformer.
	TridentVolumeReferences() TridentVolumeReferenceInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// TridentActionMirrorUpdates returns a TridentActionMirrorUpdateInformer.
func (v *version) TridentActionMirrorUpdates() TridentActionMirrorUpdateInformer {
	return &tridentActionMirrorUpdateInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TridentActionSnapshotRestores returns a TridentActionSnapshotRestoreInformer.
func (v *version) TridentActionSnapshotRestores() TridentActionSnapshotRestoreInformer {
	return &tridentActionSnapshotRestoreInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TridentBackends returns a TridentBackendInformer.
func (v *version) TridentBackends() TridentBackendInformer {
	return &tridentBackendInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TridentBackendConfigs returns a TridentBackendConfigInformer.
func (v *version) TridentBackendConfigs() TridentBackendConfigInformer {
	return &tridentBackendConfigInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TridentMirrorRelationships returns a TridentMirrorRelationshipInformer.
func (v *version) TridentMirrorRelationships() TridentMirrorRelationshipInformer {
	return &tridentMirrorRelationshipInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TridentNodes returns a TridentNodeInformer.
func (v *version) TridentNodes() TridentNodeInformer {
	return &tridentNodeInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TridentSnapshots returns a TridentSnapshotInformer.
func (v *version) TridentSnapshots() TridentSnapshotInformer {
	return &tridentSnapshotInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TridentSnapshotInfos returns a TridentSnapshotInfoInformer.
func (v *version) TridentSnapshotInfos() TridentSnapshotInfoInformer {
	return &tridentSnapshotInfoInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TridentStorageClasses returns a TridentStorageClassInformer.
func (v *version) TridentStorageClasses() TridentStorageClassInformer {
	return &tridentStorageClassInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TridentTransactions returns a TridentTransactionInformer.
func (v *version) TridentTransactions() TridentTransactionInformer {
	return &tridentTransactionInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TridentVersions returns a TridentVersionInformer.
func (v *version) TridentVersions() TridentVersionInformer {
	return &tridentVersionInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TridentVolumes returns a TridentVolumeInformer.
func (v *version) TridentVolumes() TridentVolumeInformer {
	return &tridentVolumeInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TridentVolumePublications returns a TridentVolumePublicationInformer.
func (v *version) TridentVolumePublications() TridentVolumePublicationInformer {
	return &tridentVolumePublicationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TridentVolumeReferences returns a TridentVolumeReferenceInformer.
func (v *version) TridentVolumeReferences() TridentVolumeReferenceInformer {
	return &tridentVolumeReferenceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
