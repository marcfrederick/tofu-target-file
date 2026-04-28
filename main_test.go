package main

import (
	"testing"
)

func Test_findResourcesInFile(t *testing.T) {
	tests := []struct {
		name      string
		file      string
		want      []Resource
		wantError bool
	}{
		{
			name: "single resource",
			file: "testdata/single_resource.tf",
			want: []Resource{
				{Type: "aws_s3_bucket", Name: "example"},
			},
		},
		{
			name: "multiple resources",
			file: "testdata/multiple_resources.tf",
			want: []Resource{
				{Type: "aws_s3_bucket", Name: "first"},
				{Type: "aws_s3_bucket", Name: "second"},
				{Type: "aws_iam_role", Name: "example"},
			},
		},
		{
			name: "non-resource blocks are ignored",
			file: "testdata/non_resource_blocks.tf",
			want: []Resource{
				{Type: "aws_instance", Name: "web"},
			},
		},
		{
			name: "empty file returns no resources",
			file: "testdata/empty.tf",
			want: nil,
		},
		{
			name:      "invalid HCL returns error",
			file:      "testdata/invalid.tf",
			wantError: true,
		},
		{
			name:      "missing file returns error",
			file:      "testdata/nonexistent.tf",
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := findResourcesInFile(tt.file)
			if (err != nil) != tt.wantError {
				t.Fatalf("findResourcesInFile() error = %v, wantError %v", err, tt.wantError)
			}
			if len(got) != len(tt.want) {
				t.Fatalf("findResourcesInFile() returned %d resources, want %d: %v", len(got), len(tt.want), got)
			}
			for i, r := range got {
				if r != tt.want[i] {
					t.Errorf("resource[%d] = %v, want %v", i, r, tt.want[i])
				}
			}
		})
	}
}

func TestResourceString(t *testing.T) {
	r := Resource{Type: "aws_s3_bucket", Name: "example"}
	if got := r.String(); got != "aws_s3_bucket.example" {
		t.Errorf("Resource.String() = %q, want %q", got, "aws_s3_bucket.example")
	}
}
