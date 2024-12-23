import { createWebSocketConnection, WsEventType } from './ws'
import { clearCanvas, drawLetter, get2dContext, getCanvas } from './renderer'

function main() {
  const conn = createWebSocketConnection()

  const canvas = getCanvas()
  const ctx = get2dContext(canvas)
  ctx.translate(0, canvas.height)
  ctx.scale(1, -1)

  conn.subscribe(WsEventType.snakePosition, (data) => {
    clearCanvas(canvas)
    for (const px of data.payload) {
      drawLetter(ctx, { x: px.X * 10, y: px.Y * 10 })
    }
  })
}

main()
