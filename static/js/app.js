var app = (function(){

    var showPeoples = function(err,data){
        if(err!=null){
            alert("Error");
            return;
        }
        $("#add").html("");
        var table= $("<table border='1' style='width:100%; boder:1px solid black; border-collapse:collapse;'><tr><td>Name</td><td>Age</td><td>City</td><td>Address</td></tr></table>")
        data.forEach((people)=>{
            table.append("<tr><td>"+people.name+"</td><td>"+people.age+"</td><td>"+people.city+"</td><td>"+people.address+"</td></tr>")
        });
        $("#add").append(table);
    }
    var addPeopleShow = function(err){
        if(err!=null){
            alert("Error");
            return ;
        }
        app.peoples();
    }
    return {
        peoples:function(){
            client.peoples(showPeoples);
        },
        addPeople:function(){
            var name = $("#name").val();
            var age = parseInt($("#age").val());
            var city = $("#city").val();
            var ad = $("#address").val();
            client.addPeople({name:name,age:age,city:city,address:ad},addPeopleShow);

        }
    }
})();