import { getSerializedSnakeMove } from "./snakeMoveProto"
import * as sm from "./generated/protos/snakemove"

enum WsEventType {
  sendMessage = 'send_message'
}

type WsEvent = {
  type: WsEventType,
  payload: unknown
}

class Connection {
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

  public subscribe(evName: WsEventType, cb: (ev: WsEvent) => unknown) {
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
        console.error("expected event type");
        return;
      }

      const subs = this.subscribers[ev.type]

      subs?.forEach(s => s(ev))
    }
  }
}

export function testWebSocketConnection() {
  const url = new URL("ws://" + document.location.host + "/ws")
  const conn = new Connection(url)

  conn.subscribe(WsEventType.sendMessage, (ev) => {
    console.log(ev)
  })

  const interval = setInterval(function() {
    conn.sendEvent({
      type: WsEventType.sendMessage,
      payload: getSerializedSnakeMove(sm.main.Direction.E)
    })
  }, 1000)
  return interval
}

