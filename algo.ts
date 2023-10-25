let s: string = ''

for (let i = 1; i < 6; i++) {
  s += algo(i)
}

function algo(value: number): number {
  if (value === 0) {
    return 2
  } else if (value === 1) {
    return 1
  } else {
    return algo(value - 1) + algo(value - 2)
  }
}

console.log(s)  