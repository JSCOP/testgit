module github.com/JSCOP/testgit/tree/test/hello

go 1.17

replace github.com/JSCOP/testgit/tree/test/greetings => ../greetings

require github.com/JSCOP/testgit/tree/test/greetings v0.0.0-00010101000000-000000000000
