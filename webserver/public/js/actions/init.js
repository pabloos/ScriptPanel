window.onload = setEnviorement;

function setEnviorement(){
    if(areLoginCookiesSet()) {
        removePresentation();

        setUser(loginCookies.get("name"), loginCookies.get("department"), loginCookies.get("company"));
    } 
}

function areLoginCookiesSet() {
    if (loginCookies.get("company") == null) {
        return false;
    }

    return true;
}

var userScripts;
