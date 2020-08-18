function isElectron() {
    if (typeof window !== 'undefined' && typeof window.process === 'object' && window.process.type === 'renderer') {
        return true;    // Renderer process
    }

    if (typeof process !== 'undefined' && typeof process.versions === 'object' && !!process.versions.electron) {
        return true;    // Main process
    }

    if (typeof navigator === 'object' && typeof navigator.userAgent === 'string' && navigator.userAgent.indexOf('Electron') >= 0) {
        return true;    // Detect the user agent when the `nodeIntegration` option is set to true
    }

    return false;
}

if (isElectron()) {
    const { app } = require('electron').remote
    process.env.ELECTRONVUESITE = app.getAppPath()
    process.env.ELECTRONVUESITE = process.env.ELECTRONVUESITE.replace("\\app.asar", "")

    var child = require('child_process').execFile
    let executablePath = "";
    if (process.platform === "win32") {
        executablePath = __dirname + '/server.exe'
    } else if (process.platform === "linux") {
        executablePath = __dirname + '/server'
    } else if (process.platform === "darwin") {
        executablePath = __dirname + '/servermacos'
    }

    child(executablePath, (err, data) => {
        if (err) {
            console.error(err)
            return
        }

        console.log(data.toString())
    })
}
