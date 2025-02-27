package main

import (
	"encoding/json"
	"os"
	"path"

	"github.com/GoogleCloudPlatform/cloud-foundation-toolkit/cli/bpmetadata"
	"github.com/invopop/jsonschema"
)

const schemaFileName = "bpmetadataschema.json"

// generateSchema creates a JSON Schema based on the types
// defined in the type BlueprintMetadata and it's recursive
// children. The generated schema will be used to validate
// all metadata files for consistency and will be uploaded
// to https://www.schemastore.org/ to provide IntelliSense
// VSCode for authors manually authoring the metadata.
func generateSchemaFile(o, wdPath string) error {
	sData, err := GenerateSchema()
	if err != nil {
		return err
	}

	// check if the provided output path is relative
	if !path.IsAbs(o) {
		o = path.Join(wdPath, o)
	}

	err = os.WriteFile(path.Join(o, schemaFileName), sData, 0644)
	if err != nil {
		return err
	}

	Log.Info("generated JSON schema for BlueprintMetadata", "path", path.Join(o, schemaFileName))
	return nil
}

func GenerateSchema() ([]byte, error) {
	r := &jsonschema.Reflector{}
	s := r.Reflect(&bpmetadata.BlueprintMetadata{})
	sData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return nil, err
	}

	return sData, nil
}
