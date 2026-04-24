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

func TestContainsAnnotations(t *testing.T) {
	tests := []struct {
		name        string
		object      metav1.Object
		annotations map[string]string
		want        bool
	}{
		{
			name: "object contains all annotations",
			object: &metav1.ObjectMeta{
				Annotations: map[string]string{
					"key1": "value1",
					"key2": "value2",
				},
			},
			annotations: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			want: true,
		},
		{
			name: "object contains subset of annotations",
			object: &metav1.ObjectMeta{
				Annotations: map[string]string{
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
				},
			},
			annotations: map[string]string{
				"key1": "value1",
			},
			want: true,
		},
		{
			name: "object missing one annotation",
			object: &metav1.ObjectMeta{
				Annotations: map[string]string{
					"key1": "value1",
				},
			},
			annotations: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			want: false,
		},
		{
			name: "object has annotation with different value",
			object: &metav1.ObjectMeta{
				Annotations: map[string]string{
					"key1": "wrong-value",
				},
			},
			annotations: map[string]string{
				"key1": "value1",
			},
			want: false,
		},
		{
			name: "object has no annotations",
			object: &metav1.ObjectMeta{
				Annotations: nil,
			},
			annotations: map[string]string{
				"key1": "value1",
			},
			want: false,
		},
		{
			name: "empty annotations to check always returns true",
			object: &metav1.ObjectMeta{
				Annotations: map[string]string{
					"key1": "value1",
				},
			},
			annotations: map[string]string{},
			want:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ContainsAnnotations(tt.object, tt.annotations)
			if got != tt.want {
				t.Errorf("ContainsAnnotations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeAnnotations(t *testing.T) {
	tests := []struct {
		name     string
		object   metav1.Object
		toMerge  map[string]string
		wantAnno map[string]string
	}{
		{
			name: "merge into existing annotations",
			object: &metav1.ObjectMeta{
				Annotations: map[string]string{
					"existing-key": "existing-value",
				},
			},
			toMerge: map[string]string{
				"new-key": "new-value",
			},
			wantAnno: map[string]string{
				"existing-key": "existing-value",
				"new-key":      "new-value",
			},
		},
		{
			name: "merge overwrites existing key",
			object: &metav1.ObjectMeta{
				Annotations: map[string]string{
					"key": "old-value",
				},
			},
			toMerge: map[string]string{
				"key": "new-value",
			},
			wantAnno: map[string]string{
				"key": "new-value",
			},
		},
		{
			name: "merge into nil annotations",
			object: &metav1.ObjectMeta{
				Annotations: nil,
			},
			toMerge: map[string]string{
				"key": "value",
			},
			wantAnno: map[string]string{
				"key": "value",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MergeAnnotations(tt.object, tt.toMerge)
			got := tt.object.GetAnnotations()
			for k, v := range tt.wantAnno {
				if got[k] != v {
					t.Errorf("MergeAnnotations() annotation[%s] = %v, want %v", k, got[k], v)
				}
			}
			if len(got) != len(tt.wantAnno) {
				t.Errorf("MergeAnnotations() len = %d, want %d", len(got), len(tt.wantAnno))
			}
		})
	}
}
