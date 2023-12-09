app "hello"
    packages { pf: "https://github.com/roc-lang/basic-cli/releases/download/0.7.0/bkGby8jb0tmZYsy2hg1E_B2QrCgcSTxdUlHtETwm5m4.tar.br" }
    imports [
        pf.Stdout,
        # "input" as input : Str,
        "example_input" as example : Str,
    ]
    provides [main] to pf

# This involved a _lot_ of copying from https://github.com/ostcar/aoc2023/blob/main/days/day05.roc
# since it's my first Roc program.

main =
    Stdout.line (Str.joinWith (parseInput example) "\n-------------\n")

Span : { source : U64, destination : U64, offset : U64 }
Mapping : List Span

# parseInput : Str -> (List U64, List Mapping)
parseInput : Str -> List Str
parseInput = \input ->
    inputGroups =
        input
        |> Str.trim
        |> Str.split "\n\n"

    # dbg inputGroups
    seeds =
        List.first
        |> Str.split " "
    inputGroups
