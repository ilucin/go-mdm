var request = new XMLHttpRequest();
var projects;

location.hash = '';

request.open('GET', '/projects', true);
request.onload = function() {
  if (request.status >= 200 && request.status < 400) {
    projects = JSON.parse(request.responseText).Items;
    showProjects()
  } else {
    alert("Error! It's Krešo's fault!")
  }
};

window.onhashchange = function() {
  var split;
  if (location.hash && location.hash.indexOf('projects/')) {
    split = location.hash.split('/');
    showProject(parseInt(split[split.length - 1]));
  } else {
    document.querySelector('.project').innerHTML = '';
  }
};

request.onerror = function() {
  alert('Error! Krešo did something wrong!');
};

function showProjects() {
  document.querySelector('.projects').innerHTML = projects.map(function(p) {
    return '<div><a href="#projects/' + p.ID + '">' + p.Repo + '</a></div>';
  }).join('');
}

function getProject(id) {
  for (var i = 0; i < projects.length; i++) {
    if (projects[i].ID === id) {
      return projects[i];
    }
  }
}

function showProject(id) {
  var p = getProject(id);

  var lib, projVersion, currVersion, color;

  var html = '<h2>' + p.Repo + '</h2>';
  html += '<table><tr><th> Package name </th> <th> Package version </th> <th> Project version </th></tr>';

  for (var i = 0; i < p.Libs.length; i++) {
    lib = p.Libs[i].Lib;
    projVersion = p.Libs[i].Version.trim();
    currVersion = lib.Version.trim();

    if (projVersion === currVersion) {
      color = 'blue';
    } else {
      color = 'red';
    }

    html += '<tr style="color: ' + color + '"><td>' + lib.Name + '</td><td>' + currVersion + '</td><td>' + projVersion + '</td></tr>';
  }

  html += '</table>';

  document.querySelector('.project').innerHTML = html;
}

request.send();
