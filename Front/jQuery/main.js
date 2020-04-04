
// //原生js
// console.log("Hello world");
// document.getElementById('a')
//     .style
//     .background = 'blue';
// // 用js加背景颜色


// // jQuery('#b')
// 1.选择器
// $('#b')
// .css('background','green');

// $('div')
// .css('background','brown');

// $('.b')
// .css('background','red');


// 2.过滤器
$('div').find('.child')
.css('border','2px solid black');

//选择父级
$('#child1')
.parent()
.css('border','2px solid red');

//所有大辈,一直向上找
$('#child2')
.parents('.grandpa')
.css('border','2px solid green');

//过滤child下的女孩子
$('.child')
.filter('.girl')
.css('background','black');

//3.操作样式
//可以给一个对象
var a = $('.c')
    .css({

        // color: 'red',
        // background: 'black',

    })
    .addClass('black')//添加一个类
    .removeClass('black');

console.log(a.hasClass('black'));//判断是不是有这个类

//隐藏\展示元素

// a.show();
// a.hide();
// a.fadeout();

a.fadeOut(2000);
a.fadeIn(2000);
a.slideUp(2000);
a.slideDown(2000);

//4.操作DOM
var text = $('.text').text();
console.log(text);
$('.text').text('jQuery');
console.log($('.text').html());

$('.text').append('<p>apend</p>');
$('.text').prepend('<p>prepend</p>');
// $('.text').remove();
