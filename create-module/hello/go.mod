module github.com/burkocyigit/learning-golang-yavuzlar/create-module/hello

go 1.21.3

replace github.com/burkocyigit/learning-golang-yavuzlar/create-module/greetings => ../greetings

require github.com/burkocyigit/learning-golang-yavuzlar/create-module/greetings v0.0.0-00010101000000-000000000000
