function renderChart(data, labels) {
    var ctx = document.getElementById("myChart").getContext('2d');
    var myChart = new Chart(ctx, {
        type: 'line',
        data: {
            labels: labels,
            datasets: [{
                label: 'This week',
                data: data,
            }]
        },
    });
}
$(document).ready(function () {
    // When id with Action is clicked
    $("#Action").click(function () {
        // Load ajax.php as JSON and assign to the data variable
        $.getJSON('http://localhost:5080', function (data) {
            // set the html content of the id myThing to the value contained in data
            console.log(data);
            dataPoints = []
            data.map((item) => {
                dataPoints.push({ "y": item.time, "name": item.name })
            });
            let total = 0;
            let hrnum = 0;
            dataPoints.forEach((item) => {
                hrnum = Number(item.y) * 100 * 0.6 * .06
                total += hrnum;
            });
            let timeHr = Number(String(total).split(".")[0]);
            let timeMn = Number(String((total * 100) % 60).split(".")[0]);
            document.getElementById('timeHr').innerHTML =
                String(timeHr);
            document.getElementById('timeMn').innerHTML =
                String(timeMn);
            var chart = new CanvasJS.Chart("chartContainer", {
                theme: "dark2",
                backgroundColor: "transparent",
                exportFileName: "Doughnut Chart",
                exportEnabled: true,
                animationEnabled: true,
                title: {
                },
                legend: {
                    cursor: "pointer",
                    itemclick: explodePie,
                    fontWeight: "normal",
                },
                data: [{
                    type: "doughnut",
                    innerRadius: 160,
                    radius: "100%",
                    indexLabelFontSize: 20,
                    showInLegend: true,
                    toolTipContent: "<b>{name}</b>: {y} (#percent%)",
                    indexLabel: "{name} - #percent%",

                    dataPoints: dataPoints,

                }]

            });
            chart.render();


            //$("#myThing").html(data.value);
        });
        // .then(() => {

        // }));

    });
});


function showDiv() {
    div = document.getElementById('chartContainer');
    div.style.display = "block";
    divr = document.getElementById('totalHrs');
    divr.style.display = "block";
    divh = document.getElementById('Action');
    divh.style.display = "none";
}


// window.onload = function () {



function explodePie(e) {
    if (typeof (e.dataSeries.dataPoints[e.dataPointIndex].exploded) === "undefined" || !e.dataSeries.dataPoints[e.dataPointIndex].exploded) {
        e.dataSeries.dataPoints[e.dataPointIndex].exploded = true;
    } else {
        e.dataSeries.dataPoints[e.dataPointIndex].exploded = false;
    }
    e.chart.render();
}

function updateClock() {
    var now = new Date(), // current date
        months = ['January', 'February', '...']; // you get the idea
    time = now.getHours() + ':' + now.getMinutes(), // again, you get the idea

        // a cleaner way than string concatenation
        date = [now.getDate(),
        months[now.getMonth()],
        now.getFullYear()].join(' ');

    // set the content of the element with the ID time to the formatted string
    document.getElementById('time').innerHTML = [date, time].join(' / ');

    // call this function again in 1000ms
    setTimeout(updateClock, 1000);
}
updateClock();
$("#Act").click(function () {
    if (confirm('Are you sure you want to save this thing into the database?')) {
        $.ajax({
            url: "http://localhost:5080/quit/",
            data: {}
        })
    } else {
        // Do nothing!
    }
})

var TxtType = function (el, toRotate, period) {
    this.toRotate = toRotate;
    this.el = el;
    this.loopNum = 0;
    this.period = parseInt(period, 10) || 2000;
    this.txt = '';
    this.tick();
    this.isDeleting = false;
};

TxtType.prototype.tick = function () {
    var i = this.loopNum % this.toRotate.length;
    var fullTxt = this.toRotate[i];

    if (this.isDeleting) {
        this.txt = fullTxt.substring(0, this.txt.length - 1);
    } else {
        this.txt = fullTxt.substring(0, this.txt.length + 1);
    }

    this.el.innerHTML = '<span class="wrap">' + this.txt + '</span>';

    var that = this;
    var delta = 200 - Math.random() * 100;

    if (this.isDeleting) { delta /= 2; }

    if (!this.isDeleting && this.txt === fullTxt) {
        delta = this.period;
        this.isDeleting = true;
    } else if (this.isDeleting && this.txt === '') {
        this.isDeleting = false;
        this.loopNum++;
        delta = 500;
    }

    setTimeout(function () {
        that.tick();
    }, delta);
}

window.onload = function () {
    var elements = document.getElementsByClassName('typewrite');
    for (var i = 0; i < elements.length; i++) {
        var toRotate = elements[i].getAttribute('data-type');
        var period = elements[i].getAttribute('data-period');
        if (toRotate) {
            new TxtType(elements[i], JSON.parse(toRotate), period);
        }
    }
    // INJECT CSS
    var css = document.createElement("style");
    css.type = "text/css";
    css.innerHTML = ".typewrite > .wrap { border-right: 0.08em solid #fff}";
    document.body.appendChild(css);
    document.getElementById('timera').innerHTML =
        00 + ":" + 20;
    startTimera();

    function startTimera() {
        var presentTime = document.getElementById('timera').innerHTML;
        var timeArray = presentTime.split(/[:]+/);
        var m = timeArray[0];
        var s = checkSecond((timeArray[1] - 1));
        if (s == 59) { m = m - 1 }
        //if(m<0){alert('timer completed')}

        document.getElementById('timera').innerHTML =
            m + ":" + s;
        console.log(m)
        setTimeout(startTimera, 1000);
        if (m == 0 && s == 0) {
            console.log("powli")
            $.ajax({
                url: "http://localhost:5080/disable/",
                data: {}
            })
            document.getElementById('timera').innerHTML =
                " 0:00 !You are reday to Hack!";

        }
    }

    function checkSecond(sec) {
        if (sec < 10 && sec >= 0) { sec = "0" + sec }; // add zero in front of numbers < 10
        if (sec < 0) { sec = "59" };
        return sec;
    }
}
