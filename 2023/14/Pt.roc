interface Pt
    exposes [
      n, s, e, w,
      add,
    ]
    imports []

Pt : (I64, I64)

n : Pt
n = (0, -1)

s : Pt
s = (0, 1)

e : Pt
e = (1, 0)

w : Pt
w = (-1, 0)

add : Pt, Pt -> Pt
add = \(x1, y1), (x2, y2) -> (x1+x2, y1+y2)
