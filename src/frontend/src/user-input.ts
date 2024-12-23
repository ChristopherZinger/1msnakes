export enum Direction {
  N = 0,
  E = 1,
  S = 2,
  W = 3,
}

enum ValidKeyInputs {
  W = 'w',
  A = 'a',
  S = 's',
  D = 'd',
  J = 'j',
  K = 'k',
  L = 'l',
  H = 'h'
}

const DirectionToValidKeyInputs: Record<Direction, ValidKeyInputs[]> = {
  [Direction.N]: [ValidKeyInputs.W, ValidKeyInputs.K],
  [Direction.E]: [ValidKeyInputs.D, ValidKeyInputs.L],
  [Direction.W]: [ValidKeyInputs.A, ValidKeyInputs.H],
  [Direction.S]: [ValidKeyInputs.S, ValidKeyInputs.J],
}

export function addSnakeMoveHandler(cb: (direction: Direction) => void) {
  let previousDirection = Direction.N

  window.addEventListener('keydown', (e) => {
    if (
      e instanceof KeyboardEvent
      && isValidInputKey(e.key)
      && !DirectionToValidKeyInputs[previousDirection].includes(e.key)) {
      const newDirection = Number(
        Object.entries(DirectionToValidKeyInputs)
          .find(([, v]) => v.includes(e.key as ValidKeyInputs))
          ?.map(([k]) => k)[0]
      )

      if (isDirection(newDirection)) {
        // TODO: confirmation about previous directin
        // should be returned from the backend
        previousDirection = newDirection
        cb(newDirection)
      } else {
        console.error('new direction incorrect value', newDirection)
      }
    }
  })
}

function isValidInputKey(key: unknown): key is ValidKeyInputs {
  return Object.values(ValidKeyInputs).includes(key as ValidKeyInputs)
}

function isDirection(d: unknown): d is Direction {
  return Object.values(Direction).includes(d as Direction)
}
