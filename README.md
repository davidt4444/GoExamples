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
cd ..
go mod init github.com/davidt4444/goexamples/bcs

go get github.com/gorilla/mux



 cd mysql
 go install
 (The go.mod here could be used for the gopls install)
 cd ../

