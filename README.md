# arrow-lang
productivity focused programming language, that generates use and throw python scripts.


Print:

    "hello world" -> print
    "hello world\n" -> out
    "input your name: " -> out -> in -> $name
    "your noame is $name" -> print


Reading file:

    ./myfile.txt -> $content            
    /etc/passwd -> $pwds

Writing file:

    $data -> ./file.csv        

Appending file:

    $data -> append 'file.csv'

Urls:

    http://test.com/ -> $html
    http://test.com/ -> ./html.txt

    or also:

    http://test.com/ -> $html -> ./html.txt


Infinite loop:

    =>
        "test" -> print

Iterate array:

    /etc/passwd -> lines -> $pwds
    $pwds =>
        print

    or also:

    $pwds =>
        $_ -> print

    or also

    $pwds =>
        -> print

Iterate 3 times:

    (3) =>
        $_ +1 -> $num
        "iteration num $num" -> print


loop n times:

    ($n) =>
        -> $i

loop asynchronously:

    ($n) :=>
        -> $i
    <=:


Check size:

    "this is a test" -> $var
    $var_sz -> print

    list 1,2,3 -> $arr
    $arr_sz -> print


System shell:

    !ls -> $out -> ./files.txt

    "/etc/" -> $etc -> !ls $etc -> ./config_files.txt
