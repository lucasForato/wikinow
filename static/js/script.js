// Sidebar active link
document.addEventListener("htmx:afterRequest", function(event) {
  if (!event.detail.elt.classList.contains("sidebar-btn")) return

  document.querySelectorAll(".sidebar-btn").forEach(btn => {
    btn.classList.add("text-white");
    btn.classList.remove("text-orange-400");
  });

  const btn = event.detail.elt;

  if (btn && btn.classList.contains("sidebar-btn")) {
    btn.classList.remove("text-white");
    btn.classList.add("text-orange-400");
  }
});

// Ctrl + K to focus search
document.addEventListener('keydown', function(event) {
  if (event.ctrlKey && event.key === 'k') {
    event.preventDefault();
    htmx.trigger(htmx.find('button'), 'ctrlK');
  }
});

// open search modal
document.addEventListener('keydown', function(event) {
  if (event.key !== "Escape") return
  const modal = document.getElementById("searchModal")
  console.log('modal', modal)
  modal.remove()
})

// close search modal on outside click
document.addEventListener('click', function(event) {
  if (event.target.id === "searchModal") event.target.remove()
})
