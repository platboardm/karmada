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

package annotations

import (
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestGetAnnotationValue(t *testing.T) {
	tests := []struct {
		name        string
		object      metav1.Object
		key         string
		expected    string
	}{
		{
			name: "annotation exists",
			object: &metav1.ObjectMeta{
				Annotations: map[string]string{
					"test-key": "test-value",
				},
			},
			key:      "test-key",
			expected: "test-value",
		},
		{
			name: "annotation does not exist",
			object: &metav1.ObjectMeta{
				Annotations: map[string]string{},
			},
			key:      "missing-key",
			expected: "",
		},
		{
			name:     "nil annotations",
			object:   &metav1.ObjectMeta{},
			key:      "test-key",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetAnnotationValue(tt.object, tt.key)
			if got != tt.expected {
				t.Errorf("GetAnnotationValue() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestSetAnnotation(t *testing.T) {
	tests := []struct {
		name   string
		object metav1.Object
		key    string
		value  string
	}{
		{
			name:   "set annotation on object with nil annotations",
			object: &metav1.ObjectMeta{},
			key:    "new-key",
			value:  "new-value",
		},
		{
			name: "overwrite existing annotation",
			object: &metav1.ObjectMeta{
				Annotations: map[string]string{
					"existing-key": "old-value",
				},
			},
			key:   "existing-key",
			value: "new-value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetAnnotation(tt.object, tt.key, tt.value)
			got := tt.object.GetAnnotations()[tt.key]
			if got != tt.value {
				t.Errorf("SetAnnotation() annotation value = %v, want %v", got, tt.value)
			}
		})
	}
}

func TestRemoveAnnotation(t *testing.T) {
	tests := []struct {
		name   string
		object metav1.Object
		key    string
	}{
		{
			name: "remove existing annotation",
			object: &metav1.ObjectMeta{
				Annotations: map[string]string{
					"test-key": "test-value",
				},
			},
			key: "test-key",
		},
		{
			name:   "remove annotation from nil map",
			object: &metav1.ObjectMeta{},
			key:    "test-key",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RemoveAnnotation(tt.object, tt.key)
			if _, exists := tt.object.GetAnnotations()[tt.key]; exists {
				t.Errorf("RemoveAnnotation() annotation %v still exists after removal", tt.key)
			}
		})
	}
}

func TestHasAnnotation(t *testing.T) {
	tests := []struct {
		name     string
		object   metav1.Object
		key      string
		expected bool
	}{
		{
			name: "annotation exists",
			object: &metav1.ObjectMeta{
				Annotations: map[string]string{
					"test-key": "test-value",
				},
			},
			key:      "test-key",
			expected: true,
		},
		{
			name:     "annotation does not exist",
			object:   &metav1.ObjectMeta{},
			key:      "missing-key",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HasAnnotation(tt.object, tt.key)
			if got != tt.expected {
				t.Errorf("HasAnnotation() = %v, want %v", got, tt.expected)
			}
		})
	}
}
