package calculations

import (
	"log"
	"os"
)

func Calculate(n int64, logging bool) int64 {

	if logging {
		log.SetOutput(os.Stdout)

		log.Println("Start calculations...")
		log.Printf("Calculate %d!\n", n)
	}

	result := int64(1)
	for i := int64(2); i <= n; i++ {
		result *= i
	}

	if logging {
		log.Println("Calculations complete!")
	}

	return result
}
