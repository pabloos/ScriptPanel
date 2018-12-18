function login() {
    const userObject = getFormObject(forms["login"]);

    const config = {
        method: 'POST',
        body: JSON.stringify(userObject),
        headers:{
            'Content-Type': 'application/json'
        }
    }

    let request = new Request("/login", config)

    fetch(request)
        .then(ok => ok = ok.json())
        .catch(error => console.error('Error:', error))
        .then(ok => {
            if(!ok.authorized) { //login fails case
                alert("Login failed");
            } else {
                setUser(userObject.Name, userObject.Department, userObject.Company);
                removePresentation();
            }
        })
}

function setUser(name, department, company){
    loginCookies.set('name', name);
    loginCookies.set('company', company);
    loginCookies.set('department', department);
    
    $('#loginModal').modal('hide') //the modal crashes and does not hide himself if we overwrite the eventlistener of the button
   
    signupButton.style.display = "none";

    loginbutton2userbutton(name);

    $("#presentationContainer").remove();

    $("#scriptPanel").remove();

    getUserScripts();
}

async function loginbutton2userbutton(user) {
    loginButton.style.display = "none";

    const btnGroup = document.getElementsByClassName("btn-group")[0];

    btnGroup.appendChild(createComponent(userButton(user)));
    btnGroup.appendChild(createComponent(userDropdown()));
}

function unset(){
    $("#scriptPanel").remove();

    signupButton.style.display = "block";

    userbutton2loginbutton();

    unsetCookies(loginCookies);
}

function userbutton2loginbutton() {
    $("#userButton").remove();
    loginButton.style.display = "block";

    document.getElementsByClassName("dropdown-menu")[0].remove()
}

//este es el boton de dentro del formulario
loginAction.addEventListener("click", login);