
/etc/passwd -> lines -> $users
$users =>
    -> $line -> split ":"
    $_[5] -> $home 
    [$home_sz > 0]
        !stat $home -> lines -> grep "Access" -> $out
        [$out_sz > 0]
            $out[0] -> extract "\(([^)]+)\)" -> $perms
            "$home$:    $perms[0]" -> print
