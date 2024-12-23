export function getCanvas(): HTMLCanvasElement {
  const canvas = document.getElementById("canvas")
  if (!(canvas instanceof HTMLCanvasElement)) {
    throw new Error("expected canvas html element")
  }
  return canvas
}

export function get2dContext(canvas: HTMLCanvasElement): CanvasRenderingContext2D {
  const ctx = canvas.getContext("2d")
  if (!ctx) {
    throw new Error("expected 2d context")
  }
  return ctx
}

export function drawLetter(
  ctx: CanvasRenderingContext2D,
  pos: { x: number, y: number },
) {
  ctx.fillStyle = '#000000'
  ctx.font = "30px serif";
  ctx.fillText('*', pos.x, pos.y);
}

export function clearCanvas(canvas: HTMLCanvasElement) {
  const ctx = get2dContext(canvas)
  ctx.clearRect(0, 0, canvas.height, canvas.height)
}
