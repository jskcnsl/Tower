package layer

// Layer interface for different feature layer to run
type Layer interface {
	Run(args ...interface{}) error
	Config(args ...interface{}) error
}
