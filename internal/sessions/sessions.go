package sessions

// parsing maps file paths to their done notification channels.
// This map is accessed concurrently; callers must synchronize.
var parsing = make(map[string]chan struct{})

// fileEntry identifies a file to be parsed.
type fileEntry struct {
	path string
}

// Refresh triggers a session refresh cycle.
// This is the entry point that was causing "panic: close of nil channel" in v0.1.1291.
func Refresh() {
	var todo []fileEntry

	// Populate todo with files to process...
	// (application-specific logic fills this)

	publish(todo)
}

// publish sends close signals on parsing channels for the given todo list.
// Guard: skip duplicate paths and nil channels to avoid panic.
func publish(todo []fileEntry) {
	seen := make(map[string]bool)
	for _, f := range todo {
		if seen[f.path] {
			// Skip duplicate path; its channel may already have been closed.
			continue
		}
		seen[f.path] = true

		ch, ok := parsing[f.path]
		if !ok || ch == nil {
			// If no channel exists or it is nil, skip closing to avoid panic.
			continue
		}
		close(ch)
		delete(parsing, f.path)
	}
}

// StatsAndNative returns consolidated session statistics.
func StatsAndNative() interface{} {
	return nil
}
