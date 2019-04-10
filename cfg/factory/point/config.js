{
    "name":"KRATOHVIL",	
    "pointConfigFiles":"../cfg/factory/point/config.js",
    "intervalOnOff2":"KOLBASA",
    "intervalOnOff":{
        "mahima":[
            { "pin": "7", "state": "1", "interval": "::3"},
            {"gpio": "7", "state": "0", "interval": "::2"},
            {" pin": "3", "state": "1", "interval": "::1" },
            {"gpio": "3", "state": "0", "interval": "::2" },
            {" pin": "4", "state": "1", "interval": "::3"},
            {"gpio": "4", "state": "0", "interval": "::4"}],
        "granato":[
            {
                "pin": "5",
                "state": "1",
                "interval": "::5"
            },
            {
                "pin": "5",
                "state": "0",
                "interval": "::4"
            },
            {
                "pin": "11",
                "state": "1",
                "interval": "::3"
            },
            {
                "pin": "11",
                "state": "0",
                "interval": "::2"
            }
        ]
    },

    "internalIP":"127.0.0.1",
    "internalPort":"7500",
    "externalPort":"7250",
    "webEmail":"vilis@mailof.com",
    "webAliveInterval":"7200",	
    "webEmailMutt":"../cfg/factory/muttrc.set",	
    "scriptPath":"../cfg/factory/scripts",	
    "logPath":"../data",
    "webPort":"49888",
    "eventPath":"../kika/event",
    "errorPath":"../mika/error",	
    "templatePath":"tmpl",
    "templateExt":".tmpl"
}
