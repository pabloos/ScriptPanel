function getValuesFromForm(form){
    const lengthChildsLength = form.children.length;

    var array = [];

    for(var i = 0; i < lengthChildsLength; i++) {
        child = form.children[i];

        if (child.getAttribute("class") == "form-group") {
            array.push(child.children[1].value); //the second one is the input node
        }
    }

    return JSON.stringify(array);
}