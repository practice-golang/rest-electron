# Simple CRUD app using Go([ql](https://github.com/cznic/ql)), Vue(+Electron)
For CRUD Practice  
Put [rest-client](https://github.com/practice-golang/rest-client) in Electron  
Need binary as dist/server.exe which from public/server.exe at [rest-ql-crud](https://github.com/practice-golang/rest-ql-crud)  

[Demo](https://www.dropbox.com/s/qxi5hwji6o1h76b/ElectronVue-0.0.1-win.zip?dl=1)  

## Build, Run except Electron
### Prepare js dependancies
```powershell
npm install
```

### Run Vue.js dev server for web browser
```powershell
./public/server.exe
npm run vue
```

### Build only Vue.js source, not electron
```powershell
npm run build:vue
```

## Build for Electron
### REST server
```powershell
`./public/server.exe`
```

### Build vue.js and run electron
```powershell
npm run electron
```

### Build for Windows x64
```powershell
npm run build:win64
```

### Build for Linux - Not yet
```powershell
npm run build:linux
```

### Build for Mac - Not yet
```powershell
npm run build:osx
```

## ql.db
* project_root/ql.db : Using for `Run Vue.js dev server for web browser`
* dist/ql.db : Using for else
* electron-dist/resources/ql.db : It will be created when insert a new data after `build`
