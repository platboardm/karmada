/*
Copyright 2024 The Karmada Authors.

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

// Package annotations provides utility functions for working with
// Kubernetes object annotations in the Karmada control plane.
package annotations

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetAnnotationValue retrieves the value of a specific annotation from a
// Kubernetes object's metadata. Returns the annotation value and a boolean
// indicating whether the annotation was present.
//
// Example usage:
//
//	value, exists := GetAnnotationValue(obj.GetAnnotations(), "example.io/my-annotation")
//	if exists {
//	    fmt.Println("Annotation value:", value)
//	}
func GetAnnotationValue(annotations map[string]string, annotationKey string) (string, bool) {
	if annotations == nil {
		return "", false
	}
	value, exists := annotations[annotationKey]
	return value, exists
}

// SetAnnotation sets or updates an annotation on the provided annotations map.
// If the annotations map is nil, a new map is created and returned.
func SetAnnotation(annotations map[string]string, key, value string) map[string]string {
	if annotations == nil {
		annotations = make(map[string]string)
	}
	annotations[key] = value
	return annotations
}

// RemoveAnnotation removes an annotation by key from the provided annotations map.
// Returns the updated map. If the key does not exist, the map is returned unchanged.
func RemoveAnnotation(annotations map[string]string, key string) map[string]string {
	if annotations == nil {
		return nil
	}
	delete(annotations, key)
	return annotations
}

// HasAnnotation returns true if the given annotations map contains the specified key.
func HasAnnotation(annotations map[string]string, key string) bool {
	if annotations == nil {
		return false
	}
	_, exists := annotations[key]
	return exists
}

// ContainsAnnotations checks whether all the specified annotations (key-value pairs)
// are present in the given annotations map.
// Note: if required is nil or empty, this always returns true.
func ContainsAnnotations(annotations map[string]string, required map[string]string) bool {
	for k, v := range required {
		if val, exists := annotations[k]; !exists || val != v {
			return false
		}
	}
	return true
}

// MergeAnnotations merges the source annotations into the destination annotations map.
// If a key exists in both maps, the source value takes precedence.
// Returns the merged annotations map. If both src and dst are nil, returns nil.
// Note: dst is modified in place when non-nil; a new map is allocated only when dst is nil.
func MergeAnnotations(dst, src map[string]string) map[string]string {
	if src == nil {
		return dst
	}
	if dst == nil {
		dst = make(map[string]string, len(src))
	}
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

// GetObjectAnnotations is a convenience helper that retrieves annotations
// directly from a metav1.Object interface.
func GetObjectAnnotations(obj metav1.Object) map[string]string {
	if obj == nil {
		return nil
	}
	return obj.GetAnnotations()
}
