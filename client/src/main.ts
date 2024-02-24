import './style.css'

type Message = { username: string; message: string; };

const usernameInput: HTMLInputElement = document.getElementById('username') as HTMLInputElement;
const messageInput: HTMLInputElement = document.getElementById('message') as HTMLInputElement;
const chatList: HTMLUListElement = document.getElementById('chat') as HTMLUListElement;

const socket: WebSocket = new WebSocket("ws://localhost:8080/ws");

socket.onmessage = function (event: MessageEvent) {
  const message: Message = JSON.parse(event.data);
  const listItem: HTMLLIElement = document.createElement('li');
  listItem.textContent = `${message.username}: ${message.message}`;
  listItem.classList.add('message-item');
  chatList.appendChild(listItem);
};

window.sendMessage = (): void => {
  const username: string = usernameInput.value;
  const message: string = messageInput.value;

  if (username && message) {
    const payload: Message = {
      username: username,
      message: message
    };
    socket.send(JSON.stringify(payload));
    messageInput.value = '';
  } else {
    alert('Username and message are required!');
  }
}