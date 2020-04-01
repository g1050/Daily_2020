
var name = 'Gao Xingkun';//var 关键词定义变量

var arr = [
'1',
true
]

var isAdmin = true;
isAdmin = false;

if(isAdmin){
    alert('isAdmin');
}

else{
    alert('not Admin');
}

alert('Hello js ' + name);
console.log(arr.length);//查看数组中元素的个数//查看数组中元素的个数

var day = 1;
switch(day)
{
    case 1:
        alert('Monday');
        day++;
        break;
    case 2:
        alert('Tuesday');
        break;
    case 3:
        alert('Wednesday');
        break;
    default:
        break;
}

var string = 'Hello' + 'World';
console.log(day);
console.log(string);


function add(a,b){
    console.log(a+b);
    return a+b;
}

var result = add(1,2);
console.log('result = ' + result);

function person(name){
    var sex,age;
    return {
        getName: function(){
            return name;
        },
        setAge: function(newAge){
            age = newAge;
        },
        getAge: function(){
            return age;
        }
        
    }
}

var p = person('Li Lei');
p.setAge(22);
console.log(p.getName() + ' ' + p.getAge());