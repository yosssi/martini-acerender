package acerender

import "github.com/yosssi/ace"

// Options represents options for the Ace renderer.
type Options struct {
	// AceOptions represents the options for Ace.
	AceOptions *ace.Options
}

// initializeOptions initializes the options.
func initializeOptions(opts *Options) *Options {
	if opts == nil {
		opts = &Options{}
	}

	return opts
}
