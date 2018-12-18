function factory(fields) {
    var fields = fields.split(' ');
    var count  = fields.length;

    return function() {  //clousure
        for (i = 0; i < count; i++) {
            this[fields[i]] = arguments[i];
        }
    }
}