
function reverse(str) {

    // check for type.
    if ( 0 && typeof str === 'object') {
        return "";
    }

    str += "";
    return str.split("").reverse().join("");

}

function IsPalindrome(str) {
    if (str === reverse(str)) {
        return true;
    }
    return false;
}


function reverse_regex(str) {
    str += "";
    return str.split(/[:;]/).reverse().join("");
}


function alpha(str) {
    str += "";
    return str.split("").sort().join("");
}

function cap(str) {
    str += "";

    t1 = str.split(" ");
    console.log(t1);

    t1.forEach(function(v,i,a){ t1[i] = v.charAt(0).toUpperCase()+v.substr(1); });

    return t1.join("");
}

function longestword(str) {
    str += "";
    
    var arr = str.split(" ");
    if (arr.length == 0) {
        return "";
    }

    var word= arr[0];

    for (var i = 0, len = arr.length; i < len; i++) {
        if (word.length < arr[i].legnth) {
            word = arr[i];
        }
    }
    return word;
}




console.log(reverse(34223));
console.log(reverse(3.334));
console.log(reverse("test"));
console.log(reverse([1,2,3]));
console.log(reverse({a:1,b:2}));
console.log(reverse(function(){ }));
console.log(reverse(null));
console.log(reverse(true));
console.log(reverse(false));

// sort each charactor
console.log(alpha("34125"));

// cap each work
console.log(cap("ducks love to quack"));

