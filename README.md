## Indian Poker
#### Fall 2019 - CSCI 6221 
#### Course Project
Ruocheng Shan

Guodong Xie

Kai Wang

Ruikai Zhou

##1. Install
### 1.1 Install GTK
For linux
```cassandraql
apt-get install libgtk2.0-dev libglib2.0-dev libgtksourceview2.0-dev
```
For Windows
1. install msys2
2. open msys2 terminal
3. install gtk with following commands
```cassandraql
pacman -S mingw-w64-x86_64-gtk2
pacman -S mingw-w64-x86_64-devhelp
pacman -S mingw-w64-x86_64-toolchain base-devel
```

### 1.2 Get GTK-GO
```cassandraql
go get github.com/mattn/go-gtk/gtk
```
### 1.3 Get Project
```cassandraql
git clone https://github.com/RuochengShan/repo-gwu-csci_6221.git
```

make sure your project and the gtk-go are in go_path
## 2.Run
2.1 start the server
```cassandraql
cd your_go_path/src/repo-gwu-csci_6221/server
go build ./
server
```
2.2 start the client
```cassandraql
cd your_go_path/src/repo-gwu-csci_6221/
go run main.go
```

