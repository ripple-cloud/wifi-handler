$(document).ready(function() {
  // Network
  var networks = new Bloodhound({
    datumTokenizer: Bloodhound.tokenizers.obj.whitespace('name'),
      queryTokenizer: Bloodhound.tokenizers.whitespace,
      limit: 10,
      // remote lists all available networks
      // change this to prefetch if drop down list is to match only characters entered
      remote: {
        url: '/network',
        filter: function(list) {
          return $.map(list, function(network) { return { name: network }; });
        }
      }
  });
  networks.initialize();

  $('#prefetch .typeahead').typeahead(null, {
    name: 'networks',
    displayKey: 'name',
    source: networks.ttAdapter()
  }); 

  // Security
  var security = ["None", "WEP", "WPA", "WPA2"];
  var substringMatcher = function(strs) {
    return function findMatches(q, cb) {
      matches = [];
                
      $.each(strs, function(i, str) {
          matches.push({ value: str });
      });

      cb(matches);
    };
  };

  $('#the-basics .typeahead').typeahead({
    hint: true,
    highlight: true,
    minLength: 1
  },
  {
    name: 'security',
    displayKey: 'value',
    source: substringMatcher(security)
  });
});
