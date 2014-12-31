$(document).ready(function() {
  var networks = new Bloodhound({
    datumTokenizer: Bloodhound.tokenizers.obj.whitespace('name'),
      queryTokenizer: Bloodhound.tokenizers.whitespace,
      limit: 10,
      remote: {
        url: '/network',
        filter: function(list) {
          return $.map(list, function(network) { return { name: network }; });
        }
      }
  });

  networks.initialize();

  $('.typeahead').typeahead(null, {
    name: 'networks',
    displayKey: 'name',
    source: networks.ttAdapter()
  }); 
});
