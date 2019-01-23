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
./server_prepare/server.exe
npm run vue
```

### Build only Vue.js source, not electron
```powershell
npm run build:vue
```

## Build or run for Electron
* Considered MS-Windows only. If you use other OS, you should edit `slash` of copy command in `package.json`

### Build vue.js and run electron
```powershell
npm run electron
```

### Build for 32bit OS
```sh
Never
```

### Build for Windows
* When build, `dist` folder will be created
* `server.exe` will be copied to `dist` folder
```powershell
npm run build:win
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
* `ql.db` location: `built_app_path/resources`
    * When you insert first data, this file will be created to above folder
    * If you want to include `ql.db` which is already created to the binary, copy it to above folder yourself
