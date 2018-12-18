function makeIterator(objectArray) {
    var index = 0;
    var length = objectArray.length;

    return {
        next: function() {
            if (!this.hasNext()) {
                return null;
            }

            element = objectArray[index];
            index += 1;
            return element;
        },

        hasNext: function() {
            return index < length;
        }
    }
}