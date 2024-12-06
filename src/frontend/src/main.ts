import { testWebSocketConnection } from './ws'

function main() {
  testWebSocketConnection()
  testCanvas()
}

function testCanvas() {
  const canvas = document.getElementById("canvas")
  if (!(canvas instanceof HTMLCanvasElement)) {
    console.error("expected canvas element")
    return
  }

  const ctx = canvas.getContext("2d")
  if (!ctx) {
    console.error("expected 2d context")
    return
  }
  drawTwoRectangles(ctx)
}

function drawTwoRectangles(ctx: CanvasRenderingContext2D) {
  ctx.fillStyle = "rgb(200 0 0)";
  ctx.fillRect(10, 10, 50, 50);

  ctx.fillStyle = "rgb(0 0 200 / 50%)";
  ctx.fillRect(30, 30, 50, 50);

  ctx.clearRect(50, 50, 50, 50)
}

main()
