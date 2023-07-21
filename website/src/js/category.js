import Clipboard from "./_clipboard.js";

function selectText(element) {
  if (window.getSelection) {
    if (window.getSelection().empty) {  // Chrome
      window.getSelection().empty();
    } else if (window.getSelection().removeAllRanges) {  // Firefox
      window.getSelection().removeAllRanges();
    }
  } else if (document.selection) {  // IE?
    document.selection.empty();
  }
  const range = document.createRange();
  range.selectNode(element);
  window.getSelection().addRange(range);
}

document
  .querySelectorAll(".selectable")
  .forEach((element) => {
    element.addEventListener("mouseover", () => {
      selectText(element);
    }, false)
  });

new Clipboard('.selectable');
