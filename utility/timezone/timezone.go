package timezone

import (
	"log"
	"time"

	_ "time/tzdata"
)

// Used for set env timezone (example: "Asia/Bangkok")
func SetEnvTimezone(timezone string) {
	ict, err := time.LoadLocation(timezone)
	if err != nil {
		log.Fatalf("error loading location '%s': %v\n", timezone, err)
	}
	time.Local = ict
	log.Printf("Local time zone %v", time.Now().In(ict))
}
