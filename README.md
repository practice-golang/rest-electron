# App practice & study using Go([ql](https://github.com/cznic/ql)) & Vue(+Electron)

Used [rest-client](https://github.com/practice-golang/rest-client), [rest-ql-crud](https://github.com/practice-golang/rest-ql-crud)

[Demo](https://www.dropbox.com/s/7o9cv2zfnobrvkm/ElectronVue-0.0.2-win.zip?dl=1)  

## Build, Run except Electron
### Prepare js dependancies
```powershell
npm install
```

### Run Vue.js dev server for web browser
```powershell
$ ./server_prepare/server.exe

# Run on another console
$ npm run vue
```

### Build only Vue.js source, not electron
```powershell
$ npm run build:vue
```

## Build or run for Electron

* For linux, macos : I will never do them. Try, try!! Good luck!! ^-^

### Build vue.js and run electron
```powershell
$ npm run electron
```

### Build for Windows
* When build, `dist` folder will be created
* `server.exe` will be copied to `dist` folder and will be included in `app.asar` file
```powershell
$ npm run build:win
```

## ql.db
* project_root/ql.db : `Run Vue.js dev server for web browser` use this.
* Path for electron dist : `built_app_path/resources` - See `package.json`
    * When you insert first data, this file will be created to the path.
    * Not work with `portable` (eg. `electron-builder --win portable --x64`) option of electron-builder.

----
LICENSE : Public domain.
