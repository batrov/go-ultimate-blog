var layer;

var view = new ol.View({
    center: ol.proj.fromLonLat([112.632629, -7.966620]),
    zoom: 10
})

var map = new ol.Map({
    target: 'map',
    layers: [
        new ol.layer.Tile({
            source: new ol.source.OSM()
        })
    ],
    view: view
});

var x = document.getElementById("demo");

function getLocation() {
    if (navigator.geolocation) {
        navigator.geolocation.getCurrentPosition(showPosition);
    } else {
        x.innerHTML = "Geolocation is not supported by this browser.";
    }
}

function showPosition(position) {
    x.innerHTML = "Latitude: " + position.coords.latitude +
        "<br>Longitude: " + position.coords.longitude;

    var curLoc = ol.proj.fromLonLat([position.coords.longitude, position.coords.latitude]);

    map.removeLayer(layer);
    layer = new ol.layer.Vector({
        source: new ol.source.Vector({
            features: [
                new ol.Feature({
                    geometry: new ol.geom.Point(curLoc)
                })
            ]
        })
    });
    map.addLayer(layer);

    view.animate({
        center: curLoc,
        duration: 2000,
        zoom: 20,
    });
}

$('#recenter').click(function() {
    getLocation()
});

$('#quit').click(function() {
    window.location.href = '/';
});