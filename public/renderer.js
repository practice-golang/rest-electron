// This file is required by the index.html file and will
// be executed in the renderer process for that window.
// All of the Node.js APIs are available in this process.

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

function loadScript(url) {
    var script = document.createElement('script');
    script.src = url;
    document.getElementsByTagName('head')[0].appendChild(script);
}

if (isElectron()) {
    const { app } = require('electron').remote;
    process.env.ELECTRONVUESITE = app.getAppPath()
    // console.log(process.env.ELECTRONVUESITE)
    process.env.ELECTRONVUESITE = process.env.ELECTRONVUESITE.replace("\\app.asar", "")
    // console.log(process.env.ELECTRONVUESITE)

    var child = require('child_process').execFile;
    let executablePath = "";
    if (process.platform === "win32") {
        executablePath = __dirname + '/server.exe';
    } else if (process.platform === "linux") {
        executablePath = __dirname + '/server';
    } else if (process.platform === "darwin") {
        executablePath = __dirname + '/servermacos';
    }

    child(executablePath, (err, data) => {
        if (err) {
            console.error(err);
            return;
        }

        console.log(data.toString());
    });

    // require(__dirname + "/build.js")
} else {
    // loadScript(__dirname + '/build.js');
}
