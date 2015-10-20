package runtime

// Return an implementation of Env, suitable for testing
func FakeEnv() Env {
	return NewScope(nil)
}
