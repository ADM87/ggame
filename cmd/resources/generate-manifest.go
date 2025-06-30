package resources

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	models "github.com/ADM87/ggame/resources"
	"github.com/ADM87/ggame/sys/types"
	"github.com/spf13/cobra"
)

func GenerateResourcesManifest() *cobra.Command {
	input := types.NewCmdArg("input", "i", "Input directory for resources", "./resources", false)
	output := types.NewCmdArg("output", "o", "Output file for the manifest", "./resources/static/manifest.json", false)

	cmd := &cobra.Command{
		Use:   "generate-manifest",
		Short: "Generate resources manifest",
		Long:  "Generates a manifest file for resources used in the game.",
		RunE: func(cmd *cobra.Command, args []string) error {
			inputPath, err := filepath.Abs(input.GetValue())
			if err != nil {
				return err
			}

			outputPath, err := filepath.Abs(output.GetValue())
			if err != nil {
				return err
			}

			manifest := &models.ResourceManifest{
				Dynamic: make(map[string]models.ResourceMetadata),
				Static:  make(map[string]models.ResourceMetadata),
			}

			staticPath := filepath.Join(inputPath, "static")
			dynamicPath := filepath.Join(inputPath, "dynamic")

			os.Remove(filepath.Join(staticPath, "manifest.json"))

			// Generate static resources
			staticResources, err := generateResourcesFrom(staticPath)
			if err != nil {
				return err
			}
			manifest.Static = mapResourceGroup(staticResources)

			// Generate dynamic resources
			dynamicResources, err := generateResourcesFrom(dynamicPath)
			if err != nil {
				return err
			}
			manifest.Dynamic = mapResourceGroup(dynamicResources)

			// Write the manifest to the output file
			file, err := os.Create(outputPath)
			if err != nil {
				return err
			}
			defer file.Close()

			encoder := json.NewEncoder(file)
			encoder.SetIndent("", "  ")
			if err := encoder.Encode(manifest); err != nil {
				return err
			}
			return nil
		},
	}

	input.RegisterWith(cmd.Flags().StringVarP)
	output.RegisterWith(cmd.Flags().StringVarP)

	return cmd
}

func generateResourcesFrom(path string) ([]models.ResourceMetadata, error) {
	var resources []models.ResourceMetadata

	err := filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || filepath.Ext(p) == "" {
			return nil // Skip directories and non-files
		}

		relPath, err := filepath.Rel(path, p)
		if err != nil {
			return err
		}
		resources = append(resources, models.ResourceMetadata{
			Path: relPath,
		})
		return nil
	})

	if err != nil {
		return nil, err
	}

	return resources, nil
}

func mapResourceGroup(group []models.ResourceMetadata) map[string]models.ResourceMetadata {
	resourceMap := make(map[string]models.ResourceMetadata)
	for _, resource := range group {
		name := strings.Replace(filepath.Base(resource.Path), filepath.Ext(resource.Path), "", 1)
		if _, exists := resourceMap[name]; exists {
			panic("Duplicate resource name found: " + name)
		}
		resourceMap[name] = resource

	}
	return resourceMap
}
