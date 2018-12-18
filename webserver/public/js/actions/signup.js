function signup() {    
    const config = {
        method: 'POST',
        body: JSON.stringify(getFormObject(forms["signup"])),
        headers:{
            'Content-Type': 'application/json'
        }
    }

    let request = new Request("/signup", config);

    fetch(request)
        .then(response => response = response.json())
        .catch(error => console.error('Error:', error))
        .then(response => {
            if(!response.OK) { //signup fails case
                alert("Signup failed");
            } else {
                $('#signupModal').modal('hide')
            }
        });
}

signupAction.addEventListener("click", signup);