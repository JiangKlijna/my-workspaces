
let fs = require("fs");
let path = require("path");


// foreach dir
function travelDir(dir, callback) {
    let index = 0;
    fs.readdirSync(dir).forEach(function (file) {
        let pathname = path.join(dir, file);
        if (fs.statSync(pathname).isDirectory()) {
            travelDir(pathname, callback);
        } else {
            callback(pathname, index++);
        }
    });
}

function work(dir, fun) {
    travelDir(dir, function (pathname, index) {
        let basename = path.basename(pathname)
        let pointIndex = basename.lastIndexOf(".")
        let name = basename.substring(0, pointIndex)
        let ext = basename.substr(pointIndex)
        if (pointIndex < 0) {
            name = basename;
            ext = "";
        }

        let newbasename = fun(name, ext, index);
        if (newbasename) {
            let newpathname = path.join(path.dirname(pathname), newbasename)
            // console.log(pathname, newpathname)
            try {
                fs.renameSync(pathname, newpathname);
                console.log(pathname, "rename success. newname is ", newbasename);
            } catch (err) {
                console.log(pathname, "rename failure.", err.message);
            }
        }
    })
}

// main
// node rename.js [dir] [function]
// function (name, ext, index) => string|false
if (require.main === module) {
    let argv = process.argv;
    if (argv.length < 4) {
        console.log("node rename.js [dir] [function]");
        process.exit(1);
    }
    let dir = argv[2];
    try {
        if (!fs.statSync(dir).isDirectory()) {
            console.log(dir, "is not dir");
            process.exit(1);
        }
    } catch (e) {
        console.log(e.message);
        process.exit(1);
    }

    console.log("dir is", dir);
    let fun = argv[3];
    try {
        fun = eval("(" + fun + ")");
    } catch (e) {
        console.log(fun, "is not JS code.", e.message);
        process.exit(1);
    }
    if (typeof fun !== 'function') {
        console.log(fun, "is not JS function");
        process.exit(1);
    }
    console.log("fun is", fun);
    work(dir, fun);
}
