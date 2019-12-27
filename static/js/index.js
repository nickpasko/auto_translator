setInterval(function() {
    $.get('/api/wc/status', function(data) {
        $('#wc_status').html(data);
    });
}, 1000)
