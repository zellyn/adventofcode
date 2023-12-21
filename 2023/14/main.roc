app "hello"
    packages {
        pf: "https://github.com/roc-lang/basic-cli/releases/download/0.7.0/bkGby8jb0tmZYsy2hg1E_B2QrCgcSTxdUlHtETwm5m4.tar.br",
    }
    imports [
        pf.Stdout,
        pf.File,
        pf.Task.{ Task, await },
        pf.Path.{ Path },
        pf.Stderr,
        pf.Arg,
        Charmap,
    ]
    provides [main] to pf

main : Task {} *
main =
    run
    |> Task.onErr handleErr

Error : [
    FailedToReadArgs,
    FailedToReadFile Path Str,
]

handleErr : Error -> Task {} *
handleErr = \err ->
    usage = "roc main.roc -- example"

    errorMsg = when err is
        FailedToReadArgs -> "Failed to read command line arguments, usage: \(usage)"
        FailedToReadFile path errString -> "Failed to read file \(displayPath path): \(errString)"

    Stderr.line "Error: \(errorMsg)"

run : Task {} Error
run =
    { inputPath } <- await readArgs

    {} <- Stdout.line "Reading file \(displayPath inputPath):" |> await

    contents <- readFile inputPath |> await

    m = Charmap.parse contents

    {} <- Stdout.line (Charmap.toString m "/") |> await

    Stdout.line "Done"


readFile : Path -> Task Str [FailedToReadFile Path Str]
readFile = \path ->
    File.readUtf8 path
    |> Task.mapErr \err ->
        when err is
            FileReadErr _ readErr -> FailedToReadFile path (File.readErrToStr readErr)
            FileReadUtf8Err _ _ -> FailedToReadFile path "bad utf-8 found"


readArgs : Task { inputPath: Path } [FailedToReadArgs]_
readArgs =
    args <- Arg.list |> Task.attempt

    when args is
        Ok ([ _, arg ]) ->
            Task.ok { inputPath: Path.fromStr arg }
        _ ->
            Task.err FailedToReadArgs


displayPath : Path -> Str
displayPath = \path ->
    "\"\(Path.display path)\""
