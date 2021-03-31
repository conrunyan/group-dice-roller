const input = document.querySelector("#textarea");
const messages = document.querySelector("#messages");
const username = document.querySelector("#username");
const send = document.querySelector("#send");

const dice = ["d4", "d6", "d8", "d10", "d12", "d20", "d100"];

const url = "ws://" + window.location.host + "/ws";
const ws = new WebSocket(url);

ws.onmessage = function (msg) {
  insertMessage(JSON.parse(msg.data));
};

function setOnClicks() {
  dice.map((die) => {
    let button = document.querySelector(`#${die}`);
    button.onclick = () => {
      const message = {
        dieType: die,
      };
      console.log({message})
      ws.send(JSON.stringify(message));
    };
  });
}

function insertMessage(messageObj) {
  const message = document.createElement("div");
  message.setAttribute("class", "chat-message");
  message.textContent = `${messageObj.username}: ${messageObj.content}`;

  messages.appendChild(message);
  messages.insertBefore(message, messages.firstChild);
}

document.body.onload = () => {
  console.log("Loaded!");
  setOnClicks();
};
