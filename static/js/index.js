function executeSearch() {
    var searchString = $('.search-input-desktop').val();
    $.get("/search/", 'search='+searchString)
}
