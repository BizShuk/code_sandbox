執行queue
    function printing() {
       console.log(1);
       setTimeout(function() { console.log(2); }, 1000);
       setTimeout(function() { console.log(3); }, 0);
       console.log(4);
    }
    printing();
    1
    4
    3
    2


    right sided view B-tree [待 js 實作]
        可能要考queue
        從root 進去後 每一層結束後加個null 來判斷該層最後要寫入的value
        a0 null a1 a2 null a3 a4 a5 a6 null
        如果
        http://blog.csdn.net/ljiabin/article/details/44900023




