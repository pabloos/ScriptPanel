function getUserScripts() {
    const user = {
        name: Cookies.get('name'),
        company: Cookies.get('company'),
        department: Cookies.get('department')
    }
    
    sendAndThen("POST", JSON.stringify(user), "/getUserScripts", addScriptsToPanel);    //debería ser get
}