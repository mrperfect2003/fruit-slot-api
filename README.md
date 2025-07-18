# fruit-slot-api
A Golang-based API that simulates a slot machine with random fruit spins. Built using Fiber, this project returns random fruit combinations, checks for wins, supports 10x spin mode, logs outcomes, and includes unit tests.

## Fruit Slot API 

A simple backend slot machine game built in **Go** using **Fiber**. It simulates a fruit-based spin machine and returns random results. If at least two fruits match, it's a win!

---

##  Features
- **GET /play**: Returns 3 random fruits.
- **GET /play/10**: Spins 10 times and counts how many were wins.
- Secure random generation using `crypto/rand`
- Modular folder structure with controller, routes, repo, config
- Environment configuration via `.env`
- Console logging for requests and results
- Unit tests for core win-checking logic

---

##  Tech Stack
- **Go 1.20+**
- [Fiber v2](https://github.com/gofiber/fiber)
- [godotenv](https://github.com/joho/godotenv)

---

##  Project Structure
```
fruit-slot-api/
├── config/                 # Fruit list configuration
├── controller/             # API handlers
├── model/                  # Response models
├── repo/                   # Logic and unit tests
├── routes/                 # Route definitions
├── .env                    # Environment variables
├── main.go                 # Entry point
├── go.mod
```

---

##  Getting Started

### 1️ Clone the repository
```bash
git clone https://github.com/yourusername/fruit-slot-api.git
cd fruit-slot-api
```

### 2️ Install dependencies
```bash
go mod tidy
```

### 3️ Run the server
```bash
go build -o play
./play
```

### 4️ Test in Postman or browser
- `GET http://localhost:3000/play`
- `GET http://localhost:3000/play/10`

---

##  Run Unit Tests
```bash
go test ./repo
```

---

##  Sample Response
###  Win
```json
{
  "fruits": ["Lemon", "Lemon", "Grape"],
  "message": "You win!"
}
```

###  Try Again
```json
{
  "fruits": ["Cherry", "Orange", "Watermelon"],
  "message": "Try again!"
}
```

###  Multi Spin
```json
{
  "results": [
    {"fruits": ["Lemon", "Lemon", "Cherry"], "message": "You win!"},
    {"fruits": ["Orange", "Pineapple", "Watermelon"], "message": "Try again!"}
  ],
  "win_count": 4
}
```

---

##  Demo Video
 [Google Drive Link](https://drive.google.com/link)

---

##  Author
**Keshav Raj**  
[Portfolio](https://keshav-raj.web.app/) • [GitHub](https://github.com/mrperfect2003) • [LinkedIn](https://www.linkedin.com/in/keshavraj18)

---

##  License
This project is open source under the [MIT License](LICENSE).
