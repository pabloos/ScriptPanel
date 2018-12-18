//TODO: check that the form passed is a form node

function getFormObject(form) {    
    const length = form.children.length;
    
    fields = "";

    for(var i = 0; i < length; i++){
        fields = fields.concat(" " + form.children[i].children[0].innerHTML);   //get the fields
    }

    const values = Array.from(form.children).map(child => child.children[1].value); //get the data 

    var FormType = factory(fields.substr(1)); //we need to remove the first blank space

    return new FormType(... values);
}
