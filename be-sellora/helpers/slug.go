package helpers

import "strings"
import "regexp"


func Slugify(str string) string {
	slug := strings.TrimSpace(str)
	slug = strings.ToLower(slug)

	re := regexp.MustCompile(`[^a-z0-9]+`)
	slug = re.ReplaceAllString(slug, "-")

	slug = strings.ReplaceAll(slug, " ", "-")

	reDoubleDash := regexp.MustCompile(`--+`)
	slug = reDoubleDash.ReplaceAllString(slug, "-")

	slug = strings.Trim(slug, "-")

	return slug
}
