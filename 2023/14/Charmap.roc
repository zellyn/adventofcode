interface Charmap
    exposes [
      parse,
      minmax,
      toString
    ]
    imports [
    ]

Point : (I64, I64)

parse : Str -> Dict Point Str
parse = \input ->
    input
    |> Str.trim
    |> Str.split "\n"
    |> List.map Str.graphemes
    |> List.mapWithIndex \list, y ->
        list |> List.mapWithIndex \char, x ->
            ((Num.toI64 x, Num.toI64 y), char)
    |> List.join
    |> Dict.fromList

pointsMap : Point, Point, (I64, I64 -> I64) -> Point
pointsMap = \(x1, y1), (x2, y2), f ->
    (f x1 x2, f y1 y2)

minmax : Dict Point Str -> Result ( Point, Point) [EmptyMap]
minmax = \dict ->
  when Dict.keys dict is
    [] -> Err EmptyMap
    [head, .. as tail] ->
        List.walk tail ( head, head ) \(min, max), point ->
            ( pointsMap min point Num.min, pointsMap max point Num.max )
        |> Ok

 toString : Dict Point Str, Str -> Str
 toString = \dict, default ->
    when minmax dict is
        Err EmptyMap -> ""
        Ok ((minx, miny), (maxx, maxy)) ->
            List.range { start: At miny, end: At maxy}
            |> List.map \y ->
                List.range { start: At minx, end: At maxx}
                |> List.map \x ->
                    Dict.get dict (x,y) |> Result.withDefault default
                |> Str.joinWith ""
            |> Str.joinWith "\n"
