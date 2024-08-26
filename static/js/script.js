// This event is triggered when the URL changes
// It will be responsible for changing the active button in the sidebar
document.addEventListener("htmx:replacedInHistory", function(evt) {
  document.querySelectorAll(".sidebar-btn").forEach(btn => {
    btn.classList.add("text-white");
    btn.classList.remove("text-orange-400");
  });

  const btn = document.querySelector(`[data-path='${evt.detail.path}']`)

  if (btn && btn.classList.contains("sidebar-btn")) {
    btn.classList.remove("text-white");
    btn.classList.add("text-orange-400");
  }
});



let isListenerAttached = false;
function handleCtrlK(evt) {
  if (evt.ctrlKey && evt.key === 'k') {
    evt.preventDefault();
    const modal = document.getElementById("searchModal")
    if (modal) return;
    htmx.trigger(htmx.find('button'), 'ctrlK');
  }
}
if (!isListenerAttached) {
  document.addEventListener('keydown', handleCtrlK);
  isListenerAttached = true;
}


// open search modal
document.addEventListener('keydown', function(evt) {
  if (evt.key !== "Escape") return
  const modal = document.getElementById("searchModal")
  modal.remove()
})

// close search modal on outside click
document.addEventListener('click', function(evt) {
  if (evt.target.id === "searchModal") evt.target.remove()
})

function getSearchParams() {
  const searchInput = document.getElementById("searchInput")
  return searchInput.value
}

// search option trigger to change page
document.addEventListener("htmx:afterRequest", function(evt) {
  const element = evt.detail.elt

  if (element.classList.contains("search-option")) {
    console.log('got here')
    const modal = document.getElementById("searchModal")
    modal.remove()
  }
});
