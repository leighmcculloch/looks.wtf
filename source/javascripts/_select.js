HTMLElement.prototype.selectText = function() {
  if (window.getSelection) {
    if (window.getSelection().empty) {  // Chrome
      window.getSelection().empty();
    } else if (window.getSelection().removeAllRanges) {  // Firefox
      window.getSelection().removeAllRanges();
    }
  } else if (document.selection) {  // IE?
    document.selection.empty();
  }
  var range = document.createRange();
  range.selectNode(this);
  window.getSelection().addRange(range);
}

$(function() {
  $('.selectable').mouseover(function() {
    this.selectText();
  })
});

new Clipboard('.selectable');