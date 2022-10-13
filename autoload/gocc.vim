let s:cmd = expand('<sfile>:h:h:gs!\\!/!') . '/gocc/gocc'

let s:root = expand('<sfile>:p:h:h')
let s:go_dir = s:root 

if !filereadable(c:cmd)
    call system(printf('cd %s && go get -d && go build', s:go_dir))
endif

function! gocc#run
