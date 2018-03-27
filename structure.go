package a2conf

// Location .
type Location struct {
	Value string
}

// Directory .
type Directory struct {
	Options string
}

// VirualHost .
type VirualHost struct {
	Include []string

	Location []*Location

	SSLEngine             string
	SSLCertificateFile    string
	SSLCertificateKeyFile string

	IfModule []*IfModule
}

// IfModule .
type IfModule struct {
	Children map[string]string
}
