package main
import (
	"context"
	"fmt"
	"log"
	"net/http"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
func main(){
	// THE CONNECTION: We are telling Go to look for a computer named "my-database"
	rdb := redis.NewClient(&redis.Options{
		Addr: "my-database:6379", 
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Increment a counter in the database every time someone visits
		visits, err := rdb.Incr(ctx, "page_visits").Result()
		if err != nil {
			fmt.Fprintf(w, "Database connection failed: %v", err)
			return
		}
		fmt.Fprintf(w, "Day 3: You have visited this page %d times.", visits)
	})

	fmt.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))

}