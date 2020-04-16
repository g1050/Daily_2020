var app = new Vue({
    el: '#app',
    data:{
        name: 'Gao XingKun',
        sex: 'man',
        age: null, 
        friend: ['Jack','Tony','Mary'],
        friends:[
            {
                name: 'Jack',
                age:50 
            },

            {
                name: 'Tony',
                age:65 
            },

            {
                name: 'Mary',
                age: 19
            }
        ],
        url: 'http://www.baidu.com'
    }
});