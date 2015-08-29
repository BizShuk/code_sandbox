var linknode = function(value,next,prev){
    this.value = value || 0;
    this.next = next || null;
    this.prev = prev || null;
}

var link_list_sample = [];

var i = 0;

for( i=0 ; i <=5 ; i++ ){
    var node = new linknode( i , null , null );
    link_list_sample.push( node );
    
    if( link_list_sample.length > 1 ){
        node.prev = link_list_sample[i-1];
        node.prev.next = node;
    }

}

console.log("\x1b[32m link list sample:\x1b[m\n",link_list_sample);

var reverselist_iterated = function(head){
    var prev = temp = null;
    var current = head;

    while(current != null){
        var temp = current.next;
        current.next = prev;
        prev = current;
        current.prev = temp;
        current = temp;
    }

    return prev;
}

reverselist_iterated(link_list_sample[0]);
console.log("\x1b[32m reverse link list:\x1b[m\n",link_list_sample);
