# App practice & study using Go([ql](https://github.com/cznic/ql)) & Vue(+Electron)

Used [rest-client](https://github.com/practice-golang/rest-client), [rest-ql-crud](https://github.com/practice-golang/rest-ql-crud)

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

* For linux, macos : I will never do them. Try, try!! Good luck!! ^-^;

### Build vue.js and run electron
```powershell
npm run electron
```

### Build for Windows
* When build, `dist` folder will be created
* `server.exe` will be copied to `dist` folder and will be included in `app.asar` file
```powershell
npm run build:win
```

## ql.db
* project_root/ql.db : Using for `Run Vue.js dev server for web browser`
* `ql.db` location: `built_app_path/resources` - See `package.json`
    * ~~When you insert first data, this file will be created to above folder~~

----
LICENSE : Public domain.
