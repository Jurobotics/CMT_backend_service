const websocket = new WebSocket("ws://localhost:8001");

websocket.onopen = () => {};

export enum messageType {
  MAKE = "MAKE",
  STOP = "STOP",
}

interface item {
  servoId: number;
  pumps: number;
}

interface message {
  type: messageType;
  content: Array<item>;
}

export function sendMessage(message: message): void {
  websocket.send(JSON.stringify(message));
}
