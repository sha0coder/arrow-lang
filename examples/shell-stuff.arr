; execute commands with ! 


; print all go files on current folder
!ls -l -> split '\n' -> $files
$files =>
    -> $file
    [$file =~ "\.go$"]
        "go file: $file" -> print



; show content of all the homes of all users
/etc/passwd -> split '\n' -> $pwds
$pwds =>
    -> $pwd -> split ':' -> $parts -> len -> $l
    [$l < 5]
        cont
    $parts[5] -> $home -> print
    !ls -l $home -> print

