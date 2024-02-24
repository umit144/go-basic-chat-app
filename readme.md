# Chat Application

This project consists of a simple WebSocket server written in Go and a client written in Vanilla JavaScript. This application allows users to join a chat room and communicate with each other via messages.

## Installation

### Frontend (Client)
First, navigate to the `client` directory:
```bash
cd client
```
```bash
npm install
```
```bash
npm run dev
```

You can access the chat application by going to http://localhost:3000 in your browser.

------------------------------------------------------------------------------------------------------------------------

### Backend (Server)
Navigate to the root directory of the project:
```bash
cd server
```
Start the server by running:
```bash
go run main.go
```

------------------------------------------------------------------------------------------------------------------------

## Usage

After starting both the server and the client, open your web browser and go to http://localhost:3000.

Enter your username and the message you want to send in the respective input fields.

Click the "Send" button to send the message.

You'll see your message along with messages from other users displayed in the chat room.