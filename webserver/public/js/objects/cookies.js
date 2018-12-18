FactoryCookies = function() {
    return {
        set: function(name, value){
            Cookies.set(name, value);
        },
        get: function(name){
            return Cookies.get(name);
        },
        remove: function(name){
            Cookies.remove(name);
        }
    };
};

loginCookies = new(FactoryCookies);

function unsetCookies(cookies){
    cookies.remove('name');
    cookies.remove('company');
    cookies.remove('department');
}
