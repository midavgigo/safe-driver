function send_api_order_post(div){
    var body ={
        "passenger_id":         div.getElementsByClassName("passenger_id")[0].value,
        "address_from":         div.getElementsByClassName("address_from")[0].value,
        "address_to":           div.getElementsByClassName("address_to")[0].value,
        "tariff":               div.getElementsByClassName("tariff")[0].value,
        "selected_services":    div.getElementsByClassName("selected_services")[0].value.split(","),
        "comment":              div.getElementsByClassName("comment")[0].value,
    };
    console.log(body);
    fetch('http://localhost:8080/api/order', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(body)
    });
}

function get_order_info(div){
    var order_id = div.getElementsByClassName("order_id")[0].value;
    fetch('http://localhost:8080/api/order/'+order_id, {
        method: 'GET'
    })
}

function send_order_cancel(div){
    var order_id = div.getElementsByClassName("order_id")[0].value;
    fetch('http://localhost:8080/api/order/'+order_id+'/cancel', {
        method: 'POST'
    })
}

function send_driver_status(div){
    var body ={
        "is_available":         div.getElementsByClassName("is_available")[0].value=="on",
        "current_location":{
            "lat": div.getElementsByClassName("lat")[0].value,
            "lng": div.getElementsByClassName("lng")[0].value,
        },
    };
    console.log(body);
    fetch('http://localhost:8080/api/driver/status', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(body)
    });
}

function accept_order(div){
    var order_id = div.getElementsByClassName("order_id")[0].value;
    fetch('http://localhost:8080/api/order/'+order_id+'/accept', {
        method: 'POST'
    })
}

function arrived_order(div){
    var order_id = div.getElementsByClassName("order_id")[0].value;
    fetch('http://localhost:8080/api/order/'+order_id+'/arrived', {
        method: 'POST'
    })
}

function update_order_status(div){
    var body ={
        "status": div.getElementsByClassName("status")[0].value,
    };
    var order_id = div.getElementsByClassName("order_id")[0].value;
    fetch('http://localhost:8080/api/order/'+order_id+'/status', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(body)

    })
}