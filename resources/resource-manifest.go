package resources

type ResourceManifest struct {
	Dynamic map[string]ResourceMetadata `json:"dynamic"`
	Static  map[string]ResourceMetadata `json:"static"`
}

type ResourceMetadata struct {
	Path string `json:"path"`
}
