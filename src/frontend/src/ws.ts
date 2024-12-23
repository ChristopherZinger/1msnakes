import { Direction } from "./user-input"

export enum WsEventType {
  sendMessage = 'send_message',
  snakePosition = 'snake_position'
}

type WsEvent = {
  type: WsEventType.snakePosition,
  payload: { X: number, Y: number }[]
} | {
  type: WsEventType.sendMessage,
  payload: unknown
}

export class Connection {
  private conn: WebSocket
  private subscribers: Partial<Record<WsEvent['type'], Function[]>>

  constructor(url: string | URL) {
    if (!window["WebSocket"]) {
      throw new Error('expected_window_WebSocket')
    }

    this.conn = new WebSocket(url)
    this.rounteEvent()

    this.subscribers = {}
  }

  public subscribe<T extends WsEventType>(evName: T, cb: (ev: Extract<WsEvent, { type: T }>) => unknown) {
    if (this.subscribers[evName]) {
      this.subscribers[evName].push(cb)
    } else {
      this.subscribers[evName] = [cb]
    }
  }

  public sendEvent(ev: WsEvent) {
    this.conn.send(JSON.stringify(ev));
  }

  private rounteEvent() {
    this.conn.onmessage = (message) => {
      const data = JSON.parse(message.data);
      const ev: WsEvent = data; // TODO: this should be a type error - "unsave any"; check linter settings

      if (ev.type === undefined) {
        console.error("expected event type, got", data);
        return;
      }

      const subs = this.subscribers[ev.type]

      subs?.forEach(s => s(ev))
    }
  }
}

export function createWebSocketConnection() {
  const url = new URL("ws://" + document.location.host + "/ws")
  const conn = new Connection(url)

  return conn;
}

export function postSnakeNextMove(conn: Connection, direction: Direction) {
  conn.sendEvent({
    type: WsEventType.sendMessage,
    payload: { direction }
  })
}
