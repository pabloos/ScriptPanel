uploadfile.addEventListener("change", function(event) {

    // When the control has changed, there are new files

    var i = 0,
        files = uploadfile.files,
        len = files.length;

    for (; i < len; i++) {
        console.log("Filename: " + files[i].name);
        console.log("Type: " + files[i].type);
        console.log("Size: " + files[i].size + " bytes");
    }

}, false);

function hideRunModal(object) {
    $('#uploadScriptModal').modal('hide') //the modal crashes and does not hide himself if we overwrite the eventlistener
}

    //!!!!!!!!!
    //esto de aquí es la DESCARGA de un archivo, por lo que hay que trasladarlo al callback del execScript
    // filename = "sacar de un json.bash" //hay que poner una extension o si no el chrmoe pone por defecto .txt

    // var element = document.createElement('a');
    // element.setAttribute('href', 'data:text/plain;charset=utf-8,' + encodeURIComponent(object));
    // element.setAttribute('download', filename);
  
    // element.style.display = 'none';
    // document.body.appendChild(element);
  
    // element.click();
  
    // document.body.removeChild(element);


//funcion para enviar el script
function sendFile(){
    var form = new FormData();

    form.append("name", Cookies.get('name'));
    form.append("company", Cookies.get('company'));
    form.append("department", Cookies.get('department'));

    form.append("uploadfile", uploadfile.files[0]);

    //script config
    form.append("flags", getValuesFromForm(uploadScriptFormflags));
    form.append("args", getValuesFromForm(uploadScriptFormArgs));
    form.append("description", getScriptDescription());

    sendAndThen('POST', form, "/upload", hideRunModal);
}

var optionsCounter = 0;
var argsCounter = 0;

 function addInputField(){

    let destiny;
    let name;

    //who's the caller? --> go to lookup-table
    if(this.id == "addArgFieldButton") {
        destiny = uploadScriptFormArgs;
        name = "arg-" + (argsCounter + 1);
        argsCounter++;

    } else if (this.id == "addFlagFieldButton") {
        destiny = uploadScriptFormflags;
        name = "flag-" + (optionsCounter + 1);
        optionsCounter++;
    }

    var formGroupDiv = document.createElement('div');
    formGroupDiv.setAttribute("class", "form-group");

        var label = document.createElement('label');
        label.setAttribute("for", name);
        label.innerHTML = name;

        var input = document.createElement('input');
        input.setAttribute("class", "form-control");
        input.setAttribute("type", "text");
        input.setAttribute("name", name);

    formGroupDiv.appendChild(label);
    formGroupDiv.appendChild(input);

    destiny.appendChild(formGroupDiv);
}

function removeInputField() {
    let destiny;

    //who's the caller? --> go to lookup-table
    if(this.id == "removeArgFieldButton") {
        destiny = uploadScriptFormArgs;
        argsCounter--;

    } else if (this.id == "removeFlagFieldButton") {
        destiny = uploadScriptFormflags;
        optionsCounter--;
    }

    const length = destiny.childNodes.length-1;

    if (length == -1) return;

    var lastOne = destiny.childNodes[length];
    
    destiny.removeChild(lastOne);
}

function getScriptDescription(){
    return scriptDescription.value;
}

//event registers
uploadScriptAction.addEventListener("click", sendFile);

addFlagFieldButton.addEventListener("click", addInputField);
removeFlagFieldButton.addEventListener("click", removeInputField);

addArgFieldButton.addEventListener("click", addInputField);
removeArgFieldButton.addEventListener("click", removeInputField);