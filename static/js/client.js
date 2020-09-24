var client = (function(){
    return {
        peoples:function(callback){
            $.get({url:"/people"}).
                then((data)=>callback(null,data),(err)=>callback(err,null));
        },addPeople:function(people,callback){
            console.log(people);
            $.post({url:"/people",data:JSON.stringify(people),contentType:"application/json"}).
                then(()=>callback(null),(er)=>callback(er));
        }
    }
})();