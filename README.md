# westonescan
An asset scanning tool based on the GO language for the aid of penetration test engineers 
# How to use?  
Install:  
Set GOPATH, this environment variable points to your project directory.    
Remember to put the iprange folder in your src/github environment directory .  

Example :  

westonescan.exe 192.168.0.0/16 1-65535  
westonescan.exe 192.168.0.1-254 80,3389  
westonescan.exe 192.168.0.1 3389
