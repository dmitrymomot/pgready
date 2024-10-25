package pgready

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq" // PostgreSQL driver
)

var (
	dbConnString  string
	retryAttempts int
	retryInterval time.Duration
)

func init() {
	// Define command-line flags
	flag.StringVar(&dbConnString, "db-url", "", "PostgreSQL connection string")
	flag.IntVar(&retryAttempts, "retry-attempts", 20, "Number of connection retry attempts")
	flag.DurationVar(&retryInterval, "retry-interval", time.Second, "Interval between retry attempts")
	flag.Parse()

	if dbConnString == "" {
		// Try to read the database URL from the environment
		dbConnString = os.Getenv("DATABASE_URL")
	}

	// Show an error message and usage information if the database URL is missing
	if dbConnString == "" {
		fmt.Println("Error: database URL is required")
		flag.Usage()
		os.Exit(1)
	}
}

func main() {
	for i := 0; i < retryAttempts; i++ {
		db, err := sql.Open("postgres", dbConnString)
		if err != nil {
			fmt.Printf("Failed to open connection: %v\n", err)
			time.Sleep(retryInterval)
			continue
		}

		err = db.Ping()
		db.Close()
		if err == nil {
			fmt.Println("PostgreSQL is ready for connections.")
			os.Exit(0)
		}

		fmt.Printf("PostgreSQL not ready yet: %v\n", err)
		time.Sleep(retryInterval)
	}

	fmt.Println("PostgreSQL did not become ready within the given time period.")
	os.Exit(1)
}
