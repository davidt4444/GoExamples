# GoExamples

# Install Go
https://go.dev/dl/

# Install the extension
open extensions in vscode and search, then install
After that, when you open a go.mod, the gopls will install itself 
# install mysql module

mkdir bcs
cd bcs
mkdir service
go mod init github.com/davidt4444/goexamples/bcs/service
go get github.com/go-sql-driver/mysql  
// Push the changes to the repo from goexamples base
cd ..
go mod init github.com/davidt4444/goexamples/bcs
go get github.com/gorilla/mux
go get github.com/davidt4444/goexamples/bcs/service@latest
in the go.mod that was created add
replace github.com/davidt4444/goexamples/bcs/service => ./service
to the bottom of the file to point to the local service


 cd mysql
 go install
 (The go.mod here could be used for the gopls install)
 cd ../

