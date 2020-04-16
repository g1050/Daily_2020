var app = new Vue({
    el: '#app',
    data: {
        name: 'Jack',
        sex: null,
        age: 22
    }
});

var test = new Vue({
    el: '#test',
    data: {
        math: null
    }
});

var test2 = new Vue({
    el: '#test2',
    data: {
        chinese: 100
    }

});

var test3 = new Vue({
    el: '#test3',
    data: {
        english: 141
    }
});

var test4 = new Vue({
    el: '#test4',
    data: {
        foodList: ['葱', '姜', '蒜'],
        people: [{
                name: 'Jack',
                age: 22
            },
            {
                name: 'Tom',
                age: 21
            },
            {
                name: 'Mary',
                age: 24
            }
        ]
    }
});

var test5 = new Vue({
    el: '#test5',
    data: {
        url: "http://www.baidu.com",
        isActive: true
    }
});

var test6 = new Vue({
    el: '#test6',
    methods: {
        onClick: function() {
            alert('Yo');
        }
    }
});

var test7 = new Vue({
    el: '#test7',
    data: {
        sex: 'male',
        hobby: [],
        article: 'Hello js,Hello Vue',
        from: null
    }
});

var test8 = new Vue({
    el: '#test8',
    data: {
        math: 90,
        english: 141,
        chinese: 110
    },
    computed: {
        sum: function() {
            return this.math + this.english + this.chinese;
        },
        average: function() {
            return this.sum / 3;
        }
    }
});