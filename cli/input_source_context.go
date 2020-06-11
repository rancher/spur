package cli

// InputSourceContext is an interface used to allow
// other input sources to be implemented as needed.
//
// Source returns an identifier for the input source. In case of file source
// it should return path to the file.
type InputSourceContext interface {
	Source() string
	Get(name string) (interface{}, bool)
}
