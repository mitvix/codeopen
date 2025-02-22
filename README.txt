Very Simple golang script to run vscode flatpak from shell using command "code"

To run just build and copy to your path

git clone https://github.com/mitvix/codeopen
go build -o code main.go
sudo cp code /usr/local/bin/
code README.txt

Just it! 
