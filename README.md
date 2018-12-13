# Simple CRUD app using Go([ql](https://github.com/cznic/ql)), Vue(+Electron)
For CRUD Practice  
Put [rest-client](https://github.com/practice-golang/rest-client) in Electron  
Need binary as dist/server.exe from [rest-ql-crud](https://github.com/practice-golang/rest-ql-crud)  

[Demo](https://www.dropbox.com/s/qxi5hwji6o1h76b/ElectronVue-0.0.1-win.zip?dl=1)  

## Build, Run
```powershell
# Prepare js dependancies
npm install

# Run dev server for browser, not electron
npm run dev

# Build only vue.js source, not electron
npm run build:vue

# Below are build for electron
# REST server - dist/server.exe which compiled from rest-ql-crud should be placed

# Build vue.js and run electron
npm run start

# Build for Windows x64
npm run build:win64

# Build for Linux - Not yet
npm run build:linux

# Build for Mac - Not yet
npm run build:osx
```
