var trigger = $('#trigger');
var card = $('#card');
var loaded  = false;

trigger.on('click',

    function(){
        if(card.is(':visible')){
            card.slideUp();
        }else{
            if(!loaded){
                card.load('card.html');
                loaded = true;
            }
            card.slideDown();
        }
    }

)

$.ajax('http://api.github.com/users/g1050')
    .done(function(data){
        console.log(data);
    })

var form = $('#search');
var input = $('input#username');
var result = $('#result');

form.on('submit',function(event){
    event.preventDefault();
    console.log('222');
})