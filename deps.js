var nanomorph = require('nanomorph')

function htmlToNode (html) {
  var cont = document.createElement('div')
  cont.innerHTML = html
  return cont.firstChild
}

window.gopherjsNanomorphDeps = {
  nanomorph: nanomorph,
  htmlToNode: htmlToNode
}
