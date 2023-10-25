let s: string = ''

for (let i = 1; i < 6; i++) {
  s += algo(i)
}

function algo(n: number): number {
  if (n === 0) {
    return 2
  } else if (n === 1) {
    return 1
  } else {
    return algo(n-1) + algo(n-2)
  }
}

console.log(s)  