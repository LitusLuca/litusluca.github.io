let t_cookies = 0;
let cookie = 1;

function buttonclick(){
    t_cookies += cookie;
    document.getElementById("CookieCount").innerHTML = t_cookies + " cookies";
    var elem = document.getElementById("CookiePNG");
    var size = 256;
    var y = 0
    var id = setInterval(frame, 1);
    var done = false;
    function frame(){
        if (size == 300) {
            clearInterval(id);
            done = true;
            button_an(done, size, y);  
        }
        else{
            size +=2;
            y--;
            elem.style.width = `${size}px`;
            elem.style.transform = `translateY(${y}px)`
        }      
    }
    create_cookie();
    if (t_cookies == 200){
        window.alert('Hoi, \nwordt het misschien niet tijd om iets anders te doen.')
    }

}
function button_an(done = false, size = 300, y){
    if (done == true){
        var elem = document.getElementById("CookiePNG");
        var id = setInterval(frame, 1);
        function frame(){
            if (size == 256){
                clearInterval(id);              
            }
            else{
                size -=2;
                y++;
                elem.style.width = `${size}px`;
                elem.style.transform = `translateY(${y}px)`
            }

        }
    }

}

function timeline(interval = 1, func = function(){}, begin = 0, end=1, step = 1, done_func = function(){}) {
       var id = setInterval(frame, interval);
           function frame(){
               if(begin == end){
                   clearInterval(id);
                   done_func();
               }else{
                   begin += step;
                   func(begin);
               }
       }
}
function print(line) {
    console.log(line);
    }

function create_cookie(){
    const particle = document.createElement('img');
    particle.src = 'cookie.png'
    document.getElementById("cookies").appendChild(particle);
    x = (Math.floor(Math.random() * (1920 - 257) ));
    desY = (window.innerHeight-256);

    const animation = particle.animate([{
        transform: `translate(${x}px, ${0}px)`,
         opacity: 1
    }
    ,{
        transform: `translate(${x}px, ${desY}px)`,
        opacity: 0
    }],
    {
        duration: 2000 + Math.random() * 1000
    });
    animation.onfinish = () => {
        particle.remove();
    }

}
