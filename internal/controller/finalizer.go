package controller

const AIPlatformFinalizer = "platform.mady.dev/finalizer"

func EnsureFinalizer(finalizers []string) ([]string, bool) {
	for _, f := range finalizers {
		if f == AIPlatformFinalizer {
			return finalizers, false
		}
	}
	return append(finalizers, AIPlatformFinalizer), true
}
func RemoveFinalizer(finalizers []string) ([]string, bool) {
	out := make([]string, 0, len(finalizers))
	removed := false
	for _, f := range finalizers {
		if f == AIPlatformFinalizer {
			removed = true
			continue
		}
		out = append(out, f)
	}
	return out, removed
}
