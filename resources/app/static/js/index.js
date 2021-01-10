let index = {
    about: function(aboutPayload) {
        let c = document.createElement("div");
        c.setAttribute("class", "about");

        // display appname
        let appName = document.createElement("div");
        appName.setAttribute("class", "appName");
        appName.innerHTML = aboutPayload.AppName;
        c.appendChild(appName);

        // use table to display detail info
        let detail = document.createElement("table");
        detail.setAttribute("class", "detail");

        detail.insertRow();
        detail.rows[0].insertCell().innerHTML = "Version:";
        detail.rows[0].insertCell().innerHTML = aboutPayload.Version;

        detail.insertRow();
        detail.rows[1].insertCell().innerHTML = "BuiltTime:";
        detail.rows[1].insertCell().innerHTML = aboutPayload.BuiltTime;

        detail.insertRow();
        detail.rows[2].insertCell().innerHTML = "Electron:";
        detail.rows[2].insertCell().innerHTML = aboutPayload.Electron;

        detail.insertRow();
        detail.rows[3].insertCell().innerHTML = "Astilectron:";
        detail.rows[3].insertCell().innerHTML = aboutPayload.Astilectron;

        c.appendChild(detail);

        // Github url
        let githubUrl = document.createElement("a");
        githubUrl.setAttribute("href", aboutPayload.Github);
        githubUrl.setAttribute("onclick", "require('electron').shell.openExternal('" + aboutPayload.Github + "')")
        githubUrl.innerHTML = "Github";
        c.appendChild(githubUrl);

        asticode.modaler.setContent(c);
        asticode.modaler.show();
    },
    addFolder(name, path) {
        let div = document.createElement("div");
        div.className = "dir";
        div.onclick = function() { index.explore(path) };
        div.innerHTML = `<i class="fa fa-folder"></i><span>` + name + `</span>`;
        document.getElementById("dirs").appendChild(div)
    },
    init: function() {
        // Init
        asticode.loader.init();
        asticode.modaler.init();
        asticode.notifier.init();

        // Wait for astilectron to be ready
        document.addEventListener('astilectron-ready', function() {
            // Listen
            index.listen();

            // // Explore default path
            // index.explore();
        });
    },

    /*
    explore: function (path) {
        // Create message
        let message = { "name": "explore" };
        if (typeof path !== "undefined") {
            message.payload = path
        }

        // Send message
        asticode.loader.show();
        astilectron.sendMessage(message, function (message) {
            // Init
            asticode.loader.hide();

            // Check error
            if (message.name === "error") {
                asticode.notifier.error(message.payload);
                return
            }

            // Process path
            document.getElementById("path").innerHTML = message.payload.path;

            // Process dirs
            document.getElementById("dirs").innerHTML = ""
            for (let i = 0; i < message.payload.dirs.length; i++) {
                index.addFolder(message.payload.dirs[i].name, message.payload.dirs[i].path);
            }

            // Process files
            document.getElementById("files_count").innerHTML = message.payload.files_count;
            document.getElementById("files_size").innerHTML = message.payload.files_size;
            document.getElementById("files").innerHTML = "";
            if (typeof message.payload.files !== "undefined") {
                document.getElementById("files_panel").style.display = "block";
                let canvas = document.createElement("canvas");
                document.getElementById("files").append(canvas);
                new Chart(canvas, message.payload.files);
            } else {
                document.getElementById("files_panel").style.display = "none";
            }
        })
    },*/
    listen: function() {
        astilectron.onMessage(function(message) {
            switch (message.name) {
                case "about":
                    index.about(message.payload);
                    console.log("about clicked!\n");
                    return { payload: "payload" };
                    break;
                case "check.out.menu":
                    asticode.notifier.info(message.payload);
                    break;
                case "print":
                    console.log(message.payload);
                    foldest.print(message.payload);
                    break;
            }
        });
    }
};

let foldest = {
    start: function() {
        console.log("button Start is clicked!\n");
        document.getElementById("btn-start").disabled = true;
        astilectron.sendMessage({ "name": "start" }, function() {
            console.log("finished!");
            document.getElementById("btn-start").disabled = false;
        });
    },
    print: function(html) {
        document.getElementById("outputs").innerHTML += (html + `<br>`);
    },
}