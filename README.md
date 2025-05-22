# ballroom-go






## Project Structure
```
ballroom-go/
├── main.go               # Entry point; sets up the Gin router and starts the server
├── models.go             # Request/response structs; 
├── simulation.go         # Core logic for calculating average partners
├── handlers.go           # HTTP route logic
├── go.mod / go.sum       # Dependency files
└── README.md             # Documentation
```

## Run the app
```
go run .
```

## Build and run the app

```
go build -o ballroom-go
./ballroom-go
```

## Test with curl
```
curl -X POST http://localhost:8000/calculate-partners \
  -H "Content-Type: application/json" \
  -d '{"total_leaders":2,"total_followers":2,"dance_styles":["Waltz"],"leader_knowledge":{"1":["Waltz"],"2":["Waltz"]},"follower_knowledge":{"A":["Waltz"],"B":["Waltz"]},"dance_duration_minutes":60}'
```



docker build -t ballroom-go .


docker run -p 8000:8000 ballroom-go

