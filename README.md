# 🏦 Wallet API - Golang, Gin-Gonic & SQLite

Welcome to Wallet API! 🔐💳 This API allows users to create accounts, log in securely using JWT and bcrypt, and perform transactions like withdrawals, deposits, and transfers between accounts.

## 🚀 Features
- 🔑 User registration with encrypted passwords using **bcrypt**
- 🔐 Secure authentication with **JWT**
- 💰 Deposit money into an account
- 📤 Withdraw funds
- 🔄 Transfer between accounts
- 📜 Transaction history tracking
- 🛠 Built with **Golang**, **Gin-Gonic**, and **SQLite**

## 🏗 Installation & Setup

### 1️⃣ Clone this repository:
git clone https://github.com/mufasa-dev/Wallet-flow-api-in-Golang.git

### 2️⃣ Navigate to the project folder:

cd wallet-api
### 3️⃣ Install dependencies:

go mod tidy
### 4️⃣ Run the server:
go run main.go

## ⚙️ API Endpoints
### 🔑 Authentication
Method	Endpoint	Description

POST	/sigup	Create a new user

POST	/sigin	User authentication

### 💰 Transactions
Method	Endpoint	Description

POST	/api/v1/deposit	Deposit money into account

POST	/api/v1/withdraw	Withdraw funds from account

POST	/api/v1/transfer	Transfer money between users
### 📜 Account Info
Method	Endpoint	Description

GET	/statement	View transaction history
## 🔧 Environment Variables
Create a .env file and define the following:

JWT_SECRET=your_secret_key_here
## 🛠 Built With
🔷 Golang - Efficient backend language

🔥 Gin-Gonic - Fast & lightweight framework

📂 SQLite - Embedded database for simplicity

🔐 JWT & Bcrypt - Secure authentication & password hashing

## 🤝 Contributing
Feel free to submit issues, suggestions, or pull requests to improve this project! 🛠🚀

## 📜 License
This project is under the MIT License.
