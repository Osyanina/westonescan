# westonescan
An asset scanning tool based on the GO language for the aid of penetration test engineers .  
Its main function is to identify whether the target assets are alive in batches, and to perform port external network open scanning of the surviving assets .
# Installation & Usage 
Install:  
Set GOPATH, this environment variable points to your project directory.    
Remember to put the iprange folder in your src/github environment directory .  
git clone https://github.com/Osyanina/westonescan.git  
cd westonescan  
go build westonescan.go  

Example :  

westonescan.exe 192.168.0.0/16 1-65535  
westonescan.exe 192.168.0.1-254 80,3389  
westonescan.exe 192.168.0.1 3389
