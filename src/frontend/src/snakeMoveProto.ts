import * as sm from "./generated/protos/snakemove"

export function getSerializedSnakeMove(d: sm.main.Direction) {
  const s = new sm.main.SnakeMove({
    direction: d
  })
  const result = s.serializeBinary()
  return result
}
