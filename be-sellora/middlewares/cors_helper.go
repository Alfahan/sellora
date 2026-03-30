package middlewares

import (
	"be-sellora/config"
	"strings"
)

func GetCORSOrigins() []string {
	originsEnv := config.GetEnv("FRONTEND_URL", "")
	rawOrigins := strings.Split(originsEnv, ",")
	uniqueMap := make(map[string]bool)

	var cleanOrigins []string
	for _, o := range rawOrigins {
		trimed := strings.TrimSpace(o)
		trimed = strings.Trim(trimed, "\"'")
		trimed = strings.TrimRight(trimed, "/")

		if trimed == "" {
			continue
		}

		if !uniqueMap[trimed] {
			uniqueMap[trimed] = true
			cleanOrigins = append(cleanOrigins, trimed)
		}
	}

	return cleanOrigins
}
