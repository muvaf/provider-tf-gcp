/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ActionObservation struct {
}

type ActionParameters struct {

	// The target Storage Class of objects affected by this Lifecycle Rule. Supported values include: MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE, ARCHIVE.
	// +kubebuilder:validation:Optional
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`

	// The type of the action of this Lifecycle Rule. Supported values include: Delete and SetStorageClass.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type BucketObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type BucketParameters struct {

	// The bucket's Cross-Origin Resource Sharing (CORS) configuration.
	// +kubebuilder:validation:Optional
	Cors []CorsParameters `json:"cors,omitempty" tf:"cors,omitempty"`

	// Whether or not to automatically apply an eventBasedHold to new objects added to the bucket.
	// +kubebuilder:validation:Optional
	DefaultEventBasedHold *bool `json:"defaultEventBasedHold,omitempty" tf:"default_event_based_hold,omitempty"`

	// The bucket's encryption configuration.
	// +kubebuilder:validation:Optional
	Encryption []EncryptionParameters `json:"encryption,omitempty" tf:"encryption,omitempty"`

	// When deleting a bucket, this boolean option will delete all contained objects. If you try to delete a bucket that contains objects, Terraform will fail that run.
	// +kubebuilder:validation:Optional
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// A set of key/value label pairs to assign to the bucket.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The bucket's Lifecycle Rules configuration.
	// +kubebuilder:validation:Optional
	LifecycleRule []LifecycleRuleParameters `json:"lifecycleRule,omitempty" tf:"lifecycle_rule,omitempty"`

	// The Google Cloud Storage location
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// The bucket's Access & Storage Logs configuration.
	// +kubebuilder:validation:Optional
	Logging []LoggingParameters `json:"logging,omitempty" tf:"logging,omitempty"`

	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Enables Requester Pays on a storage bucket.
	// +kubebuilder:validation:Optional
	RequesterPays *bool `json:"requesterPays,omitempty" tf:"requester_pays,omitempty"`

	// Configuration of the bucket's data retention policy for how long objects in the bucket should be retained.
	// +kubebuilder:validation:Optional
	RetentionPolicy []RetentionPolicyParameters `json:"retentionPolicy,omitempty" tf:"retention_policy,omitempty"`

	// The Storage Class of the new bucket. Supported values include: STANDARD, MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE, ARCHIVE.
	// +kubebuilder:validation:Optional
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`

	// Enables uniform bucket-level access on a bucket.
	// +kubebuilder:validation:Optional
	UniformBucketLevelAccess *bool `json:"uniformBucketLevelAccess,omitempty" tf:"uniform_bucket_level_access,omitempty"`

	// The bucket's Versioning configuration.
	// +kubebuilder:validation:Optional
	Versioning []VersioningParameters `json:"versioning,omitempty" tf:"versioning,omitempty"`

	// Configuration if the bucket acts as a website.
	// +kubebuilder:validation:Optional
	Website []WebsiteParameters `json:"website,omitempty" tf:"website,omitempty"`
}

type ConditionObservation struct {
}

type ConditionParameters struct {

	// Minimum age of an object in days to satisfy this condition.
	// +kubebuilder:validation:Optional
	Age *float64 `json:"age,omitempty" tf:"age,omitempty"`

	// Creation date of an object in RFC 3339 (e.g. 2017-06-13) to satisfy this condition.
	// +kubebuilder:validation:Optional
	CreatedBefore *string `json:"createdBefore,omitempty" tf:"created_before,omitempty"`

	// Creation date of an object in RFC 3339 (e.g. 2017-06-13) to satisfy this condition.
	// +kubebuilder:validation:Optional
	CustomTimeBefore *string `json:"customTimeBefore,omitempty" tf:"custom_time_before,omitempty"`

	// Number of days elapsed since the user-specified timestamp set on an object.
	// +kubebuilder:validation:Optional
	DaysSinceCustomTime *float64 `json:"daysSinceCustomTime,omitempty" tf:"days_since_custom_time,omitempty"`

	// Number of days elapsed since the noncurrent timestamp of an object. This
	// condition is relevant only for versioned objects.
	// +kubebuilder:validation:Optional
	DaysSinceNoncurrentTime *float64 `json:"daysSinceNoncurrentTime,omitempty" tf:"days_since_noncurrent_time,omitempty"`

	// Storage Class of objects to satisfy this condition. Supported values include: MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE, ARCHIVE, STANDARD, DURABLE_REDUCED_AVAILABILITY.
	// +kubebuilder:validation:Optional
	MatchesStorageClass []*string `json:"matchesStorageClass,omitempty" tf:"matches_storage_class,omitempty"`

	// Creation date of an object in RFC 3339 (e.g. 2017-06-13) to satisfy this condition.
	// +kubebuilder:validation:Optional
	NoncurrentTimeBefore *string `json:"noncurrentTimeBefore,omitempty" tf:"noncurrent_time_before,omitempty"`

	// Relevant only for versioned objects. The number of newer versions of an object to satisfy this condition.
	// +kubebuilder:validation:Optional
	NumNewerVersions *float64 `json:"numNewerVersions,omitempty" tf:"num_newer_versions,omitempty"`

	// Match to live and/or archived objects. Unversioned buckets have only live objects. Supported values include: "LIVE", "ARCHIVED", "ANY".
	// +kubebuilder:validation:Optional
	WithState *string `json:"withState,omitempty" tf:"with_state,omitempty"`
}

type CorsObservation struct {
}

type CorsParameters struct {

	// The value, in seconds, to return in the Access-Control-Max-Age header used in preflight responses.
	// +kubebuilder:validation:Optional
	MaxAgeSeconds *float64 `json:"maxAgeSeconds,omitempty" tf:"max_age_seconds,omitempty"`

	// The list of HTTP methods on which to include CORS response headers, (GET, OPTIONS, POST, etc) Note: "*" is permitted in the list of methods, and means "any method".
	// +kubebuilder:validation:Optional
	Method []*string `json:"method,omitempty" tf:"method,omitempty"`

	// The list of Origins eligible to receive CORS response headers. Note: "*" is permitted in the list of origins, and means "any Origin".
	// +kubebuilder:validation:Optional
	Origin []*string `json:"origin,omitempty" tf:"origin,omitempty"`

	// The list of HTTP headers other than the simple response headers to give permission for the user-agent to share across domains.
	// +kubebuilder:validation:Optional
	ResponseHeader []*string `json:"responseHeader,omitempty" tf:"response_header,omitempty"`
}

type EncryptionObservation struct {
}

type EncryptionParameters struct {

	// A Cloud KMS key that will be used to encrypt objects inserted into this bucket, if no encryption method is specified. You must pay attention to whether the crypto key is available in the location that this bucket is created in. See the docs for more details.
	// +kubebuilder:validation:Required
	DefaultKMSKeyName *string `json:"defaultKmsKeyName" tf:"default_kms_key_name,omitempty"`
}

type LifecycleRuleObservation struct {
}

type LifecycleRuleParameters struct {

	// The Lifecycle Rule's action configuration. A single block of this type is supported.
	// +kubebuilder:validation:Required
	Action []ActionParameters `json:"action" tf:"action,omitempty"`

	// The Lifecycle Rule's condition configuration.
	// +kubebuilder:validation:Required
	Condition []ConditionParameters `json:"condition" tf:"condition,omitempty"`
}

type LoggingObservation struct {
}

type LoggingParameters struct {

	// The bucket that will receive log objects.
	// +kubebuilder:validation:Required
	LogBucket *string `json:"logBucket" tf:"log_bucket,omitempty"`

	// The object prefix for log objects. If it's not provided, by default Google Cloud Storage sets this to this bucket's name.
	// +kubebuilder:validation:Optional
	LogObjectPrefix *string `json:"logObjectPrefix,omitempty" tf:"log_object_prefix,omitempty"`
}

type RetentionPolicyObservation struct {
}

type RetentionPolicyParameters struct {

	// If set to true, the bucket will be locked and permanently restrict edits to the bucket's retention policy.  Caution: Locking a bucket is an irreversible action.
	// +kubebuilder:validation:Optional
	IsLocked *bool `json:"isLocked,omitempty" tf:"is_locked,omitempty"`

	// The period of time, in seconds, that objects in the bucket must be retained and cannot be deleted, overwritten, or archived. The value must be less than 3,155,760,000 seconds.
	// +kubebuilder:validation:Required
	RetentionPeriod *float64 `json:"retentionPeriod" tf:"retention_period,omitempty"`
}

type VersioningObservation struct {
}

type VersioningParameters struct {

	// While set to true, versioning is fully enabled for this bucket.
	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`
}

type WebsiteObservation struct {
}

type WebsiteParameters struct {

	// Behaves as the bucket's directory index where missing objects are treated as potential directories.
	// +kubebuilder:validation:Optional
	MainPageSuffix *string `json:"mainPageSuffix,omitempty" tf:"main_page_suffix,omitempty"`

	// The custom object to return when a requested resource is not found.
	// +kubebuilder:validation:Optional
	NotFoundPage *string `json:"notFoundPage,omitempty" tf:"not_found_page,omitempty"`
}

// BucketSpec defines the desired state of Bucket
type BucketSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketParameters `json:"forProvider"`
}

// BucketStatus defines the observed state of Bucket.
type BucketStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Bucket is the Schema for the Buckets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type Bucket struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BucketSpec   `json:"spec"`
	Status            BucketStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketList contains a list of Buckets
type BucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Bucket `json:"items"`
}

// Repository type metadata.
var (
	Bucket_Kind             = "Bucket"
	Bucket_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Bucket_Kind}.String()
	Bucket_KindAPIVersion   = Bucket_Kind + "." + CRDGroupVersion.String()
	Bucket_GroupVersionKind = CRDGroupVersion.WithKind(Bucket_Kind)
)

func init() {
	SchemeBuilder.Register(&Bucket{}, &BucketList{})
}
