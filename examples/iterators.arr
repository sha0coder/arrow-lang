; loops made easy

;iterate from 0 to 3 (interval variable $_ is setted always on arrow lang)
(3) =>
    "3 times" -> print

; lets get the magic var $_ on $i
(10) =>
    -> $i
    ;same as $_ -> $i
    ;also can do: -> print

    ;display
    "iteartion number $i" -> print


; from 0 to n
4 -> $n
($n) =>
    "4 times" -> print


; commented just in case:
;=>
;    "endless loop, press control+C" -> print


/etc/passwd -> $users
$users =>
    -> $user -> print