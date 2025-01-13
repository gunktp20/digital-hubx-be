package utils

import "strings"

func formatAllowedTypes(allowedTypes []string) string {
	extensions := make([]string, len(allowedTypes))
	for i, t := range allowedTypes {
		switch t {
		default:
			if strings.Contains(t, "image/") {
				t = strings.Split(t, "image/")[1]
			}
			extensions[i] = t

		}
	}
	return strings.Join(extensions, " , ")
}
