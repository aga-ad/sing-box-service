# sing-box-service
Windows Service wrapper for [sing-box v1.12.12](https://github.com/SagerNet/sing-box).

## Build
You need Go version 1.25.3 or later.

Just run ```make``` or ```go build``` command from Makefile.

## Usage
1. Build ```sing-box-service.exe``` or download release.

2. Make your own ```config.json``` for sing-box. See https://sing-box.sagernet.org/configuration/ for documentation. Specify only absolute paths, for example for the log file.

3. Place ```config.json``` and ```sing-box-service.exe``` in any folder (let's say ```C:\path\to\singboxdir\```). 

4. Run ```sing-box-service.exe``` and check that the program is working properly.

5. Start Command Prompt as Administrator and run with correct binpath:
```
sc create "sing-box-service" binpath= "C:\path\to\singboxdir\sing-box-service.exe" DisplayName= "sing-box-service" start= auto
net start sing-box-service
```



## Disclaimer
This project is an independent work based on the [sing-box](https://github.com/SagerNet/sing-box), 
which is licensed under the GNU GPLv3.  
The authors of the original project are not affiliated with, do not endorse, 
and are not responsible for this project.
